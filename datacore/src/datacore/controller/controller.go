package controller

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"io"
	"log"

	"datacore/models"
	"datacore/mongo"
	"github.com/gorilla/mux"
	"strconv"
	"strings"
	"fmt"
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

// UpdateDataset PUT /
func (c *Controller) UpdateDataset(w http.ResponseWriter, r *http.Request) {
	var dataset models.Dataset
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Fatalln("Error UpdateDataset", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}



	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error UpdateDataset", err)
	}

	if err := json.Unmarshal(body, &dataset); err != nil { // unmarshall body contents as a type Candidate
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error UpdateDataset unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	fmt.Println("TO BE UPDATED Dataset ID - ", dataset.ID)
	success := c.Storer.UpdateDataset(dataset) // updates the dataset in the DB

	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

// GetDataset GET - Gets a single dataset by ID /
func (c *Controller) GetDataset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)

	id := vars["id"] // param id
	log.Println(id);
	datasetid, err := strconv.Atoi(id);

	if err != nil {
		log.Fatalln("Error GetDataset", err)
	}

	dataset := c.Storer.GetDatasetById(datasetid)
	data, _ := json.Marshal(dataset)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// DeleteDataset DELETE /
func (c *Controller) DeleteDataset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)
	id := vars["id"] // param id
	log.Println(id);

	datasetid, err := strconv.Atoi(id);

	if err != nil {
		log.Fatalln("Error GetDataset", err)
	}

	if err := c.Storer.DeleteDataset(datasetid); err != "" { // delete a dataset by id
		log.Println(err);
		if strings.Contains(err, "404") {
			w.WriteHeader(http.StatusNotFound)
		} else if strings.Contains(err, "500") {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

// SearchDataset GET /
func (c *Controller) SearchDataset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)

	query := vars["query"] // param query
	log.Println("Search Query - " + query);

	datasets := c.Storer.GetDatasetsByString(query)
	data, _ := json.Marshal(datasets)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
