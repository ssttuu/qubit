package node

type Node struct {
	Id string `json:"id" datastore:"id"`
	Version int `json:"version" datastore:"version"`
	Digest string `json:"digest" datastore:"digest"`
	Name string `json:"name" datastore:"name"`
	Type string `json:"type" datastore:"type"`
	Inputs []string `json:"inputs" datastore:"inputs"`
	Outputs []string `json:"outputs" datastore:"outputs"`
	//Params map[string]interface{} `json:"params" datastore:"params"`
}
