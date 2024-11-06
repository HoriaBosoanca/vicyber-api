package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleImage(router *mux.Router) {
	router.HandleFunc("/image", createImage).Methods("POST")
}

type Image struct {
	Data string `json:"data"`
}

func createImage(w http.ResponseWriter, r *http.Request) {
	if !CheckApiKey(w, r) {
		return
	}

	// decode json into image var
	var base64Img Image
	err := json.NewDecoder(r.Body).Decode(&base64Img)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// TODO:
	imageData, err := decodeBase64Img(base64Img.Data)
	if err != nil {
		http.Error(w, "Failed to decode image", http.StatusBadRequest)
		return
	}
	log.Println(len(imageData))
	imageurl, err := AddImageToBucket(imageData)
	if err != nil {
		http.Error(w, "Failed to create image", http.StatusInternalServerError)
		return
	}
	log.Println(imageurl)

	imgurl := Image{Data: imageurl}
	if DB.Create(&imgurl).Error != nil {
		http.Error(w, "Failed to create image", http.StatusInternalServerError)
		return
	}

	// send base64 image back in json
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(imgurl)
}

func decodeBase64Img(base64ImgData string) ([]byte, error) {
	imgData, err := base64.StdEncoding.DecodeString(base64ImgData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image base64 data: %w", err)
	}
	return imgData, nil
}
