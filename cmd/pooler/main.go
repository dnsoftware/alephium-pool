// Обработка заданий и полученных решений блоков (тут нужна максимальная скорость обмена между нодой монеты и майнерами)
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Docker 3 !")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8090")
	http.ListenAndServe(":8090", nil)

	//app.Run()
}
