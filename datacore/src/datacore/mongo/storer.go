package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strings"

	"datacore/models"
)

//Repository ...
type Storer struct{
	Sesion Session
}

// SERVER the DB server
const SERVER = "http://127.0.0.1:27017"

// DBNAME the name of the DB instance
const DBNAME = "dataCore"

// COLLECTION is the name of the collection in DB
const COLLECTION = "datasets"

const SERVICES = "services"

var datasetId = 10


// GetDatasets returns the list of Datasets
func (s Storer) GetDatasets() models.Datasets {
	session, err := NewSession(SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	//defer session.Close()

	c := session.Copy().session.DB(DBNAME).C(COLLECTION)
	results := models.Datasets{}

	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}
	session.session.Close()
	return results
}

// GetDatasettById returns a unique Dataset
func (s Storer) GetDatasetById(id string ) models.Dataset {
	session, err := NewSession(SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	//defer session.Close()

	c := session.Copy().session.DB(DBNAME).C(COLLECTION)
	var result models.Dataset

	fmt.Println("ID in GetDatasetById", id)
	//bson.M{"_id": id}
	query:=bson.M{"_id": id}
	if err := c.Find(query).One(&result); err != nil {
		fmt.Println("Failed to write result:", err)
	}

	//result.ID=bson.ObjectId(result.ID).Hex()
	session.Close()
	return result
}

// GetDatasetbyString takes a search string as input and returns Datasets
func (s Storer) GetDatasetsByString(query string) models.Datasets {
	session, err := NewSession(SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	//defer session.Close()

	c := session.Copy().session.DB(DBNAME).C(COLLECTION)
	result := models.Datasets{}

	// Logic to create filter
	qs := strings.Split(query, " ")
	and := make([]bson.M, len(qs))
	for i, q := range qs {
		and[i] = bson.M{"title": bson.M{
			"$regex": bson.RegEx{Pattern: ".*" + q + ".*", Options: "i"},
		}}
	}
	filter := bson.M{"$and": and}

	if err := c.Find(&filter).Limit(5).All(&result); err != nil {
		fmt.Println("Failed to write result:", err)
	}

	//for _,r := range result{
	//	r.ID=bson.ObjectId(r.ID).Hex()
	//}
	session.Close()
	return result
}

// AddDataset adds a Dataset in the DB
func (s Storer) AddDataset(dataset models.Dataset) bool {
	session, err := NewSession(SERVER)
	//defer session.Close()

	i := bson.NewObjectId()
	dataset.ID=i.Hex()
	session.Copy().session.DB(DBNAME).C(COLLECTION).Insert(&dataset)

	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Added New  Dataset Title- ", dataset.Title)
	session.Close()
	return true
}

// UpdateDataset updates a dataset in the DB
func (s Storer) UpdateDataset(dataset models.Dataset) bool {
	session, err := NewSession(SERVER)
	//defer session.Close()


	//session.Copy().session.DB(DBNAME).C(COLLECTION).UpdateId(dataset.ID, dataset)
	err = session.Copy().session.DB(DBNAME).C(COLLECTION).UpdateId(dataset.ID, dataset)

	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Updated Dataset ID - ", dataset.ID)
	session.Close()
	return true
}

// DeleteDataset deletes an Dataset
func (s Storer) DeleteDataset(id string) string {
	session, err := NewSession(SERVER)
	//defer session.Close()

	// Remove dataset
	query:=bson.M{"_id": id}
	if err = session.Copy().session.DB(DBNAME).C(COLLECTION).Remove(query); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}

	fmt.Println("Deleted Dataset ID - ", id)
	// Write status
	session.Close()
	return "OK"
}

// AddDataset adds a Dataset in the DB
func (s Storer) AddService(service models.Service) bool {
	session, err := NewSession(SERVER)
	//defer session.Close()

	i := bson.NewObjectId()
	service.ID=i.Hex()
	session.Copy().session.DB(DBNAME).C(SERVICES).Insert(&service)

	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Added New  Service - ", service.Title)
	session.Close()
	return true
}

// GetDatasets returns the list of Datasets
func (s Storer) ListServices() models.Services {
	session, err := NewSession(SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	//defer session.Close()

	c := session.Copy().session.DB(DBNAME).C(SERVICES)
	results := models.Services{}

	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}
	session.Close()
	return results
}

func (s Storer) GetServicesByString(query string) models.Services {
	session, err := NewSession(SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	//defer session.Close()

	c := session.Copy().session.DB(DBNAME).C(SERVICES)
	result := models.Services{}

	// Logic to create filter
	qs := strings.Split(query, " ")
	and := make([]bson.M, len(qs))
	for i, q := range qs {
		and[i] = bson.M{"title": bson.M{
			"$regex": bson.RegEx{Pattern: ".*" + q + ".*", Options: "i"},
		}}
	}
	filter := bson.M{"$and": and}

	if err := c.Find(&filter).Limit(5).All(&result); err != nil {
		fmt.Println("Failed to write result:", err)
	}
	session.Close()
	return result
}

// DeleteDataset deletes an Dataset
func (s Storer) DeleteService(id string) string {
	session, err := NewSession(SERVER)
	//defer session.Close()

	// Remove dataset
	query:=bson.M{"_id": id}
	if err = session.Copy().session.DB(DBNAME).C(SERVICES).Remove(query); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}

	fmt.Println("Deleted service ID - ", id)
	// Write status
	session.Close()
	return "OK"
}
