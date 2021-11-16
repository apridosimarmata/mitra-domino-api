package Models

import "time"

type Order struct {
	Id        int       `json:"id"`
	BuyerId   string    `json:"buyerId"`
	CreatedAt time.Time `json:"createdAt"`

	UserId int
}
