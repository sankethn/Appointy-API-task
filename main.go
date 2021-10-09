package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var mutex sync.Mutex

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UID      string             `json:"uid" bson:"uid"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
}

type Post struct {
	ID              primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	PID             string              `json:"pid" bson:"pid"`
	UID             string              `json:"uid" bson:"uid"`
	Caption         string              `json:"caption" bson:"caption"`
	ImageUrl        string              `json:"imageurl" bson:"imageurl"`
	PostedTimestamp primitive.Timestamp `json:"postedtimestamp" bson:"postedtimestamp"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && r.URL.Path == "/" {
		fmt.Fprintf(w, "Welcome home!")
	} else {
		notFound(w, r)
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("content-type", "application/json")
		var user User
		_ = json.NewDecoder(r.Body).Decode(&user)
		hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
		if err != nil {
			log.Panic(err)
		}
		mutex.Lock()
		user.Password = string(hashedPwd)
		collection := client.Database("instagram").Collection("users")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		result, _ := collection.InsertOne(ctx, user)
		mutex.Unlock()
		json.NewEncoder(w).Encode(result)
	} else {
		notFound(w, r)
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Add("content-type", "application/json")
		id := strings.TrimPrefix(r.URL.Path, "/users/")
		var user User
		collection := client.Database("instagram").Collection("users")
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		mutex.Lock()
		err := collection.FindOne(ctx, bson.D{{"uid", id}}).Decode(&user)
		mutex.Unlock()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
			return
		}
		json.NewEncoder(w).Encode(user)
	} else {
		notFound(w, r)
		return
	}

}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("content-type", "application/json")
		var post Post
		_ = json.NewDecoder(r.Body).Decode(&post)
		collection := client.Database("instagram").Collection("posts")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		mutex.Lock()
		result, _ := collection.InsertOne(ctx, post)
		mutex.Unlock()
		json.NewEncoder(w).Encode(result)
	} else {
		notFound(w, r)
		return
	}
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Add("content-type", "application/json")
		pid := strings.TrimPrefix(r.URL.Path, "/posts/")
		var post Post
		collection := client.Database("instagram").Collection("posts")
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		mutex.Lock()
		err := collection.FindOne(ctx, bson.D{{"pid", pid}}).Decode(&post)
		mutex.Unlock()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
			return
		}
		json.NewEncoder(w).Encode(post)
	} else {
		notFound(w, r)
		return
	}

}

func ListUserPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("content-type", "application/json")
		w.Header().Add("limit", "10")
		uid := strings.TrimPrefix(r.URL.Path, "/posts/users/")
		var posts []Post
		collection := client.Database("instagram").Collection("posts")
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		mutex.Lock()
		cursor, err := collection.Find(ctx, bson.M{"uid": uid})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
			return
		}
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var post Post
			cursor.Decode(&post)
			posts = append(posts, post)
		}
		if err := cursor.Err(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
			return
		}
		mutex.Unlock()
		json.NewEncoder(w).Encode(posts)
	} else {
		notFound(w, r)
		return
	}

}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/users", CreateUser)
	mux.HandleFunc("/users/", GetUser)
	mux.HandleFunc("/posts", CreatePost)
	mux.HandleFunc("/posts/", GetPost)
	mux.HandleFunc("/posts/users/", ListUserPost)
	http.ListenAndServe("localhost:8080", mux)
}
