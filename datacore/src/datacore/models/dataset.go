package models

// Product represents an e-comm item
type Dataset struct {
	ID     int 	     	 `bson:"_id"`
	Title  string        `json:"title"`
	Data   string        `json:"data"`
}

// Products is an array of Product objects
type Datasets []Dataset
