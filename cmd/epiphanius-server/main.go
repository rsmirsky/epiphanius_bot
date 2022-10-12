package main

import (
	//"fmt"
	//"log"
	//"os"
	//"time"
	"epiphanius_bot/internal/telebot"
	"epiphanius_bot/pkg/controller"
	"epiphanius_bot/pkg/db"
	"epiphanius_bot/pkg/storage"
	//"fmt"

	//"fmt"

	//"epiphanius_bot/pkg/types"
	//"fmt"

	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	//"embed"
	//"io/fs"
	//"epiphanius_bot/pkg/logging"
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


	// Serve static files from the frontend/dist directory.
	// fs := http.FileServer(http.Dir("./epiphanius-frontend/dist"))
	// http.Handle("/", fs)

	// // Start the server.
	// fmt.Println("Server listening on port 9080    ")
	// log.Panic(
	// 	http.ListenAndServe(":9080", nil),
	// )

    //connect to frontend
	// var frontend embed.FS

	// stripped, err := fs.Sub(frontend, "epiphanius-frontend/dist")
    // if err != nil {
    //     log.Fatalln(err)
    // }

    // frontendFS := http.FileServer(http.FS(stripped))
	// fmt.Println("frontendFS=frontendFS=")
	// fmt.Println(frontendFS)
	// fmt.Println("frontendFS=frontendFS=")
    // http.Handle("/", frontendFS)

	// log.Info("frontend connected")


	//viper
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error read in config: ", err)
		return
	}

	log.Info("viper configured")

	//initial data base
	DB := db.Connect(viper.GetString("DSN_POSTGRES"))

	log.Info("database connection - successfully!")

	//migration
	db.RunMigrates()

	log.Info("migrations - successfully!")

	// init telebot
	go storage.Init(DB)

	
	//telebot initial
	go telebot.Telebot()
	log.Info("telebot - successfully!")


	router := mux.NewRouter()

	router.HandleFunc("/holidays", controller.GetHolidays).Methods(http.MethodGet)
	router.HandleFunc("/holidays/{id}", controller.GetHoliday).Methods(http.MethodGet)
	router.HandleFunc("/holidays", controller.CreateHoliday).Methods(http.MethodPost)
	router.HandleFunc("/holidays/{id}", controller.UpdateHoliday).Methods(http.MethodPut)
	router.HandleFunc("/holidays/{id}", controller.DeleteHoliday).Methods(http.MethodDelete)
	http.ListenAndServe(":9080",
		&CORSRouterDecorator{router})

	log.Info("API is running!")
	//http.ListenAndServe(":4000", router)

}

/***************************************************/

// CORSRouterDecorator applies CORS headers to a mux.Router
type CORSRouterDecorator struct {
	R *mux.Router
}

func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter,
	req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods",
			"POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Accept-Language,"+
				" Content-Type, YourOwnHeader")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}

	c.R.ServeHTTP(rw, req)
}
