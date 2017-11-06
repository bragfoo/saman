package common

const (
	EQ = "eq"
	GT = "gt"
	LT = "lt"
)
type Payload struct {
	rows map[string]string
	cols []Cols
}

type Cols struct {
	colName string
	value string
	cmp string
}


func process()  {
}