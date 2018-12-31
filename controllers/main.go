package controller

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/choa-mb/monster-db-api/db"
  "github.com/choa-mb/monster-db-api/models"
)

func GetMonsters(c *gin.Context) {
  monsters := []models.Monster{}
  err := db.GetDB().Select(&monsters, "SELECT * FROM monsters")

  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, monsters)
}

func GetMonsterById(c *gin.Context) {

}

func AddMonster(c *gin.Context) {

}

