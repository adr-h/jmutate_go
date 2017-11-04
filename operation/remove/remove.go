package remove

import (
	"errors"
	"jmutate_go/operation/common"
	"fmt"
)

type Remove struct{
	arg RemoveArg
}

type RemoveArg struct {
	At    int
	Count int
}

func New(arg interface{}) (r Remove,err error) {
	r = Remove{ RemoveArg{} }
	mapValues, ok := arg.(map[string]interface{})
	if !ok {
		err = errors.New("argument was not a map")
		return
	}

	if val, ok := mapValues["at"]; ok {
		from, ok := val.(int)
		if (!ok || from < 0){
			err = errors.New("the 'at' field must be an integer with a value equal to or higher than  0")
			return
		}
		r.arg.At = from
	} else {
		err = errors.New("the 'at' field is required for REMOVE operations")
		return
	}

	if val, ok := mapValues["count"]; ok {
		count, ok := val.(int)
		if (!ok || count <= 0) {
			err = errors.New("the 'count' field must be an integer with a value higher than 0")
			return
		}

		r.arg.Count = count
	} else {
		r.arg.Count = 1
	}

	return
}

func (r Remove) Apply(jsonPointer string, targetDocument map[string]interface{}) (mutatedDocument map[string]interface{}, err error) {
	pointer, err := common.GetPointer(jsonPointer)
	if (err != nil) {
		return
	}

	arr,_,err := pointer.Get(targetDocument)

	switch value := arr.(type){
		case []interface{} :{
			at := r.arg.At
			count := r.arg.Count
			newSlice := append( value[0:at], value[at+count:]... )

			var result interface{}
			result, err = pointer.Set(targetDocument,newSlice)
			if (err != nil) {
				return
			}

			mutatedDocument, _ = result.(map[string]interface{})
			return
		}
		default : {
			err = fmt.Errorf("value at '%s' was not an array", jsonPointer)
			return
		}
	}

	return
}
