package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
    
	"github.com/TejaTrinay/api-go-lang/helper"
	"github.com/TejaTrinay/api-go-lang/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = helper.ConnectDB()

//Create Article
func createArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var article models.Articles
	_ = json.NewDecoder(r.Body).Decode(&article)
	result, err := collection.InsertOne(context.TODO(), article)

	if err != nil {
		helper.GetError(err, w)
		return
	}
	

	json.NewEncoder(w).Encode(result)
}

//Get Article
func getArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article models.Articles
	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&article)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(article)
}

// Get articles
func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var articles []models.Articles

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var article models.Articles
		err := cur.Decode(&article) 
		if err != nil {
			log.Fatal(err)
		}
		articles = append(articles, article)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(articles) 
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/articles", getArticles).Methods("GET")
	r.HandleFunc("/api/articles/{id}", getArticle).Methods("GET")
	r.HandleFunc("/api/articles", createArticle).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
