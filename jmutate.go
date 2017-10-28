package jmutate_go

import (
	jpointer "github.com/xeipuuv/gojsonpointer"
	"jmutate_go/operation"
	"github.com/pquerna/ffjson/ffjson"
)

type JsonMutation struct {
	operations map[string]operation.Document
}

func New(mutationDocument []byte) (JsonMutation, error){
	mutation := JsonMutation{}

	jsonMap := make(map[string]operation.Document)
	if err := ffjson.Unmarshal(mutationDocument, &jsonMap); err != nil {
		 return mutation, err
	}

	mutation.operations = jsonMap

	return mutation, nil
}


func (j JsonMutation) Apply(document []byte) (newDocument []byte,err error) {
	var tempDoc interface{}
	err = ffjson.Unmarshal(document, &tempDoc)
	if err != nil {
		return newDocument, nil
	}

	for pointerString, operationDoc := range j.operations {
		pointer, err := jpointer.NewJsonPointer(pointerString)
		if err != nil {
			return newDocument, err
		}

		operationReceiver, _, err := pointer.Get(tempDoc)
		if (err != nil) {
			return newDocument, err
		}

		operationResult, err := operationDoc.Run(operationReceiver)
		if (err != nil){
			return newDocument, err
		}

		tempDoc, err = pointer.Set(tempDoc, operationResult)
		if (err != nil) {
			return newDocument, err
		}
	}

	newDocument,err = ffjson.Marshal(tempDoc)
	return newDocument, err
}
