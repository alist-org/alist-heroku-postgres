package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strconv"
	"path/filepath"

	"github.com/alist-org/alist/v3/cmd/flags"
	"github.com/alist-org/alist/v3/pkg/utils/random"
)

type Database struct {
	Type        string `json:"type" env:"DB_TYPE"`
	Host        string `json:"host" env:"DB_HOST"`
	Port        int    `json:"port" env:"DB_PORT"`
	User        string `json:"user" env:"DB_USER"`
	Password    string `json:"password" env:"DB_PASS"`
	Name        string `json:"name" env:"DB_NAME"`
	DBFile      string `json:"db_file" env:"DB_FILE"`
	TablePrefix string `json:"table_prefix" env:"DB_TABLE_PREFIX"`
	SSLMode     string `json:"ssl_mode" env:"DB_SSL_MODE"`
}

type Scheme struct {
	Https    bool   `json:"https" env:"HTTPS"`
	CertFile string `json:"cert_file" env:"CERT_FILE"`
	KeyFile  string `json:"key_file" env:"KEY_FILE"`
}

type LogConfig struct {
	Enable     bool   `json:"enable" env:"LOG_ENABLE"`
	Name       string `json:"name" env:"LOG_NAME"`
	MaxSize    int    `json:"max_size" env:"MAX_SIZE"`
	MaxBackups int    `json:"max_backups" env:"MAX_BACKUPS"`
	MaxAge     int    `json:"max_age" env:"MAX_AGE"`
	Compress   bool   `json:"compress" env:"COMPRESS"`
}

type Config struct {
	Force          bool      `json:"force" env:"FORCE"`
	Address        string    `json:"address" env:"ADDR"`
	Port           int       `json:"port" env:"PORT"`
	SiteURL        string    `json:"site_url" env:"SITE_URL"`
	Cdn            string    `json:"cdn" env:"CDN"`
	JwtSecret      string    `json:"jwt_secret" env:"JWT_SECRET"`
	TokenExpiresIn int       `json:"token_expires_in" env:"TOKEN_EXPIRES_IN"`
	Database       Database  `json:"database"`
	Scheme         Scheme    `json:"scheme"`
	TempDir        string    `json:"temp_dir" env:"TEMP_DIR"`
	BleveDir       string    `json:"bleve_dir" env:"BLEVE_DIR"`
	Log            LogConfig `json:"log"`
	MaxConnections int       `json:"max_connections" env:"MAX_CONNECTIONS"`
}

func main() {
	tempDir := filepath.Join(flags.DataDir, "temp")
	indexDir := filepath.Join(flags.DataDir, "bleve")
	logPath := filepath.Join(flags.DataDir, "log/log.log")
	dbPath := filepath.Join(flags.DataDir, "data.db")
	DATABASE_URL := os.Getenv("DATABASE_URL")
	fmt.Println("DatabaseUrl", DATABASE_URL)
	//DATABASE_URL = "postgres://hfhgpvbymdzusj:39d7f6f3ee4288103e382d5dec22ce668c4e5cb65120f64d574b808775674eb4@ec2-3-218-171-44.compute-1.amazonaws.com:5432/d4o07n33pf6ot7"
	u, err := url.Parse(DATABASE_URL)
	if err != nil {
		fmt.Println(err)
	}
	user := u.User.Username()
	pass, _ := u.User.Password()
	host := u.Hostname()
	port, _ := strconv.Atoi(u.Port())
	name := u.Path[1:]
	config := Config{
		Address:        "0.0.0.0",
		JwtSecret:      random.String(16),
		TokenExpiresIn: 48,
		TempDir:        tempDir,
		Database: Database{
			User:        user,
			Password:    pass,
			Host:        host,
			Port:        port,
			Name:        name,
			TablePrefix: "x_",
			DBFile:      dbPath,
		},
		BleveDir: indexDir,
		Log: LogConfig{
			Enable:     true,
			Name:       logPath,
			MaxSize:    10,
			MaxBackups: 5,
			MaxAge:     28,
		},
		MaxConnections: 0,
	}
	confBody, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatalf("failed marshal json: %s", err.Error())
	}
	err = ioutil.WriteFile("/opt/alist/data/config.json", confBody, 0777)
	if err != nil {
		log.Fatalf("failed write json: %s", err.Error())
	}
}
