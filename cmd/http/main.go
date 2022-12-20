package main

import (
	"github.com/gin-gonic/gin"
	"github.com/victoorraphael/poc-hex-shelf/internal/core/service/hellosrv"
	"github.com/victoorraphael/poc-hex-shelf/internal/handlers/hellohdl"
	"github.com/victoorraphael/poc-hex-shelf/internal/repositories/hellorepo"
)

func main() {
	repo := hellorepo.NewRepo()
	srv := hellosrv.New(repo)
	handler := hellohdl.New(srv)

	router := gin.New()
	router.GET("/", handler.Get)

	router.Run(":8080")
}
