package syntax

import (
	"fmt"
	"os"
	"testing"
	"net"
)

/**
go error错误处理
 */
func TestException(t *testing.T){
	f, err := os.Open("/test.txt")

	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			fmt.Println("generic error: ", err)
		}
		return
	}

	if err != nil {
		fmt.Println("error:",err)
		return
	}
	fmt.Println(f.Name(), "open successfully")
}

type CustomError struct{
	Code int
	Msg string
}

func (ce CustomError) Error() string{
	return string(ce.Code) + "" + ce.Msg
}

func TestException2(t *testing.T){
	fmt.Println("division begin")
	_,err := divisionZeroError(1,0)
	if err != nil{
		fmt.Println("division error")
	}else{
		fmt.Println("division success")
	}

	if err,ok := err.(CustomError); ok{//判断err是否是CustomError类型，是的话类型转换为CustomError
		fmt.Println(err.Code,err.Msg)
	}

}

func divisionZeroError(a int,b int) (int,error){
	if b == 0 {
		return 0,CustomError{1,"被除数不能为0"}
	}
	return a/b,nil;
}