package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ant0ine/go-json-rest/rest"
	"fmt"
	"net/http"
	"log"
)


func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router,err :=rest.MakeRouter(
		rest.Post("/register",Register),
		rest.Get("/query/:id", QueryUserById),
		rest.Delete("/delete/:id", DeleteUser),
		rest.Put("/update",Update),
	)
	if err != nil {
		fmt.Println("err=>", err)
	}
	api.SetApp(router)
	http.Handle("/user/",http.StripPrefix("/user", api.MakeHandler()))
	log.Fatal(http.ListenAndServe(":3000", nil))
}
