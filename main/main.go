package main

import (
	"fmt"
	"github.com/zhujiejun/golang-base/greet"
	"os"
)

func main() {
	message := greet.Hello("Golang")
	if len(os.Args) > 1 {
		message = greet.Hello(os.Args[1])
	}
	fmt.Println(message)
}

/*func main() {
	http.HandleFunc("/", server.Handler)
	http.HandleFunc("/count", server.Counter)
	log.Fatal(http.ListenAndServe("192.168.100.100:18080", nil))
}*/
