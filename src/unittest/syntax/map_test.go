package syntax

import "testing"

func TestBase(t *testing.T){
	var userMap map[int] string //定义map
	userMap = make(map[int]string) //初始化map

	userMap[1] = "tom"
	userMap[3] = "jerry"
	userMap[44] = "home"

	for k,v := range userMap{
		println(k,v)
	}

}
