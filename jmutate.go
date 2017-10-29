package jmutate_go

import (
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


func (j JsonMutation) Apply(targetDocument []byte) (mutatedDocument []byte,err error) {
	tempDoc := make(map[string]interface{})
	err = json.Unmarshal(targetDocument, &tempDoc)
	if err != nil {
		return mutatedDocument, nil
	}

	for pointerString, operationDoc := range j.operations {
		operation.RunOperationDocument(pointerString,operationDoc,tempDoc)
	}

	mutatedDocument,err = json.Marshal(tempDoc)
	return mutatedDocument, err
}
