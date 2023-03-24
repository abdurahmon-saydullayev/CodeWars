package tocamelcase

import (
	"fmt"
	"strings"
)


func ToCamelCase(){
	sample:="salom_hi"
	var newstr string
	boolean := false
	for i:=0;i<len(sample);i++{
		if string(sample[i])=="-" || string(sample[i])=="_"{
			boolean=true
		} else{
			if boolean {
				newstr += strings.ToUpper(string(sample[i]))
				boolean=false
			} else {
				newstr+=string(sample[i])
			}
		}
	}
	fmt.Println(newstr)
}