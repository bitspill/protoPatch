package patch

import (
	"errors"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
)

type Patch struct {
	NewValues proto.Message
	Ops       []Op
}

type Op struct {
	Path []Step
}

type Step struct {
	Tag      int32
	Name     string
	JsonName string

	Action   Action
	SrcIndex int
	DstIndex int
	MapKey   interface{} // integral or string
}

type Action int8

const (
	ActionInvalid Action = iota
	ActionReplace
	ActionAppend
	ActionRemove
	ActionRemoveOne  // Only valid for repeated or map fields
	ActionReplaceOne // Only valid for repeated or map fields
	ActionAppendOne  // Only valid for repeated or map fields
	ActionStrPatch   // ToDo: implement diff-match-patch/unix-diff for string primitive type
	ActionStepInto
)

func ApplyPatch(patch Patch, dst proto.Message) (proto.Message, error) {
	srcDynamic, err := toDynamic(patch.NewValues)
	if err != nil {
		return nil, err
	}
	dstDynamic, err := toDynamic(dst)
	if err != nil {
		return nil, err
	}

	for _, op := range patch.Ops {
		var iterStack []*stackEntry
		dIterator := dstDynamic
		sIterator := srcDynamic
		for _, step := range op.Path {
			fmt.Println(step)
			fieldDescriptor, err := getFieldDescriptor(&step, dIterator)
			if err != nil {
				return nil, err
			}

			t := fieldDescriptor.GetType()

			if t == descriptor.FieldDescriptorProto_TYPE_GROUP {
				// https://developers.google.com/protocol-buffers/docs/proto#groups
				// Note that this feature is deprecated and should not be used
				// when creating new message types – use nested message types instead.
				return nil, errors.New("group type unsupported")
			}

			if t == descriptor.FieldDescriptorProto_TYPE_MESSAGE && step.Action == ActionStepInto {
				var se *stackEntry
				se, sIterator, dIterator, err = stepIntoMessage(fieldDescriptor, step, sIterator, dIterator)
				if err != nil {
					return nil, err
				}

				iterStack = append(iterStack, se)
			} else {
				err = processPrimitive(step, dIterator, fieldDescriptor, sIterator)
				if err != nil {
					return nil, err
				}
			}

			if len(iterStack) > 0 {
				iterStack[len(iterStack)-1].Res = dIterator
			}
		}

		err = applyStack(iterStack)
		if err != nil {
			return nil, err
		}
	}

	clone := proto.Clone(dst)
	err = dstDynamic.ConvertTo(clone)
	if err != nil {
		return nil, err
	}
	return clone, nil
}

func stepIntoMessage(fieldDescriptor *desc.FieldDescriptor, step Step, sIterator *dynamic.Message, dIterator *dynamic.Message) (*stackEntry, *dynamic.Message, *dynamic.Message, error) {
	se := &stackEntry{
		Iter:  dIterator,
		Fd:    fieldDescriptor,
		Key:   step.MapKey,
		Index: step.DstIndex - 1,
	}

	sfv, dfv, err := getFieldValue(fieldDescriptor, step, sIterator, dIterator)
	if err != nil {
		return nil, nil, nil, err
	}

	if fdIsAny(fieldDescriptor) {
		// if the field is an any.Any we want the message
		// contained within not the any.Any itself
		sfv, err = extractInnerMessageFromAny(sfv)
		if err != nil {
			return nil, nil, nil, err
		}
		dfv, err = extractInnerMessageFromAny(dfv)
		if err != nil {
			return nil, nil, nil, err
		}
	}

	sIterator, err = toDynamic(sfv)
	if err != nil {
		return nil, nil, nil, err
	}

	dIterator, err = toDynamic(dfv)
	if err != nil {
		return nil, nil, nil, err
	}

	return se, sIterator, dIterator, nil
}

func getFieldValue(fieldDescriptor *desc.FieldDescriptor, step Step, sIterator *dynamic.Message, dIterator *dynamic.Message) (interface{}, interface{}, error) {
	var sf interface{}
	var df interface{}
	var err error

	switch {
	case fieldDescriptor.IsMap():
		sf, err = sIterator.TryGetMapField(fieldDescriptor, step.MapKey)
		if err != nil {
			return nil, nil, err
		}
		df, err = dIterator.TryGetMapField(fieldDescriptor, step.MapKey)
		if err != nil {
			return nil, nil, err
		}
	case fieldDescriptor.IsRepeated():
		sf, err = sIterator.TryGetRepeatedField(fieldDescriptor, step.SrcIndex-1)
		if err != nil {
			return nil, nil, err
		}
		df, err = dIterator.TryGetRepeatedField(fieldDescriptor, step.DstIndex-1)
		if err != nil {
			return nil, nil, err
		}
	default:
		sf, err = sIterator.TryGetField(fieldDescriptor)
		if err != nil {
			return nil, nil, err
		}
		df, err = dIterator.TryGetField(fieldDescriptor)
		if err != nil {
			return nil, nil, err
		}
	}
	return sf, df, err
}

