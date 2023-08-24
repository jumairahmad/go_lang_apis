package models

type Address struct {
	State   string `json:"state" bson:"user_state"`
	city    string `json:"city" bson:"user_city"`
	pincode int    `json:"pincode" bson:"user_pincode"`
}

type User struct {
	name    string  `json:"name" bson:"user_name"`
	age     int     `json:"age" bson:"user_age"`
	address Address `json: "address" bson:"user_address"`
}
