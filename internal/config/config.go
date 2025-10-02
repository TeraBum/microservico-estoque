package config

import (
	"api-estoque/internal/utils"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port               string `envconfig:"PORT"`
	SupabaseConnString string `envconfig:"SUPABASE_CONN_STRING"`
}

var Env Config
var logger = utils.SetupLogger()

func PostgresConn(maxConns int, maxIdleTime time.Duration, maxLifetime time.Duration) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(Env.SupabaseConnString)
	if err != nil {
		logger.Fatalf("Erro ao parsear configuracao do Postgres: %s", err.Error())
	}

	config.MaxConns = int32(maxConns)
	config.MaxConnIdleTime = maxIdleTime
	config.MaxConnLifetime = maxLifetime

	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		logger.Fatalf("Erro ao conectar no Postgres: %s", err.Error())
	}

	return dbpool
}

func Load() {
	if err := godotenv.Load(".env"); err != nil {
		logger.Info("Arquivo .env nao encontrado. Carregando variaveis do ambiente de execucao.")
	}

	if err := envconfig.Process("", &Env); err != nil {
		logger.Fatal("Erro ao processar variaveis de ambiente", err)
	}
}
