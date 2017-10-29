package del

import "jmutate_go/operation/common"

type Del struct {}

func New() (Del,error){
	return Del{},nil
}

func (d Del) Apply(jsonPointer string, targetDocument map[string]interface{}) (mutatedDocument map[string]interface{}, err error){
	pointer, err := common.GetPointer(jsonPointer)
	if (err != nil) {
		return
	}

	return common.DeletePointer(pointer,targetDocument)
}