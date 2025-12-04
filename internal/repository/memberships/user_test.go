package memberships

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/robin238/fastcampus-golang-3-music-catalog/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_repository_CreateUser(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	
	gormDB,err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

	type args struct{
		model memberships.User
	}

	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		args args
		// db *gorm.DB
		// Named input parameters for target function.
		// model   memberships.User
		wantErr bool
		mockFn func(args args)
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				model:memberships.User{ 
					Email: "test@gmail.com",
					Username: "testusername",
					Password: "password" ,
					CreatedBy: "test@gmail.com",
					UpdatedBy: "test@gmail.com",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(`INSERT INTO "users" (.+) VALUES (.+)`).
				WithArgs(
					sqlmock.AnyArg(), // created_at
					sqlmock.AnyArg(), // updated_at
					sqlmock.AnyArg(), // deleted_at (NULL)
					args.model.Email,
					args.model.Username,
					args.model.Password,
					args.model.CreatedBy,
					args.model.UpdatedBy,
				).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectCommit()
			},
		
		},

		{
			name: "error",
			args: args{
				model:memberships.User{ 
					Email: "test@gmail.com",
					Username: "testusername",
					Password: "password" ,
					CreatedBy: "test@gmail.com",
					UpdatedBy: "test@gmail.com",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(`INSERT INTO "users" (.+) VALUES (.+)`).
				WithArgs(
					sqlmock.AnyArg(), // created_at
					sqlmock.AnyArg(), // updated_at
					sqlmock.AnyArg(), // deleted_at (NULL)
					args.model.Email,
					args.model.Username,
					args.model.Password,
					args.model.CreatedBy,
					args.model.UpdatedBy,
				).
				WillReturnError(assert.AnError)

				mock.ExpectRollback()
			},
		
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &repository{db:gormDB}
			
			if err := r.CreateUser(tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("repository.CreateUser() error = %v, wantErr %v", err,tt.wantErr)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func Test_repository_GetUser(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	
	gormDB,err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

	now := time.Now()
	type args struct{
		email string
		username string
		id int
	}

	tests := []struct {
		name string // description of this test case
		// Named input parameters for   constructor.
		// db *gorm.DB
		// Named input parameters for target function.
		args args
		want     *memberships.User
		wantErr  bool
		mockFn func(args args)
	}{

		{
			name: "success",
			args: args{ 
					email: "test@gmail.com",
					username: "testusername",
			},
			want: &memberships.User{
				Model : gorm.Model{
					ID: 1,
					CreatedAt: now,
					UpdatedAt: now,
				},
				Email:"test@gmail.com" ,
				Username: "testusername",
				Password: "password",
				CreatedBy: "test@gmail.com",
				UpdatedBy: "test@gmail.com",

			},
			wantErr: false,
			mockFn: func(args args) {
				// mock.ExpectBegin()

				mock.ExpectQuery(`SELECT \* FROM "users" .+`).
				WithArgs(
					args.email,
					args.username,
					args.id, 
					1).
				WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "email" ,
				"username", "password" , "created_by", "updated_by"}).AddRow(1 , now , now, "test@gmail.com", "testusername","password" , "test@gmail.com", "test@gmail.com"))

				// mock.ExpectCommit()
			},
		
		},
		{
			name: "failed",
			args: args{ 
					email: "test@gmail.com",
					username: "testusername",
			},
			want: nil,
			wantErr: true,
			mockFn: func(args args) {
				// mock.ExpectBegin()

				mock.ExpectQuery(`SELECT \* FROM "users" .+`).
				WithArgs(
					args.email,
					args.username,
					args.id, 
					1).
				WillReturnError(assert.AnError)
				// mock.ExpectCommit()
			},
		
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &repository{db:gormDB}
			got, gotErr := r.GetUser(tt.args.email, tt.args.username, tt.args.id)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetUser() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetUser() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			assert.Equal(t, tt.want, got)

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}


