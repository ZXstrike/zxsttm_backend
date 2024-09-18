package config

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServePort  string
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
	MySQL      MySQLConfig
}

type MySQLConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

var Config *AppConfig

func LoadConfig() (*AppConfig, error) {
	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			return nil, fmt.Errorf("error loading .env file: %w", err)
		}
	}

	privateKey, err := loadECDSAPrivateKey(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return nil, fmt.Errorf("error loading private key: %w", err)
	}

	publicKey, err := loadECDSAPublicKey(os.Getenv("PUBLIC_KEY"))
	if err != nil {
		return nil, fmt.Errorf("error loading public key: %w", err)
	}

	Config = &AppConfig{
		ServePort:  os.Getenv("SERVE_PORT"),
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		MySQL: MySQLConfig{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Database: os.Getenv("DB_DATABASE"),
		},
	}

	return Config, nil
}

// Helper function to load an ECDSA private key from a PEM encoded string
func loadECDSAPrivateKey(pemEncodedKey string) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemEncodedKey))
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing the key")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse EC private key: %v", err)
	}

	return privateKey, nil
}

// Helper function to load an ECDSA public key from a PEM encoded string
func loadECDSAPublicKey(pemEncodedKey string) (*ecdsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pemEncodedKey))
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing the key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse EC public key: %v", err)
	}

	return publicKey.(*ecdsa.PublicKey), nil
}
