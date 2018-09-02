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

var datasetId = 10


// GetDatasets returns the list of Datasets
func (s Storer) GetDatasets() models.Datasets {
	session, err := NewSession(SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}


	defer session.Close()

	c := session.Copy().session.DB(DBNAME).C(COLLECTION)
	results := models.Datasets{}

	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

// GetDatasettById returns a unique Dataset
func (s Storer) GetDatasetById(id int) models.Dataset {
	session, err := NewSession(SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.Copy().session.DB(DBNAME).C(COLLECTION)
	var result models.Dataset

	fmt.Println("ID in GetDatasetById", id)

	if err := c.FindId(id).One(&result); err != nil {
		fmt.Println("Failed to write result:", err)
	}

	return result
}

// GetDatasetbyString takes a search string as input and returns Datasets
func (s Storer) GetDatasetsByString(query string) models.Datasets {
	session, err := NewSession(SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

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

	return result
}

// AddDataset adds a Dataset in the DB
func (s Storer) AddDataset(dataset models.Dataset) bool {
	session, err := NewSession(SERVER)
	defer session.Close()

	datasetId += 1
	dataset.ID = datasetId
	session.Copy().session.DB(DBNAME).C(COLLECTION).Insert(dataset)

	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Added New  Dataset ID- ", dataset.ID)

	return true
}

// UpdateDataset updates a dataset in the DB
func (s Storer) UpdateDataset(dataset models.Dataset) bool {
	session, err := NewSession(SERVER)
	defer session.Close()


	//session.Copy().session.DB(DBNAME).C(COLLECTION).UpdateId(dataset.ID, dataset)
	err = session.Copy().session.DB(DBNAME).C(COLLECTION).UpdateId(dataset.ID, dataset)

	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Updated Dataset ID - ", dataset.ID)

	return true
}

// DeleteDataset deletes an Dataset
func (s Storer) DeleteDataset(id int) string {
	session, err := NewSession(SERVER)
	defer session.Close()

	// Remove dataset
	if err = session.Copy().session.DB(DBNAME).C(COLLECTION).RemoveId(id); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}

	fmt.Println("Deleted Dataset ID - ", id)
	// Write status
	return "OK"
}
