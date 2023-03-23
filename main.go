package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/graduation-fci/service-graph/dependencies"
	"github.com/graduation-fci/service-graph/proto"
	"github.com/graduation-fci/service-graph/server"
	"google.golang.org/grpc"
)

func main() {
	dp := dependencies.NewDependencyInjection().WithNeo4j()
	defer dp.Shutdown()

	go healthCheck()
	StartGRPC(dp)
}

func StartGRPC(dp *dependencies.DP) {
	listner, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	graphServer := server.NewGraphServer(dp)
	proto.RegisterGraphServiceServer(grpcServer, graphServer)
	log.Printf("server listening at %v", listner.Addr())
	if err := grpcServer.Serve(listner); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func healthCheck() {
	mode := gin.ReleaseMode
	if os.Getenv("env") != "prod" {
		mode = gin.DebugMode
	}

	gin.SetMode(mode)
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})

	router.Run(":" + os.Getenv("PORT"))
}
