package services

import (
	"log"
	"moovio-v3/utils"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type MoovioHandler struct {
	svc *MoovioSvc
}

func NewHandler(svc *MoovioSvc) *MoovioHandler {
	return &MoovioHandler{
		svc: svc,
	}
}

func (m *MoovioHandler) getMovieList(w http.ResponseWriter, r *http.Request) {
	datas, err := m.svc.GetMovieList()
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
	}

	utils.WriteJSON(w, http.StatusOK, "", datas)
}

func (m *MoovioHandler) registerHandler() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/getmovielist", m.getMovieList)

	return router
}

func (m *MoovioHandler) Start() error {
	listenAddr := os.Getenv("API_URL")
	router := m.registerHandler()

	handler := cors.Default().Handler(router)
	server := new(http.Server)
	server.Handler = handler
	server.Addr = listenAddr

	log.Println("Moovio services running on", listenAddr)
	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
