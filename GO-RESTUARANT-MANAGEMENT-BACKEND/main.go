package main

import (
	"github.com/gin-gonic/gin"
	"github.com/phahladira-moja/restaurant-management-system/database"
	"github.com/phahladira-moja/restaurant-management-system/middleware"
	"github.com/phahladira-moja/restaurant-management-system/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {

	// use os to get environment variables
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}
