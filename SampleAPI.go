package main

import ("github.com/gorilla/mux"
        "log"
        "net/http"
        "encoding/json"
      )

// Person Data
type Person struct{
 ID        string `json:"id,omitempty"`
 Firstname string `json:"firstname,omitempty"`
 Lastname  string `json:"lastname,omitempty"`
 Address   *Address `json:"address,omitempty"`
}

// Address
type Address struct{
  City string `json:"city,omitempty"`
  Country string `json:"country,omitempty"`
}

// Array of person
var people []Person

// Gets All people GET Request : localhost:8080/people
func getPeople(w http.ResponseWriter, r *http.Request){
  json.NewEncoder(w).Encode(people)
}

// Gets One person from people GET Request : localhost:8080/person/{id}
func getPerson(w http.ResponseWriter, r *http.Request)  {
    param:=mux.Vars(r);
    for _,item:=range people{
    if item.ID == param["id"] {
      json.NewEncoder(w).Encode(item)
      return
      }
    }
    json.NewEncoder(w).Encode(&Person{})
}

// Creates One person POST Request : localhost:8080/person
/*
post body raw data{
"id":"5",
"firstname":"john",
"lastname":"doe"
}*/
func createPerson(w http.ResponseWriter, r *http.Request)  {
      params:=mux.Vars(r)
      var person Person
      _=json.NewDecoder(r.Body).Decode(&person)
      person.ID = params["id"]
      people = append(people,person)
      json.NewEncoder(w).Encode(people)
}

// Deletes One person DELETE Request : localhost:8080/person/{id}
func deletePerson(w http.ResponseWriter, r *http.Request)  {
  params:=mux.Vars(r)
  for index,item := range people{
    if item.ID == params["id"]{
      people  = append(people[:index],people[index+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(people)
}

// Routes Function and Sample Data
func main() {
    router := mux.NewRouter()
    people = append(people,Person{ID:"1",Firstname:"ahmed",Lastname:"hagag",Address: &Address{Country:"Egypt",City:"Cairo"}})
    people = append(people,Person{ID:"2",Firstname:"nehal",Lastname:"hagag"})
    people = append(people,Person{ID:"3",Firstname:"adam",Lastname:"hagag"})
    //Routes
    router.HandleFunc("/people",getPeople).Methods("GET");
    router.HandleFunc("/person/{id}",getPerson).Methods("GET");
    router.HandleFunc("/person",createPerson).Methods("POST");
    router.HandleFunc("/person/{id}",deletePerson).Methods("DELETE");
    // Server Start on port :8080
    log.Fatal(http.ListenAndServe(":8080",router))
}
