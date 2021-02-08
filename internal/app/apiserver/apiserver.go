package apiserver

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/AlastorTh/avTask/internal/app/store"
	"github.com/AlastorTh/avTask/model"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

//APIServer ...
type APIServer struct {
	config *Config
	logger *zap.Logger
	router *mux.Router
	store  *store.Store
}

func zapInit() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalln(err)
	}
	return logger
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: zapInit(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *APIServer) Start() error {
	s.logger.Info("starting api server")
	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	http.ListenAndServe(s.config.BindAddr, s.router)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.Use(loggingMiddleware(s.logger))

	api := s.router.PathPrefix("/api/v1").Subrouter()
	api.Path("/getAd").Methods("GET").HandlerFunc(s.handleGetAd())
	api.Path("/getAdList").Methods("GET").HandlerFunc(s.handleGetAdList())
	api.Path("/createAd").Methods("POST").HandlerFunc(s.handlePostAd())
}

func loggingMiddleware(log *zap.Logger) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(
			func(writer http.ResponseWriter, request *http.Request) {
				log.Info(
					request.Method,
					zap.String("url", request.URL.String()),
					zap.Any("query", request.URL.Query()),
				)
				handler.ServeHTTP(writer, request)
			},
		)
	}
}
func (s *APIServer) handleGetAd() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "PLACEHOLDER")
	}
}

func (s *APIServer) handleGetAdList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "sup")
	}
}

func (s *APIServer) handlePostAd() http.HandlerFunc {

	type request struct {
		Name     string   `json:"name"`
		Descript string   `json:"descript"`
		Price    float64  `json:"price"`
		Piclinks []string `json:"piclinks"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		ad := &model.Ad{
			Name:     req.Name,
			Descript: req.Descript,
			Price:    req.Price,
			PicLinks: req.Piclinks,
		}
		if _, err := s.store.Ad().Create(ad); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusCreated, ad)
	}
}

func (s *APIServer) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *APIServer) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}
