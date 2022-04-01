package main
import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	DatabaseUrl := os.Getenv("DATABASE_URL")
	fmt.Println("DatabaseUrl",DatabaseUrl)
	// DATABASE_URL = postgres://user3123:passkja83kd8@ec2-117-21-174-214.compute-1.amazonaws.com:6212/db982398
	u, err := url.Parse(DatabaseUrl)
	if err != nil {
		fmt.Println(err)
	}
	user := u.User.Username()
	pass, _ := u.User.Password()
	host := u.Hostname()
	port := u.Port()
	name := u.Path[1:]
	_ = os.Setenv("DB_HOST", host)
	_ = os.Setenv("DB_PORT", port)
	_ = os.Setenv("DB_USER", user)
	_ = os.Setenv("DB_PASS", pass)
	_ = os.Setenv("DB_NAME", name)
}