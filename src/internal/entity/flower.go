package entity

import "github.com/google/uuid"

// FlowerList - List flowers.
type FlowerList struct {
	Flowers []*Flower
}

// Flower - Struct.
type Flower struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price float64   `json:"price"`
}
