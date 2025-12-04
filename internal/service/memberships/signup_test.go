package memberships

import (
	"testing"

	"github.com/robin238/fastcampus-golang-3-music-catalog/internal/configs"
	"github.com/robin238/fastcampus-golang-3-music-catalog/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_service_SignUp(t *testing.T) {

	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)

	type args struct {
		request memberships.SignUpRequest
	}

	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		// cfg        *configs.Config
		// repository repository
		args args
		// Named input parameters for target function.
		// request memberships.SignUpRequest
		wantErr bool
		mockFn func(args args)
	}{
		{
			name : "success",
			args: args{
				request: memberships.SignUpRequest{
					Email : "test@gmail.com", 
					Username : "testusername",
					Password : "password",
				},
			},
			wantErr: false,
			mockFn:func(args args){
				mockRepo.EXPECT().GetUser(args.request.Email , args.request.Username, uint(0)).Return(nil , gorm.ErrRecordNotFound)
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(nil)
			},
		},
		// TODO: Add test cases.

		{
			name : "failed when get user",
			args: args{
				request: memberships.SignUpRequest{
					Email : "test@gmail.com", 
					Username : "testusername",
					Password : "password",
				},
			},
			wantErr: true,
			mockFn:func(args args){
				mockRepo.EXPECT().GetUser(args.request.Email , args.request.Username, uint(0)).Return(nil , assert.AnError)
				// mockRepo.EXPECT().CreateUser(gomock.Any()).Return(nil)

			},
		},

		{
			name : "failed when create user",
			args: args{
				request: memberships.SignUpRequest{
					Email : "test@gmail.com", 
					Username : "testusername",
					Password : "password",
				},
			},
			wantErr: true,
			mockFn:func(args args){
				mockRepo.EXPECT().GetUser(args.request.Email , args.request.Username, uint(0)).Return(nil , gorm.ErrRecordNotFound)
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(assert.AnError)

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// >>> FIX HERE <<<
            if tt.mockFn != nil {
                tt.mockFn(tt.args)
            }

			s := &service{
				cfg: &configs.Config{},
			 	repository: mockRepo}
			gotErr := s.SignUp(tt.args.request)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("SignUp() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("SignUp() succeeded unexpectedly")
			}
		})
	}
}
