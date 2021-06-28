package jet

import (
	. "github.com/go-jet/jet/v2/internal/jet"
)

func _UnwindRowFromModel(columns []Column, data interface{}) []Serializer{
	return UnwindRowFromModel(columns, data)
}
