package apiserver

import (
	"io"
	"net/http"

	"github.com/Ressley/hacknu/internal/app/apiserver/controllers"
	"github.com/Ressley/hacknu/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//APIserver ...
type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

//New ...
func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

//Start ...
func (s *APIserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("starting api server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIserver) configureRouter() {

	s.router.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	s.router.HandleFunc("/login", controllers.Login).Methods("POST")
}

func (s *APIserver) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}

func (s *APIserver) handleSignup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "henlo")
	}
}

/*
func (s *APIserver) handleSignin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := middleware.Authentication(w, r)
		if err != nil {
			return
		}
		io.WriteString(w, "auth")
	}
}*/
