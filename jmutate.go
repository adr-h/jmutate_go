package jmutate_go

import (
	jpointer "github.com/xeipuuv/gojsonpointer"
	"jmutate_go/operation"
	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type JsonMutation struct {
	operations map[string]operation.Document
}

func New(mutationDocument []byte) (JsonMutation, error){
	mutation := JsonMutation{}

	jsonMap := make(map[string]operation.Document)
	if err := json.Unmarshal(mutationDocument, &jsonMap); err != nil {
		 return mutation, err
	}

	mutation.operations = jsonMap

	return mutation, nil
}


func (j JsonMutation) Apply(document []byte) (newDocument []byte,err error) {
	tempDoc := make(map[string]interface{})
	err = json.Unmarshal(document, &tempDoc)
	if err != nil {
		return newDocument, nil
	}


	for pointerString, operationDoc := range j.operations {
		pointer, err := jpointer.NewJsonPointer(pointerString)
		if err != nil {
			return newDocument, err
		}

		/*
		TODO: certain operations (e.g: SET or DEL) don't require an operation receiver.
		Refactor this so it doesn't unnecessarily retrieve the operationReceiver in those cases
		*/
		operationReceiver, _, err := pointer.Get(tempDoc)
		if (err != nil) {
			return newDocument, err
		}

		operationResult, err := operationDoc.Run(operationReceiver)
		if (err != nil){
			return newDocument, err
		}

		setResult, err := pointer.Set(tempDoc, operationResult)
		if (err != nil) {
			return newDocument, err
		}
		tempDoc = setResult.(map[string]interface{})
	}

	newDocument,err = json.Marshal(tempDoc)
	return newDocument, err
}
