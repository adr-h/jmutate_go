package jmutate_go

import (
	jpointer "github.com/xeipuuv/gojsonpointer"
	"encoding/json"
	"jmutate_go/operation"
)

type JsonMutation struct {
	operations map[jpointer.JsonPointer]operation.Document
}

func New(mutationDocument []byte) (JsonMutation, error){
	mutation := JsonMutation{}

	jsonMap := make(map[string]operation.Document)
	if err := json.Unmarshal(mutationDocument, &jsonMap); err != nil {
		 return mutation, err
	}

	for key, value := range jsonMap {
		pointer, err := jpointer.NewJsonPointer(key)
		if err != nil {
			return mutation, err
		}
		mutation.operations[pointer] = value
	}

	return mutation, nil
}


func (j JsonMutation) Apply(document []byte) (newDocument []byte, err error) {
	newDocument = make([]byte, len(document))
	copy(newDocument, document)

	for pointer, operationDoc := range j.operations {
		operationReceiver, _, err := pointer.Get(document)
		if (err != nil) {
			return newDocument, err
		}

		operationResult, err := operationDoc.Run(operationReceiver)
		if (err != nil){
			return newDocument, err
		}

		setResult, err := pointer.Set(document, operationResult)
		if (err != nil) {
			return newDocument, err
		}
		newDocument, _ = setResult.([]byte)
	}

	return newDocument, nil
}
