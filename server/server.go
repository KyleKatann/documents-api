package server

import (
	"log"
	"net/http"

	userHandler "github.com/nepp-tumsat/documents-api/server/handler/user"
)

func Serve(addr string) {
	http.HandleFunc("/users", get(userHandler.HandleUserList()))

	log.Println("Server running...")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Printf("Listen and serve failed. %+v", err)
	}
}

func get(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodGet)
}

func post(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodPost)
}

func httpMethod(apiFunc http.HandlerFunc, method string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		writer.Header().Set("Access-Control-Allow-Origin", "*")

		if request.Method == http.MethodOptions {
			headers := request.Header.Get("Access-Control-Request-Headers")
			writer.Header().Set("Access-Control-Allow-Headers", headers)
			return
		}

		// 指定のHTTPメソッドでない場合はエラー
		if request.Method != method {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte("Method Not Allowed"))
			return
		}

		// 共通のレスポンスヘッダを設定
		writer.Header().Set("Content-Type", "interfaces/json")
		apiFunc(writer, request)
	}
}
