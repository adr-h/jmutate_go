package set

import (
	"jmutate_go/operation/common"
)

type Set struct {
	argument interface{}
}

func New(argument interface{}) (Set, error){
	return Set {argument}, nil
}

func (s Set) Apply(jsonPointer string, targetDocument map[string]interface{}) (mutatedDocument map[string]interface{}, err error) {
	pointer, err := common.GetPointer(jsonPointer)
	return common.SetPointer(pointer,targetDocument, s.argument)
}

