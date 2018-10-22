package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestStartStop(t *testing.T) {

	//remove existing video
	err := os.Remove("video.avi")
	if err != nil {
		//here we are just printing because not having a video is not an error
		fmt.Println("There are no videos to remove on test setup.")
	}

	//make a channel of type comm
	okChan := make(chan comm)

	t.Run("False global variable", func(t *testing.T) {
		if recording {
			t.Error("'recording' global variable should be false.")
		}
	})

	t.Run("Start", func(t *testing.T) {
		go startRecording(okChan)
		testResult := <-okChan
		if !testResult.success {
			t.Error("Expected success message, received: ", testResult.message)
		}
	})

	//capture at least 5 seconds of video
	time.Sleep(time.Second * 5)

	t.Run("Stop", func(t *testing.T) {

		//stop recording sends a 'close' signal to a channel of empty structs
		//there is no default way to check if a channel is closed or not
		//the signal will trigger a select case that closes the video writer, the camera and changes the 'recording' global variable
		stopRecording()

		//give some time for the sending of the signal
		time.Sleep(time.Second * 2)

		if recording {
			t.Error("global variable 'recording' should be false, but received: ", recording)
		}

		_, err := os.Open("video.avi")

		if err != nil {
			t.Error("Could not find 'video.avi'.")
		}
	})

	//erase test video
	os.Remove("video.avi")
}
