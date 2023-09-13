package structs

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main/config"
	mongof "main/functions/mongo"
)

type PremiumUser struct {
	ID           string       `json:"id"`
	Customer     string       `json:"customer"`
	PremiumLevel PremiumLevel `json:"premium_level"`
	ExpiresAt    int64        `json:"expires_at"`
}
type PremiumLevel int

const (
	MONTH = "month"
	YEAR  = "year"
)

func NewPremium(userId, customer string, createdAt int64, yearMonth string) *PremiumUser {
	//expires := createdAt
	switch yearMonth {
	case MONTH:
	case YEAR:

	}
	return &PremiumUser{ID: userId, Customer: customer, PremiumLevel: 1}
}
func (puser *PremiumUser) RegisterNewCustomer() {
	if puser.Exists() {
		return
	}

}
func (puser *PremiumUser) Exists() bool {
	_, err := mongof.FindOne(bson.M{"id": puser.ID}, options.FindOne(), config.MongoDatabase, "premium")
	if err != nil || err == mongo.ErrNoDocuments {
		return false
	}
	return true
}
