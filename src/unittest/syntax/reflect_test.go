package syntax

import (
	"testing"
	"reflect"
	"fmt"
)

type bean struct{
	Id int "5"
	name string "tom"
	father string `species:"gopher" color:"blue"`
}
//GoLang利用反射获取struct的tag
func printReflectInfo(){
	bn := &bean{1,"jerry","kiurs"}
	elem := reflect.TypeOf(bn).Elem();
	for i := 0; i<elem.NumField() ; i++ {
		field := elem.Field(i);
		fmt.Println(field,field.Tag.Get("color"), field.Tag.Get("species"))
	}
}

func TestReflect(t *testing.T){
	printReflectInfo()
}
