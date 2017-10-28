package operation

//represents the JSON document that describes an operation
type Document struct {
	op  string
	arg interface{}
	Operation `json:"-"`
}

func (d Document) Run(receiver interface{}) (result interface{},err error){
	operation, err := OperationFactory(d.op, d.arg)
	if (err != nil){
		return
	}

	return operation.Apply(receiver)
}