package ring

import (
	"testing"
)

func TestAddNode(t *testing.T) {
	ring := NewConsistentHasRing(10)
	nodeOneId := "123"

	ring.AddNode(nodeOneId)

	if len(ring.SortedKeys) != 1 {
		t.Errorf("expected length 1 got %v", len(ring.SortedKeys))
	}

	if ring.Ring[5] != buildNodeId(nodeOneId) {
		t.Errorf("expected ring[5] to be %v got %v", buildNodeId(nodeOneId), ring.Ring[5])
	}

	nodeTwoId := "456"

	ring.AddNode(nodeTwoId)

	if len(ring.SortedKeys) != 2 {
		t.Errorf("expected length 2 got %v", len(ring.SortedKeys))
	}

	if ring.Ring[8] != buildNodeId(nodeTwoId) {
		t.Errorf("expected ring[8] to be %v got %v", buildNodeId(nodeTwoId), ring.Ring[8])
	}

	nodeThreeId := "789"

	ring.AddNode(nodeThreeId)

	if len(ring.SortedKeys) != 3 {
		t.Errorf("expected length 3 got %v", len(ring.SortedKeys))
	}

	if ring.Ring[9] != buildNodeId(nodeThreeId) {
		t.Errorf("expected ring[9] to be %v got %v", buildNodeId(nodeThreeId), ring.Ring[9])
	}
}

func TestAddingNodeAtTakenPosition(t *testing.T) {
	ring := NewConsistentHasRing(10)
	nodeOneId := "123"

	ring.AddNode(nodeOneId)

	if len(ring.SortedKeys) != 1 {
		t.Errorf("expected length 1 got %v", len(ring.SortedKeys))
	}

	if ring.Ring[5] != buildNodeId(nodeOneId) {
		t.Errorf("expected ring[5] to be %v got %v", buildNodeId(nodeOneId), ring.Ring[5])
	}

	_, err := ring.AddNode(nodeOneId)

	if err == nil {
		t.Error("expected adding a node at a taken position to return an error")
	}
}

func TestRemoveNodeLengthOne(t *testing.T) {
	ring := NewConsistentHasRing(10)
	nodeOneId := "123"

	ring.AddNode(nodeOneId)

	if len(ring.SortedKeys) != 1 {
		t.Errorf("expected length 1 got %v", len(ring.SortedKeys))
	}

	if ring.Ring[5] != buildNodeId(nodeOneId) {
		t.Errorf("expected ring[5] to be %v got %v", buildNodeId(nodeOneId), ring.Ring[5])
	}

	ring.RemoveNode(nodeOneId)

	if len(ring.SortedKeys) != 0 {
		t.Errorf("expected length 0 got %v", len(ring.SortedKeys))
	}

	if ring.Ring[5] != "" {
		t.Errorf("expected ring[5] to be an empty string got %v", ring.Ring[5])
	}
}

func TestRemoveNodeLengthGreaterThanOne(t *testing.T) {
	ring := NewConsistentHasRing(10)
	nodeOneId := "123"

	ring.AddNode(nodeOneId)

	if len(ring.SortedKeys) != 1 {
		t.Errorf("expected length 1 got %v", len(ring.SortedKeys))
	}

	if ring.Ring[5] != buildNodeId(nodeOneId) {
		t.Errorf("expected ring[5] to be %v got %v", buildNodeId(nodeOneId), ring.Ring[5])
	}

	nodeTwoId := "456"

	ring.AddNode(nodeTwoId)

	if len(ring.SortedKeys) != 2 {
		t.Errorf("expected length 2 got %v", len(ring.SortedKeys))
	}

	if ring.Ring[8] != buildNodeId(nodeTwoId) {
		t.Errorf("expected ring[8] to be %v got %v", buildNodeId(nodeTwoId), ring.Ring[8])
	}

	ring.RemoveNode(nodeOneId)

	if len(ring.SortedKeys) != 1 {
		t.Errorf("expected length 1 got %v", len(ring.SortedKeys))
	}

	if ring.Ring[5] != "" {
		t.Errorf("expected ring[5] to be an empty string got %v", ring.Ring[5])
	}
}

func TestEmptyRing(t *testing.T) {
	ring := NewConsistentHasRing(10)
	key := "123"

	node, err := ring.GetNode(key)

	if err != nil {
		t.Errorf("expected err to be nil got %v", err.Error())
	}

	if node != "" {
		t.Errorf("expected node to be empty got %v", node)
	}
}

func TestGetNode(t *testing.T) {
	ring := NewConsistentHasRing(10)
	nodeOneId := "123"

	ring.AddNode(nodeOneId)

	if len(ring.SortedKeys) != 1 {
		t.Errorf("expected length 1 got %v", len(ring.SortedKeys))
	}

	if ring.Ring[5] != buildNodeId(nodeOneId) {
		t.Errorf("expected ring[5] to be %v got %v", buildNodeId(nodeOneId), ring.Ring[5])
	}

	nodeTwoId := "456"

	ring.AddNode(nodeTwoId)

	if len(ring.SortedKeys) != 2 {
		t.Errorf("expected length 2 got %v", len(ring.SortedKeys))
	}

	if ring.Ring[8] != buildNodeId(nodeTwoId) {
		t.Errorf("expected ring[8] to be %v got %v", buildNodeId(nodeTwoId), ring.Ring[8])
	}

	nodeThreeId := "789"

	ring.AddNode(nodeThreeId)

	if len(ring.SortedKeys) != 3 {
		t.Errorf("expected length 3 got %v", len(ring.SortedKeys))
	}

	if ring.Ring[9] != buildNodeId(nodeThreeId) {
		t.Errorf("expected ring[9] to be %v got %v", buildNodeId(nodeThreeId), ring.Ring[9])
	}

	keyOne := "foo"
	node, err := ring.GetNode(keyOne)

	if err != nil {
		t.Errorf("expected err to be nil got %v", err.Error())
	}

	if node != buildNodeId(nodeOneId) {
		t.Errorf("expected node to be %v got %v", buildNodeId(nodeOneId), err.Error())
	}

	keyTwo := "derek"
	node, err = ring.GetNode(keyTwo)

	if err != nil {
		t.Errorf("expected err to be nil got %v", err.Error())
	}

	if node != buildNodeId(nodeThreeId) {
		t.Errorf("expected node to be %v got %v", buildNodeId(nodeThreeId), err.Error())
	}

	keyThree := "a4a2a13a-5582-4666-a5b6-6958da9fcd53"
	node, err = ring.GetNode(keyThree)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err.Error())
	}

	if node != buildNodeId(nodeOneId) {
		t.Errorf("expected node to be %v got %v", buildNodeId(nodeOneId), err.Error())
	}

}
