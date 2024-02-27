package main

import (
	"fmt"
	"os"
)

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
	JWTSecret  string
}

var Envs = initConfig()

func initConfig() Config {
	DBAddress := "db:3306"
	fmt.Println(DBAddress)
	return Config{
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "mysql"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBAddress:  DBAddress,
		DBName:     getEnv("DB_NAME", "projectmanager"),
		JWTSecret:  getEnv("JWT_SECRET", "randomjwtsecretkey"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
