package auth

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/nidyaonur/stickverse-backend/internal/logic"
	"github.com/nidyaonur/stickverse-backend/pkg/env"
	"github.com/nidyaonur/stickverse-backend/pkg/pb"
	"github.com/nidyaonur/stickverse-backend/pkg/repository"
	"github.com/nidyaonur/stickverse-backend/pkg/utils"
)

func TestRegister(t *testing.T) {
	env := env.GetEnvironment()
	repo := repository.NewRepository(env.WriterDBUrl)
	defer repo.CloseDBConnection()

	jwtManager := NewJWTManager(env.JwtSecret, 30*24*time.Hour)
	logicService := logic.NewService(repo)
	server := NewAuthServer(repo, logicService, jwtManager)
	for i := 0; i < 2; i++ {
		randomString := utils.RandString(5)
		req := &pb.RegisterRequest{
			Username: "test" + randomString,
			Email:    "test" + randomString + "@test.com",
			Password: "test",
		}
		res, err := server.Register(context.Background(), req)
		if err != nil {
			t.Errorf("error: %v", err)
		}
		if res.AccessToken == "" {
			t.Errorf("access token is empty")
		}
		fmt.Println(res.AccessToken)
	}
}
