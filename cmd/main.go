package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/robin238/fastcampus-golang-3-music-catalog/internal/configs"
	membershipsHandler "github.com/robin238/fastcampus-golang-3-music-catalog/internal/handler/memberships"
	"github.com/robin238/fastcampus-golang-3-music-catalog/internal/models/memberships"
	membershipsRepo "github.com/robin238/fastcampus-golang-3-music-catalog/internal/repository/memberships"
	membershipsSvc "github.com/robin238/fastcampus-golang-3-music-catalog/internal/service/memberships"
	"github.com/robin238/fastcampus-golang-3-music-catalog/pkg/internalsql"
)

func main() {
	var (
		cfg *configs.Config
	)

	err:=configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil{
		log.Fatal("gagal inisiasi config" , err)
	}

	cfg =configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err!=nil {
		log.Fatalf("failed to connect to database, err: %+v",err)
	}

	db.AutoMigrate(&memberships.User{})
	r:= gin.Default()

	membershipRepo := membershipsRepo.NewRepository(db)
	membershipSvc := membershipsSvc.NewService(cfg , membershipRepo)
	membershipHandler := membershipsHandler.NewHandler(r , membershipSvc)
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)

}