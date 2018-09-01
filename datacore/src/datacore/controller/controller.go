package controller

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"io"
	"log"

	"datacore/models"
	"datacore/mongo"
)
//Controller ...
type Controller struct {
	Storer mongo.Storer
}

// Index GET /
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	datasets := c.Storer.GetDatasets() // list of all datasets
	data, _ := json.Marshal(datasets)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// AddDataset POST /
func (c *Controller) AddDataset(w http.ResponseWriter, r *http.Request) {
	var dataset models.Dataset
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request

	log.Println(body)

	if err != nil {
		log.Fatalln("Error AddDataset", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddDataset", err)
	}

	if err := json.Unmarshal(body, &dataset); err != nil { // unmarshall body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddDataset unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	log.Println(dataset)
	success := c.Storer.AddDataset(dataset) // adds the dataset to the DB
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}

