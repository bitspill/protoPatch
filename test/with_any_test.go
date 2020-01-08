package test

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"

	patch "github.com/bitspill/protoPatch"
	"github.com/bitspill/protoPatch/test/proto"
)

func TestAny(t *testing.T) {
	origA := &test_proto.TestMessage{Num: 1}
	origRa1 := &test_proto.TestMessage{Num: 6}
	origRa2 := &test_proto.TestMessage{Num: 15}
	origRa3 := &test_proto.TestMessage{Num: 19}

	anyOrigA, err := ptypes.MarshalAny(origA)
	if err != nil {
		t.Fatal(err)
	}
	anyOrigRa1, err := ptypes.MarshalAny(origRa1)
	if err != nil {
		t.Fatal(err)
	}
	anyOrigRa2, err := ptypes.MarshalAny(origRa2)
	if err != nil {
		t.Fatal(err)
	}
	anyOrigRa3, err := ptypes.MarshalAny(origRa3)
	if err != nil {
		t.Fatal(err)
	}

	origVal := &test_proto.WithAny{
		A:  anyOrigA,
		Ra: []*any.Any{anyOrigRa1, anyOrigRa2, anyOrigRa3},
	}

	newA := &test_proto.TestMessage{Num: 2}
	newRa := &test_proto.TestMessage{Num: 6}

	anyNewA, err := ptypes.MarshalAny(newA)
	if err != nil {
		t.Fatal(err)
	}
	anyNewRa, err := ptypes.MarshalAny(newRa)
	if err != nil {
		t.Fatal(err)
	}

	newVal := &test_proto.WithAny{
		A:  anyNewA,
		Ra: []*any.Any{anyNewRa},
	}

	p := patch.Patch{
		NewValues: newVal,
		Ops: []patch.Op{
			{
				Path: []patch.Step{
					{
						Tag:    1,
						Action: patch.ActionStepInto,
					},
					{
						Tag:    1,
						Action: patch.ActionReplace,
					},
				},
			},
			{
				Path: []patch.Step{
					{
						Name:     "ra",
						Action:   patch.ActionStepInto,
						DstIndex: 2,
						SrcIndex: 1,
					},
					{
						Name:   "num",
						Action: patch.ActionReplace,
					},
				},
			},
			{
				Path: []patch.Step{
					{
						Name:     "ra",
						Action:   patch.ActionStepInto,
						DstIndex: 3,
						SrcIndex: 1,
					},
					{
						Name:   "num",
						Action: patch.ActionReplace,
					},
				},
			},
		},
	}

	res, err := patch.ApplyPatch(p, origVal)
	if err != nil {
		t.Fatal(err)
	}

	s, err := (&jsonpb.Marshaler{}).MarshalToString(res)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(s)
	// fmt.Println(res)

	wa := res.(*test_proto.WithAny)
	da := &ptypes.DynamicAny{}
	err = ptypes.UnmarshalAny(wa.A, da)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(da)

	expectedResult := &test_proto.WithAny{
		A:  anyNewA,
		Ra: []*any.Any{anyNewRa, anyNewRa, anyNewRa},
	}

	es, err := (&jsonpb.Marshaler{}).MarshalToString(expectedResult)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(es)

	if !proto.Equal(res, expectedResult) {
		t.Fatal("!=")
	}
}

func TestNestedAny(t *testing.T) {
	t.SkipNow()
	origA := &test_proto.TestMessage{Num: 1}
	// origRa1 := &test_proto.TestMessage{Num: 6}
	// origRa2 := &test_proto.TestMessage{Num: 15}
	// origRa3 := &test_proto.TestMessage{Num: 19}

	anyOrigA, err := ptypes.MarshalAny(origA)
	if err != nil {
		t.Fatal(err)
	}
	// anyOrigRa1, err := ptypes.MarshalAny(origRa1)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// anyOrigRa2, err := ptypes.MarshalAny(origRa2)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// anyOrigRa3, err := ptypes.MarshalAny(origRa3)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	nestVal := &test_proto.WithAny{
		A: anyOrigA,
		// Ra: []*any.Any{anyOrigRa1, anyOrigRa2, anyOrigRa3},
	}

	nestValAny, err := ptypes.MarshalAny(nestVal)
	if err != nil {
		t.Fatal(err)
	}

	origVal := &test_proto.WithAny{
		A: nestValAny,
	}

	newA := &test_proto.TestMessage{Num: 2}
	// newRa := &test_proto.TestMessage{Num: 6}

	anyNewA, err := ptypes.MarshalAny(newA)
	if err != nil {
		t.Fatal(err)
	}
	// anyNewRa, err := ptypes.MarshalAny(newRa)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	nestNewVal := &test_proto.WithAny{
		A: anyNewA,
		// Ra: []*any.Any{anyNewRa},
	}
	nestNewValAny, err := ptypes.MarshalAny(nestNewVal)
	if err != nil {
		t.Fatal(err)
	}

	newVal := &test_proto.WithAny{
		A: nestNewValAny,
	}

	p := patch.Patch{
		NewValues: newVal,
		Ops: []patch.Op{
			{
				Path: []patch.Step{
					{
						Tag:    1, // a
						Action: patch.ActionStepInto,
					},
					{
						Tag:    1, // a
						Action: patch.ActionStepInto,
					},
					{
						Tag:    1, // num
						Action: patch.ActionReplace,
					},
				},
			},
			// {
			// 	Path: []patch.Step{
			// 		{
			// 			Tag:    1, // a
			// 			Action: patch.ActionStepInto,
			// 		},
			// 		{
			// 			Name:     "ra",
			// 			Action:   patch.ActionStepInto,
			// 			DstIndex: 2,
			// 			SrcIndex: 1,
			// 		},
			// 		{
			// 			Name:   "num",
			// 			Action: patch.ActionReplace,
			// 		},
			// 	},
			// },
			// {
			// 	Path: []patch.Step{
			// 		{
			// 			Tag:    1, // a
			// 			Action: patch.ActionStepInto,
			// 		},
			// 		{
			// 			Name:     "ra",
			// 			Action:   patch.ActionStepInto,
			// 			DstIndex: 3,
			// 			SrcIndex: 1,
			// 		},
			// 		{
			// 			Name:   "num",
			// 			Action: patch.ActionReplace,
			// 		},
			// 	},
			// },
		},
	}

	res, err := patch.ApplyPatch(p, origVal)
	if err != nil {
		t.Fatal(err)
	}

	s, err := (&jsonpb.Marshaler{}).MarshalToString(res)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(s)
	// fmt.Println(res)

	wa := res.(*test_proto.WithAny)
	da := &ptypes.DynamicAny{}
	err = ptypes.UnmarshalAny(wa.A, da)
	if err != nil {
		t.Fatal(err)
	}

	expectedNestedResult := &test_proto.WithAny{
		A: anyNewA,
		// Ra: []*any.Any{anyNewRa, anyNewRa, anyNewRa},
		// Ra: []*any.Any{anyOrigRa1, anyOrigRa2, anyOrigRa3},
	}

	expectedNestedAny, err := ptypes.MarshalAny(expectedNestedResult)
	if err != nil {
		t.Fatal(err)
	}

	expectedResult := &test_proto.WithAny{
		A: expectedNestedAny,
	}

	es, err := (&jsonpb.Marshaler{}).MarshalToString(expectedResult)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(es)

	if !proto.Equal(res, expectedResult) {
		t.Fatal("!=")
	}
}
