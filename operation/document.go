//go:generate ffjson $GOFILE
package operation

//represents the JSON document that describes an operation
type Document struct {
	OperationName string      `json:"op"`
	Argument      interface{} `json:"arg"`
}

func (d Document) Run(receiver interface{}) (result interface{},err error){
	operation, err := OperationFactory(d.OperationName, d.Argument)
	if (err != nil){
		return
	}

	return operation.Apply(receiver)
}