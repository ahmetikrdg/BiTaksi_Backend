package tools_convert_Package_Test

import (
	"BiTaksi_Backend_Driver/tools/convert"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CsvToStruct_Total_Value(t *testing.T) {
	a := len(convert.CsvToStruct())

	assert.NotNil(t, a)
	assert.Equal(t, a, 101000)
}

func Test_CsvToStruct_Total_Count(t *testing.T) {
	a := len(convert.CsvToStruct())

	assert.NotNil(t, a)
	assert.NotEqual(t, a, 0)
}
