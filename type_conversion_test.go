package reflect

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToType(t *testing.T) {
	var (
		samples = []struct {
			name   string
			input  interface{}
			output interface{}
			t      Type
		}{
			{
				name:   "string slice to bool slice",
				input:  []string{"true", "false", "true"},
				output: []bool{true, false, true},
				t:      reflect.SliceOf(TypeBool),
			},
			{
				name:   "int slice to string slice",
				input:  []int{1, 2, 3},
				output: []string{"1", "2", "3"},
				t:      reflect.SliceOf(TypeString),
			},
			{
				name:   "string slice to int slice",
				input:  []string{"1", "2", "3"},
				output: []int{1, 2, 3},
				t:      reflect.SliceOf(TypeInt),
			},

			{
				name:   "string:int map to bool:string map",
				input:  map[string]int{"false": 0, "true": 1},
				output: map[bool]string{true: "1", false: "0"},
				t:      reflect.MapOf(TypeBool, TypeString),
			},
			{
				name:   "float64:bool map to string:int8 map",
				input:  map[float64]bool{3.14: true, 56.17: false},
				output: map[string]int8{"3.14": 1, "56.17": 0},
				t:      reflect.MapOf(TypeString, TypeInt8),
			},
		}
	)

	for _, sample := range samples {
		t.Run(
			sample.name,
			func(t *testing.T) {
				var (
					res interface{}
					err error
				)

				res, err = ConvertToType(sample.input, sample.t)
				assert.Equal(t, sample.output, res)
				assert.Equal(t, nil, err, fmt.Sprintf("%s", err))
			},
		)
	}
}
