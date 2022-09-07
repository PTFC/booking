package main

import (
	"booking/api/flight-api/handlers"
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
	conn, err := grpc.Dial(":2222", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	//Singleton
	flightClient := pb.NewFPTFlightClient(conn)

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	//Handler for GIN Gonic
	h := handlers.NewFlightHandler(flightClient)
	os.Setenv("GIN_MODE", "debug")
	g := gin.Default()
	g.Use(middleware.LoggingMiddleware(logger))

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("minage", custom_validator.ValidAgeValidator(int64(18)))
	}

	//Create routes
	gr := g.Group("/v1/api")

	gr.POST("/create", h.CreateFlight)
	gr.POST("/update", h.UpdateFlight)
	gr.GET("/search", h.SearchFlight)

	//Listen and serve
	http.ListenAndServe(":3332", g)
}
