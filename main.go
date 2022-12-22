package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	inf := log.New(os.Stdout, "INFO ", log.Ldate)
	err := log.New(os.Stderr, "ERROR ", log.Ldate)

	port := os.Getenv("CPORT")
	if len(port) == 0 {
		err.Println("CPORT env is required")
		os.Exit(1)
	}

	inf.Println("Starting..")

	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusOK, "msg: "+os.Getenv("CMSG")+"\n")
	})
	r.Run(":" + port)
}
