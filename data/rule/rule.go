package rule

import (
	"context"
	"math"
	"time"

	"github.com/gurkankaymak/hocon"
	"github.com/onheap/eval"
	"github.com/onheap/eval_lab/data/model"
)

const (
	UserID eval.VariableKey = iota
	Gender
	Age
	UserTags
	Interests
	BirthDate

	Address
	AddressCountry
	AddressState
	AddressCity
	Language

	Credit
	CreditLimit
	Discount
	Balance

	CreatedAt
	UpdatedAt
	Platform
	OSVersion
	AppVersion

	IsStudent
	IsVip
)

func ToEvalCtx(ctx context.Context, u *model.User) *eval.Ctx {
	return &eval.Ctx{
		Ctx: ctx,
		VariableFetcher: eval.SliceVarFetcher(
			[]eval.Value{
				UserID:         u.ID,
				Gender:         int64(u.Gender),
				Age:            int64(u.Age),
				IsStudent:      u.IsStudent,
				IsVip:          u.IsVip,
				UserTags:       u.UserTags,
				Interests:      u.Interests,
				CreatedAt:      u.CreatedAt.Unix(),
				Address:        u.Address,
				AddressCountry: u.Address.Country,
				AddressState:   u.Address.State,
				AddressCity:    u.Address.City,
				Credit:         int64(u.Credit),
				CreditLimit:    int64(u.CreditLimit),
				Discount:       int64(u.Discount * 100),
				Balance:        int64(u.Balance),
				Language:       u.Language,
				UpdatedAt:      u.UpdatedAt.Unix(),
				Platform:       u.CurrentDevice.Platform,
				OSVersion:      u.CurrentDevice.OSVersion,
				AppVersion:     u.CurrentDevice.AppVersion,
				BirthDate:      u.BirthDate.Unix(),
			},
		),
	}
}

func CompileConfig() *eval.Config {
	return &eval.Config{
		CostsMap:       make(map[string]float64),
		CompileOptions: make(map[eval.CompileOption]bool),

		ConstantMap:        ConstantMap(),
		VariableKeyMap:     VariableKeyMap(),
		OperatorMap:        OperatorMap(),
		StatelessOperators: []string{"is_birthday", "distance"},
	}
}

func VariableKeyMap() map[string]eval.VariableKey {
	return map[string]eval.VariableKey{
		"user_id":         UserID,
		"gender":          Gender,
		"age":             Age,
		"is_student":      IsStudent,
		"is_vip":          IsVip,
		"user_tags":       UserTags,
		"interests":       Interests,
		"created_at":      CreatedAt,
		"address":         Address,
		"address.country": AddressCountry,
		"address.state":   AddressState,
		"address.city":    AddressCity,
		"credit":          Credit,
		"credit_limit":    CreditLimit,
		"discount":        Discount,
		"balance":         Balance,
		"language":        Language,
		"updated_at":      UpdatedAt,
		"platform":        Platform,
		"os_version":      OSVersion,
		"app_version":     AppVersion,
		"birth_date":      BirthDate,
	}
}

func ConstantMap() map[string]eval.Value {
	return map[string]eval.Value{
		// genders
		"Male":   int64(model.GenderMale),
		"Female": int64(model.GenderFemale),
		"Other":  int64(model.GenderOther),

		// credits
		"Excellent": int64(model.CreditExcellent),
		"Great":     int64(model.CreditGreat),
		"Good":      int64(model.CreditGood),
		"OK":        int64(model.CreditOK),
		"Bad":       int64(model.CreditBad),
		"Terrible":  int64(model.CreditTerrible),

		// duration
		"Day":   int64(60 * 60 * 24),
		"Week":  int64(60 * 60 * 24 * 7),
		"Month": int64(60 * 60 * 24 * 30),
		"Year":  int64(60 * 60 * 24 * 365),

		"Headquarters": &model.Address{
			Country:   "US",
			State:     "CA",
			City:      "San Francisco",
			Latitude:  37.77172396341641,
			Longitude: -122.40537627272957,
		},
	}
}

func OperatorMap() map[string]eval.Operator {
	return map[string]eval.Operator{
		"now": func(_ *eval.Ctx, _ []eval.Value) (eval.Value, error) {
			return model.FakeNow().Unix(), nil
		},

		"is_birthday": func(_ *eval.Ctx, params []eval.Value) (eval.Value, error) {
			a, b, e := eval.DestructParamsInt2("is_birthday", params)
			if e != nil {
				return nil, e
			}
			birthday := time.Unix(a, 0)
			today := time.Unix(b, 0)

			_, m1, d1 := birthday.Date()
			_, m2, d2 := today.Date()

			return m1 == m2 && d1 == d2, nil
		},

		"distance": func(_ *eval.Ctx, params []eval.Value) (eval.Value, error) {
			if len(params) != 2 {
				return nil, eval.ParamsCountError("distance", 2, len(params))
			}
			a1, ok := params[0].(*model.Address)
			if !ok {
				return nil, eval.ParamTypeError("distance", "*model.Address", params[0])
			}

			a2, ok := params[1].(*model.Address)
			if !ok {
				return nil, eval.ParamTypeError("distance", "*model.Address", params[2])
			}

			radius := 6371000.0
			rad := math.Pi / 180.0
			lat1 := a1.Latitude * rad
			lng1 := a1.Longitude * rad
			lat2 := a2.Latitude * rad
			lng2 := a2.Longitude * rad
			theta := lng2 - lng1
			dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
			return int64(dist * radius / 1000), nil
		},
	}
}

func LoadRules() ([]string, error) {
	conf, err := hocon.ParseResource("data/rule/rules.conf")
	if err != nil {
		return nil, err
	}
	var rules []string
	for _, rule := range conf.GetRoot().(hocon.Array) {
		r := rule.(hocon.Object)["rule"].String()
		rules = append(rules, r)
	}
	return rules, nil
}
