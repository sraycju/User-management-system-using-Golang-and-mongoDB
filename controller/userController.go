package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"fmt"
	"log"
	"crud/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://srayc:guddulal92@cluster0.ql9yvp2.mongodb.net/?retryWrites=true&w=majority"
const dbName = "Users"
const colName = "UserList"

var collection *mongo.Collection

// connection with monogoDB

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected successfully!!")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")
}

// Dao layer

// insert record
func insertUser(user model.Users){
	inserted, err := collection.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted user in database with id: ", inserted.InsertedID)

}
// update record
func updateUser(userId string) {
	id, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"desc": "Description updated"}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

//get users
func getAllUsers() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var users []primitive.M

	for cur.Next(context.Background()) {
		var user bson.M
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	defer cur.Close(context.Background())
	return users
}

// deleting a record
func deleteUser(userId string) {
	id, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User got delete with delete count: ", deleteCount)
}

// calls functions from the dao layer

// Create 
func CreateUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var user model.Users
	_ = json.NewDecoder(r.Body).Decode(&user)
	insertUser(user)
	json.NewEncoder(w).Encode(user)

}

// Retrieve
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllUsers()
	json.NewEncoder(w).Encode(allMovies)
}


// Delete
func DeleteAUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteUser(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// Update
func UpdateDescription(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateUser(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}