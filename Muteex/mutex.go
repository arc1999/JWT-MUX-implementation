package Muteex

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)
type Employee struct {
	Name string `json:"name"`
	Age int `json:"age"`
}


type E1 []Employee

var (
	mutex    sync.Mutex
	S1       E1
	myKey         = []byte("iamstilllearning")
	myrouter  =mux.NewRouter().StrictSlash(true)
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return myKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				next.ServeHTTP(w, r)
			}
		} else {

			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func update(s4 E1,s3 Employee,a int,w http.ResponseWriter){
	mutex.Lock()
	defer mutex.Unlock()
	for index,articles :=range S1 {
		if(articles.Age==a){
			s4=append(S1[:index],s3)
			//json.NewEncoder(w).Encode(s4v)
			s4=append(s4, S1[index+1:]...)
		}
	}

	json.NewEncoder(w).Encode(s4)
}

func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}


func fetchemployee(w http.ResponseWriter, r *http.Request){

	json.NewEncoder(w).Encode(S1)
	fmt.Println("endpoint hi Homepage")

}
func returnsingleemployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["age"]
	a, _ := strconv.Atoi(key)

	for _, articles := range S1 {
		if (articles.Age==a) {
			json.NewEncoder(w).Encode(articles)
		}

	}

	//fmt.Fprintf(w, "Key: "+key)
}
func addemployee( w http.ResponseWriter, r *http.Request){
	Str,_:=ioutil.ReadAll(r.Body)
	var s2 Employee
	json.Unmarshal(Str,&s2)
	S1 =append(S1,s2)
	fmt.Println(S1)
	json.NewEncoder(w).Encode(s2)


}
func deleteemployee(w http.ResponseWriter ,r *http.Request){
	vars:=mux.Vars(r)
	key:=vars["age"]
	a, _ := strconv.Atoi(key)
	fmt.Printf("delete hit")
	for index,articles:= range S1 {

		if(articles.Age==a){

			S1 =append(S1[:index], S1[index+1:]...)
			if(index==0){
				fmt.Fprint(w,"last element deletion query")
			}
		}
	}
}
func updateemployee(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	key:=vars["age"]
	a,_:=strconv.Atoi(key)
	str,_:=ioutil.ReadAll(r.Body)
	var s3 Employee
	var s4 E1
	json.Unmarshal(str,&s3)
	update(s4,s3,a,w)
}
func Handlerequests() {
	myrouter.Use(JwtAuthentication)

	myrouter.HandleFunc("/", homepage)
	myrouter.HandleFunc("/fetch", fetchemployee)
	myrouter.HandleFunc("/add", addemployee).Methods("POST")
	myrouter.HandleFunc("/search/{age}", returnsingleemployee)
	myrouter.HandleFunc("/delete/{age}", deleteemployee)
	myrouter.HandleFunc("/update/{age}", updateemployee).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8081", myrouter))

}
