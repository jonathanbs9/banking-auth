package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func Start(){
	// Sanity check for ENV var
	sanityCheck()
	router := mux.NewRouter()
	//authRepository := domain.NewAuthRepository(getDbClient())
	//ah := AuthHandler{ service.NewLoginService(authRepository, domain.GetRolePermiss)}

	router.HandleFunc("/auth/login", ah.Login).Methods(http.MethodPost)
	//router.HandleFunc("/auth/register", ah.NotImplementedHandler).Methods(http.MethodPost)
	router.HandleFunc("/auth/verify", ah.Verify).Methods(http.MethodGet)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Println(fmt.Sprintf("Starting OAuth server on #{address}:#{port} | ", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("#{address}:#{port}"),router))
}

func getDbClient() *sqlx.DB{
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client

}

// Recorro la lista para ver si alguna está vacía. Si al menos 1 es "", termina la app!
func sanityCheck(){
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASS",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}
	for _ , k := range envProps {
		if os.Getenv(k) == "" {
			log.Println(fmt.Sprintf("Variable de entorno %s no está definida. Terminando app!", k))
		}
	}
}