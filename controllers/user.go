package controllers

import (
	"net/http"
	"yukevent/lib/database"
	mid "yukevent/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterUserController(c echo.Context) error {

	var userReq RequestUserRegister

	c.Bind(&userReq)

	result, err := database.RegisterUser(userReq.toModel())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, newResponse(*result))
}

func LoginUserController(c echo.Context) error {
	var userReq RequestUserLogin

	c.Bind(&userReq)

	result, err := database.LoginUser(userReq.toModel())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	token, err := mid.CreateTokenJWT(result.ID, result.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	// If no error will return the Token
	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func GetUser(c echo.Context) error {

	users, err := database.GetUser()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, newResponseArray(*users))
}

// func LoginUserController(c echo.Context) error {

// 	var userReq RequestLoginUser
// 	var users model.User

// 	c.Bind(&userReq)
// 	// Find The Email
// 	err := config.DB.Where("email = ?", userReq.Email).First(&users).Error

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"message": err.Error(),
// 		})
// 	}
// 	// Compare Password
// 	err = helper.CheckPasswordHash(userReq.Password, users.Password)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"message": err.Error(),
// 		})
// 	}
// 	// Create JWT Token
// 	token, err := mid.CreateTokenJWT(users.ID, users.Name)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"message": err.Error(),
// 		})
// 	}
// 	// If no error will return the Token
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"token": token,
// 	})
// }

// func LoginUser(userInput model.User) (string, error) {

// 	var users model.User

// 	err := config.DB.Where("email = ?", userInput.Email).First(&users).Error

// 	if err != nil {
// 		return &model.User{}, err
// 	}

// 	err = helper.CheckPasswordHash(userInput.Password, users.Password)

// 	if err != nil {
// 		return &model.User{}, err
// 	}

// 	users.Token, err = mid.CreateTokenJWT(users.ID, users.Name)

// 	if err != nil {
// 		return &model.User{}, err
// 	}

// 	return &users, nil
// }
