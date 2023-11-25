package handler

import (
	"miniproject/config"
	"miniproject/entity"
	"miniproject/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c echo.Context) error {

	input := new(entity.User)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Gagal melakukan hash password"})
	}

	user := entity.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     input.Role,
	}
	if err := config.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Gagal mendaftar pengguna"})
	}

	user.Password = ""

	return c.JSON(http.StatusCreated, map[string]interface{}{"message": "Pendaftaran berhasil", "user": user})
}

func LoginUser(c echo.Context) error {

	input := new(entity.User)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}

	var user entity.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "Kredensial tidak valid"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "Kredensial tidak valid"})
	}

	tokenString, err := middleware.GenerateToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Gagal menghasilkan token"})
	}

	user.Password = ""

	return c.JSON(http.StatusOK, map[string]interface{}{"token": tokenString, "message": "Login berhasil"})
}
