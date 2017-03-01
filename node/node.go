package node

type Node struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Inputs []string `json:"inputs"`
	Params map[string]interface{} `json:"params"`
}
