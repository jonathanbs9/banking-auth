package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func Start(){
	// Sanity check for ENV var
	sanityCheck()
	router := mux.NewRouter()
	//authRepository := domain.NewAuthRepository(getDbClient())
	//ah := AuthHandler{ service.NewLoginService(authRepository, domain.GetRolePermiss)}

	router.HandleFunc("/auth/login", ah.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/register", ah.NotImplementedHandler).Methods(http.MethodPost)
	router.HandleFunc("/auth/verify", ah.Verify).Methods(http.MethodGet)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Println(fmt.Sprintf("Starting OAuth server on #{address}:#{port} | "))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("#{address}:#{port}"),router))
}

func getDbClient() *sqlx.DB{}

func sanityCheck(){}