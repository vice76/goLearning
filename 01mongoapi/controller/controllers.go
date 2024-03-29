package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vice76/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://yadavaj582000:Arpit8090@cluster0.nrnkphh.mongodb.net/"
const dbName = "netflix"
const colName = "watchlist"

// Most important
var collection *mongo.Collection

// connect with mongoDb
func init() {
	//client option
	// follow this syntax for connection

	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	// context.TODO() it has context type ,
	// for how long you are going to work
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongodb connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("collection instance is ready")
}

//db helpers-file

func insertOneRecord(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db ", inserted.InsertedID)
}

//update1record

func updateOneRecord(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	//its not id it always_id
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("no of records updated", result.ModifiedCount)
}

//delete a record

func deleteOneRecord(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	//its not id it always_id
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted record", result.DeletedCount)
}

//delete many records

func deleteAllRecord() int64 {
	// id, _ := primitive.ObjectIDFromHex(movieId)
	//its not id it always_id
	filter := bson.D{{}}

	result, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted record", result.DeletedCount)
	return result.DeletedCount
}

//read all records

func getAllMovies() []primitive.M {
	//Passing the bson.D{{}} as the filter matches  documents in the collection
	result, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M

	for result.Next(context.Background()) {
		var movie bson.M
		err := result.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer result.Close(context.Background())
	fmt.Println("Toatl no of records")
	return movies
}

// Actual controllers - file

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneRecord(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkedAsWatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	updateOneRecord(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneRecord(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllRecord()
	json.NewEncoder(w).Encode(count)
}
