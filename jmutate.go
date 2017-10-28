package jmutate_go

import (
	jpointer "github.com/xeipuuv/gojsonpointer"
	"encoding/json"
	"jmutate_go/operation"
)

type JsonMutation struct {
	operations map[string]operation.Document
}

func New(mutationDocument []byte) (JsonMutation, error){
	mutation := JsonMutation{
		operations : make(map[string]operation.Document),
	}

	jsonMap := make(map[string]json.RawMessage)
	if err := json.Unmarshal(mutationDocument, &jsonMap); err != nil {
		 return mutation, err
	}

	for key, value := range jsonMap {
		//operationDoc, ok := value.(operation.Document)
		//if !ok {
		//	return mutation, errors.New("Invalid operation was provided!")
		//}
		operationDoc := operation.Document{}
		err := json.Unmarshal(value,&operationDoc)
		if (err != nil){
			return mutation, err
		}

		mutation.operations[key] = operationDoc
	}

	return mutation, nil
}


func (j JsonMutation) Apply(document []byte) (newDocument []byte,err error) {
	var tempDoc interface{}
	err = json.Unmarshal(document, &tempDoc)
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

	newDocument,err = json.Marshal(tempDoc)
	return newDocument, err
}
