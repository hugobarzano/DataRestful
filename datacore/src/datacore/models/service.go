package models

// Product represents an e-comm item
type Service struct {
	ID     string     	 `bson:"_id" json:"_id,omitempty"`
	Title  string        `json:"title"`
	Url	   string        `json:"url"`
}

// Products is an array of Product objects
type Services []Service
