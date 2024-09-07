package main

import (
	postgresql "hotelservice/database/postgreSQL"
	proto "hotelservice/proto"
	"hotelservice/server/service"
	"log"
	"net"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := os.Getenv("HOTEL_PORT")
	dataSourceName := os.Getenv("dataSourceName")

	db, err := postgresql.PostgreSQL("postgres", dataSourceName)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", "hotelservice"+port)
	if err != nil {
		log.Println(err)
	}
	defer lis.Close()

	server := service.NewHotelService(db)
	s := grpc.NewServer()
	proto.RegisterHotelServiceServer(s, server)
	reflection.Register(s)

	log.Println("Server is listening on port", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error to Server - %v", err)
	}
}
