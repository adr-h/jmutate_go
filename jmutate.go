package jmutate_go

import (
	jpointer "github.com/xeipuuv/gojsonpointer"
	"encoding/json"
	"jmutate_go/operation"
)



type JsonMutation struct {
	changes map[jpointer.JsonPointer]map[string]MutationDescription
}

type MutationDescription struct {
	op  string
	arg interface{}
}

func (d MutationDescription) Validate() bool{
	return operation.IsValidOperation(d.op, d.arg)
}


func New(mutationDocument []byte) (JsonMutation, error){
	mutation := JsonMutation{}

	jsonMap := make(map[string]map[string]MutationDescription)
	if err := json.Unmarshal(mutationDocument, &jsonMap); err != nil {
		 return mutation, err
	}

	for key, value := range jsonMap {
		pointer, err := jpointer.NewJsonPointer(key)
		if err != nil{
			return mutation, err
		}
		mutation.changes[pointer] = value
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
