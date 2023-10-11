package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strconv"
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
	Address      string `json:"address" env:"ADDR"`
	HttpPort     int    `json:"http_port" env:"HTTP_PORT"`
	HttpsPort    int    `json:"https_port" env:"HTTPS_PORT"`
	ForceHttps   bool   `json:"force_https" env:"FORCE_HTTPS"`
	CertFile     string `json:"cert_file" env:"CERT_FILE"`
	KeyFile      string `json:"key_file" env:"KEY_FILE"`
	UnixFile     string `json:"unix_file" env:"UNIX_FILE"`
	UnixFilePerm string `json:"unix_file_perm" env:"UNIX_FILE_PERM"`
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
	Force                 bool      `json:"force" env:"FORCE"`
	SiteURL               string    `json:"site_url" env:"SITE_URL"`
	Cdn                   string    `json:"cdn" env:"CDN"`
	JwtSecret             string    `json:"jwt_secret" env:"JWT_SECRET"`
	TokenExpiresIn        int       `json:"token_expires_in" env:"TOKEN_EXPIRES_IN"`
	Database              Database  `json:"database"`
	Scheme                Scheme    `json:"scheme"`
	TempDir               string    `json:"temp_dir" env:"TEMP_DIR"`
	BleveDir              string    `json:"bleve_dir" env:"BLEVE_DIR"`
	Log                   LogConfig `json:"log"`
	DelayedStart          int       `json:"delayed_start" env:"DELAYED_START"`
	MaxConnections        int       `json:"max_connections" env:"MAX_CONNECTIONS"`
	TlsInsecureSkipVerify bool      `json:"tls_insecure_skip_verify" env:"TLS_INSECURE_SKIP_VERIFY"`
}

func main() {
	configFilePath := "/opt/alist/data/config.json"
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
	var jwtSecret string
	if _, err := os.Stat(configFilePath); err == nil {
		data, err := ioutil.ReadFile(configFilePath)
		if err != nil {
			log.Fatalf("failed to read %s: %s", configFilePath, err.Error())
		}
		var existingConfig Config
		err = json.Unmarshal(data, &existingConfig)
		if err != nil {
			log.Fatalf("failed to unmarshal existing config: %s", err.Error())
		}
		jwtSecret = existingConfig.JwtSecret
	}
	if jwtSecret == "" {
		jwtSecret = random.String(16)
	}
	config := Config{
		Scheme: Scheme{
			Address:    "0.0.0.0",
			UnixFile:   "",
			HttpsPort:  -1,
			ForceHttps: false,
			CertFile:   "",
			KeyFile:    "",
		},
		JwtSecret:      jwtSecret,
		TokenExpiresIn: 48,
		TempDir:        "data/temp",
		Database: Database{
			User:        user,
			Password:    pass,
			Host:        host,
			Port:        port,
			Name:        name,
			TablePrefix: "x_",
			DBFile:      "data/data.db",
		},
		BleveDir: "data/bleve",
		Log: LogConfig{
			Enable:     true,
			Name:       "data/log/log.log",
			MaxSize:    50,
			MaxBackups: 30,
			MaxAge:     28,
		},
		MaxConnections:        0,
		TlsInsecureSkipVerify: true,
	}
	confBody, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatalf("failed marshal json: %s", err.Error())
	}
	err = ioutil.WriteFile(configFilePath, confBody, 0644)
	if err != nil {
		log.Fatalf("failed write json: %s", err.Error())
	}
}
