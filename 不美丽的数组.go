package main
import(
	"fmt"
)
func main(){
	var a,b,c,k int
	fmt.Scan(&a)
	for i:=0;i<a;i++{
		fmt.Scan(&b)
		c+=b
	}
	for k>=0{
        k++
		if (c/k)<a {break}
	}
	p:=c-a*(k-1)
    fmt.Println(p*(a-p))
}