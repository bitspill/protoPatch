package patch

import (
	"errors"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

func ToProto(patch *Patch) (proto.Message, error) {
	newValAny, err := ptypes.MarshalAny(patch.NewValues)
	if err != nil {
		return nil, err
	}

	pp := &ProtoPatch{
		NewValues: newValAny,
		Ops:       make([]*ProtoOp, len(patch.Ops)),
	}

	for _, op := range patch.Ops {
		pop := &ProtoOp{
			Path: make([]*ProtoStep, len(op.Path)),
		}
		for _, step := range op.Path {
			pStep := &ProtoStep{
				Tag:      step.Tag,
				Name:     step.Name,
				JsonName: step.JsonName,
				SrcIndex: int32(step.SrcIndex),
				DstIndex: int32(step.DstIndex),
			}

			pStep.Action = actionToProtoAction(step.Action)

			switch k := step.MapKey.(type) {
			case int32:
				pStep.MapKey = &ProtoStep_IntMapKey{IntMapKey: int64(k)}
			case int64:
				pStep.MapKey = &ProtoStep_IntMapKey{IntMapKey: k}
			case uint32:
				pStep.MapKey = &ProtoStep_UIntMapKey{UIntMapKey: uint64(k)}
			case uint64:
				pStep.MapKey = &ProtoStep_UIntMapKey{UIntMapKey: k}
			case bool:
				pStep.MapKey = &ProtoStep_BoolMapKey{BoolMapKey: k}
			case string:
				pStep.MapKey = &ProtoStep_StrMapKey{StrMapKey: k}
			}

			pop.Path = append(pop.Path, pStep)
		}
		pp.Ops = append(pp.Ops, pop)
	}

	return pp, nil
}

func FromProto(message proto.Message) (*Patch, error) {
	pp, ok := message.(*ProtoPatch)
	if !ok {
		return nil, errors.New("not a proto patch")
	}

	dynAny := &ptypes.DynamicAny{}
	err := ptypes.UnmarshalAny(pp.NewValues, dynAny)
	if err != nil {
		return nil, err
	}

	patch := &Patch{
		NewValues: dynAny.Message,
		Ops:       make([]Op, len(pp.Ops)),
	}

	for _, protoOp := range pp.Ops {
		op := Op{
			Path: make([]Step, len(protoOp.Path)),
		}
		for _, protoStep := range protoOp.Path {
			step := Step{
				Tag:      protoStep.Tag,
				Name:     protoStep.Name,
				JsonName: protoStep.JsonName,
				Action:   protoActionToAction(protoStep.Action),
				SrcIndex: int(protoStep.SrcIndex),
				DstIndex: int(protoStep.DstIndex),
				MapKey:   getMapKey(protoStep),
			}
			op.Path = append(op.Path, step)
		}
		patch.Ops = append(patch.Ops, op)
	}

	return patch, nil
}

func actionToProtoAction(action Action) ProtoAction {
	switch action {
	case ActionInvalid:
		return ProtoAction_ActionInvalid
	case ActionReplace:
		return ProtoAction_ActionReplace
	case ActionAppend:
		return ProtoAction_ActionAppend
	case ActionRemove:
		return ProtoAction_ActionRemove
	case ActionRemoveOne:
		return ProtoAction_ActionRemoveOne
	case ActionReplaceOne:
		return ProtoAction_ActionReplaceOne
	case ActionAppendOne:
		return ProtoAction_ActionAppendOne
	case ActionStrPatch:
		return ProtoAction_ActionStrPatch
	case ActionStepInto:
		return ProtoAction_ActionStepInto
	default:
		return ProtoAction_ActionInvalid
	}
}

func protoActionToAction(action ProtoAction) Action {
	switch action {
	case ProtoAction_ActionInvalid:
		return ActionInvalid
	case ProtoAction_ActionReplace:
		return ActionReplace
	case ProtoAction_ActionAppend:
		return ActionAppend
	case ProtoAction_ActionRemove:
		return ActionRemove
	case ProtoAction_ActionRemoveOne:
		return ActionRemoveOne
	case ProtoAction_ActionReplaceOne:
		return ActionReplaceOne
	case ProtoAction_ActionAppendOne:
		return ActionAppendOne
	case ProtoAction_ActionStrPatch:
		return ActionStrPatch
	case ProtoAction_ActionStepInto:
		return ActionStepInto
	default:
		return ActionInvalid
	}
}

func getMapKey(step *ProtoStep) interface{} {
	mapKey := step.GetMapKey()
	if mapKey == nil {
		return nil
	}

	switch key := mapKey.(type) {
	case *ProtoStep_IntMapKey:
		return key.IntMapKey
	case *ProtoStep_BoolMapKey:
		return key.BoolMapKey
	case *ProtoStep_StrMapKey:
		return key.StrMapKey
	case *ProtoStep_UIntMapKey:
		return key.UIntMapKey
	default:
		return nil
	}
}
