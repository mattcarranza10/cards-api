package main

import (
	"log"

	"cards-api/internal/auth"
	"cards-api/internal/cards"
	cardsHandler "cards-api/internal/cards/handler"
	cardsInfra "cards-api/internal/cards/infrastructure"
	"cards-api/internal/drivers"
	"cards-api/internal/encryption"
	"cards-api/internal/login"
	loginHandler "cards-api/internal/login/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	loginSvc := login.NewService()
	loginHand := loginHandler.NewHandler(loginSvc)

	cardRepo := cardsInfra.NewGormRepository(drivers.GetGormDB())
	cardSvc := cards.NewService(cardRepo, encryption.GetAES256EncryptionService())
	cardHand := cardsHandler.NewHandler(cardSvc)

	r := gin.Default()

	api := r.Group("cards-api")
	{
		loginGroup := api.Group("/v1/login", loginHand.Login)
		loginGroup.POST("")

		cardGroup := api.Group("/v1/cards")
		cardGroup.Use(auth.Middleware())
		cardGroup.POST("", cardHand.AddCard)
		cardGroup.GET("/:id", cardHand.GetCardDetails)
		cardGroup.PUT("/:id", cardHand.UpdateCardInformation)
		cardGroup.DELETE("/:id", cardHand.DeleteCard)

		eventGroup := api.Group("/v1/events")
		eventGroup.Use(auth.Middleware())
		eventGroup.POST("/cards", cardHand.UpdateCardsConcurrently)
	}

	r.Run(":8080")
}
