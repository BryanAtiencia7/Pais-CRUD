package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Pais struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Nombre string `json:"nombre" bson:"nombre"`
}

var client *mongo.Client

func MostrarPais(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	objID, _ := primitive.ObjectIDFromHex(id)
	var pais Pais
	collection := client.Database("paisesdb").Collection("paises")
	err := collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&pais)
	if err != nil {
		http.Error(w, "Pais no encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(pais)
}

func main() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/MostrarPais", MostrarPais)
	fmt.Println("Microservicio MostrarPais corriendo en el puerto 8082")
	http.ListenAndServe(":8082", nil)
}