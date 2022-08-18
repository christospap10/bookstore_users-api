package app

import "github.com/gin-gonic/gin"

var (
	router     = gin.Default()
	portNumber = ":8080"
)

func StartApplication() {
	mapUrls()
	router.Run(portNumber)
}
