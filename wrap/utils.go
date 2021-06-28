package wrap

import (
	"github.com/go-jet/jet/v2/internal/jet"
)

func UnwindRowFromValues(value interface{}, values []interface{}) []jet.Serializer {
	row := []jet.Serializer{}

	allValues := append([]interface{}{value}, values...)

	for _, val := range allValues {
		row = append(row, jet.ToSerializerValue(val))
	}

	return row
}
