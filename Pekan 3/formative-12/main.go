package main

import (
	"fmt"
	"formative-12/methods"
	"formative-12/middleware"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr: ":8080",
	}

	http.Handle("/post_grade", middleware.Auth(http.HandlerFunc(methods.PostGrade)))
	http.HandleFunc("/grades", methods.GetGrades)

	fmt.Println("server running at http://localhost:8080")
	server.ListenAndServe()
}
