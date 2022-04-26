package models

type Consignment struct {
	ID          string       `bson:"id"`
	Weight      string       `bson:"weight"`
	Description string       `bson:"description"`
	Containers  []*Container `bson:"containers"`
	VesselID    string       `bson:"vessel_id"`
}
type Container struct {
	ID         string `bson:"id"`
	CustomerID string `bson:"customer_id"`
	UserID     string `bson:"user_id"`
	Origin     string `bson:"origin"`
}
