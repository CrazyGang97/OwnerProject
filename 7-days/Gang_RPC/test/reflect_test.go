package test

import (
	"fmt"
	"reflect"
	"testing"
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
	//fmt.Println(env.Region())
}

func Test_3(t *testing.T) {
	str := "asdasdas"
	str_p := &str
	str_type := reflect.TypeOf(str)
	//fmt.Println(str_type)
	//fmt.Println(str_type.Kind())
	str_p_type := reflect.TypeOf(str_p)
	//fmt.Println(str_p_type)
	//fmt.Println(str_p_type.Kind())

	new1 := reflect.New(str_type)
	fmt.Println(new1.Kind(), new1.Type())
	new2 := reflect.New(str_p_type.Elem())
	fmt.Println(new2.Kind(), new2.Type())
	fmt.Println(new2.Elem().Kind(), new2.Elem().Type())
}

func Test_4(t *testing.T) {
	p := People{
		Name: "",
		Age:  0,
		name: "",
	}
	fmt.Println(reflect.TypeOf(p).Name())
	//err := new(error)
	fmt.Println(reflect.TypeOf(new(error)))
	fmt.Println(reflect.TypeOf(new(People)))
	fmt.Println(reflect.TypeOf((*error)(nil)))
}
