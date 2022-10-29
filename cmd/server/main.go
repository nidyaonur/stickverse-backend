package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/k0kubun/pp"
	"github.com/nidyaonur/stickverse-backend/chat"
	"github.com/nidyaonur/stickverse-backend/internal/auth"
	"github.com/nidyaonur/stickverse-backend/internal/data"
	"github.com/nidyaonur/stickverse-backend/internal/logic"
	"github.com/nidyaonur/stickverse-backend/internal/maplist"
	"github.com/nidyaonur/stickverse-backend/internal/structure"
	"github.com/nidyaonur/stickverse-backend/internal/user"
	"github.com/nidyaonur/stickverse-backend/pkg/env"
	"github.com/nidyaonur/stickverse-backend/pkg/pb"
	"github.com/nidyaonur/stickverse-backend/pkg/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var Repo repository.Repository

func accessibleRoles() map[string][]string {
	const laptopServicePath = "/techschool.pcbook.LaptopService/"

	return map[string][]string{
		laptopServicePath + "CreateLaptop": {"admin"},
		laptopServicePath + "UploadImage":  {"admin"},
		laptopServicePath + "RateLaptop":   {"admin", "user"},
	}
}

func main() {
	env := env.GetEnvironment()
	pp.Print(env)

	// Database
	repo := repository.NewRepository(env.WriterDBUrl)
	defer repo.CloseDBConnection()

	jwtManager := auth.NewJWTManager(env.JwtSecret, 30*24*time.Hour)
	//moduleFlag := flag.String("module", "server", "module argument")
	//flag.Parse()
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)

	}
	logicService := logic.NewService(repo)
	dataGenerator := data.NewDataGenerator(repo)
	dataGenerator.InitDataValues()

	authServer := auth.NewAuthServer(repo, logicService, jwtManager)
	userServer := user.NewUserServer(repo, logicService, jwtManager)
	maplistServer := maplist.NewMaplistServer(repo, logicService, jwtManager)
	structureServer := structure.NewStructureServer(repo, logicService, jwtManager)

	interceptor := auth.NewAuthInterceptor(jwtManager, accessibleRoles())
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	}
	grpcServer := grpc.NewServer(serverOptions...)
	pb.RegisterAuthServiceServer(grpcServer, authServer)
	pb.RegisterUserServiceServer(grpcServer, userServer)
	pb.RegisterMaplistServiceServer(grpcServer, maplistServer)
	pb.RegisterStructureServiceServer(grpcServer, structureServer)
	reflection.Register(grpcServer)

	fmt.Println("Starting server...")
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}

func client() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := chat.NewChatServiceClient(conn)
	message := chat.Message{
		Body: "Hello, world!",
	}
	response, err := c.SendMessage(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error while calling SendMessage: %v", err)
	}
	fmt.Printf("Response from server: %s\n", response.Body)

}
