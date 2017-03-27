package node

type Node struct {
	Id string `json:"id" datastore:"id"`
	Version int `json:"version" datastore:"version"`
	Name string `json:"name" datastore:"name"`
	Type string `json:"type" datastore:"type"`
	Inputs []string `json:"inputs" datastore:"inputs"`
	Outputs []string `json:"outputs" datastore:"outputs"`
}
