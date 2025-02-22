package ring

import (
	"consistent-hash-ring/internal/hash"
	"fmt"
	"slices"
)

type ConsistentHashRing struct {
	Ring       map[int]string
	SortedKeys []int
	Size       int
}

type RingData struct {
	Position int
	Index    int
}

func NewConsistentHasRing(size int) ConsistentHashRing {
	return ConsistentHashRing{
		Ring:       map[int]string{}, // position: node id
		SortedKeys: []int{},
		Size:       size,
	}
}

func buildNodeId(id string) string {
	return id + "_node"
}

func (c *ConsistentHashRing) AddNode(id string) error {
	nodeId := buildNodeId(id)
	position, err := hash.Hash(nodeId, c.Size)
	if err != nil {
		return err
	}

	if _, ok := c.Ring[position]; ok {
		return fmt.Errorf("a node at position %v already exists", position)
	}

	c.Ring[position] = nodeId
	c.SortedKeys = append(c.SortedKeys, position)
	slices.Sort(c.SortedKeys)

	return nil
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0, len(s)-1)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func findIndex(s []int, value int) int {
	idx := -1

	for i, v := range s {
		if v == value {
			idx = i
			break
		}
	}

	return idx
}

func (c *ConsistentHashRing) RemoveNode(id string) error {
	nodeId := id + "_node"
	position, err := hash.Hash(nodeId, c.Size)
	if err != nil {
		return err
	}

	nodeIdx := findIndex(c.SortedKeys, position)
	if nodeIdx < 0 {
		return fmt.Errorf("no node with value %v found in sorted keys", position)
	}
	c.SortedKeys = removeIndex(c.SortedKeys, nodeIdx)
	delete(c.Ring, position)

	return nil
}

func (c ConsistentHashRing) GetNode(key string) (string, error) {
	if len(c.SortedKeys) == 0 {
		return "", nil
	}

	position, err := hash.Hash(key, c.Size)
	if err != nil {
		return "", err
	}

	var node int

	for idx := range c.SortedKeys {
		if c.SortedKeys[idx] == position || c.SortedKeys[idx] > position {
			node = c.SortedKeys[idx]
			break
		}

		// wrap around to the first node
		if idx == len(c.SortedKeys) {
			node = c.SortedKeys[0]
		}
	}

	return c.Ring[node], nil
}
