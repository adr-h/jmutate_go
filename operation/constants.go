package operation

// Allowed operations
const (
	SET = "SET"
	DEL = "DEL"
	INCR = "INCR"
	MULTI = "MULTI"
)
func IsValidOperation(operation string, argument interface{}) bool {
	switch(operation){
	case SET:
		return true
	case DEL:
		return true
	case INCR:
		return true
	case MULTI:
		return true
	default:
		return false
	}
}
