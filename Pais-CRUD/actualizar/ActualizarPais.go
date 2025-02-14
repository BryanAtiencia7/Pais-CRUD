package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Pais struct {
	Nombre string `json:"nombre" bson:"nombre"`
}

var client *mongo.Client

func ActualizarPais(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	objID, _ := primitive.ObjectIDFromHex(id)
	var pais Pais
	json.NewDecoder(r.Body).Decode(&pais)
	collection := client.Database("paisesdb").Collection("paises")
	update := bson.M{"$set": pais}
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Pais actualizado")
}

func main() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/ActualizarPais", ActualizarPais)
	fmt.Println("Microservicio ActualizarPais corriendo en el puerto 8083")
	http.ListenAndServe(":8083", nil)
}