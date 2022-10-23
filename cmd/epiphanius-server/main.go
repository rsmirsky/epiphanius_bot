package main

import (
	"epiphanius_bot/internal/telebot"
	"epiphanius_bot/pkg/controller"
	"epiphanius_bot/pkg/db"
	"epiphanius_bot/pkg/storage"

	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	//log.SetLevel(log.WarnLevel)
}

func main() {

	//viper
	log.Info("viper open .env")
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error read in config: ", err)
		return
	}
	

	//initial data base
	log.Info("start connect to database")
	DB := db.Connect(viper.GetString("DSN_POSTGRES"))



	//migration
	log.Info("run migrates")

	db.RunMigrates()

	

	// init telebot
	log.Info("database initialization")
	go storage.Init(DB)
	

	//telebot initial
	log.Info("connection to telebot")
	go telebot.Telebot()
	

    //create router
	log.Info("create router")
	router := mux.NewRouter()
	router.HandleFunc("/holidays", controller.GetHolidays).Methods(http.MethodGet)
	router.HandleFunc("/holidays/{id}", controller.GetHoliday).Methods(http.MethodGet)
	router.HandleFunc("/holidays", controller.CreateHoliday).Methods(http.MethodPost)
	router.HandleFunc("/holidays/{id}", controller.UpdateHoliday).Methods(http.MethodPut)
	router.HandleFunc("/holidays/{id}", controller.DeleteHoliday).Methods(http.MethodDelete)

	//create server for vue
	log.Info("start creating a server for vue")
	fs := http.FileServer(http.Dir("./dist"))
	router.PathPrefix("/").Handler(fs)
	

	//cors

	http.Handle("/", &corsRouterDecorator{router})

	log.Panicf("panic on server creation",
		http.ListenAndServe(":3000", nil),
	)

}

/***************************************************/

// CORSRouterDecorator applies CORS headers to a mux.Router
type corsRouterDecorator struct {
	R *mux.Router
}

func (c *corsRouterDecorator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, PATCH")
		rw.Header().Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}

	c.R.ServeHTTP(rw, req)

}
