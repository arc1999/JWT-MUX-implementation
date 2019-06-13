package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var myKey = []byte("iamstilllearning")

func hPage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Println("Failed to generate token")
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8081/", nil)
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(body))
}
func fetch(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Println("Failed to generate token")
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8081/fetch", nil)
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(body))
}

func add(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Println("Failed to generate token")
	}

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://localhost:8081/add", r.Body)
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(body))
}

func search(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Println("Failed to generate token")
	}
//	keys,_  := r.URL.Query()["age"]
	vars:=mux.Vars(r)
	key:=vars["age"]
	//a,_:=strconv.Atoi(key)
	//str,_:=ioutil.ReadAll(r.Body)
//	fmt.Println(key)
	client := &http.Client{}
	req, _ := http.NewRequest("PUT", "http://localhost:8081/search/"+key, r.Body)
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)
	fmt.Println(res)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(body))
}
func update(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Println("Failed to generate token")
	}
	//	keys,_  := r.URL.Query()["age"]
	vars:=mux.Vars(r)
	key:=vars["age"]
	//a,_:=strconv.Atoi(key)
	//str,_:=ioutil.ReadAll(r.Body)
	//fmt.Println(key)
	client := &http.Client{}
	req, _ := http.NewRequest("PUT", "http://localhost:8081/update/"+key, r.Body)
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)
	fmt.Println(res)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(body))
}
func delete(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Println("Failed to generate token")
	}
	//	keys,_  := r.URL.Query()["age"]
	vars:=mux.Vars(r)
	key:=vars["age"]
	//a,_:=strconv.Atoi(key)
	//str,_:=ioutil.ReadAll(r.Body)
//	fmt.Println(key)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8081/delete/"+key, r.Body)
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)
	fmt.Println(res)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(body))
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "Ayush Chauhan"
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()
	claims["authorization_scope"]= "*"

	tokenString, err := token.SignedString(myKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func handleRequests() {
	myrouter :=mux.NewRouter().StrictSlash(true)

	myrouter.HandleFunc("/", hPage)
	myrouter.HandleFunc("/fetch", fetch)
	myrouter.HandleFunc("/add", add).Methods("POST")
	myrouter.HandleFunc("/search/{age}", search)
	myrouter.HandleFunc("/delete/{age}", delete)
	myrouter.HandleFunc("/update/{age}", update).Methods("PUT")

	//http.HandleFunc("/", hPage)
	//http.HandleFunc("/fetch", fetch)
	//http.HandleFunc("/add", add)
	//http.HandleFunc("/update/{age}", update)
	log.Fatal(http.ListenAndServe(":8082", myrouter))
}

func main() {
	handleRequests()
}