package syntax

import (
	"fmt"
	"testing"
)

func TestPoint(t *testing.T){
	str := "aaa";
	var pstr *string;
	pstr = &str;
	fmt.Printf("*pstr原始值：%s\n",*pstr)
	fmt.Println("&str指针内存地址：",&str)
	fmt.Println("pstr指针内存地址：",pstr)
	fmt.Println("&pstr指针内存值：",&pstr)
	changePointString(pstr);
	fmt.Printf("*pstr改变之后值：%s\n",*pstr)
}

func changePointString(pstr *string){
	*pstr = "bbb";
	fmt.Printf("changePointString *pstr：%s\n",*pstr)
}