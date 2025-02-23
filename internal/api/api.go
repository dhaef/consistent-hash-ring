package api

import (
	"consistent-hash-ring/internal/db"
	"consistent-hash-ring/internal/model"
	"consistent-hash-ring/internal/ring"
	"encoding/json"
	"fmt"
	"net/http"
)

type API struct {
	Ring     ring.ConsistentHashRing
	NodeToDb map[string]db.Database
}

func NewAPI(ringSize int) API {
	return API{
		Ring:     ring.NewConsistentHasRing(ringSize),
		NodeToDb: map[string]db.Database{},
	}
}

func (api *API) Setup() error {
	for i := 1; i <= 4; i++ {
		id := fmt.Sprintf("db-%d", i)
		d := db.NewDatabase(id)
		err := d.Start("./nodes/node-" + id + ".db")
		if err != nil {
			return fmt.Errorf("failed to start node: "+id, err)
		}

		err = model.CreateItemsTable(d.Conn)
		if err != nil {
			return fmt.Errorf("failed to setup db: "+id, err)
		}

		nodeId, err := api.Ring.AddNode(id)
		if err != nil {
			return fmt.Errorf("failed to add node: "+id, err)
		}

		api.NodeToDb[nodeId] = d
	}

	return nil
}

func (api API) RegisterRoutes() *http.ServeMux {
	h := http.NewServeMux()
	h.HandleFunc("GET /{id}", api.getItem)
	h.HandleFunc("POST /", api.createItem)

	return h
}

func decode[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}

func encode[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func (api API) createItem(w http.ResponseWriter, r *http.Request) {
	body, err := decode[model.Item](r)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("could not parse request body"))
		return
	}

	// TODO: validation

	nodeIdToWriteTo, err := api.Ring.GetNode(body.Id)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not find node"))
		return
	}

	node := api.NodeToDb[nodeIdToWriteTo]

	// write to node
	_, err = model.Create(node.Conn, body.Id, body.Value)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to write data"))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (api API) getItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	nodeIdToReadFrom, err := api.Ring.GetNode(id)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not find node"))
		return
	}

	node := api.NodeToDb[nodeIdToReadFrom]

	// read from node
	item, err := model.Get(node.Conn, id)
	if err != nil {
		fmt.Println(err.Error())
	}

	encode(w, http.StatusOK, item)
}
