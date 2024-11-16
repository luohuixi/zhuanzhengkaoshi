package main
import (
	"fmt"
)
func main(){
	var a,b,c,o int
	var d,e []int
    fmt.Scan(&a,&b)
	for i:=0;i<a;i++{
		fmt.Scan(&c)
		var k,kk int
		k=c
		kk=c
		for j:=0;j<i;j++{
			if k>d[j] {
				k=d[j]
				d[j]=kk
                kk=k
			}
		}
		d=append(d,k)
	}
	e=append(e,d[0])
	for i:=1;i<a;i++{
        e=append(e,e[i-1]+d[i])
	}
	for i:=0;i<b;i++{
        fmt.Scan(&c)
		o=0
		for j:=0;j<a;j++{
			if c<=e[j] {
				fmt.Println(j+1)
				o=1
				break
			}
		}
		if o==0 {fmt.Println(-1)}
	}
}