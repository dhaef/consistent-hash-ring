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

	ring.RemoveNode(nodeOneId)

	if len(ring.SortedKeys) != 0 {
		t.Errorf("expected length 0 got %v", len(ring.SortedKeys))
	}

	if ring.Ring[5] != "" {
		t.Errorf("expected ring[5] to be an empty string got %v", ring.Ring[5])
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
