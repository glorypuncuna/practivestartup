package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/payment"
	"bwastartup/transaction"
	"bwastartup/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:micupowi@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	userService := user.NewService(userRepository)
	authService := auth.NewService()
	campaignService := campaign.NewService(campaignRepository)
	paymentService := payment.NewService()
	transactionService := transaction.NewService(transactionRepository, campaignRepository, paymentService)

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService, userService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	router := gin.Default()
	router.Static("/images", "./images")
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.LoginUser)
	api.POST("/check-email", userHandler.CheckEmail)
	api.POST("/upload-avatar", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.POST("/campaign", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.POST("/campaign-image", authMiddleware(authService, userService), campaignHandler.CreateImage)
	api.POST("/transaction", authMiddleware(authService, userService), transactionHandler.CreateTransaction)

	api.PUT("/campaign/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)

	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/detail-campaign/:id", campaignHandler.GetCampaignById)
	api.GET("/campaigns/:campaignId/transactions", authMiddleware(authService, userService), transactionHandler.GetTransactionsByCampaignId)
	api.GET("/transactions", authMiddleware(authService, userService), transactionHandler.GetTransactionsByUserId)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		if !strings.Contains(header, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		var token string
		arrayToken := strings.Split(header, " ")
		if len(arrayToken) == 2 {
			token = arrayToken[1]
		}

		validToken, err := authService.ValidateToken(token)
		if err != nil || !validToken.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		decodedToken := validToken.Claims.(jwt.MapClaims)
		userId := int(decodedToken["user_id"].(float64))

		user, err := userService.GetUserById(userId)
		if err != nil || !validToken.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}

}
