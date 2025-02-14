package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func EliminarPais(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	objID, _ := primitive.ObjectIDFromHex(id)
	collection := client.Database("paisesdb").Collection("paises")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Pais eliminado")
}

func main() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/EliminarPais", EliminarPais)
	fmt.Println("Microservicio EliminarPais corriendo en el puerto 8084")
	http.ListenAndServe(":8084", nil)
}