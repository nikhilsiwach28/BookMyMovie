package config

import "os"

type Config struct{
	port  string 
	host  string
}

func GetConfigFromEnv() *Config{
	return &Config{
		port : getValueFromEnv("port","8080"),
		host : getValueFromEnv("host","localhost"),
		
	}
}

func getValueFromEnv(key ,defaultValue string) string{
	if value := os.Getenv(key); value == ""{
		return defaultValue; 
	}else{return key}; 
}