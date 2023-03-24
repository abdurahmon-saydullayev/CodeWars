package nonrepetead

import "fmt"

func NonRepetead(){
	str:="test"
	counter:=0

	for i:=0;i<len(str);i++{
		for j:=i+1;j<len(str);j++{
			if str[i]==str[j]{
				counter++
			}
		}
	}
	if counter>0{
		fmt.Println("null")
	} else{
		
	}
}