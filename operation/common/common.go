package common

import (
	jpointer "github.com/xeipuuv/gojsonpointer"
	"errors"
)

func GetPointer(jsonPointer string) (jpointer.JsonPointer, error){
	pointer, err := jpointer.NewJsonPointer(jsonPointer)
	return pointer, err
}

func GetPointerAndReceiver(
	jsonPointer string, targetDocument map[string]interface{},
) (jpointer.JsonPointer, interface{}, error) {
	var receiver interface{}
	pointer, err := GetPointer(jsonPointer)
	if(err != nil) {
		return pointer,receiver,err
	}

	receiver, _, err = pointer.Get(targetDocument)
	return pointer,receiver,err
}

func SetPointer(
	pointer jpointer.JsonPointer, targetDocument map[string]interface{}, value interface{},
) (mutatedDocument map[string]interface{}, err error) {
	setResult, err := pointer.Set(targetDocument, value)
	if (err != nil) {
		return
	}
	mutatedDocument, ok := setResult.(map[string]interface{})
	if (!ok) {
		err = errors.New("Error occured while casting document")
		return
	}

	return
}
