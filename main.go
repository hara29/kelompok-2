package main

import (
	"Kelompok2/controllers"
	"net/http" //package representasi web server di golang
)

func main() {
	//fungsi ini digunakan untuk routing, parameter pertama rootnya, parameter kedua handlernya 
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/logout", controllers.Logout)
	http.HandleFunc("/nilai-mahasiswa", controllers.Index) 
	http.HandleFunc("/nilai-mahasiswa/input-nilai", controllers.Input) 
	// http.HandleFunc("/nilai-mahasiswa/edit-nilai", controllers.Edit) 
	//membuat sekaligus menjalankan server di port 5000
	http.ListenAndServe(":5000", nil) 
}
