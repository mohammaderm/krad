package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/mohammaderm/krad/config"
	producthandler "github.com/mohammaderm/krad/internal/presentation/http"
	"github.com/mohammaderm/krad/internal/presentation/http/authentication"
	productrepository "github.com/mohammaderm/krad/internal/repository/product"
	userrepository "github.com/mohammaderm/krad/internal/repository/user"
	productservice "github.com/mohammaderm/krad/internal/service/product"
	userservice "github.com/mohammaderm/krad/internal/service/user"
	"github.com/mohammaderm/krad/log"
)

func dbconnection(logger log.Logger, config *config.Database) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Mysql.Username,
		config.Mysql.Password,
		config.Mysql.Host,
		config.Mysql.Port,
		config.Mysql.Database,
	)
	con, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		logger.Panic(&log.Field{
			Package:  "main",
			Function: "dbconnection",
			Params:   "logger,config",
			Message:  "failed db connection.",
		})
		return nil, err
	}
	return con, nil

}

func main() {
	config, _ := config.NewConfig("./config/config.yaml")
	logger, err := log.New(&log.Logconfig{
		Path:         config.Logger.Internal_Path,
		Pattern:      config.Logger.Filename_Pattern,
		MaxAge:       config.Logger.Max_Age,
		RotationTime: config.Logger.Rotation_Time,
		RotationSize: config.Logger.Max_Size,
	})
	if err != nil {
		logger.Panic(&log.Field{
			Package:  "main",
			Function: "main",
			Params:   "_",
			Message:  err.Error(),
		})
	}
	db, err := dbconnection(logger, &config.Database)
	if err != nil {
		logger.Panic(&log.Field{
			Package:  "main",
			Function: "main",
			Params:   "_",
			Message:  err.Error(),
		})
	}
	// Product
	productRepository := productrepository.NewRepository(db, logger)
	productService := productservice.NewService(logger, productRepository)
	productHandler := producthandler.NewProductHanlder(logger, productService)

	// User Auth
	userRepository := userrepository.NewRepository(db, logger)
	userService := userservice.NewService(logger, userRepository)
	AuthHandler := authentication.NewAuthHanlder(logger, userService)

	r := mux.NewRouter()
	route := r.PathPrefix("/api/v1/").Subrouter()

	route.HandleFunc("/auth/login", AuthHandler.Login).Methods("Post")
	route.HandleFunc("/auth/register", AuthHandler.Register).Methods("Post")

	route.HandleFunc("/product/lastproduct", productHandler.GLTProduct).Methods("GET")
	route.HandleFunc("/product/{id}", productHandler.GetByID).Methods("GET")
	route.HandleFunc("/product/category={categoryid}/", productHandler.GetByCategoryId).Methods("GET")

	fmt.Printf("server is runing on port %s... \n", config.Server.Port)
	http.ListenAndServe(":8089", r)
}
