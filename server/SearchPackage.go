package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/gorilla/mux"
	// "github.com/jinzhu/gorm"
)

func searchPackage(c *gin.Context) {
	fmt.Println("inside search package")
	location := c.Param("location")
	fmt.Println(location)
	tokenAuth, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	fmt.Println("passed authorization check")
	fmt.Println(tokenAuth)
	// role := FetchRole(tokenAuth)
	// fmt.Println(role)
	// if  != nil {
	// 	c.JSON(http.StatusUnauthorized, "unauthorized")
	// 	fmt.Println(err)
	// 	return
	// }

	// if role == "Guide" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": "User needs to be a tourist. Please register as a tourist.",
	// 	})
	// 	return
	// }

	if location == " " {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please enter a valid location name.",
		})
		return

	}
	fmt.Println("passed location check")
	//email := c.Params.ByName("email"

	type PackageResponse struct {
		Id           string
		Email        string
		Duration     string
		Location     string
		Accomodation string
		Itinerary    string
		Included     string
		Price        string
	}
	pkg := []PackageResponse{}
	packages := DB.Table("packages").Where("location = ?", location).Find(&pkg)
	if packages.Error != nil {
		fmt.Println("failed DB check")
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No package currently available for this location.",
		})
		return
	} else {
		fmt.Println("passed DB check")
		c.JSON(200, pkg)
	}
}
