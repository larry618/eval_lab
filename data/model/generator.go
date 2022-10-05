package model

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

var (
	interests   = []string{"painting", "video_games", "travel", "reading", "sports", "technology", "cooking", "writing"}
	phoneBrands = []string{"apple", "samsung", "google", "oneplus", "motorola", "xiaomi", "huawei", "sony", "oppo", "lg"}

	userTags = [][]string{
		{"new", "active", "inactive", "return"},
		{"normal", "professional", "celebrity"},
		{"top", "high_value", "medium_value", "low_value", "lost"},
		{"A", "B", "C", "D", "E", "F"},
	}

	categoryIDs = []int64{1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000}

	FakeNow = func() time.Time { return time.Unix(2000000000, 0) }
)

func GenUser(f gofakeit.Faker) *User {
	u := new(User)
	u.ID = int64(f.IntRange(0, math.MaxInt))
	u.Name = f.Name()
	u.BirthDate = f.DateRange(time.Unix(0, 0), FakeNow())
	u.Age = FakeNow().Year() - u.BirthDate.Year()
	u.Gender = chooseWithP(f, 0.1, GenderOther, Gender(f.RandomInt([]int{1, 2})))
	u.Address = GenAddress(f)
	u.Language = f.LanguageBCP()
	u.CreatedAt = f.DateRange(u.BirthDate, FakeNow())
	u.UpdatedAt = f.DateRange(u.CreatedAt, FakeNow())
	u.Credit = Credit(f.IntRange(-1, 3))
	u.CreditLimit = chooseWithP(f, 0.8, 0, f.Float64Range(0, 5000))
	u.Discount = chooseWithP(f, 0.5, 1, f.Float64Range(0.5, 1))
	u.Balance = chooseWithP(f, 0.6, 0, f.Float64Range(0, 5000))
	u.IsStudent = chooseWithP(f, 0.1, true, false)
	u.IsVip = chooseWithP(f, 0.3, true, false)
	u.CurrentDevice = GenDevice(f)
	u.PaymentFeatures = GenUserPaymentFeatures(f)
	u.Interests = sampleArray(f, interests, f.IntRange(0, 6))

	for i := 0; i < f.IntRange(0, 6); i++ {
		u.RecentDevices = append(u.RecentDevices, GenDevice(f))
	}

	for _, tags := range userTags {
		u.UserTags = append(u.UserTags, f.RandomString(tags))
	}

	return u
}

func GenItem(f gofakeit.Faker) *Item {
	u := new(Item)
	u.ID = f.Int64()
	u.Name = f.Emoji()
	u.Category = f.RandomString(interests)
	u.CategoryID = sampleArray(f, categoryIDs, 1)[0]
	u.Discount = chooseWithP(f, 0.5, 1, f.Float64Range(0.7, 1))
	u.Rating = f.Float64Range(3, 5)
	u.Status = uint8(f.RandomInt([]int{0, 1, 2}))
	u.Price = f.Price(5, 2000)
	u.WarehouseAddress = GenAddress(f)
	u.PaymentFeatures = GenItemPaymentFeatures(f)
	return u
}

func GenAddress(f gofakeit.Faker) *Address {
	return &Address{
		Country:   f.CountryAbr(),
		State:     f.StateAbr(),
		City:      f.City(),
		Street:    f.Street(),
		Latitude:  f.Latitude(),
		Longitude: f.Longitude(),
	}
}

func GenItemPaymentFeatures(f gofakeit.Faker) *ItemPaymentFeatures {
	u := new(ItemPaymentFeatures)
	u.MtdCount = int64(f.IntRange(0, 1000))
	u.YtdCount = int64(f.IntRange(int(u.MtdCount), 100000))
	u.TotalCount = int64(f.IntRange(int(u.YtdCount), 5000000))

	u.CountPerAge = make(map[int]int64)
	u.CountPerGender = make(map[Gender]int64)
	u.RatingPerAge = make(map[int]float64)
	u.RatingPerGender = make(map[Gender]float64)

	for i := 15; i < 60; i++ {
		u.CountPerAge[i] = int64(f.IntRange(0, 20000))
		u.RatingPerAge[i] = f.Float64Range(3, 5)
	}

	u.CountPerGender = map[Gender]int64{
		GenderFemale: int64(f.IntRange(0, int(u.TotalCount)/3)),
		GenderMale:   int64(f.IntRange(0, int(u.TotalCount)/3)),
		GenderOther:  int64(f.IntRange(0, int(u.TotalCount)/3)),
	}

	u.RatingPerGender = map[Gender]float64{
		GenderMale:   f.Float64Range(3, 5),
		GenderFemale: f.Float64Range(3, 5),
		GenderOther:  f.Float64Range(3, 5),
	}

	for i := 0; i < f.IntRange(0, 23); i++ {
		u.PreferredAges = append(u.PreferredAges, f.IntRange(1, 100))
	}

	u.PreferredGenders = sampleArray(f, []Gender{GenderOther, GenderFemale, GenderMale}, f.IntRange(0, 3))
	return u
}

func GenUserPaymentFeatures(f gofakeit.Faker) *UserPaymentFeatures {
	u := new(UserPaymentFeatures)

	u.TotalCount = int64(f.IntRange(0, 1000))
	if u.TotalCount != 0 {
		u.TotalAmount = f.Float64Range(5000, 500000)
	}

	if u.TotalCount != 0 {
		u.YtdCount = int64(f.IntRange(0, int(u.TotalCount)))
		if u.YtdCount != 0 {
			u.YtdAmount = f.Float64Range(1, u.TotalAmount)
		}
	}

	if u.YtdCount != 0 {
		u.MtdCount = int64(f.IntRange(0, int(u.YtdCount)))
		if u.MtdCount != 0 {
			u.MtdAmount = f.Float64Range(1, u.YtdAmount)
		}
	}

	u.AvgPrice = u.TotalAmount / float64(u.TotalCount)

	u.LatestPurchaseItem = GenItem(f)
	u.LatestPurchasedAt = f.DateRange(time.Unix(0, 0), FakeNow())

	u.PreferCategoryIDs = sampleArray(f, categoryIDs, f.IntRange(0, 4))
	return u
}

func GenDevice(f gofakeit.Faker) *Device {
	d := new(Device)
	d.Brand = f.RandomString(phoneBrands)

	if d.Brand == "apple" {
		d.Platform = "ios"
	} else {
		d.Platform = "android"
	}

	if d.Platform == "ios" {
		d.OSVersion = fmt.Sprintf("%d.%d", f.IntRange(13, 16), f.IntRange(0, 5))
	} else {
		d.OSVersion = strconv.Itoa(f.IntRange(8, 14))
	}

	d.AppVersion = fmt.Sprintf("%d.%d.%d", f.IntRange(5, 10), f.IntRange(0, 10), f.IntRange(0, 10))
	return d
}

func sampleArray[T any](f gofakeit.Faker, a []T, size int) []T {
	l := len(a)
	size = IF(l > size, size, l)

	res := make([]T, 0, size)

	set := make(map[int]bool, size)

	for len(res) < size {
		i := f.IntRange(0, l-1)
		if !set[i] {
			res = append(res, a[i])
			set[i] = true
		}
	}

	return res
}

func chooseWithP[T any](f gofakeit.Faker, probability float64, a, b T) T {
	return IF(f.Float64Range(0, 1) < probability, a, b)
}

func IF[V any](b bool, v1, v2 V) V {
	if b {
		return v1
	}
	return v2
}
