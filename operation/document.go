package operation

//represents the JSON document that describes an operation
type Document struct {
	op  string
	arg interface{}
	Operation `json:"-"`
}

func NewDocument(op string, arg interface{}) (doc Document, err error) {
	doc.op = op
	doc.arg = arg
	err = doc.Validate()
	return doc, err
}

func (d Document) Validate() error{
	op, err := OperationFactory(d.op, d.arg)
	if(err != nil){
		return err
	}

	d.Operation = op
	return nil
}