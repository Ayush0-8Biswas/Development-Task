package http

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
)

func MakeComment(w http.ResponseWriter, r *http.Request) {
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
	for k, restaurant := range data.Eateries {
		if restaurant.Id == restId {
			w.Header().Set("Content-Type", "application/json")
			name, text := r.FormValue("name"), r.FormValue("text")
			restaurant.Comments = append(restaurant.Comments, Comment{Name: name, Text: text})
			data.Eateries[k] = restaurant

			finalData, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(http.StatusNotModified)
				log.Fatalln(err)
			}

			err = ioutil.WriteFile("./pkg/db/data.json", finalData, fs.ModeAppend)
			if err != nil {
				w.WriteHeader(http.StatusNotModified)
				log.Fatalln(err)
			}

			w.WriteHeader(http.StatusOK)
			return
		}
	}

	w.WriteHeader(http.StatusNotModified)
}
