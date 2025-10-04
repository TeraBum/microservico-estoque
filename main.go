package main

import (
	"api-estoque/internal/config"
	"api-estoque/internal/controllers"
	"api-estoque/internal/repositories"
	"api-estoque/internal/router"
	"api-estoque/internal/services"
	"api-estoque/internal/utils"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title API Estoque
// @version 1.0
// @description Documentação API de estoque TeraBum
// @BasePath /api/v1/estoque
// @schemes http
func main() {
	// Setup
	logger := utils.SetupLogger()
	logger.Info("Iniciando api-estoque...")
	config.Load()

	// Repositories
	repos := repositories.InstanciateRepositories()

	// Services
	srvcs := services.InstanciateServices(repos, logger)

	// Controllers
	ctrls := controllers.InstanciateControllers(srvcs, logger)

	// Router
	router := router.New(logger, ctrls)
	router.Run()

	// Iniciar servidor http
	server := &http.Server{
		Addr:    ":8080",
		Handler: router.Router,
	}
	go func() {
		logger.Info("App iniciado. Servindo api-estoque na porta: 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Falha ao iniciar servidor na porta 8080: %v", err)
		}
	}()

	// Hook de shutdown:
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Esperar notificação de shutdown.
	<-quit
	logger.Info("Sinal de shutdown recebido. Iniciando shutdown gracioso...")

	// Contexto para timeout caso o shutdown trave
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Fechar servidor HTTP
	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Errorf("Erro durante shutdown do servidor HTTP: %v", err)
	} else {
		logger.Info("Servidor HTTP finalizado com sucesso.")
	}

	logger.Info("Shutdown completo. Saindo.")
}
