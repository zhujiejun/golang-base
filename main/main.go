package main

import (
	"log"
	"net/http"
	"github.com/zhujiejun/Server"
	//"github.com/gin-gonic/gin"
)

/*func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("hello world!")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}*/
func main() {
	http.HandleFunc("/", Server.handler)
	http.HandleFunc("/count", Server.counter)
	log.Fatal(http.ListenAndServe("192.168.100.100:18080", nil))
}
