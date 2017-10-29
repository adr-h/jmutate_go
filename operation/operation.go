package operation

import (
	"jmutate_go/operation/incr"
	"errors"
	"jmutate_go/operation/set"
)

// Allowed operations
const (
	SET = "SET"
	DEL = "DEL"
	INCR = "INCR"
	MULTI = "MULTI"
)

type Operation interface{
	Apply(jsonPointer string, targetDocument map[string]interface{}) (mutatedDocument map[string]interface{}, err error)
}

func OperationFactory(document Document) (Operation, error) {
	switch(document.OperationName){
	case SET:
		return set.New(document.Argument)
	case DEL:
		return nil, nil
	case INCR:
		return incr.New(document.Argument)
	case MULTI:
		return nil, nil
	default:
		return nil, errors.New("Unknown JSON mutation operation: " + document.OperationName)
	}
}
