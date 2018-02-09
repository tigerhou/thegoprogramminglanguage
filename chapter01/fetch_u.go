package main
import(
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)
func main(){
	for _,占位符 := range os.Args[1:]{
		if !strings.HasPrefix(占位符,"http://"){
			占位符 = "http://"+占位符
		}
		response,err := http.Get(占位符);
		if err!=nil{
			fmt.Fprintf(os.Stderr,"fetch:%v\n",err)
			os.Exit(1)
		}
		//body,err := ioutil.ReadAll(response.Body)
		_,err1 := io.Copy(os.Stdout,response.Body)
		response.Body.Close()
		if err1!=nil{
			fmt.Fprintf(os.Stderr,"fetch: reading %s:%v\n",占位符,err1)
			os.Exit(1)
		}
		//fmt.Printf("%s",body)
	}
}