package test

import (
	"fmt"
	"reflect"
	"testing"

	"code.byted.org/gopkg/env"
)

type People struct {
	Name string
	Age  int64
	name string
}

func Test_1(t *testing.T) {
	str := reflect.New(reflect.TypeOf(""))
	str.Elem().SetString("dsadas")
	fmt.Println(str.Elem())
	fmt.Println(str)
}

func Test_2(t *testing.T) {
	fmt.Println(env.Region())
}
