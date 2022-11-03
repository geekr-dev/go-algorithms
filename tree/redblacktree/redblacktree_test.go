package redblacktree_test

import (
	"fmt"
	"testing"

	"github.com/geekr-dev/go-algorithms/tree/redblacktree"
)

// IntComparator provides a basic comparison on int
func IntComparator(a, b interface{}) int {
	aAsserted := a.(int)
	bAsserted := b.(int)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

func TestRedBlackTreePut(t *testing.T) {
	tree := redblacktree.NewTree(IntComparator)
	tree.Insert(5, "e")
	tree.Insert(6, "f")
	tree.Insert(7, "g")
	tree.Insert(3, "c")
	tree.Insert(4, "d")
	tree.Insert(1, "x")
	tree.Insert(2, "b")
	tree.Insert(1, "a") // overwrite

	if actualValue := tree.Size(); actualValue != 7 {
		t.Errorf("Got %v expected %v", actualValue, 7)
	}
	if actualValue, expectedValue := fmt.Sprintf("%d%d%d%d%d%d%d", tree.Keys()...), "1234567"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := fmt.Sprintf("%s%s%s%s%s%s%s", tree.Values()...), "abcdefg"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	tests1 := [][]interface{}{
		{1, "a"},
		{2, "b"},
		{3, "c"},
		{4, "d"},
		{5, "e"},
		{6, "f"},
		{7, "g"},
		{8, nil},
	}

	for _, test := range tests1 {
		// retrievals
		actualNode := tree.Get(test[0])
		if actualNode != nil && actualNode.Value != test[1] {
			t.Errorf("Got %v expected %v", actualNode.Value, test[1])
		}
	}
}

func TestRedBlackTreeDelete(t *testing.T) {
	tree := redblacktree.NewTree(IntComparator)
	tree.Insert(5, "e")
	tree.Insert(6, "f")
	tree.Insert(7, "g")
	tree.Insert(3, "c")
	tree.Insert(4, "d")
	tree.Insert(1, "x")
	tree.Insert(2, "b")
	tree.Insert(1, "a") // overwrite

	tree.Delete(5)
	tree.Delete(6)
	tree.Delete(7)
	tree.Delete(8)
	tree.Delete(5)

	if actualValue, expectedValue := fmt.Sprintf("%d%d%d%d", tree.Keys()...), "1234"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := fmt.Sprintf("%s%s%s%s", tree.Values()...), "abcd"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := fmt.Sprintf("%s%s%s%s", tree.Values()...), "abcd"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := tree.Size(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 7)
	}

	tests2 := [][]interface{}{
		{1, "a"},
		{2, "b"},
		{3, "c"},
		{4, "d"},
		{5, nil},
		{6, nil},
		{7, nil},
		{8, nil},
	}

	for _, test := range tests2 {
		actualNode := tree.Get(test[0])
		if actualNode != nil && actualNode.Value != test[1] {
			t.Errorf("Got %v expected %v", actualNode.Value, test[1])
		}
	}

	tree.Delete(1)
	tree.Delete(4)
	tree.Delete(2)
	tree.Delete(3)
	tree.Delete(2)
	tree.Delete(2)

	if actualValue, expectedValue := fmt.Sprintf("%s", tree.Keys()), "[]"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := fmt.Sprintf("%s", tree.Values()), "[]"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if empty, size := tree.Empty(), tree.Size(); empty != true || size != -0 {
		t.Errorf("Got %v expected %v", empty, true)
	}

}
