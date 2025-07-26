package model

import "time"

type Account struct {
	Id        int       `bson:"_id" json:"id" example:"1"`
	Name      string    `bson:"name" json:"name" example:"Evgenij"`
	FirstName string    `bson:"firstName" json:"first_name" example:"Kobzar"`
	LastName  string    `bson:"lastName" json:"last_name" example:""`
	Login     string    `bson:"login" json:"login" example:"ekobzar"`
	Password  string    `bson:"password" json:"password" example:"123456"`
	CreatedAt time.Time `bson:"createdAt" json:"created_at"`
}

func NewAccount() *Account {
	return &Account{}
}

func (a *Account) GetId() int {
	return a.Id
}

func (a *Account) SetId(id int) {
	a.Id = id
}
