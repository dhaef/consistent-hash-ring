package hash

import (
	"fmt"
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

	fmt.Println(nodeOnePosition)
	fmt.Println(nodeTwoPosition)
	fmt.Println(nodeThreePosition)
	fmt.Println(nodeFourPosition)
}
