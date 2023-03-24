package trimstring

import "fmt"

func TrimString(){
	num:=8
	str:="salomsalom"
	fmt.Println(str)
	newstr:=""
	for i:=0;i<num-3;i++{
		newstr+=string(str[i])
	}
	fmt.Println(newstr+"...")
}