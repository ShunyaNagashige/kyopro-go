package main

import "fmt"

func main(){
	const masuNum=3

	//いいね！
	var num string
	count := 0

	fmt.Scanf("%s",&num)

	for i:=0;i<masuNum;i++{
		if num[i]=='1'{
			count++
		}
	}

	fmt.Println(count)
}