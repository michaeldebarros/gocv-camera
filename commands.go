package main

import (
	"log"
	"time"
	"os"
	"strconv"

	"gocv.io/x/gocv"
)

func startRecording(okChan chan comm) {
	//the idea is to capture video from ONE camera, so we pass in the first deviceID
	deviceID := 0

	index := os.Getenv("CAMERA_INDEX")
	if index == "" {
		index = "0"
		i,err := strconv.Atoi(index)
		if err != nil{
			log.Println("Erro ao converter CAMERA_INDEX para int.")
		}
		deviceID = i
	} 
	
	saveFile := "video.avi"

	camera, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		log.Println("Erro ao abrir a câmera, ID: ", deviceID)
		newComm := comm{false, "Não foi possível abrir a câmera."}
		okChan <- newComm
		return
	}

	img := gocv.NewMat()
	defer img.Close()
	if ok := camera.Read(&img); !ok {
		log.Println("Error while reading images from device: ", deviceID)
		newComm := comm{false, "Não foi possível ler as imagens da câmera."}
		okChan <- newComm
		return
	}

	//here we calculate the fps from the camera, because simply using the Get method does not give the accurate fps.
	calculatedFPS := calculateFPS(camera)

	writer, err := gocv.VideoWriterFile(saveFile, "MJPG", calculatedFPS, img.Cols(), img.Rows(), true)
	if err != nil {
		log.Println("Error while opening recording device.")
		newComm := comm{false, "Não foi possível abrir o dispositivo de gravacão de video."}
		okChan <- newComm
		return
	}

	recording = true
	//this newComm is in the function scope
	newComm := comm{true, "Gravação iniciada com successo."}

	okChan <- newComm
	for {
		select {
		//just to stop, this will be taken out
		case <-stop:
			writer.Close()
			camera.Close()
			recording = false
			return
		default:
			if ok := camera.Read(&img); !ok {
				log.Printf("Dispositivo fechado: %v\n", deviceID)
				return
			}
			if img.Empty() {
				continue
			}
			writer.Write(img)
		}
	}
}

func stopRecording() {
	close(stop)
}

func calculateFPS(camera *gocv.VideoCapture) float64 {
	img := gocv.NewMat()
	defer img.Close()
	if ok := camera.Read(&img); !ok {
		log.Println("Problema ao calcular FPS pois não foi possível efetuar a letura do dispositivo")
	}
	start := time.Now()
	for i := 0; i < 20; i++ {
		if ok := camera.Read(&img); !ok {
			log.Println("Dispositivo fechado dentro de calculateFPS")
			return 0
		}
	}
	fps := 20 / time.Since(start).Seconds()
	return fps
}
