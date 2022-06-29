package main

import (
	"database/sql"
	"fmt"
	"log"
	"my-todo/configs"
	"my-todo/internal/domain"
	"my-todo/internal/model"
	"my-todo/internal/model/postgres"

	"net/http"
	"time"

	_ "github.com/bmizerany/pq"
)

type timeHandler struct {
	format string
}

func (th timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + tm))
}

var (
	cfg      *configs.Config
	dbClient *sql.DB

	// store
	taskStore      model.TaskStore
	taskCountStore model.TaskCountStore
	userStore      model.UserStore

	// domain
	taskDomain domain.TaskDomain
	authDomain domain.AuthDomain
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

func loadDatabase() error {
	var err error
	dbClient, err = postgres.NewPostgresClient(cfg.DBAddress)
	if err == nil {
		fmt.Println("connect database successful", cfg.DBAddress)
	}
	return err
}

func loadStores() {
	log.Println("Loading store function is called")
	taskStore = postgres.NewTaskStore(dbClient)
	//taskCountStore = rd.NewTaskCountStore(redisClient)
	userStore = postgres.NewUserStore(dbClient)
}

func loadDomain() {
	taskDomain = domain.NewTaskDomain(taskCountStore, taskStore, userStore)
	authDomain = domain.NewAuthDomain(userStore, cfg.JwtKey)
}

func main() {

	if err := loadConfig(); err != nil {
		panic(err)
	}

	if err := loadDatabase(); err != nil {
		panic(err)
	}

	loadStores()
	loadDomain()

	log.Println("config value is ", cfg)

	mux := http.NewServeMux()

	th := timeHandler{format: time.RFC1123}

	mux.Handle("/time", th)

	log.Print("Listening...")
	http.ListenAndServe(":3000", mux)
}
