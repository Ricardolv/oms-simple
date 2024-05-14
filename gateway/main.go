package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"

	common "github.com/Ricardolv/commons"
	_ "github.com/joho/godotenv/autoload"

	pb "github.com/Ricardolv/commons/api"
)

var (
	httpAddr          = common.EnvString("HTTP_ADDR", ":3000")
	ordersServiceAddr = "localhost:2000"
)

func main() {

	conn, err := grpc.Dial(ordersServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	log.Println("Dialog orders service at ", ordersServiceAddr)

	client := pb.NewOrderServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHandler(client)
	handler.registerRoutes(mux)

	log.Println("Starting http server on", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server:", err)
	}

}
