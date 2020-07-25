package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shakilbd009/go-users-api/domain/users"
	"github.com/shakilbd009/go-users-api/services"
	"github.com/shakilbd009/go-users-api/utils/errors"
)

func getUserId(userId string) (int64, *errors.RestErr) {

	userID, userErr := strconv.ParseInt(userId, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userID, nil
}

func Get(c *gin.Context) {

	userID, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	result, getErr := services.UsersService.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshal(c.GetHeader("X-Public") == "true"))
}

func Create(c *gin.Context) {

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshal(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {

	userID, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.Id = userID
	isPartial := c.Request.Method == http.MethodPatch
	result, err := services.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshal(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {

	userID, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	if err := services.UsersService.DeleteUser(userID); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {

	status := c.Query("status")
	users, err := services.UsersService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshal(c.GetHeader("X-Public") == "true"))
}
