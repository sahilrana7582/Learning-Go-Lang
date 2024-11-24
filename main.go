package main

import (
	"fmt"
	"log"
	"net/http"
)

func hanleHome(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/home"{

		http.Error(w, "404 Abort Path is not Found!", http.StatusNotFound)
		return;
	}
	http.ServeFile(w, r,"./static/home.html")
	
}

func handleForm(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/form"{
		http.Error(w, "404 Path not Found !", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, "./static/form.html")

}


func handleAbout(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/about"{
		http.Error(w, "404! The Path not found", http.StatusNotFound)
	}
	

	http.ServeFile(w, r, "./static/about.html")
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	// Check if the URL path is correct
	if r.URL.Path != "/contact-me" {
		http.Error(w, "404! Path not Found", http.StatusNotFound)
		return // Stop further processing
	}

	// Only handle POST requests for form submission
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() Error: %v", err)
		return
	}

	// Get form values
	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")

	// Respond with the form data or a success message
	fmt.Fprintf(w, "Form submitted successfully!<br>")
	fmt.Fprintf(w, "Name: %v<br>", name)
	fmt.Fprintf(w, "Email: %v<br>", email)
	fmt.Fprintf(w, "Message: %v<br>", message)
}


func main(){

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	http.HandleFunc("/home", hanleHome)
	http.HandleFunc("/form", handleForm)
	http.HandleFunc("/about", handleAbout)
	http.HandleFunc("/submit", handlePost)




	PORT := ":8080"

	 err:=http.ListenAndServe(PORT, nil)
	 
	if err != nil{
		log.Fatal("Server is not able to run")
		return
	}


	fmt.Printf("Running Server On *8080")

}