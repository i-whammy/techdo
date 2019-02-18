package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// 問題1 
func FizzBuzz(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	num, _ := strconv.Atoi(p.ByName("num"))

	if num <= 0 {
		http.Error(w, "Input must be positive number", http.StatusBadRequest)
		return
	}

	var resp string

	switch {
		case num % 15 == 0:
			resp = "FizzBuzz!"
		case num % 5 == 0:
			resp = "Buzz"
		case num % 3 == 0:
			resp = "Fizz"
		default:
			resp = strconv.Itoa(num)
	}
	fmt.Fprintln(w, resp)
}

// 問題4
func Name(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("name")

	if name != savedProfile.Name {
		http.Error(w, "Name is not registered", http.StatusBadRequest)
		return
	}
	
	jsonResp, _ := json.Marshal(savedProfile)
	fmt.Fprintln(w, fmt.Sprintf("%+v", string(jsonResp)))
}

type ProfileJSON struct {
	Name			string 	`json:"name"`
	Age				int 	`json:"age"`
	Gender			string  `json:"gender"`
	FavoriteFoods	[]string  `json:"favoirte_foods"`		
}
var savedProfile = ProfileJSON{}
	
// 問題3
func Profile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var param ProfileJSON

	if err := json.Unmarshal(bodyBytes, &param) ; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if param.Name == savedProfile.Name {
		http.Error(w, "Name is duplicated.", http.StatusBadRequest)
		return
	} else {
		savedProfile.Name = param.Name
		savedProfile.Age = param.Age
		savedProfile.Gender = param.Gender
		savedProfile.FavoriteFoods = param.FavoriteFoods
	}
}

func main() {
	router := httprouter.New()

	router.GET("/FizzBuzz/:num", FizzBuzz)
	router.GET("/Profile/:name", Name)
	router.POST("/Profile", Profile)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}