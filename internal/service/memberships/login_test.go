package memberships

import (
	"fmt"
	"testing"

	"github.com/robin238/fastcampus-golang-3-music-catalog/internal/configs"
	"github.com/robin238/fastcampus-golang-3-music-catalog/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_service_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)

	type args struct {
		request memberships.LoginRequest
	}
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		args args
		// Named input parameters for target function.
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				request: memberships.LoginRequest{
					Email:    "test@example.com",
					Password: "password",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(&memberships.User{
					Model:    gorm.Model{ID: 1},
					Email:    "test@example.com",
					Password: "$2a$10$MeEVJduCLQJ5ffjANcSS8efvD6FadUtXfDh2B6QJPGXHHE6wpIOF6",
					Username: "robin",
				}, nil)
			},
		},
		{
			name: "failed - when get user error",
			args: args{
				request: memberships.LoginRequest{
					Email:    "test@example.com",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				cfg: &configs.Config{
					Service: configs.Service{SecretJWT: "abc"},
				},
				repository: mockRepo,
			}

			if tt.mockFn != nil {
				tt.mockFn(tt.args)
			}

			got, gotErr := s.Login(tt.args.request)

			if (gotErr != nil) != tt.wantErr {
				t.Errorf("Login() failed: %v", gotErr)
				return
			}

			if !tt.wantErr {
				fmt.Printf("test case: %s, got: %s\n", tt.name, got)
				assert.NotEmpty(t, got)
			} else {
				fmt.Printf("test case: %s, got: %s\n", tt.name, got)
				assert.Empty(t, got)
			}
		})
	}
}
