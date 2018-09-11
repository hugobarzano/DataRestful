package models

type Operation struct {
	Dataset_id  string	 	 `json:"dataset_id"`
	Value    string      `json:"value"`
	Operator string      `json:"operator"`
	Service_url  string  `json:"service_url"`

}
