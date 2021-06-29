package main

import "fmt"

func main(){
	var n int

	fmt.Scanf("%d",&n)

	a:=make([]int,n)
	
	for i:=0;i<n;i++{
		fmt.Scanf("%d",&a[i])
	}

	count:=0

	for isEveryAEven(a){
		for i:=0;i<len(a);i++{
			a[i]=a[i]/2
		}
		count++
	}

	fmt.Println(count)
}

func isEveryAEven(a []int)bool{
	for i:=0;i<len(a);i++{
		if a[i]%2!=0{
			return false
		}
	}

	return true
}