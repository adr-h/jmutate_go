package operation

import (
	"jmutate_go/operation/incr"
	"errors"
)

// Allowed operations
const (
	SET = "SET"
	DEL = "DEL"
	INCR = "INCR"
	MULTI = "MULTI"
)

type Operation interface{
	Apply(receiver interface{}) (result interface{}, err error)
}

func OperationFactory(operationName string, argument interface{}) (Operation, error) {
	switch(operationName){
	case SET:
		return nil, nil
	case DEL:
		return nil, nil
	case INCR:
		return incr.New(argument)
	case MULTI:
		return nil, nil
	default:
		return nil, errors.New("Unknown JSON mutation operation: " + operationName)
	}
}
