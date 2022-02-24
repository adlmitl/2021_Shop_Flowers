package entity

import "github.com/google/uuid"

// BasketList - Список корзин.
type BasketList struct {
	Baskets []Basket
}

// Basket - Корзина.
type Basket struct {
	Id       uuid.UUID
	Count    int
	IdUser   *User
	IdFlower *[]Flower
}
