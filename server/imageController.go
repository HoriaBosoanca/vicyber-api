package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleImage(router *mux.Router) {
	router.HandleFunc("/image", createImage).Methods("POST")
}

type Image struct {
	ID uint `gorm:"primaryKey"`
	Data string `json:"data"`
}

func createImage(w http.ResponseWriter, r *http.Request) {
	if !CheckApiKey(w, r) {
		return
	}

	// decode json into image struct
	var imageStruct Image
	err := json.NewDecoder(r.Body).Decode(&imageStruct)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	imageDataBytes, err := decodeimageStruct(imageStruct.Data)
	if err != nil {
		http.Error(w, "Failed to decode image", http.StatusBadRequest)
		return
	}
	imageurl, err := AddImageToBucket(imageDataBytes)
	if err != nil {
		http.Error(w, "Failed to create image", http.StatusInternalServerError)
		return
	}

	imageStruct.Data = imageurl
	if DB.Create(&imageStruct).Error != nil {
		http.Error(w, "Failed to create image", http.StatusInternalServerError)
		return
	}

	// send base64 image back in json
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(imageStruct)
}

func decodeimageStruct(imageStructData string) ([]byte, error) {
	imgData, err := base64.StdEncoding.DecodeString(imageStructData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image base64 data: %w", err)
	}
	return imgData, nil
}
