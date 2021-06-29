package main

import "fmt"

func main(){
	var a,b,c,x int

	fmt.Scanf("%d\n%d\n%d\n%d\n",&a,&b,&c,&x)

	count:=0

	for i:=0;i<=a;i++{
		for j:=0;j<=b;j++{
			for k:=0;k<=c;k++{
				if x==500*i+100*j+50*k{
					count++
				}
			}
		}
	}

	fmt.Println(count)
}