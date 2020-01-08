package test

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/jsonpb"

	patch "github.com/bitspill/protoPatch"
	"github.com/bitspill/protoPatch/test/proto"
)

func TestPrimitive(t *testing.T) {
	origVal := &test_proto.Primitives{
		One:      "foo",
		Two:      2,
		Three:    1.3,
		Four:     true,
		Corpus:   test_proto.Primitives_NEWS,
		Rep:      []int64{1, 2, 3},
		MapField: map[int64]string{1: "hello", 2: "world"},
		Sm:       &test_proto.SimpleMessage{One: 2},
		Rsm:      []*test_proto.SimpleMessage{{One: 4}, {One: 5}, {One: 6}, {One: 7}},
		Msm:      map[string]*test_proto.SimpleMessage{"one-one": {One: 1}, "one-two": {One: 1}},
	}

	newVal := &test_proto.Primitives{
		One:      "bar",
		Two:      4,
		Three:    3.14,
		Four:     false,
		Corpus:   test_proto.Primitives_VIDEO,
		Rep:      []int64{4, 5, 6},
		Sm:       &test_proto.SimpleMessage{One: 1},
		MapField: map[int64]string{2: "darkness"},
		Rsm:      []*test_proto.SimpleMessage{{One: 1}},
		Msm:      map[string]*test_proto.SimpleMessage{"one-two": {One: 2}},
	}

	p := patch.Patch{
		NewValues: newVal,
		Ops: []patch.Op{
			{
				Path: []patch.Step{
					{
						Tag:    10,
						Action: patch.ActionReplaceOne,
						MapKey: "one-two",
					},
				},
			},
			{
				Path: []patch.Step{
					{
						Action: patch.ActionReplace,
						Name:   "three",
						// Tag: 3,
					},
				},
			},
			{
				Path: []patch.Step{
					{
						Tag:      9,
						Action:   patch.ActionStepInto,
						SrcIndex: 1,
						DstIndex: 2,
					},
					{
						Action: patch.ActionReplace,
						Tag:    1,
					},
				},
			},
			{
				Path: []patch.Step{
					{
						Action: patch.ActionReplace,
						Tag:    1,
					},
				},
			},
			{
				Path: []patch.Step{
					{
						Action: patch.ActionReplace,
						Tag:    2,
					},
				},
			},
			{
				Path: []patch.Step{
					{
						Action: patch.ActionReplace,
						Tag:    3,
					},
				},
			},
			{
				Path: []patch.Step{
					{
						Action: patch.ActionReplace,
						Tag:    4,
					},
				},
			},
			{
				Path: []patch.Step{
					{
						Action: patch.ActionReplace,
						Tag:    5,
					},
				},
			},
			{
				Path: []patch.Step{
					{
						Tag:    6,
						Action: patch.ActionReplace,
					},
				},
			},
			{
				Path: []patch.Step{
					{
						// Tag:      7,
						Name: "map_field",
						// JsonName: "mapField",
						Action: patch.ActionReplaceOne,
						MapKey: int32(2),
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
}

func TestNest(t *testing.T) {
	con := &test_proto.ContainerMessage{
		Cm: &test_proto.SimpleMessage{
			One: 2,
		},
		Two: 1,
	}

	newCon := &test_proto.ContainerMessage{
		Cm: &test_proto.SimpleMessage{
			One: 1,
		},
		Two: 2,
	}

	p := patch.Patch{
		NewValues: newCon,
		Ops: []patch.Op{
			{
				Path: []patch.Step{
					{
						Name:   "two",
						Action: patch.ActionReplace,
					},
				},
			},
			{
				Path: []patch.Step{
					{
						Name:   "cm",
						Action: patch.ActionStepInto,
					},
					{
						Name:   "one",
						Action: patch.ActionReplace,
					},
				},
			},
		},
	}

	r, err := patch.ApplyPatch(p, con)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(con)
	fmt.Println(newCon)
	fmt.Println(r)
}
