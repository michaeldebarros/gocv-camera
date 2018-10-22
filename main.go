package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

//comm is a type that communicates if there was a success or error
type comm struct {
	success bool
	message string
}

//here we use the signal pattern; it is global because this system seems to be closed, meaning probably it will be only one user in a single session.
var stop = make(chan struct{})

//global state variable indicating if there is a recording going on
var recording bool

//func main is basically the server
func main() {
	r := gin.Default()
	//mudar para POST
	r.GET("/api/v1/record/start", startHandler)
	r.GET("/api/v1/record/stop", stopHandler)
	r.GET("/api/v1/record", getRecordingHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	r.Run(port)
}

func startHandler(c *gin.Context) {
	if recording {
		c.JSON(500, gin.H{
			"message": "Já há uma gravação em curso.",
		})
		return
	}
	okChan := make(chan comm)
	go startRecording(okChan)
	ok := <-okChan
	if !ok.success {
		c.JSON(500, gin.H{
			"message": ok.message,
		})
	}
	c.JSON(200, gin.H{
		"message": ok.message,
	})
}

func stopHandler(c *gin.Context) {
	if !recording {
		c.JSON(500, gin.H{
			"message": "Não há nenhuma gravação em curso.",
		})
		return
	}
	stopRecording()
	c.JSON(200, gin.H{
		"message": "Gravação finalizada com sucesso",
	})
}

func getRecordingHandler(c *gin.Context) {
	c.File("video.avi")
}
