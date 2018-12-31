package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/choa-mb/monster-db-api/db"
  MonstersController "github.com/choa-mb/monster-db-api/controllers"
)

func main() {
  fmt.Println("Starting DnD monster REST API service")

  db.Init()

  r := gin.Default()

  v1 := r.Group("/api/v1")
  {
    monsters := v1.Group("/monsters")
    {
      monsters.GET("/", MonstersController.GetMonsters)
      monsters.POST("/", MonstersController.AddMonster)
      monsters.GET("/:id", MonstersController.GetMonsterById)
    }
  }

  r.Run()
}

