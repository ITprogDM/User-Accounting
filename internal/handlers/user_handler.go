package handlers

import (
	"UchetUsers/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateUserHandler создаёт нового пользователя
// @Summary Создать пользователя
// @Description Регистрирует нового пользователя в системе
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "Данные пользователя"
// @Success 201 {object} map[string]string "message: Пользователь успешно создан"
// @Failure 400 {object} map[string]string "error: Некорректный формат данных"
// @Failure 500 {object} map[string]string "error: Ошибка создания пользователя"
// @Router /users [post]
func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		h.logger.WithError(err).Error("Ошибка бинда JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}

	err := h.service.CreateUser(c.Request.Context(), u)
	if err != nil {
		h.logger.WithError(err).Error("Ошибка создания пользователя")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания пользователя", "details": err.Error()})
		return
	}

	h.logger.Infof("Пользователь %s успешно создан", u.Email)
	c.JSON(http.StatusCreated, gin.H{"message": "Пользователь успешно создан"})
}

// GetUserHandler получает информацию о пользователе
// @Summary Получить пользователя
// @Description Получает данные пользователя по ID
// @Tags users
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string "error: Некорректный формат ID"
// @Failure 404 {object} map[string]string "error: Пользователь не найден"
// @Router /users/{id} [get]
func (h *UserHandler) GetUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.WithError(err).Warn("Некорректный формат id")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат id"})
		return
	}

	user, err := h.service.GetUser(c.Request.Context(), id)
	if err != nil {
		h.logger.WithError(err).Error("Ошибка поиска пользователя")
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserHandler обновляет данные пользователя
// @Summary Обновить пользователя
// @Description Обновляет информацию о пользователе
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "Обновлённые данные пользователя"
// @Success 200 {object} map[string]string "message: Пользователь успешно обновлён"
// @Failure 400 {object} map[string]string "error: Некорректный формат данных"
// @Failure 500 {object} map[string]string "error: Ошибка обновления пользователя"
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUserHandler(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		h.logger.WithError(err).Error("Ошибка бинда JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}

	if err := h.service.UpdateUser(c.Request.Context(), u.ID, u.Name, u.Email, u.Age); err != nil {
		h.logger.WithError(err).Error("Ошибка обновления данных пользователя")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления пользователя", "details": err.Error()})
		return
	}

	h.logger.Infof("Пользователь с id %d успешно обновлён", u.ID)
	c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно обновлён"})
}

// DeleteUserHandler удаляет пользователя
// @Summary Удалить пользователя
// @Description Удаляет пользователя из базы данных
// @Tags users
// @Param id path int true "ID пользователя"
// @Success 200 {object} map[string]string "message: Пользователь успешно удалён"
// @Failure 400 {object} map[string]string "error: Некорректный формат ID"
// @Failure 500 {object} map[string]string "error: Ошибка удаления пользователя"
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.WithError(err).Error("Некорректный формат ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат ID"})
		return
	}

	if err = h.service.DeleteUser(c.Request.Context(), id); err != nil {
		h.logger.WithError(err).Error("Ошибка удаления пользователя!")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления пользователя", "details": err.Error()})
		return
	}

	h.logger.Infof("Пользователь с id %d успешно удалён!", id)
	c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно удалён!"})
}
