package common

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct{
	name string
	value float64
	exp float64
}

func TestCalFactorial(t *testing.T) {
	exp := float64(120)
	actual := CalFactorial(5)

	if exp != actual {
		fmt.Printf("FAILED Exp : %v, actual : %v\n", exp,actual)
	}
}
func TestCalFactorialBulk(t *testing.T) {
	
	values := []TestStruct{
		{name: "Ist RECORD", value: 5, exp: 120},
		{name: "2nd RECORD", value: 4, exp: 24},
		{name: "3rd RECORD", value: 10, exp: 3628800},
	}

	for _,v := range values {
		t.Run(v.name, func(t *testing.T) {
			if act := CalFactorial(v.value); act != v.exp {
				fmt.Printf("FAILED Exp : %v, actual : %v\n", v.exp,act)
			}
		})
	}
}

func TestCalFactorialTestify(t *testing.T) {
	
	values := []TestStruct{
		{name: "Ist RECORD", value: 5, exp: 120},
		{name: "2nd RECORD", value: 4, exp: 24},
		{name: "3rd RECORD", value: 10, exp: 3628800},
	}

	for _,v := range values {
		t.Run(v.name, func(t *testing.T) {
			if act := CalFactorial(v.value); assert.NotEqual(t, v.exp, act) {
				fmt.Printf("FAILED Exp : %v, actual : %v\n", v.exp,act)
			}
		})
	}

}