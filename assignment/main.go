package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"protoimport/assignment"
	"syscall"
	"time"
)

const ApiKeyPreviewChars = 5

func main() {
	dbUrl := os.Getenv("POSTGRES_URL")
	if dbUrl == "" {
		dbUrl = "postgres://postgres:test@localhost:3333/blocksize?sslmode=disable&pool_max_conns=10"
	}
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "block!size23487flsdkjfh*&^%8G#5m"
	}

	if len(secret) != 32 {
		_, _ = fmt.Fprintf(os.Stderr, "Incorrect SECRET length %d. Must be 32 symbols", len(secret))
		os.Exit(1)
	}

	dbPool, err := initDbPool(dbUrl)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer dbPool.Close()

	err = Migrate(dbPool)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to execute migrations: %v\n", err)
		os.Exit(1)
	}

	// Starting listener
	listener, _ := net.Listen("tcp", ":50051")

	// Create new gRPC implementation
	impl := service{
		apikeyRepository: &apiKeyRepository{db: dbPool},
		crypt:            NewCrypt(secret)}

	// Create default gRPC server and register implementation with it
	server := grpc.NewServer()
	assignment.RegisterApikeyServiceServer(server, &impl)

	// Serve the gRPC server using the previously generated listener
	go func() {
		_ = server.Serve(listener)
	}()
	log.Println("Serving @:50051")

	// Termination handling
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGTERM)
	signal.Notify(sigc, syscall.SIGINT)
	select {
	case <-sigc:
		log.Println("Terminating gracefully...")
		server.GracefulStop()
		log.Println("bye.")
	}
}

func initDbPool(dbUrl string) (*pgxpool.Pool, error) {
	var dbPool *pgxpool.Pool
	var start = time.Now()
	for {
		var err error
		dbPool, err = pgxpool.Connect(context.Background(), dbUrl)
		if err == nil {
			break
		}

		if time.Now().Sub(start).Seconds() > 30 {
			return nil, err
		}
		time.Sleep(100)
	}
	_, _ = fmt.Fprintf(os.Stdout, "Connected to database: %v\n", dbUrl)
	return dbPool, nil
}
