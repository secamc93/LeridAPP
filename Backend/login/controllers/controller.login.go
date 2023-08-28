package controllers

import (
	"login/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Usuario    string `json:"usuario" binding:"required"`
	Contrasena string `json:"contrasena" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var usuario models.Usuario
	if err := models.GetUsuarioByUsuario(&usuario, input.Usuario); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
		return
	}

	if err := models.CheckPassword(&usuario, input.Contrasena); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Inicio de sesión exitoso",
		"id_usuario": usuario.ID,
		"nombre":     usuario.Nombre,
		"id_negocio": usuario.FkNegocio,
	})
}
