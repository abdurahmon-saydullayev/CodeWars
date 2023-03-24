package stringrepeat

import "fmt"

func StrRepeat(){
	j:=6
	str:="salom"
	var newstr string
	for i:=0;i<j;i++{
		newstr+=str
	}
	fmt.Println(newstr)
}