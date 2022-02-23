package main

import (
	"github.com/gorilla/mux"
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"os"
)

type person struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
}
type feedback struct{
	ID string `json:"id"`
	Description string `json:"description"`
}

var feedbacks []feedback;
var persons []person;

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func getAllPerson(w http.ResponseWriter, r *http.Request)  {
	
	w.Header().Set("Content-Type","application/json");
	json.NewEncoder(w).Encode(persons);
}

func getPersonByid(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type","application/json");
	
	params := mux.Vars(r);
	fmt.Println(params);
	for _, item := range persons{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item);
			return
		}
	}

}

func getPersonByName(w http.ResponseWriter, r *http.Request){	
	w.Header().Set("Content-Type","application/json");
	query:=r.URL.Query();
	fmt.Println(query["name"])
	for _, item := range persons{
		if item.Name==query["name"][0]{
			json.NewEncoder(w).Encode(item)
		}
	}

}

func createPerson(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var Person person
	_ = json.NewDecoder(r.Body).Decode(&Person);
	Person.ID = strconv.Itoa(rand.Intn(10000000000));
	persons = append(persons, Person);
	json.NewEncoder(w).Encode(Person);
	fmt.Println(Person);

}

func getAllFeedback(w http.ResponseWriter, r *http.Request)  {
	enableCors(&w);
	w.Header().Set("Content-Type","application/json");
	json.NewEncoder(w).Encode(feedbacks);
}
func createFeedback(w http.ResponseWriter, r *http.Request){
	// setupResponse(&w,r);
	// header := w.Header()
	// header.Add("Access-Control-Allow-Origin", "*")
	// header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	// header.Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	w.Header().Set("Content-Type","application/json")
	// if r.Method == "OPTIONS" {
	// 	w.WriteHeader(http.StatusOK)
	// 	return
	// }
	
	var Feedback feedback
	_ = json.NewDecoder(r.Body).Decode(&Feedback);
	Feedback.ID = strconv.Itoa(rand.Intn(10000000000));
	feedbacks = append(feedbacks, Feedback);
	json.NewEncoder(w).Encode(Feedback);
	fmt.Println(Feedback);
}

func main()  {
	r:=mux.NewRouter();
	port := os.Getenv("PORT")	
	persons = append(persons, person{ID: "1",Name: "Abubakar",Address: "H.no 285"});
	persons = append(persons, person{ID: "2",Name: "Saad",Address: "H.no 282"});

	feedbacks = append(feedbacks, feedback{ID: "1",Description: "first feedback check"});
	feedbacks = append(feedbacks, feedback{ID: "2",Description: "second feedback check"});

	r.HandleFunc("/getAllperson",getAllPerson).Methods("GET");
	r.HandleFunc("/getpersonbyID/{id}",getPersonByid).Methods("GET");
	r.HandleFunc("/createPerson",createPerson).Methods("Post");
	r.HandleFunc("/byName",getPersonByName).Methods("GET");

	r.HandleFunc("/getallfeedback",getAllFeedback).Methods("GET");
	r.HandleFunc("/createfeedback",createFeedback).Methods("Post");

	fmt.Print("Server Starting at port: "+port);
	log.Fatal(http.ListenAndServe(":"+port,r))
}