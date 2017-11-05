package operation

import (
	"jmutate_go/operation/incr"
	"errors"
	"jmutate_go/operation/set"
	"jmutate_go/operation/del"
	"jmutate_go/operation/insert"
	"jmutate_go/operation/remove"
)

// Allowed operations
const (
	SET = "SET"
	DEL = "DEL"
	INCR = "INCR"
	INSERT = "INSERT"
	MULTI = "MULTI"
	REMOVE = "REMOVE"
)

type Operation interface{
	Apply(jsonPointer string, targetDocument map[string]interface{}) (mutatedDocument map[string]interface{}, err error)
}

func OperationFactory(document Document) (Operation, error) {
	switch(document.OperationName){
	case SET:
		return set.New(document.Argument)
	case DEL:
		return del.New()
	case INCR:
		return incr.New(document.Argument)
	case INSERT:
		return insert.New(document.Argument)
	case REMOVE:
		return remove.New(document.Argument)
	default:
		return nil, errors.New("Unknown JSON mutation operation: " + document.OperationName)
	}
}
