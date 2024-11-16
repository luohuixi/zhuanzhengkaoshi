package main
import(
	"fmt"
	"time"
	"sync"
)
func main(){
	var ww,w int
	var lock1 sync.Mutex
	m:=2000
	n:=10
	k:=200
	for {
		o:=0
		if ww<=n{ 
		go func(){
			for {
			  if m==0 {
				lock1.Lock()
				w++
				fmt.Println("car ",w,"  move ",o)
				ww--
				lock1.Unlock()
				return
			  }
			  if m!=0 {
				o++
			    m--
			    ww++
                time.Sleep(100 * time.Millisecond)
			    if o==k {
				  lock1.Lock()
				  w++
				  fmt.Println("car ",w,"  move ",o)
				  ww--
				  lock1.Unlock()
				  return
			    }
			  }
		    }
		}()
		}
	}
}