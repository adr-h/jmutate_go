package incr

import (
	"fmt"
)

type NumType uint8

const (
	FLOAT = NumType(iota)
	INT
)

type Incr struct {
	//no unions in golang, so this will have to do for now
	intArg   int
	floatArg float64
	whichArg NumType

	intReceiver int
	floatReceiver float64
	whichReceiver NumType
}

func New(argument interface{}) (incr Incr,err error) {
	err = incr.setArgument(argument)
	return incr, err
}

//the argument for the INCR operation can only be an int or float64
func (i Incr) setArgument(argument interface{}) error{
	intArg, ok := argument.(int)
	if (ok) {
		i.intArg = intArg
		i.whichReceiver = INT
		return nil
	}

	floatArg, ok := argument.(float64)
	if (ok) {
		i.floatArg = floatArg
		i.whichReceiver = FLOAT
		return nil
	}

	return fmt.Errorf("argument for INCR must be a valid integer or float; Got `%s` instead", argument);
}

//the receiver of the INCR operation can only be an int or float64
func (i Incr) setReceiver(receiver interface{}) error {
	intReceiver, ok := receiver.(int)
	if (ok) {
		i.intReceiver = intReceiver
		i.whichReceiver = INT
		return nil
	}

	floatReceiver, ok := receiver.(float64)
	if (ok) {
		i.floatReceiver = floatReceiver
		i.whichReceiver = FLOAT
		return nil
	}

	return fmt.Errorf("receiver of INCR must be a valid integer or float; Got `%s` instead", receiver);
}

func (i Incr) Apply(receiver interface{}) (interface{}, error) {
	if err := i.setReceiver(receiver); err != nil {
		return nil, err
	}

	//if either the receiver or the argument is a float, both must be converted to float first
	if (i.whichReceiver == FLOAT || i.whichArg == FLOAT ){
		var receiver float64
		var argument float64

		if(i.whichReceiver == FLOAT) {
			receiver = i.floatReceiver
		} else {
			receiver = float64(i.intReceiver)
		}

		if (i.whichArg == FLOAT) {
			argument = i.floatArg
		}else {
			argument = float64(i.intArg)
		}

		return receiver + argument, nil
	}

	return i.intArg + i.intReceiver, nil
}