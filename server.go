package main

import (
	"CharterGateway/backend/handlers"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

/*
 *TODO:
 * 1. сделать фронт (форма поиска билетов, описание сервиса, страница текущих поисков)
 * 3. сделать авторизацию по логину и паролю
 * 4. парсер метапоисковиков
 * 5. канал телеграма и взаимодействие с ним
 * 6. отчет парсера о работе за день по телеграму админу
 */
func main() {
	// инициализируем сервер
	var router = mux.NewRouter()
	// список адресов
	router.HandleFunc("/", handlers.WebsiteHandler).Methods("GET")
	router.HandleFunc("/flights", handlers.FlightHandler).Methods("POST")
	/*router.HandleFunc("/flights", handler).Methods("PUT")
	router.HandleFunc("/flights", handler).Methods("DELETE")
	*/
	http.Handle("/", router)
	// middleware auth
	router.Use(authMiddleware)
	// middleware logging
	router.Use(loggingMiddleware)

	http.ListenAndServe(":9999", nil)
}

/*
 Логирование запросов
*/
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, err := os.OpenFile("./logs/server_logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		} else {
			postForm, _ := json.Marshal(r.PostForm)
			file.WriteString(time.Now().String() + ": получен запрос с " + r.Method + " " + r.RequestURI +
				" с телом запроса: " + string(postForm) + "\n")
			file.Close()
		}
		next.ServeHTTP(w, r)
	})
}

/*
 TODO Аутентификация пользователя
*/
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	})
}
