package structs

type PremiumUser struct {
	ID           string
	Customer     string
	PremiumLevel PremiumLevel
	ExpiresAt    int64
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
