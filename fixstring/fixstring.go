package fixstring

import (
	"fmt"
	"strings"
)

func Fixstring(){
	str:="SaHoB"
	lower:=0
	upper:=0
	for i:=0;i<len(str);i++{
		if str[i]>='a' && str[i]<='z'{
			lower++
		} else {
			upper++
		}
	}

	if upper>lower{
		fmt.Println(strings.ToUpper(str))
	} else{
		fmt.Println(strings.ToLower(str))
	}
}