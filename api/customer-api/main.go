package main

import (
	"booking/api/customer-api/handlers"
	"booking/middleware"
	"booking/pb"
	custom_validator "booking/validator"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	//Create grpc client connect
	conn, err := grpc.Dial(":2223", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	//Singleton
	customerClient := pb.NewFPTCustomerClient(conn)

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	//Handler for GIN Gonic
	h := handlers.NewCustomerHandler(customerClient)
	os.Setenv("GIN_MODE", "debug")
	g := gin.Default()
	g.Use(middleware.LoggingMiddleware(logger))

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("minage", custom_validator.ValidAgeValidator(int64(18)))
	}

	//Create routes
	gr := g.Group("/v1/api")

	gr.POST("/create", h.CreateCustomer)
	gr.POST("/update", h.UpdateCustomer)
	gr.POST("/change-password", h.ChangePassword)
	gr.GET("/history", h.BookingHistory)

	//Listen and serve
	http.ListenAndServe(":3333", g)
}
