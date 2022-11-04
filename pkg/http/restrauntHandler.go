package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Restaurant struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Image    string `json:"image"`
}

type Restaurants struct {
	Eateries []Restaurant `json:"eateries"`
}

func GetRestaurants(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("./pkg/db/data.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	data := Restaurants{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Println(data)

	x, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = w.Write(x)
	if err != nil {
		log.Fatalln(err)
	}
}
