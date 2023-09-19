package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Apiserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *Apiserver {
	return &Apiserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *Apiserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info("starting api server")

	return nil
}

func (s *Apiserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.configureRouter()

	logrus.SetLevel(level)

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Apiserver) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *Apiserver) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
}