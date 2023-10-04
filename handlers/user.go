package handlers

import (
	dto "myapp/dto/result"
	usersdto "myapp/dto/users"
	"myapp/models"
	"myapp/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handler struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handler {
	return &handler{UserRepository}
}

func (h *handler) FindUsers(c echo.Context) error {
	users, err := h.UserRepository.FindUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: users})
}

func (h *handler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(user)})
}

func (h *handler) AddUser(c echo.Context) error {
	// Mendapatkan data pengguna dari permintaan
	var userRequest usersdto.CreateUserRequest
	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
	}

	// Membuat objek User dari data yang diterima
	user := models.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	// Menambahkan pengguna ke repository
	err := h.UserRepository.AddUser(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
	}

	// Mengembalikan respons berhasil
	return c.JSON(http.StatusCreated, dto.SuccessResult{Code: http.StatusCreated, Data: convertResponse(user)})
}

func (h *handler) UpdateUser(c echo.Context) error {
	// Mendapatkan ID pengguna dari parameter URL
	id, _ := strconv.Atoi(c.Param("id"))

	// Mendapatkan data pengguna dari permintaan
	var userRequest usersdto.UpdateUserRequest
	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
	}

	// Membuat objek User dari data yang diterima
	user := models.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	// Memperbarui pengguna ke dalam repository
	err := h.UserRepository.UpdateUser(id, &user)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
	}

	// Mengembalikan respons berhasil
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "User updated successfully"})
}

func (h *handler) DeleteUser(c echo.Context) error {
	// Mendapatkan ID pengguna dari parameter URL
	id, _ := strconv.Atoi(c.Param("id"))

	// Menghapus pengguna dari repository
	err := h.UserRepository.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
	}

	// Mengembalikan respons berhasil
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "User deleted successfully"})
}

func convertResponse(u models.User) *usersdto.UserResponse {
	return &usersdto.UserResponse{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}
