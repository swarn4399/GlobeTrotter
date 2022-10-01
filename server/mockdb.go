package main

import "github.com/jinzhu/gorm"

func seed(db *gorm.DB) {
	// reg := []Register{
	// 	{Email: "saduvishesha@gmail.com", Password: "vish", Role: "Tourist"},
	// 	{Email: "visheshasadu@gmail.com", Password: "sadu", Role: "Guide"},
	// 	{Email: "vish@gmail.com", Password: "vish", Role: "Guide"},
	// 	{Email: "harshitha@gmail.com", Password: "harsh", Role: "Tourist"},
	// 	{Email: "megh@gmail.com", Password: "megh", Role: "Tourist"},
	// 	{Email: "arijit@gmail.com", Password: "arijit", Role: "Tourist"},
	// }
	// for _, c := range reg {
	// 	db.Create(&c)
	// }

	profiles := []UserProfile{
		{Email: "saduvishesha@gmail.com", Name: "vishesha", About: "travel freak", Age: 23, Fav1: "Gainesville", Fav2: "Orlando", Fav3: "tampa"}, //User: sv, Register: sv_reg
	}
	for _, m := range profiles {
		db.Create(&m)
	}

	profiles2 := []GuideProfile{
		{Email: "visheshasadu@gmail.com", Name: "visheshaSadu", About: "travel freak", Age: 23, Location: "Gainesville", Vehicle: "Ferrari car"}, //User: sv, Register: sv_reg
		{Email: "meghna@gmail.com", Name: "meghamala", About: "Friendly and funny tour guide", Age: 24, Location: "Gainesville", Vehicle: "Ferrari car"},
		//{Email: "meghna@gmail.com", Name: "meghamala", About: "Friendly and funny tour guide", Age: 26, Location: "Gainesville", Vehicle: "Ferrari car"},
	}

	for _, pr2 := range profiles2 {
		db.Create(&pr2)
	}

	// comments := []Comment{
	// 	{Comment: "Gainesville is peaceful and has nice greenery", Email: "visheshasadu@gmail.com", Location: "Gainesville"},   //User: sv, Register: sv_reg
	// 	{Comment: "Orlando is well developed and has many nice places to visit", Email: "vish@gmail.com", Location: "Orlando"}, // User: sv, Register: sv_reg
	// }
	// for _, cmnt := range comments {
	// 	db.Create(&cmnt)

	// }
	// type Req struct {
	// 	GuideEmail   string `json:"guideEmail" gorm:"not null"`
	// 	Duration     string `json: "duration"`
	// 	Location     string `json:"location"`
	// 	Accomodation string `json:"accomodation"`
	// 	Itinerary    string `json:"itinerary"`
	// 	Included     string `json:"included"`
	// 	Price        string `json:"price"`
	// }
	packages := []Package{
		{Email: "visheshas@gmail.com", Duration: "2 weeks", Location: "Florida", Accomodation: "Hyatt Regency", Itinerary: "Day1:Miami, Day2: Fort Lauderdale, Day3: Orlando", Included: "Breakfast,Dinner, Sight-seeing", Price: "$100"},
		{Email: "arijitd@gmail.com", Duration: "1 weeks", Location: "New Delhi", Accomodation: "Hyatt Regency", Itinerary: "Day1:Red Fort, Day2: Fatehpur Sikhri, Day3: Agra", Included: "Breakfast,Dinner, Sight-seeing", Price: "Rs. 10000"},
		{Email: "arijitd@gmail.com", Duration: "1 weeks", Location: "Lucknow", Accomodation: "Hyatt Regency", Itinerary: "Day1:Red Fort, Day2: Fatehpur Sikhri, Day3: Agra", Included: "Breakfast,Dinner, Sight-seeing", Price: "Rs. 10000"},
		{Email: "arijitd@gmail.com", Duration: "1 weeks", Location: "Kolkata", Accomodation: "Hyatt Regency", Itinerary: "Day1:Red Fort, Day2: Fatehpur Sikhri, Day3: Agra", Included: "Breakfast,Dinner, Sight-seeing", Price: "Rs. 10000"},
	}
	for _, pkg := range packages {
		db.Create(&pkg)

	}

	// bookings := []Booking{
	// 	{Email: "meghamalagupta@gmail.com", PackageId: "133"},
	// 	{Email: "meghamalagupta@gmail.com", PackageId: "134"}, // User: sv, Register: sv_reg
	// }
	// for _, m := range bookings {
	// 	db.Create(&m)
	// }
}
