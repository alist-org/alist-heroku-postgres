package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strconv"
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
	SslMode     string `json:"ssl_mode" env:"DB_SLL_MODE"`
}

type Scheme struct {
	Https    bool   `json:"https" env:"HTTPS"`
	CertFile string `json:"cert_file" env:"CERT_FILE"`
	KeyFile  string `json:"key_file" env:"KEY_FILE"`
}

type CacheConfig struct {
	Expiration      int64 `json:"expiration" env:"CACHE_EXPIRATION"`
	CleanupInterval int64 `json:"cleanup_interval" env:"CLEANUP_INTERVAL"`
}

type Config struct {
	Force    bool        `json:"force"`
	Address  string      `json:"address" env:"ADDR"`
	Port     int         `json:"port" env:"PORT"`
	Assets   string      `json:"assets" env:"ASSETS"`
	Database Database    `json:"database"`
	Scheme   Scheme      `json:"scheme"`
	Cache    CacheConfig `json:"cache"`
	TempDir  string      `json:"temp_dir" env:"TEMP_DIR"`
}

func main() {
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
		Address: "0.0.0.0",
		TempDir: "data/temp",
		Database: Database{
			User:     user,
			Password: pass,
			Host:     host,
			Port:     port,
			Name:     name,
		},
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
