package main
import(
	"io/ioutil"
	"os"
	"fmt"
	"strings"
)

func main(){
	count := make(map[string]int)
	for _,filename := range os.Args[1:]{
		data,err := ioutil.ReadFile(filename)
		if err!=nil{
			fmt.Fprintf(os.Stderr,"dup3:%v\n",err)
		}

		for _,line:= range strings.Split(string(data),"\n"){
			count[line]++
		}
	}
	for line,n := range count{
		if n>1{
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}