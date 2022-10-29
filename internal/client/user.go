package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nidyaonur/stickverse-backend/pkg/pb"
	"google.golang.org/grpc"
)

// UserClient is a client to call user service RPCs
type UserClient struct {
	service pb.UserServiceClient
}

// NewUserClient returns a new user client
func NewUserClient(cc *grpc.ClientConn) *UserClient {
	service := pb.NewUserServiceClient(cc)
	return &UserClient{service}
}

// CreateUser calls create user RPC
/*
func (userClient *UserClient) CreateUser(user *pb.User) {
	req := &pb.CreateUserRequest{
		User: user,
	}

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := userClient.service.CreateUser(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			// not a big deal
			log.Print("user already exists")
		} else {
			log.Fatal("cannot create user: ", err)
		}
		return
	}

	log.Printf("created user with id: %s", res.Id)
}
*/
// SearchUser calls search user RPC
func (userClient *UserClient) GetUserData(userID uint64) (*pb.UserDataResponse, error) {
	log.Print("user id", userID)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.UserDataRequest{}
	userData, err := userClient.service.GetUserData(ctx, req)
	if err != nil {
		log.Fatal("cannot search user: ", err)
	}

	fmt.Println("user data", userData.GetEmail())
	return userData, nil

}

/*
// UploadImage calls upload image RPC
func (userClient *UserClient) UploadImage(userID string, imagePath string) {
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := userClient.service.UploadImage(ctx)
	if err != nil {
		log.Fatal("cannot upload image: ", err)
	}

	req := &pb.UploadImageRequest{
		Data: &pb.UploadImageRequest_Info{
			Info: &pb.ImageInfo{
				UserId:    userID,
				ImageType: filepath.Ext(imagePath),
			},
		},
	}

	err = stream.Send(req)
	if err != nil {
		log.Fatal("cannot send image info to server: ", err, stream.RecvMsg(nil))
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot read chunk to buffer: ", err)
		}

		req := &pb.UploadImageRequest{
			Data: &pb.UploadImageRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			log.Fatal("cannot send chunk to server: ", err, stream.RecvMsg(nil))
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("cannot receive response: ", err)
	}

	log.Printf("image uploaded with id: %s, size: %d", res.GetId(), res.GetSize())
}

// RateUser calls rate user RPC
func (userClient *UserClient) RateUser(userIDs []string, scores []float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := userClient.service.RateUser(ctx)
	if err != nil {
		return fmt.Errorf("cannot rate user: %v", err)
	}

	waitResponse := make(chan error)
	// go routine to receive responses
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Print("no more responses")
				waitResponse <- nil
				return
			}
			if err != nil {
				waitResponse <- fmt.Errorf("cannot receive stream response: %v", err)
				return
			}

			log.Print("received response: ", res)
		}
	}()

	// send requests
	for i, userID := range userIDs {
		req := &pb.RateUserRequest{
			UserId: userID,
			Score:  scores[i],
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send stream request: %v - %v", err, stream.RecvMsg(nil))
		}

		log.Print("sent request: ", req)
	}

	err = stream.CloseSend()
	if err != nil {
		return fmt.Errorf("cannot close send: %v", err)
	}

	err = <-waitResponse
	return err
}

*/
