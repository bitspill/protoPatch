package patch

import (
	"errors"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
)

type Patch struct {
	NewValues proto.Message
	Ops       []*Op
}

type Op struct {
	Path []Step
}

type Step struct {
	Tag int32

	Action   Action
	SrcIndex int
	DstIndex int
	MapKey   interface{} // integral or string
}

type Action int8

const (
	ActionInvalid Action = iota
	ActionReplaceAll
	ActionAppendAll
	ActionRemoveAll
	ActionRemoveOne
	ActionReplaceOne
	ActionAppendOne
	ActionStrPatch // ToDo: implement diff-match-patch/unix-diff for string primitive type
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
		var iterStack []stackEntry
		dIterator := dstDynamic
		sIterator := srcDynamic
		for _, step := range op.Path {
			fmt.Println(step)
			fieldDescriptor := dIterator.FindFieldDescriptor(step.Tag)
			if fieldDescriptor == nil {
				return nil, err
			}

			t := fieldDescriptor.GetType()

			if t == descriptor.FieldDescriptorProto_TYPE_GROUP {
				return nil, errors.New("group type unsupported")
			}

			if t == descriptor.FieldDescriptorProto_TYPE_MESSAGE && step.Action == ActionStepInto {
				se := stackEntry{
					Iter: dIterator,
					Fd:   fieldDescriptor,
				}

				var sf interface{}
				var df interface{}

				switch {
				case fieldDescriptor.IsMap():
					sf, err = sIterator.TryGetMapField(fieldDescriptor, step.MapKey)
					if err != nil {
						return nil, err
					}
					df, err = dIterator.TryGetMapField(fieldDescriptor, step.MapKey)
					if err != nil {
						return nil, err
					}
					se.Key = step.MapKey
				case fieldDescriptor.IsRepeated():
					sf, err = sIterator.TryGetRepeatedField(fieldDescriptor, step.SrcIndex)
					if err != nil {
						return nil, err
					}
					df, err = dIterator.TryGetRepeatedField(fieldDescriptor, step.DstIndex)
					if err != nil {
						return nil, err
					}
					se.Index = step.DstIndex
				default:
					sf, err = sIterator.TryGetField(fieldDescriptor)
					if err != nil {
						return nil, err
					}
					df, err = dIterator.TryGetField(fieldDescriptor)
					if err != nil {
						return nil, err
					}
				}

				psf, ok := sf.(proto.Message)
				if !ok {
					return nil, errors.New("field not a message")
				}

				sIterator, err = toDynamic(psf)
				if err != nil {
					return nil, err
				}

				pdf, ok := df.(proto.Message)
				if !ok {
					return nil, errors.New("field not a message")
				}

				dIterator, err = toDynamic(pdf)
				if err != nil {
					return nil, err
				}

				iterStack = append(iterStack, se)

				continue
			}

			err := processPrimitive(step, dIterator, fieldDescriptor, sIterator)
			if err != nil {
				return nil, err
			}

			if len(iterStack) > 0 {
				iterStack[len(iterStack)-1].Res = dIterator
			}
		}

		// Unwind iter stack of modified nested messages
		for _, stackVal := range iterStack {
			var err error
			switch {
			case stackVal.Fd.IsMap():
				err = stackVal.Iter.TryPutMapField(stackVal.Fd, stackVal.Key, stackVal.Res)
			case stackVal.Fd.IsRepeated():
				err = stackVal.Iter.TrySetRepeatedField(stackVal.Fd, stackVal.Index, stackVal.Res)
			default:
				err = stackVal.Iter.TrySetField(stackVal.Fd, stackVal.Res)
			}
			if err != nil {
				return nil, err
			}
		}
	}

	clone := proto.Clone(dst)
	err = dstDynamic.ConvertTo(clone)
	if err != nil {
		return nil, err
	}
	return clone, nil
}

func processPrimitive(step Step, dIterator *dynamic.Message, fd *desc.FieldDescriptor, sIterator *dynamic.Message) error {
	var err error
	switch {
	case step.Action == ActionRemoveAll:
		err = dIterator.TryClearField(fd)
	case step.Action == ActionReplaceAll:
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
	case ActionAppendAll:
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
	case ActionReplaceAll:
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
	case ActionAppendAll:
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

func toDynamic(p proto.Message) (*dynamic.Message, error) {
	dm, err := dynamic.AsDynamicMessage(p)
	if err != nil {
		return nil, err
	}
	return dm, nil
}

type stackEntry struct {
	Iter  *dynamic.Message
	Fd    *desc.FieldDescriptor
	Res   *dynamic.Message
	Index int
	Key   interface{}
}
