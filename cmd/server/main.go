package main

import (
	"consistent-hash-ring/internal/db"
	"consistent-hash-ring/internal/model"
	"consistent-hash-ring/internal/ring"
	"fmt"
)

var nodeToDb map[string]db.Database
var r ring.ConsistentHashRing

func main() {
	nodeToDb := map[string]db.Database{}
	r := ring.NewConsistentHasRing(10)

	// nodes := []db.Database{}
	for i := 1; i <= 4; i++ {
		id := fmt.Sprintf("node-%d", i)
		d := db.NewDatabase(id)
		err := d.Start("./nodes/node-" + id + ".db")
		if err != nil {
			fmt.Println("failed to start node: "+id, err.Error())
			return
		}

		err = model.CreateItemsTable(d.Conn)
		if err != nil {
			fmt.Println("failed to setup node: "+id, err.Error())
		}

		// nodes = append(nodes, d)
	}
	// fmt.Println(nodes)

	// itemId := "123"
	// value := "someKey"

	// nodeIdToWriteTo := hash.Hash(itemId, 4)
	// fmt.Println("writing to node", nodeIdToWriteTo)
	// node := nodes[nodeIdToWriteTo]

	// // write to node
	// _, err := model.Create(node.Conn, itemId, value)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// nodeIdToReadFrom := hash.Hash(itemId, 4)
	// fmt.Println("reading from node", nodeIdToReadFrom)
	// node = nodes[nodeIdToReadFrom]

	// // read from node
	// item, err := model.Get(node.Conn, itemId)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(item)
}
