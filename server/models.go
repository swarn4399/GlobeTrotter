package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/gorm"
)

type GeoResponse struct {
	Country    string  `json:"country"`
	Timezone   string  `json:"timezone"`
	Name       string  `json:"name"`
	Lon        float64 `json:"lon"`
	Lat        float64 `json:"lat"`
	Population int64   `json:"population"`
}

type RadiusResponse struct {
	Xid      string  `json:"xid"`
	Name     string  `json:"name"`
	Dist     float64 `json:"dist"`
	Rate     string  `json:"rate"`
	Wikidata string  `json:"wikidata"`
	Kinds    string  `json:"kinds"`
	Osm      string  `json:"osm"`
	P        Point   `json:"point"`
}

type Point struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type ObjectResponse struct {
	Xid       string             `json:"xid"`
	Name      string             `json:"name"`
	Kinds     string             `json:"kinds"`
	Osm       string             `json:"osm"`
	Wikidata  string             `json:"wikidata"`
	Rate      string             `json:"rate"`
	Image     string             `json:"image"`
	Prev      Preview            `json:"preview"`
	Wikipedia string             `json:"wikipedia"`
	Wiki_ext  WikiPedia_Extracts `json:"wikipedia_extracts"`
	Voyage    string             `json:"voyage"`
	Url       string             `json:"url"`
	Otm       string             `json:"otm"`
	Sources   Source             `json:"sources"`
	Inf       Info               `json:"info"`
	Bbox      BBox               `json:"bbox"`
	P         Point              `json:"point"`
	Addr      Address            `json:"address"`
}

type Preview struct {
	Source string `json:"source"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type WikiPedia_Extracts struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Html  string `json:"html"`
}

type Source struct {
	Geometry   string   `json:"geometry"`
	Attributes []string `json:"attributes"`
}
type Info struct {
	Src    string `json:"src"`
	Src_id int64  `json:"src_id"`
	Descr  string `json:"descr"`
}
type BBox struct {
	Lon_min int64 `json:"lon_min"`
	Lon_max int64 `json:"lon_max"`
	Lat_min int64 `json:"lat_min"`
	Lat_max int64 `json:"lat_max"`
}
type Address struct {
	City          string `json:"city"`
	Road          string `json:"road"`
	State         string `json:"state"`
	County        string `json:"county"`
	Country       string `json:"country"`
	Postcode      string `json:"postcode"`
	Country_code  string `json:"country_code"`
	House_number  string `json:"house_number"`
	Neighbourhood string `json:"neighbourhood"`
}
type SearchPlacesResponse struct {
	Xid      string             `json:"xid"`
	Name     string             `json:"name"`
	Kinds    string             `json:"kinds"`
	Prev     Preview            `json:"preview"`
	Wiki_ext WikiPedia_Extracts `json:"wikipedia_extracts"`
	Inf      Info               `json:"info"`
}

// type Register struct {
// 	gorm.Model
// 	Username string `json:"username"`
// 	Email    string `gorm:"primaryKey" json:"email"`
// 	Password string `json:"password" binding:"required,min=8,max=20"`
// 	Role     string `gorm:"primaryKey" json:"role"`
// }

type Register struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required,min=8,max=20"`
	Role     string `json:"role"`
}

type UserProfile struct {
	gorm.Model
	Email string `gorm:"unique" json:"email"`
	Name  string `json:"name"`
	//Password string `json:"password" binding:"required,min=8,max=20"`
	Role     string `json:"role"`
	About    string `json:"about"`
	Age      uint   `json:"age"`
	Mobile   string `json:"mobile"`
	Location string `json:"location"`
	Fav1     string `json:"fav1"`
	Fav2     string `json:"fav2"`
	Fav3     string `json:"fav3"`
}

type GuideProfile struct {
	gorm.Model
	Email string `gorm:"unique" json:"email"`
	//Password string `json:"name"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	About    string `json:"about"`
	Age      uint   `json:"age"`
	Address  string `json:"address"`
	Location string `json:"location"`
	Vehicle  string `json:"vehicle"`
}

type Comment struct {
	gorm.Model
	Title     string `json:"title"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	PackageId string `json:"packageId"`
	Review    string `json:"review"`
	//Register Register `gorm:"foreignKey:Email; json:"register"`
}

type Package struct {
	gorm.Model
	// PackageId    string `gorm:"primaryKey;autoIncrement:true"`
	Email        string `json:"email"`
	Duration     string `json:"duration"`
	Location     string `json:"location"`
	Accomodation string `json:"accomodation"`
	Itinerary    string `json:"itinerary"`
	Included     string `json:"included"`
	Price        string `json:"price"`
}

type Booking struct {
	gorm.Model
	Email     string `gorm:"primaryKey" json:"email"`
	PackageId string `gorm:"primaryKey" json:"packageId"`
}
