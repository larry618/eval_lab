package data

import "time"

type Item struct {
	ID               int64
	Name             string
	Category         string
	CategoryID       int64
	Discount         float64
	Rating           float64
	Status           uint8
	Price            float64
	WarehouseAddress *Address
	PaymentFeatures  *ItemPaymentFeatures
}

type User struct {
	ID                int64
	Name              string
	Age               int
	BirthDate         time.Time
	Gender            Gender
	Address           *Address
	Language          string
	Interests         []string
	UserTags          []string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Credit            Credit
	CreditLimit       float64
	Discount          float64
	Balance           float64
	isVip             bool
	isStudent         bool
	CurrentDevice     *Device
	RecentDevices     []*Device
	PreferCategoryIDs []int64
	PaymentFeatures   *UserPaymentFeatures
}

type Device struct {
	Brand      string
	Platform   string
	OSVersion  string
	AppVersion string
}

type Address struct {
	Country   string
	State     string
	City      string
	Street    string
	Latitude  float64
	Longitude float64
}

type UserPaymentFeatures struct {
	MtdCount  int64
	MtdAmount float64

	YtdCount  int64
	YtdAmount float64

	AvgPrice    float64
	TotalCount  int64
	TotalAmount float64

	LatestPurchaseItem *Item
	LatestPurchasedAt  time.Time
}

type ItemPaymentFeatures struct {
	MtdCount   int64
	YtdCount   int64
	TotalCount int64

	CountPerAge    map[int]int64
	CountPerGender map[Gender]int64

	RatingPerAge    map[int]float64
	RatingPerGender map[Gender]float64

	PreferredAges    []int
	PreferredGenders []Gender
}

type Gender int8

const (
	GenderMale   Gender = 1
	GenderFemale Gender = 2
	GenderOther  Gender = 3
)

func (g Gender) String() string {
	switch g {
	case GenderMale:
		return "male"
	case GenderFemale:
		return "female"
	case GenderOther:
		return "other"
	default:
		return "unknown"
	}
}

type Credit int8

const (
	CreditExcellent Credit = 3
	CreditGreat     Credit = 2
	CreditGood      Credit = 1
	CreditOK        Credit = 0
	CreditBad       Credit = -1
	CreditTerrible  Credit = -2
)

func (c Credit) String() string {
	switch c {
	case CreditExcellent:
		return "excellent"
	case CreditGreat:
		return "great"
	case CreditGood:
		return "good"
	case CreditOK:
		return "ok"
	case CreditBad:
		return "bad"
	case CreditTerrible:
		return "terrible"
	default:
		return "ok"
	}
}