func applyStack(iterStack []*stackEntry) error {
	for _, stackVal := range iterStack {
		res := stackVal.Res

		if fdIsAny(stackVal.Fd) {
			resAny, err := ptypes.MarshalAny(stackVal.Res)
			if err != nil {
				return err
			}
			// res, err = toDynamic(resAny)
			// if err != nil {
			// 	return err
			// }
			res = resAny
		}

		var err error
		switch {
		case stackVal.Fd.IsMap():
			err = stackVal.Iter.TryPutMapField(stackVal.Fd, stackVal.Key, res)
		case stackVal.Fd.IsRepeated():
			err = stackVal.Iter.TrySetRepeatedField(stackVal.Fd, stackVal.Index, res)
		default:
			err = stackVal.Iter.TrySetField(stackVal.Fd, res)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func extractInnerMessageFromAny(a interface{}) (proto.Message, error) {
	sAny, ok := a.(*any.Any)
	if !ok {
		return nil, errors.New("not an any")
	}

	dynAny := &ptypes.DynamicAny{}
	err := ptypes.UnmarshalAny(sAny, dynAny)
	if err != nil {
		return nil, err
	}

	return dynAny.Message, nil
}

func getFieldDescriptor(step *Step, dIterator *dynamic.Message) (*desc.FieldDescriptor, error) {
	var fieldDescriptor *desc.FieldDescriptor
	if step.Tag != 0 {
		fieldDescriptor = dIterator.FindFieldDescriptor(step.Tag)
	} else if step.Name != "" {
		fieldDescriptor = dIterator.FindFieldDescriptorByName(step.Name)
	} else {
		fieldDescriptor = dIterator.FindFieldDescriptorByJSONName(step.JsonName)
	}

	if fieldDescriptor == nil {
		return nil, errors.New("field not found")
	}

	return fieldDescriptor, nil
}

func processPrimitive(step Step, dIterator *dynamic.Message, fd *desc.FieldDescriptor, sIterator *dynamic.Message) error {
	var err error
	switch {
	case step.Action == ActionRemove:
		err = dIterator.TryClearField(fd)
	case step.Action == ActionReplace:
		var v interface{}
		v, err = sIterator.TryGetField(fd)
		if err == nil {
			err = dIterator.TrySetField(fd, v)
		}
	case fd.IsMap():
		err = processMapPrimitive(step, sIterator, fd, dIterator)
	case fd.IsRepeated():
		err = processRepeatedPrimitive(step, sIterator, fd, dIterator)
	}

	if err != nil {
		return err
	}

	return nil
}

func processRepeatedPrimitive(step Step, sIterator *dynamic.Message, fd *desc.FieldDescriptor, dIterator *dynamic.Message) error {
	switch step.Action {
	case ActionAppend:
		interfaceVal, err := sIterator.TryGetField(fd)
		if err != nil {
			return err
		}
		repeatedVal, ok := interfaceVal.([]interface{})
		if !ok {
			return errors.New("expected slice")
		}
		for _, v := range repeatedVal {
			err = dIterator.TryAddRepeatedField(fd, v)
			if err != nil {
				return err
			}
		}
	case ActionAppendOne:
		v, err := sIterator.TryGetRepeatedField(fd, step.SrcIndex)
		if err != nil {
			return err
		}
		err = dIterator.TryAddRepeatedField(fd, v)
		if err != nil {
			return err
		}
	case ActionRemoveOne:
		l, err := dIterator.TryFieldLength(fd)
		if err != nil {
			return err
		}
		if step.DstIndex >= l || step.DstIndex < 0 {
			return errors.New("dst out of range")
		}
		interfaceVal, err := sIterator.TryGetField(fd)
		if err != nil {
			return err
		}
		repeatedVal, ok := interfaceVal.([]interface{})
		if !ok {
			return errors.New("expected slice")
		}
		newVal := append(repeatedVal[:step.DstIndex], repeatedVal[step.DstIndex+1:]...)
		err = dIterator.TrySetField(fd, newVal)
		if err != nil {
			return err
		}
	case ActionReplaceOne:
		v, err := sIterator.TryGetRepeatedField(fd, step.SrcIndex)
		if err != nil {
			return err
		}
		dl, err := dIterator.TryFieldLength(fd)
		if err != nil {
			return err
		}
		if step.DstIndex >= dl || step.DstIndex < 0 {
			return errors.New("dst out of range")
		}
		err = dIterator.TrySetRepeatedField(fd, step.DstIndex-1, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func processMapPrimitive(step Step, sIterator *dynamic.Message, fd *desc.FieldDescriptor, dIterator *dynamic.Message) error {
	switch step.Action {
	case ActionReplace:
		v, err := sIterator.TryGetField(fd)
		if err != nil {
			return err
		}
		err = dIterator.TrySetField(fd, v)
		if err != nil {
			return err
		}
	case ActionReplaceOne,
		ActionAppendOne:
		v, err := sIterator.TryGetMapField(fd, step.MapKey)
		if err != nil {
			return err
		}
		err = dIterator.TryPutMapField(fd, step.MapKey, v)
		if err != nil {
			return err
		}
	case ActionRemoveOne:
		err := dIterator.TryRemoveMapField(fd, step.MapKey)
		if err != nil {
			return err
		}
	case ActionAppend:
		iv, err := sIterator.TryGetField(fd)
		if err != nil {
			return err
		}
		mv, ok := iv.(map[interface{}]interface{})
		if !ok {
			return err
		}
		for key, value := range mv {
			err := dIterator.TryPutMapField(fd, key, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func toDynamic(i interface{}) (*dynamic.Message, error) {
	if p, ok := i.(proto.Message); ok {
		return dynamic.AsDynamicMessage(p)
	} else {
		return nil, errors.New("not a proto.Message")
	}
}

type stackEntry struct {
	Iter  *dynamic.Message
	Fd    *desc.FieldDescriptor
	Res   proto.Message
	Index int
	Key   interface{}
}

func fdIsAny(fd *desc.FieldDescriptor) bool {
	return fd.GetMessageType().
		GetFullyQualifiedName() == "google.protobuf.Any"
}
