package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Pais struct {
	Nombre string `json:"nombre" bson:"nombre"`
}

var client *mongo.Client

func CrearPais(w http.ResponseWriter, r *http.Request) {
	var pais Pais
	json.NewDecoder(r.Body).Decode(&pais)
	collection := client.Database("paisesdb").Collection("paises")
	result, err := collection.InsertOne(context.TODO(), pais)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func main() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/CrearPais", CrearPais)
	fmt.Println("Microservicio CrearPais corriendo en el puerto 8081")
	http.ListenAndServe(":8081", nil)
}