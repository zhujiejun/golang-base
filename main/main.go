package main

import (
	"github.com/zhujiejun/golang-base/server"
	"log"
	"net/http"
)

/*func main() {
	message := greet.Hello("Golang")
	if len(os.Args) > 1 {
		message = greet.Hello(os.Args[1])
	}
	fmt.Println(message)
}*/

func main() {
	http.HandleFunc("/", server.Handler)
	http.HandleFunc("/count", server.Counter)
	log.Fatal(http.ListenAndServe("192.168.100.100:18080", nil))
}
