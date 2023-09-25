package main

import (
	"encoding/json" // Import the JSON package for encoding and decoding JSON.
	"fmt"           // Import the fmt package for printing to the console.
	"log"           // Import the log package for logging errors.
	"math/rand"     // Import the math/rand package for generating random numbers.
	"net/http"      // Import the net/http package for handling HTTP requests.
	"strconv"       // Import the strconv package for string conversions.

	"github.com/gorilla/mux" // Import the Gorilla Mux router package for routing.
)

// Define a struct for representing a Movie with JSON tags.
// JSON tags specify the field names in the JSON representation.
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// Define a struct for representing a Director with JSON tags.
type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie // Declare a slice to hold Movie objects.

// Get list of all movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set the HTTP response content type to JSON.

	json.NewEncoder(w).Encode(movies) // Encode the 'movies' slice as JSON and write it to the response writer.
}

// Delete a movie
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set the HTTP response content type to JSON.

	params := mux.Vars(r) // Get the route parameters from the request.

	for index, item := range movies { // Iterate over the 'movies' slice.
		if item.ID == params["id"] { // Check if the Movie ID matches the route parameter.
			movies = append(movies[:index], movies[index+1:]...) // Remove the movie from the slice.
			break
		}
	}

	json.NewEncoder(w).Encode(movies) // Encode the updated 'movies' slice as JSON and write it to the response writer.
}

// Get a movie
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set the HTTP response content type to JSON.

	params := mux.Vars(r) // Get the route parameters from the request.

	for _, item := range movies { // Iterate over the 'movies' slice.
		if item.ID == params["id"] { // Check if the Movie ID matches the route parameter.
			json.NewEncoder(w).Encode(item) // Encode the Movie as JSON and write it to the response writer.
			return
		}
	}
}

// Create movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set the HTTP response content type to JSON.

	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie) // Decode the JSON request body into the 'movie' variable.

	movie.ID = strconv.Itoa(rand.Intn(10000000)) // Generate a random ID for the movie.

	movies = append(movies, movie) // Append the new movie to the 'movies' slice.

	json.NewEncoder(w).Encode(movie) // Encode the created movie as JSON and write it to the response writer.
}

// Update movie
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set the HTTP response content type to JSON.

	params := mux.Vars(r) // Get the route parameters from the request.

	for index, item := range movies { // Iterate over the 'movies' slice.
		if item.ID == params["id"] { // Check if the Movie ID matches the route parameter.
			movies = append(movies[:index], movies[index+1:]...) // Remove the movie from the slice.

			var movie Movie

			_ = json.NewDecoder(r.Body).Decode(&movie) // Decode the JSON request body into the 'movie' variable.

			movie.ID = params["id"] // Set the ID of the updated movie.

			movies = append(movies, movie) // Append the updated movie to the 'movies' slice.

			json.NewEncoder(w).Encode(movie) // Encode the updated movie as JSON and write it to the response writer.
			return
		}
	}
}

func main() {
	r := mux.NewRouter() // Create a new Gorilla Mux router instance.

	// Add two movies by default
	movies = append(movies, Movie{ID: "1", Isbn: "438277", Title: "Movie 1", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "438276", Title: "Movie 2", Director: &Director{FirstName: "Foo", LastName: "Bar"}})

	// Define routes and associate them with handler functions.
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting sever at port 8080\n") // Print a message indicating the server is starting.

	log.Fatal(http.ListenAndServe(":8080", r)) // Start the HTTP server and log any errors.
}
