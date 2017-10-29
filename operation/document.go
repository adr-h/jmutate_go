//go:generate ffjson $GOFILE
package operation

//represents the JSON document that describes an operation
type Document struct {
	OperationName string      `json:"op"`
	Argument      interface{} `json:"arg"`
}

func RunOperationDocument(
	pointerString string, operationDoc Document, receivingDoc map[string]interface{},
) (mutatedDocument map[string]interface{}, err error)  {
	operation, err := OperationFactory(operationDoc)
	if (err != nil) {
		return
	}
	return operation.Apply(pointerString,receivingDoc)
}