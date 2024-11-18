package controllers

import (
	"net/http"
	"quiz-go/databases/migration"
	"quiz-go/helpers/structs"
	"quiz-go/repository"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(c *gin.Context) {
	var result gin.H

	users, err := repository.GetAllUser(migration.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": users,
		}
	}

	c.JSON(http.StatusOK, result)
}

// InsertUser inserts a new user into the database
func InsertUser(c *gin.Context) {
	var user structs.User

	// Bind JSON input to User struct
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// Insert the user into the database
	err = repository.InsertUser(migration.DbConnection, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not insert user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
}

// UpdateUser updates an existing user in the database
func UpdateUser(c *gin.Context) {
	var user structs.User
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&user)
	if err != nil {
		panic(err)
	}

	user.ID = id

	err = repository.UpdateUser(migration.DbConnection, user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user from the database
func DeleteUser(c *gin.Context) {
	var user structs.User
	id, _ := strconv.Atoi(c.Param("id"))

	user.ID = id
	err := repository.DeleteUser(migration.DbConnection, user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, user)
}
