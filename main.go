package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/result", func(w http.ResponseWriter, r *http.Request) {
		// Парсим HTML
		tmpl, err := template.ParseFiles("temp/profilPage.html")
		if err != nil {
			log.Fatalf("Error profilPage: %s", err)
		}
		// Собираем значение Query параметров
		name := r.URL.Query().Get("name")
		status := r.URL.Query().Get("status")

		requestURL := linkCollector(name, status) // создаем get запрос
		resBody := parseHTTP(requestURL)          // забираем тело ответа
		resStructArray := parsFatJson(resBody)    // парсим []byte в структуру

		// проверка на полноту ответа
		if len(resStructArray.Result) == 0 {
			fmt.Fprintf(w, "<b>Данные некоректны, попробуй еще раз</b><br><br><a href=\"http://127.0.0.1:8080\"><button>Еще раз</button></a>")
			return
		}

		tmpl.Execute(w, resStructArray) // передаем структуру в HTML шаблон; Развертывает HTML документ на сервер
	})

	//// Запуск сервера на порту 8080
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
