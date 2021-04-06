package reflect

import (
	"fmt"
	"reflect"
	"testing"
)


type ReflectData struct {
	Data string
	name string
	age int
	attr struct{
		big int
	}
}


var Data = ReflectData{
	name: "sfr",
	age: 19,
	attr: struct{ big int }{big: 20},
}

func ShowValye(v *reflect.Value)  {
	if v.Kind() != reflect.Struct{
		if v.Kind() == reflect.String{
			fmt.Println(v.String())
		}else if v.Kind() == reflect.Int {
			fmt.Println(v.Int())
		}
	}else {



		num := v.NumField()
		for i := 0; i < num; i++{
			vChild := v.Field(i)
			ShowValye(&vChild)
		}
	}
}

func TestReflect(t *testing.T) {
	fmt.Println(reflect.TypeOf(Data))
	v := reflect.ValueOf(Data)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	if v.Kind() == reflect.Struct{
		ShowValye(&v)
	}
}

func TestReflectSetValue(t *testing.T) {
	var a int
	v := reflect.ValueOf(&a)
	v.Elem()
	fmt.Println(a)
	v.Elem().SetInt(50)
	fmt.Println(a)
}

func TestElem(t *testing.T) {
	var d = ReflectData{
		name: "",
		age:  0,
		attr: struct {
			big int
		}{},
	}

	//直接报错,  panic: reflect: call of reflect.Value.Elem on struct Value
	// 不能直接修改 struct
	//el := reflect.ValueOf(d).Elem()
	el := reflect.ValueOf(&d).Elem()

	//v1 := el.FieldByName("Data")
	v1 :=el.Field(0)
	v1.SetString("321")
	fmt.Println(d)

}

type TagStr struct {
	Name string `xml:"name"`
	JName string `json:"jname""`
}

func TestTag(t *testing.T) {
	ts := TagStr{
		Name:  "",
		JName: "",
	}

	v := reflect.ValueOf(&ts).Elem()
	vt := reflect.TypeOf(&ts).Elem()

	num  := v.NumField()
	for i :=0; i < num; i++{
		vchild := v.Field(i)
		//vc := vchild.El
		vchild.SetString("222")

		//tag属于 类型 而不是值, 只能使用 typeof来获取
		//fmt.Println(vchild.Tag.Get("json"))

		vtChild := vt.Field(i)
		fmt.Println(vtChild.Tag.Get("xml"))
	}
	fmt.Println(ts)
}

type A struct {
	name string
}

func (a *A) init(data interface{})  {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Interface || val.Kind() == reflect.Ptr{
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		fmt.Println("")
	}

	for i:=0; i< val.NumField(); i++{
		fmt.Println(val.Type().Field(i).Name)
	}
}

type B struct {
	*A
	age int
}

func TestStruct(t *testing.T) {
	a := A{
		name: "is A",
	}

	b := B{
		A:   &a,
		age: 18,
	}

	b.init(&b)
}