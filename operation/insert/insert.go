package insert

import "jmutate_go/operation/common"
import (
	"github.com/mitchellh/mapstructure"
	"fmt"
	"errors"
)

type InsertArg struct {
	At *int	//optional; if not provided, the value will be pushed into the end of the array
	Values []interface{}
}

type Insert struct {
	arg InsertArg
}

func New(arg interface{}) (op Insert,err error){
	insertArg := InsertArg{}
	op = Insert{}
	err = mapstructure.Decode(arg,&insertArg)
	if (err != nil){
		return
	} else if (insertArg.Values == nil){
		err = errors.New("required field 'values' was missing")
		return
	}
	op.arg = insertArg
	return
}

func (i Insert) Apply(jsonPointer string, targetDocument map[string]interface{}) (mutatedDocument map[string]interface{}, err error){
	pointer, err := common.GetPointer(jsonPointer)
	if (err != nil) {
		return
	}

	arr,_,err := pointer.Get(targetDocument)

	switch val := arr.(type) {
	case []interface{} :
		if (i.arg.At == nil){
			//if no "at" was provided, treat it as a "push" by adding it to the end of the slice
			val = append(val, i.arg.Values...)
		}else {
			index := *i.arg.At
			val = append(val[:index], append(i.arg.Values, val[index:]...)...)
		}

		var result interface{}
		result, err = pointer.Set(targetDocument, val)
		if (err != nil){
			return
		}
		mutatedDocument = result.(map[string]interface{})
		return
	default :
		err = fmt.Errorf("value at '%s' was not an array", jsonPointer)
		return
	}
}