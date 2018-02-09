package main
import(
	"fmt"
	"os"
	"time"
	"strings"
)

func main(){
	s,sep:="",""
	loop_start := time.Now()
	for _,arg:=range os.Args[1:]{
		s+=sep+arg
		sep=" "
	}
	fmt.Println(s)
	fmt.Printf("%.2fs elapsed\n", time.Since(loop_start).Nanoseconds())
	join_start := time.Now()
	fmt.Println(strings.Join(os.Args[1:]," "))
	fmt.Printf("%.2fs elapsed\n", time.Since(join_start).Nanoseconds())
	//for index,arg:=range os.Args[1:]{
		//fmt.Print(index)
		//fmt.Println(arg)
	//}
}