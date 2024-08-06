package dto

type CustomerResponse struct {
	Id          string `json:"customer_id"`
	Name        string `json:"full_name"`
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}
