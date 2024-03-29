package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/nidyaonur/stickverse-backend/internal/client"
	"github.com/nidyaonur/stickverse-backend/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func testRegisterUser(authClient *pb.AuthServiceClient) {
	user := pb.RegisterRequest{
		Email:    "test@test.com",
		Username: "test",
		Password: "test",
	}
	resp, err := (*authClient).Register(context.Background(), &user)
	if err != nil {
		log.Fatal("cannot register user: ", err)
	}
	log.Printf("register user response: %+v", resp)

}

/*
func testSearchLaptop(laptopClient *client.LaptopClient) {
	for i := 0; i < 10; i++ {
		laptopClient.CreateLaptop(sample.NewLaptop())
	}

	filter := &pb.Filter{
		MaxPriceUsd: 3000,
		MinCpuCores: 4,
		MinCpuGhz:   2.5,
		MinRam:      &pb.Memory{Value: 8, Unit: pb.Memory_GIGABYTE},
		D}

	laptopClient.SearchLaptop(filter)
}

func testUploadImage(laptopClient *client.LaptopClient) {
	laptop := sample.NewLaptop()
	laptopClient.CreateLaptop(laptop)
	laptopClient.UploadImage(laptop.GetId(), "tmp/laptop.jpg")
}

func testRateLaptop(laptopClient *client.LaptopClient) {
	n := 3
	laptopIDs := make([]string, n)

	for i := 0; i < n; i++ {
		laptop := sample.NewLaptop()
		laptopIDs[i] = laptop.GetId()
		laptopClient.CreateLaptop(laptop)
	}

	scores := make([]float64, n)
	for {
		fmt.Print("rate laptop (y/n)? ")
		var answer string
		fmt.Scan(&answer)

		if strings.ToLower(answer) != "y" {
			break
		}

		for i := 0; i < n; i++ {
			scores[i] = sample.RandomLaptopScore()
		}

		err := laptopClient.RateLaptop(laptopIDs, scores)
		if err != nil {
			log.Fatal(err)
		}
	}
}

*/
const (
	username        = "nidya2"
	email           = "nidya.onur2@gmail.com"
	password        = "secret"
	refreshDuration = 30 * time.Second
)

func authMethods() map[string]bool {
	const (
		userPath    = "/user.UserService/"
		maplistPath = "/maplist.MaplistService/"
	)

	return map[string]bool{
		userPath + "GetUserData":   true,
		maplistPath + "GetTable":   true,
		maplistPath + "GetLibrary": true,
	}
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	serverAddress := flag.String("address", "127.0.0.1:8080", "the server address")
	enableTLS := flag.Bool("tls", false, "enable SSL/TLS")
	flag.Parse()
	log.Printf("dial server %s, TLS = %t", *serverAddress, *enableTLS)

	transportOption := grpc.WithInsecure()

	if *enableTLS {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			log.Fatal("cannot load TLS credentials: ", err)
		}

		transportOption = grpc.WithTransportCredentials(tlsCredentials)
	}

	cc1, err := grpc.Dial(*serverAddress, transportOption)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	authClient := client.NewAuthClient(cc1, email, username, password)
	// regiter 300 users
	/*
			for i := 0; i < 300; i++ {

				randString := utils.RandString(6)
				authClient := client.NewAuthClient(cc1, email, username, password)
				accesToken, err := authClient.Register()
				if err != nil {
					log.Fatal("cannot register user: ", err)
				}
				log.Printf("register user response: %+v", accesToken)
			}
		authClient := client.NewAuthClient(cc1, email, username, password)
		accesToken, err := authClient.Register()

		if err != nil {
			log.Fatal("cannot register user: ", err)
		}
		log.Printf("register user response: %+v", accesToken)
	*/
	interceptor, err := client.NewAuthInterceptor(authClient, authMethods(), refreshDuration)
	if err != nil {
		log.Fatal("cannot create auth interceptor: ", err)
	}

	cc2, err := grpc.Dial(
		*serverAddress,
		transportOption,
		grpc.WithUnaryInterceptor(interceptor.Unary()),
		grpc.WithStreamInterceptor(interceptor.Stream()),
	)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	userClient := client.NewUserClient(cc2)
	userClient.GetUserData(1)
	maplistClient := client.NewMaplistClient(cc2)
	maplistClient.GetTable(172)
	maplistClient.GetLibrary(0, 0, 5, 5)
}
