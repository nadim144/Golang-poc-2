package Handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wipro-toyota-poc/Dbconfig"
	"wipro-toyota-poc/Models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = Dbconfig.GetCollection(Dbconfig.DB, "product")

// Create a product - POST
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-from-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var product Models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	insertProduct(product)
	json.NewEncoder(w).Encode(product)

}

// Get all products - GET
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-from-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	allProducts := getAllProducts()
	json.NewEncoder(w).Encode(allProducts)
}

// Get a product - GET
func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-from-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	params := mux.Vars(r)
	getOneProduct(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

// Update a product - PUT
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-from-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneProduct(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// Delete a product - DELETE
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-from-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneProduct(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// Delete all products
func DeleteAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-from-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllProducts()
	json.NewEncoder(w).Encode(count)
}

// MonfoDB Helpers

// insert 1 product
func insertProduct(product Models.Product) {
	inserted, err := productCollection.InsertOne(context.Background(), product)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 product in db with id", inserted.InsertedID)
}

// Get 1 record
func getOneProduct(productID string) {
	id, _ := primitive.ObjectIDFromHex(productID)
	filter := bson.M{"_id": id}
	deletedCount, err := productCollection.Find(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Product got deleted with delete count", deletedCount)
}

// update 1 product
func updateOneProduct(productID string) {
	id, _ := primitive.ObjectIDFromHex(productID)
	filter := bson.M{"_id": id}
	update := bson.M{"$Set": bson.M{"price": 100.00}}

	result, err := productCollection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count: ", result.ModifiedCount)
}

// delete 1 record
func deleteOneProduct(productID string) {
	id, _ := primitive.ObjectIDFromHex(productID)
	filter := bson.M{"_id": id}
	deletedCount, err := productCollection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Product got deleted with delete count", deletedCount)
}

// delete all records

func deleteAllProducts() int64 {
	deletedResult, err := productCollection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of product dleted:", deletedResult.DeletedCount)
	return deletedResult.DeletedCount
}

//get all product from database

func getAllProducts() []primitive.M {
	cur, err := productCollection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var products []primitive.M
	for cur.Next(context.Background()) {
		var product bson.M
		err := cur.Decode(&product)

		if err != nil {
			log.Fatal(err)
		}

		products = append(products, product)
	}
	defer cur.Close(context.Background())
	return products
}
