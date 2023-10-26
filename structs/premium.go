package structs

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main/config"
	mongof "main/functions/mongo"
	"time"
)

type PremiumUser struct {
	ID           string       `json:"id"`
	Subscription string       `json:"subscription"`
	PremiumLevel PremiumLevel `json:"premium_level"`
	ExpiresAt    float64      `json:"expires_at"`
}

type PremiumLevel int

const (
	MONTH = "month"
	YEAR  = "year"
	prem  = "premium"
)

func NewPremium(userId, subscription string, ExpiresAt float64 /*, yearMonth string*/) *PremiumUser {
	//expires := createdAt
	/*	switch yearMonth {
		case MONTH:
		case YEAR:

		}*/
	return &PremiumUser{ID: userId, Subscription: subscription, PremiumLevel: 1, ExpiresAt: ExpiresAt}
}
func (puser *PremiumUser) RegisterNewCustomer() {
	if puser.Exists() {
		return
	}
	puser.new()
}
func (puser *PremiumUser) upgradePremiumLevel(newLevel PremiumLevel, expirationDate float64) {
	puser.PremiumLevel = newLevel
	puser.ExpiresAt = expirationDate
	// Update the premium user's data in the database
	puser.Update(false)
}

func (puser *PremiumUser) IsPremiumExpired() bool {
	return puser.ExpiresAt < float64(time.Now().Unix())
}
func (puser *PremiumUser) RenewPremiumSubscription(expirationDate float64) {
	puser.ExpiresAt = expirationDate
	// Update the premium user's data in the database
	puser.Update(false)
}

func (puser *PremiumUser) CancelPremiumSubscription() {
	// You may want to archive the user's data or perform any necessary cleanup
	// Mark the user as non-premium and set the expiration date to zero
	puser.PremiumLevel = 0
	puser.ExpiresAt = 0
	// Update the premium user's data in the database
	puser.Update(false)
}

func (puser *PremiumUser) FetchPremiumDetails() error {
	data, err := mongof.FindOne(bson.M{"id": puser.ID}, options.FindOne(), config.MongoDatabase, prem)
	if err != nil {
		return err
	}
	puser.FromMap(data)
	return nil
}
func (puser *PremiumUser) new() error {
	_, err := mongof.InsertOne(puser, options.InsertOne(), config.MongoDatabase, prem)
	return err
}
func (puser *PremiumUser) Update(upsert bool) error {
	_, err := mongof.UpdateOne(puser, bson.M{
		"id": puser.ID,
	}, options.Update().SetUpsert(upsert), config.MongoDatabase, prem)
	return err
}

func (puser *PremiumUser) Exists() bool {
	_, err := mongof.FindOne(bson.M{"id": puser.ID}, options.FindOne(), config.MongoDatabase, prem)
	if err != nil || err == mongo.ErrNoDocuments {
		return false
	}
	return true
}
func (puser *PremiumUser) ToMap() (data map[string]any) {
	b, _ := json.Marshal(puser)
	json.Unmarshal(b, &data)
	return
}
func (puser *PremiumUser) FromMap(data any) {
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &puser)
}
