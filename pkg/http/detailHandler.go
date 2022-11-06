package http

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type RestDetails struct {
	Eateries []RestaurantDetails `json:"eateries"`
}

type RestaurantDetails struct {
	Id       string     `json:"id"`
	Name     string     `json:"name"`
	Location string     `json:"location"`
	Details  string     `json:"details"`
	Image    string     `json:"image"`
	Menu     []menuItem `json:"menu"`
	Comments []Comment  `json:"comments"`
}

type menuItem struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

type Comment struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

func GetRestDetails(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		b, _ := json.Marshal(err)
		w.Write(b)
	}

	file, err := ioutil.ReadFile("./pkg/db/data.json")
	if err != nil {
		log.Fatalln(err)
	}

	data := RestDetails{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatalln(err)
	}

	restId := r.FormValue("id")
	for _, restaurant := range data.Eateries {
		if restaurant.Id == restId {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			f, err := json.Marshal(restaurant)
			if err != nil {
				log.Fatalln(err)
			}
			w.Write(f)
			return
		}
	}

	w.WriteHeader(http.StatusNotAcceptable)
}

type myId struct {
	Id string `json:"id"`
}

func GetDetails(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		b, _ := json.Marshal(err)
		w.Write(b)
	}

	tmpl, err := template.ParseFiles("./pkg/templates/details.html")
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "text/html")
	err = tmpl.Execute(w, myId{Id: r.FormValue("id")})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
		return
	}

	//w.WriteHeader(http.StatusOK)
}
