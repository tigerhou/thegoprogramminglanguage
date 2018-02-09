package main
import(
	"fmt"
	"net/http"
	"sync"
	"log"
)

var mu sync.Mutex
var count int

func main(){
	http.HandleFunc("/",handler)
	http.HandleFunc("/count",counter)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}
func handler(resp http.ResponseWriter,req *http.Request){
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(resp,"URL.path = %q\n",req.URL.Path)
}

func counter(resp http.ResponseWriter,req *http.Request){
	mu.Lock()
	fmt.Fprintf(resp,"Count: %d\n",count)
	mu.Unlock()
}