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
		if err != nil{
			return mutation, err
		}
		mutation.operations[pointer] = value
	}

	return mutation, nil
}


//func (j JsonMutation) Apply(document []byte) (newDocument []byte, err error) {
//	for pointer, mutationDesc := range j.changes {
//		val, kind, err := pointer.Get(document)
//		if (err != nil){
//			return
//		}
//	}
//
//}
