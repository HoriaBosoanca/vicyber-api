package server

import (
	// stdlib
	"strconv"

	// endpoint
	"encoding/json"
	"net/http"

	// router
	"github.com/gorilla/mux"
)

func HandleArticle(router *mux.Router) {
	router.HandleFunc("/article", createArticle).Methods("POST")
	router.HandleFunc("/article", getArticle).Methods("GET")
	router.HandleFunc("/article/{id}", getArticleByID).Methods("GET")
	router.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	router.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	router.HandleFunc("/article", OptionsHandler).Methods("OPTIONS")
	router.HandleFunc("/article/{id}", OptionsHandler).Methods("OPTIONS")
}

type Article struct {
	ID       		uint   `gorm:"primaryKey"`
	Title    		string `gorm:"size:100"`
	ImageID  		uint   `gorm:"foreignKey:Image"`
	VideoURL 		string `gorm:"size:1000"`
	VideoWidth		uint   `json:"videowidth"`
	Content  		string `gorm:"size:10000"`
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	// check api key
	if !CheckApiKey(w, r) {
		return
	}
	// get article from req body
	var article Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	// add article to database
	if DB.Create(&article).Error != nil {
		http.Error(w, "Failed to create article", http.StatusInternalServerError)
		return
	}
	// send
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(article)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	// get all articles from DB
	var articles []Article
	if err := DB.Find(&articles).Error; err != nil {
		http.Error(w, "Could not fetch articles", http.StatusInternalServerError)
		return
	}

	// BBCode
	for i := range articles {
        articles[i].Content = parseBBCode(articles[i].Content)
    }

	// send response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(articles)
}

func getArticleByID(w http.ResponseWriter, r *http.Request) {
	// get id url parameter
	urlParams := mux.Vars(r) // returns map of string urlParam indentifiers to string urlParam values
	id, err := strconv.Atoi(urlParams["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	// find the desired article using that ID
	var article Article
	if err := DB.First(&article, id).Error; err != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	// BBCode
	article.Content = parseBBCode(article.Content)

	// send response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(article)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	// check api key
	if !CheckApiKey(w, r) {
		return
	}
	// get id url param
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	// find first (and only) article with matching id
	var article Article
	if DB.First(&article, id).Error != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}
	// overwrite article from req body on top of existing DB article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	DB.Save(&article)
	// send response with updated article
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	// check api key
	if !CheckApiKey(w, r) {
		return
	}
	// get id from url params
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	// find article with said id and delete it
	if DB.Delete(&Article{}, id).Error != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}
	// give empty response
	w.WriteHeader(http.StatusNoContent)
}
