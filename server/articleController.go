package server

import (
	// stdlib
	"fmt"
	"io"
	"strconv"

	// endpoint
	"encoding/json"
	"net/http"

	// router
	"github.com/gorilla/mux"

	// markdown parser
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	// "github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
)

func HandleArticle(router *mux.Router) {
	router.HandleFunc("/article", createArticle).Methods("POST")
	router.HandleFunc("/article/{category}", getArticleByCategory).Methods("GET")
	router.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")

	router.HandleFunc("/article", OptionsHandler).Methods("OPTIONS")
	router.HandleFunc("/article/{category}", OptionsHandler).Methods("OPTIONS")
	router.HandleFunc("/article/{id}", OptionsHandler).Methods("OPTIONS")
}

type ArticlePost struct {
	ID   uint   `gorm:"primaryKey"`
	Tags string
	HTML string 
}

type Tags struct {
	Tags string `json:"tags"`
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	// check api key
	if !CheckApiKey(w, r) {
		return
	}

	parseFormErr := r.ParseMultipartForm(10 << 20) // Limit file size to 10MB
    if parseFormErr != nil {
        http.Error(w, "Failed to parse multipart form", http.StatusBadRequest)
		fmt.Println(parseFormErr)
        return
    }

	file, _, retrieveFileErr := r.FormFile("file")
	if retrieveFileErr != nil {
		http.Error(w, "Error retreiving the file", http.StatusBadRequest)
		fmt.Println(retrieveFileErr)
		return
	}
	defer file.Close()

	fileData, readFileErr := io.ReadAll(file)
	if readFileErr != nil {
		http.Error(w, "Could not read file", http.StatusInternalServerError)
		fmt.Println(readFileErr)
		return
	}

	myHtml := string(mdToHTML(fileData))

	metadataStr := r.FormValue("metadata")
	tags := Tags{}
	metadataErr := json.Unmarshal([]byte(metadataStr), &tags)
	if metadataErr != nil {
		http.Error(w, "Invalid json metadata", http.StatusBadRequest)
		return
	}

	addFileErr := DB.Create(&ArticlePost{
		HTML: myHtml,
		Tags: tags.Tags,
	}).Error
	if addFileErr != nil {
		http.Error(w, "Error adding file to db", http.StatusInternalServerError)
		fmt.Println(addFileErr)
		return
	}

	fmt.Println(string(myHtml))

	w.WriteHeader(http.StatusCreated)
}

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func getArticleByCategory(w http.ResponseWriter, r *http.Request) {
	tags := mux.Vars(r)["category"]

	var articles []ArticlePost
	result := DB.Where("tags = ?", tags).Find(&articles)
	if(tags == "any") {
		result = DB.Find(&articles)
	}
	if result.Error != nil {
		http.Error(w, "Articles not found", http.StatusNotFound)
		fmt.Println(result.Error)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&articles)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	if !CheckApiKey(w, r) {
		return
	}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if DB.Delete(&ArticlePost{}, id).Error != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
