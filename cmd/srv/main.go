package main

import (
	"log"
	"my-todo/configs"
	"net/http"
	"time"
)

type timeHandler struct {
	format string
}

func (th timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + tm))
}

var (
	cfg *configs.Config
)

func loadConfig() error {
	Port := 5050

	cfg = &configs.Config{
		DBAddress:    "host=localhost port=5432 user=admin password=admin dbname=togo sslmode=disable",
		RedisAddress: "localhost:6379", //os.Getenv("REDIS_ADDRESS"),
		Port:         Port,
		JwtKey:       "TestJWT", //os.Getenv("JWT_KEY"),
	}
	return nil
}

func main() {

	if err := loadConfig(); err != nil {
		panic(err)
	}

	log.Println("config value is ", cfg)

	mux := http.NewServeMux()

	th := timeHandler{format: time.RFC1123}

	mux.Handle("/time", th)

	log.Print("Listening...")
	http.ListenAndServe(":3000", mux)
}
