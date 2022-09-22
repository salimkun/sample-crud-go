package service

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/salimkun/sample-crud-go/common/util"
	"github.com/salimkun/sample-crud-go/model"
	"github.com/salimkun/sample-crud-go/repository"
)

func RegisterUser(c *gin.Context) {
	var request model.User

	if err := c.ShouldBindJSON(&request); err != nil {
		msg := err.Error()
		if strings.Contains(err.Error(), "json: cannot unmarshal") {
			msg = "invalid parameter body please check request again"
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	// validation name
	if request.Name == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "name can't be null or empty"})
		return
	}

	// validation email
	if request.Email == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "email can't be null or empty"})
		return
	} else if addr, ok := util.ValidateMailAddress(request.Email); !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("email %s not valid", addr)})
		return
	}

	// validation phone number
	if request.Phone == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "phone_number can't be null or empty"})
		return
	} else if request.Phone[0:1] != "0" && request.Phone[0:3] != "+62" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": request.Phone[0:3]})
		return
	} else {
		users, err := repository.ReadFile()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		if len(users) > 0 {
			// validate if user exist
			for _, i := range users {
				if i.Phone == request.Phone {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "phone number already register"})
					return
				}
			}
		}
	}

	// validation linkein url
	if request.LinkedInUrl != "" {
		if !strings.Contains(request.LinkedInUrl, "linkedin.com/in/") && !util.ValidateUrl(request.LinkedInUrl) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "linkedin_url not valid"})
			return
		}
	}

	// validate portofolio url
	if request.PortofolioUrl != "" {
		if !util.ValidateUrl(request.PortofolioUrl) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "portofolio_url not valid"})
			return
		}
	}
	// validate occupation
	if len(request.Occupations) > 0 {
		for idx, i := range request.Occupations {
			err := ValidateOccupation(i)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("%e in index %d", err, idx)})
				return
			}

		}
	}

	// validate education
	if len(request.Educations) > 0 {
		for idx, i := range request.Educations {
			err := ValidateEducation(i)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("%e in index %d", err, idx)})
				return
			}
		}
	}

	e := repository.CreateUser(&request)
	if e != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": e})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": "success register user"})
}

func ValidateOccupation(request model.Occupation) error {

	// validation start_date
	if request.StartDate != "" {
		if !util.ValidateDate(request.StartDate) {
			return errors.New("occupation_start invalid format")
		}
	}

	// validate end_date
	if request.EndDate != "" {
		if !util.ValidateDate(request.EndDate) {
			return errors.New("occupation_end invalid format")
		}
	}

	return nil
}

func ValidateEducation(request model.Education) error {
	// validation start_date
	if request.StartDate != "" {
		if !util.ValidateDate(request.StartDate) {
			return errors.New("education_start invalid format")
		}
	}

	// validate end_date
	if request.EndDate != "" {
		if !util.ValidateDate(request.EndDate) {
			return errors.New("education_end invalid format")
		}
	}

	// validate score
	if request.Score > 0 {
		if request.Score > 4 {
			return errors.New("education_score invalid")
		}
	}
	return nil
}

func ListUser(c *gin.Context) {
	data, e := repository.ReadFile()
	if e != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": e})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func GetUserByID(c *gin.Context) {
	userIDStr := c.Param("userID")
	userID, _ := strconv.Atoi(userIDStr)
	data, e := repository.ReadFile()
	if e != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": e})
		return
	}

	if len(data) > 0 {
		// validate if user exist
		for _, i := range data {
			if i.ID == int64(userID) {
				c.JSON(http.StatusOK, gin.H{"data": i})
				return
			}
		}
	}

	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user not found"})
}

func UpdateUser(c *gin.Context) {
	var request model.User

	if err := c.ShouldBindJSON(&request); err != nil {
		msg := err.Error()
		if strings.Contains(err.Error(), "json: cannot unmarshal") {
			msg = "invalid parameter body please check request again"
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	// get data exisinng
	data, e := repository.ReadFile()
	if e != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": e})
		return
	}

	if request.ID == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "id must be fill"})
		return
	}

	var indx int
	var match bool
	if len(data) > 0 {
		// validate if user exist
		for idx, i := range data {
			if i.ID == request.ID {
				indx = idx
				match = true
			}
		}
	}

	if !match {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	// validation name
	if request.Name == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "name can't be null or empty"})
		return
	}

	// validation email
	if request.Email == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "email can't be null or empty"})
		return
	} else if addr, ok := util.ValidateMailAddress(request.Email); !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("email %s not valid", addr)})
		return
	}

	// validation phone number
	if request.Phone == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "phone_number can't be null or empty"})
		return
	} else if request.Phone[0:1] != "0" && request.Phone[0:3] != "+62" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": request.Phone[0:3]})
		return
	} else {
		if len(data) > 0 {
			// validate if user exist
			for _, i := range data {
				if i.Phone == request.Phone && i.ID != request.ID {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "phone number already register"})
					return
				}
			}
		}
	}

	// validation linkein url
	if request.LinkedInUrl != "" {
		if !strings.Contains(request.LinkedInUrl, "linkedin.com/in/") && !util.ValidateUrl(request.LinkedInUrl) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "linkedin_url not valid"})
			return
		}
	}

	// validate portofolio url
	if request.PortofolioUrl != "" {
		if !util.ValidateUrl(request.PortofolioUrl) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "portofolio_url not valid"})
			return
		}
	}
	// validate occupation
	if len(request.Occupations) > 0 {
		for idx, i := range request.Occupations {
			err := ValidateOccupation(i)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("%e in index %d", err, idx)})
				return
			}

		}
	}

	// validate education
	if len(request.Educations) > 0 {
		for idx, i := range request.Educations {
			err := ValidateEducation(i)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("%e in index %d", err, idx)})
				return
			}
		}
	}

	data[indx] = &request
	e = repository.UpdateUser(data)
	if e != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": e})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "success update data"})
}

func DeleteUser(c *gin.Context) {
	userIDStr := c.Param("userID")
	userID, _ := strconv.Atoi(userIDStr)
	data, e := repository.ReadFile()
	if e != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": e})
		return
	}

	var match bool
	var newList []*model.User
	if len(data) > 0 {
		// validate if user exist
		for _, i := range data {
			if i.ID == int64(userID) {
				match = true
			} else {
				newList = append(newList, i)
			}
		}
	}

	if !match {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	e = repository.UpdateUser(newList)
	if e != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": e})
		return
	}

	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "success delete user"})
}
