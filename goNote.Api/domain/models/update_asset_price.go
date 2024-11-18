package models

type UpdateAssetPrice struct {
	Price float64 `json:"price" bson:"price"`
}
