package hash

import (
	"testing"
)

func TestHash(t *testing.T) {
	key := "hello"
	hash, err := Hash(key, 20)

	if err != nil {
		t.Errorf("got error %v", err.Error())
	}

	if hash != 14 {
		t.Errorf("expected n to be 14 got %v", hash)
	}

	key = "world"
	hash, err = Hash(key, 20)

	if err != nil {
		t.Errorf("got error %v", err.Error())
	}

	if hash != 11 {
		t.Errorf("expected n to be 11 got %v", hash)
	}
}

func TestNodeHash(t *testing.T) {
	nodeOne := "node_1"
	nodeTwo := "node_2"
	nodeThree := "node_3"
	nodeFour := "node_4"

	ringSize := 20

	nodeOnePosition, _ := Hash(nodeOne, ringSize)
	nodeTwoPosition, _ := Hash(nodeTwo, ringSize)
	nodeThreePosition, _ := Hash(nodeThree, ringSize)
	nodeFourPosition, _ := Hash(nodeFour, ringSize)

	if nodeOnePosition > ringSize {
		t.Errorf("expected nodeOnePosition to be less than %v got %v", ringSize, nodeOnePosition)
	}
	if nodeTwoPosition > ringSize {
		t.Errorf("expected nodeTwoPosition to be less than %v got %v", ringSize, nodeTwoPosition)
	}
	if nodeThreePosition > ringSize {
		t.Errorf("expected nodeThreePosition to be less than %v got %v", ringSize, nodeThreePosition)
	}
	if nodeFourPosition > ringSize {
		t.Errorf("expected nodeFourPosition to be less than %v got %v", ringSize, nodeFourPosition)
	}
}
