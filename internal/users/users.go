package users

import (
	"github.com/gin-gonic/gin"
)

type Users struct{}

func NewUsers() *Users {
	return &Users{}
}

// TODO: validation step for each handler
func (u Users) CreateHandler(c *gin.Context) {
}

func (u Users) ByIdHandler(c *gin.Context) {
}

func (u Users) ListHandler(c *gin.Context) {
}

func (u Users) UpdateHandler(c *gin.Context) {
}

func (u Users) DeleteHandler(c *gin.Context) {
}
