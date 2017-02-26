package operator


func ReadOperation(op Operator) string {
	return "Read Operation on " + op.Name
}


func init() {
	RegisterOperation("Read", ReadOperation)
}
