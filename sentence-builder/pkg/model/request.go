package model

// Data is the Data passed from
// step function to function
// in an AWS State Machine
type Data struct {
	Subject string `json:"subject"`
}
