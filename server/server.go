package server

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

// Server ...
type Server struct {
	r       *mux.Router
	port    uint
	command string
	secret  string
}

// Config ...
type Config struct {
	Port    uint
	Command string
	Path    string
	Secret  string
}

// NewServer ...
func NewServer(config *Config) *Server {
	r := mux.NewRouter()
	s := &Server{
		r:       r,
		port:    config.Port,
		command: config.Command,
		secret:  config.Secret,
	}
	r.HandleFunc(config.Path, s.Handler)
	fmt.Printf("Registered path %s to run command %q\n", config.Path, config.Command)
	return s
}

// Start ...
func (s *Server) Start() {
	fmt.Printf("Listening on port %d\n", s.port)
	http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.r)
}

// Handler ...
func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	if s.secret != "" {
		secret := r.Header.Get("X-Hub-Signature")
		if s.secret != secret {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, http.StatusText(http.StatusUnauthorized))
			return
		}
	}

	out, err := exec.Command("bash", "-c", s.command).Output()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
