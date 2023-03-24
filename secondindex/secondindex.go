package secondindex

import "fmt"

func SecondIndex(){
	a:="saalama"
	b:="a"
	counter:=0
	for i:=0;i<len(a);i++{
		if string(a[i])==string(b){
			counter++
			if counter ==2 {
				fmt.Println(i)
				break
			}
		} 
	}
}