package server

// server is created for worker only.
import (
	"net/http"
	"time"

	"github.com/raghukul01/groot/src/api"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Server represents a server mux
type Server struct {
	*mux.Router
	Address string
}

func (s Server) SetupComponents() {
	apiMux := s.PathPrefix("/api").Subrouter()
	api.SetUp(apiMux)
}

// New setups & returns a server
func New() *Server {
	router := mux.NewRouter()
	hostName := viper.GetString("HOSTNAME")
	addr := hostName + ":" + viper.GetString("PORT")
	server := Server{router, addr}
	server.SetupComponents()
	return &server
}

func (server Server) ServeHTTP() {
	srv := &http.Server{
		Handler: server.Router,
		Addr:    server.Address,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: time.Minute,
		ReadTimeout:  time.Minute,
	}

	logrus.Info("Server starting at addr: ", server.Address)
	logrus.Fatal(srv.ListenAndServe())
}