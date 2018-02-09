package main
import(
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe("localhost:8080",nil))
}
func handler(resp http.ResponseWriter,req *http.Request){
	fmt.Fprintf(resp,"URL.path=%q\n",req.URL.Path)
}