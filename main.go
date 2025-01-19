package main

import (
	"core/config"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"database_host": cfg.Database.Host,
			"database_port": cfg.Database.Port,
			"server_port":   cfg.Server.Port,
		})
	})
	r.Run()

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
