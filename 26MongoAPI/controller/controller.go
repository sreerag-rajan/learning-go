package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.coms/sreerag-rajan/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString string = "mongodb+srv://sreeragrajan5:mTGZxTKQo1Umkbo1@cluster0.1hhrcat.mongodb.net/?retryWrites=true&w=majority"
const dbName string = "netflix"
const collectionName string = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// connect with mongo DB

func init() { //Initialization method that only runs once when project is run
	//client option
	clientOptions := options.Client().ApplyURI(connectionString)

	//connect to mongo db
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo DB Connection Success")

	collection = client.Database(dbName).Collection(collectionName)

	// if collection instance is ready
	fmt.Println("Connection reference is ready")

}

//MONGODB Helpers - file

//insert 1 record

func insertOne(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", inserted)
}

//update 1 record
func updateOne(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", result)
	fmt.Printf("Updated %v", id)

}

// delete one record
func deleteOne(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", deleteCount)
	fmt.Printf("Deleted %v and deleteCount is %v", id, deleteCount)

}

func deleteAll() {
	deleteCount, err := collection.DeleteMany(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", deleteCount)
}

// get all movies from database
func getAll() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())

	return movies
}

//bson.M VS bson.D

//ACTUAL CONTROLLERS - file
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAll()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix

	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOne(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOne(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOne(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteAll()
	json.NewEncoder(w).Encode("All movies Deleted")
}
