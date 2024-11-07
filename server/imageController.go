package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandleImage(router *mux.Router) {
	router.HandleFunc("/image", createImage).Methods("POST")
	router.HandleFunc("/image/{id}", getImageByID).Methods("GET") 
	router.HandleFunc("/image", OptionsHandler).Methods("OPTIONS")
	router.HandleFunc("/image/{id}", OptionsHandler).Methods("OPTIONS")
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

	imageDataBytes, err := decodeImage(imageStruct.Data)
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

func decodeImage(imageStructData string) ([]byte, error) {
	imgData, err := base64.StdEncoding.DecodeString(imageStructData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image base64 data: %w", err)
	}
	return imgData, nil
}

func getImageByID(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	id, err := strconv.Atoi(urlParams["id"])
	if err != nil {
		http.Error(w, "Wrong url param", http.StatusBadRequest)
		return
	}

	var image Image
	if err := DB.First(&image, id).Error; err != nil {
		http.Error(w, "Image not found in DB", http.StatusBadRequest)
		return
	}
	imageUrl := image.Data

	imageDataBytes, err2 := GetImageFromBucket(imageUrl)
	if err2 != nil {
		http.Error(w, "Image not found in bucket", http.StatusBadRequest)
		return
	}

	image.Data = encodeImage(imageDataBytes)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(image)
}

func encodeImage(imageData []byte) string {
	base64Image := base64.StdEncoding.EncodeToString(imageData)
	return base64Image
}
