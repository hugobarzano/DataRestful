package models

type RequestOperation struct {
	Title  	string	 	 `json:"title"`
	Data   []string      `json:"data"`
	Operator string      `json:"operator"`
	Value    string      `json:"value"`
}

