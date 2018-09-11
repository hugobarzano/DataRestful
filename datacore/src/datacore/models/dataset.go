package models

//
type Dataset struct {
	ID     string     	 `bson:"_id" json:"_id,omitempty"`
	Title  string        `json:"title"`
	Data   []string      `json:"data"`
}

//
type Datasets []Dataset


