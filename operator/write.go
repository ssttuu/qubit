package operator

func WriteOperation(op Operator) string {
	return "Write Operation on " + op.Name
}

func init() {
	RegisterOperation("Write", WriteOperation)
}
