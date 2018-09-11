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
	"strings"
	"fmt"
	"bytes"
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


	dataset := c.Storer.GetDatasetById(id)
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


	if err := c.Storer.DeleteDataset(id); err != "" { // delete a dataset by id
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

//Operatiions

// Index GET /
func (c *Controller) PerformsOperation(w http.ResponseWriter, r *http.Request) {

	var op models.Operation
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request

	if err != nil {
		log.Fatalln("Error performsOperation", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error performsOperation", err)
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &op); err != nil { // unmarshall body contents as a struct
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error performsOperation unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	dataset := c.Storer.GetDatasetById(op.Dataset_id)
	fmt.Println("Dataset - ",dataset)

	req_op:=models.RequestOperation{Title:dataset.Title,Data:dataset.Data,Operator:op.Operator,Value:op.Value}
	fmt.Println("OPeration - ",req_op)


	operation_body := new(bytes.Buffer)
	json.NewEncoder(operation_body).Encode(req_op)
	//service_respond,err := http.Post("http://127.0.0.1:5000/basicOperator", "application/json; charset=utf-8", operation_body)
	service_respond,err := http.Post(op.Service_url, "application/json; charset=utf-8", operation_body)
	if err != nil{
		fmt.Print(err)
	}

	if service_respond.StatusCode!=http.StatusOK{
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	b, err := ioutil.ReadAll(service_respond.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var dataset_generated models.Dataset
	err = json.Unmarshal(b, &dataset_generated)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	success:=c.Storer.AddDataset(dataset_generated) // adds the dataset to the DB
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//
	data, _ := json.Marshal(&dataset_generated)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return

}

// AddService POST /
func (c *Controller) AddService(w http.ResponseWriter, r *http.Request) {
	var service models.Service
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request

	log.Println(body)

	if err != nil {
		log.Fatalln("Error AddService", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddService", err)
	}

	if err := json.Unmarshal(body, &service); err != nil { // unmarshall body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddService unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	log.Println(service)
	success := c.Storer.AddService(service) // adds the dataset to the DB
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}


func (c *Controller) ListServices(w http.ResponseWriter, r *http.Request) {
	services := c.Storer.ListServices() // list of all register services
	data, _ := json.Marshal(services)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}


// SearchDataset GET /
func (c *Controller) DeleteService(w http.ResponseWriter, r *http.Request) {

	fmt.Println("DELETE SERVICE")
	vars := mux.Vars(r)
	log.Println(vars)

	query := vars["query"] // param query
	log.Println("Search Query - " + query);

	services := c.Storer.GetServicesByString(query)

	for _,service := range(services){
		c.Storer.DeleteService(service.ID)
	}

	services = c.Storer.ListServices() // list of all register services
	data, _ := json.Marshal(services)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}