package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type Car struct {
	ID          int64  `json:"id"`
	Brand       string `json:"brand"`
	Model       string `json:"model"`
	Mileage     int    `json:"mileage"`
	OwnersCount int    `json:"owners_count"`
}

type Furniture struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	Manufacturer string  `json:"manufacturer"`
	Height       float64 `json:"height"`
	Width        float64 `json:"width"`
	Length       float64 `json:"length"`
}

type Flower struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	ArrivalDate string  `json:"arrival_date"`
}

type DataStore struct {
	Cars      []Car       `json:"cars"`
	Furniture []Furniture `json:"furniture"`
	Flowers   []Flower    `json:"flowers"`
}

var (
	store DataStore
	mu    sync.Mutex
)

const dataFile = "data.json"

// Load data from JSON file
func loadData() {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		log.Fatalf("Error opening data file: %v", err)
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading data file: %v", err)
	}

	if err := json.Unmarshal(byteValue, &store); err != nil {
		log.Fatalf("Error unmarshaling data: %v", err)
	}
}

// Save data to JSON file
func saveData() {
	mu.Lock()
	defer mu.Unlock()

	file, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		log.Printf("Error marshaling data: %v", err)
		return
	}

	if err := ioutil.WriteFile(dataFile, file, 0644); err != nil {
		log.Printf("Error writing to data file: %v", err)
	}
}

// Create new car
func createCar(w http.ResponseWriter, r *http.Request) {
	var car Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	mu.Lock()
	car.ID = int64(len(store.Cars) + 1)
	store.Cars = append(store.Cars, car)
	mu.Unlock()
	saveData()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}

// Get list of cars
func getCars(w http.ResponseWriter, r *http.Request) {
	log.Println("Received GET request to /cars")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(store.Cars)
}

// Get a single car by ID
func getCar(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/cars/"):] // Extract ID from URL
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for _, car := range store.Cars {
		if car.ID == id {
			json.NewEncoder(w).Encode(car)
			return
		}
	}

	http.Error(w, "Car not found", http.StatusNotFound)
}

// Update a car by ID
func updateCar(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/cars/"):] // Extract ID from URL
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedCar Car
	if err := json.NewDecoder(r.Body).Decode(&updatedCar); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, car := range store.Cars {
		if car.ID == id {
			updatedCar.ID = id
			store.Cars[i] = updatedCar
			saveData()
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updatedCar)
			return
		}
	}

	http.Error(w, "Car not found", http.StatusNotFound)
}

// Delete a car by ID
func deleteCar(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/cars/"):] // Extract ID from URL
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, car := range store.Cars {
		if car.ID == id {
			store.Cars = append(store.Cars[:i], store.Cars[i+1:]...)
			saveData()
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Car not found", http.StatusNotFound)
}

// Create new furniture
func createFurniture(w http.ResponseWriter, r *http.Request) {
	var furniture Furniture
	if err := json.NewDecoder(r.Body).Decode(&furniture); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	mu.Lock()
	furniture.ID = int64(len(store.Furniture) + 1)
	store.Furniture = append(store.Furniture, furniture)
	mu.Unlock()
	saveData()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(furniture)
}

// Get list of furniture
func getFurniture(w http.ResponseWriter, r *http.Request) {
	log.Println("Received GET request to /furniture")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(store.Furniture)
}

// Get a single furniture by ID
func getFurnitureByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/furniture/"):]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for _, furniture := range store.Furniture {
		if furniture.ID == id {
			json.NewEncoder(w).Encode(furniture)
			return
		}
	}

	http.Error(w, "Furniture not found", http.StatusNotFound)
}

// Update furniture by ID
func updateFurniture(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/furniture/"):]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedFurniture Furniture
	if err := json.NewDecoder(r.Body).Decode(&updatedFurniture); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, furniture := range store.Furniture {
		if furniture.ID == id {
			updatedFurniture.ID = id
			store.Furniture[i] = updatedFurniture
			saveData()
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updatedFurniture)
			return
		}
	}

	http.Error(w, "Furniture not found", http.StatusNotFound)
}

// Delete furniture by ID
func deleteFurniture(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/furniture/"):]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, furniture := range store.Furniture {
		if furniture.ID == id {
			store.Furniture = append(store.Furniture[:i], store.Furniture[i+1:]...)
			saveData()
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Furniture not found", http.StatusNotFound)
}

// Create new flower
func createFlower(w http.ResponseWriter, r *http.Request) {
	var flower Flower
	if err := json.NewDecoder(r.Body).Decode(&flower); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	mu.Lock()
	flower.ID = int64(len(store.Flowers) + 1)
	store.Flowers = append(store.Flowers, flower)
	mu.Unlock()
	saveData()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(flower)
}

// Get list of flowers
func getFlowers(w http.ResponseWriter, r *http.Request) {
	log.Println("Received GET request to /flowers")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(store.Flowers)
}

// Get a single flower by ID
func getFlowerByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/flowers/"):]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for _, flower := range store.Flowers {
		if flower.ID == id {
			json.NewEncoder(w).Encode(flower)
			return
		}
	}

	http.Error(w, "Flower not found", http.StatusNotFound)
}

func main() {
	loadData()

	// Car routes
	http.HandleFunc("/cars", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getCars(w, r)
		case http.MethodPost:
			createCar(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/cars/", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[len("/cars/"):]

		// Check if ID is provided in the URL
		if idStr == "" {
			switch r.Method {
			case http.MethodGet:
				getCars(w, r)
			case http.MethodPost:
				createCar(w, r)
			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
			return
		}

		switch r.Method {
		case http.MethodGet:
			getCar(w, r)
		case http.MethodPut:
			updateCar(w, r)
		case http.MethodDelete:
			deleteCar(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// Furniture routes
	http.HandleFunc("/furniture", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getFurniture(w, r)
		case http.MethodPost:
			createFurniture(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/furniture/", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[len("/furniture/"):]

		// Check if ID is provided in the URL
		if idStr == "" {
			switch r.Method {
			case http.MethodGet:
				getFurniture(w, r)
			case http.MethodPost:
				createFurniture(w, r)
			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
			return
		}

		switch r.Method {
		case http.MethodGet:
			getFurnitureByID(w, r)
		case http.MethodPut:
			updateFurniture(w, r)
		case http.MethodDelete:
			deleteFurniture(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// Flower routes
	http.HandleFunc("/flowers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getFlowers(w, r)
		case http.MethodPost:
			createFlower(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/flowers/", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[len("/flowers/"):]

		// Check if ID is provided in the URL
		if idStr == "" {
			switch r.Method {
			case http.MethodGet:
				getFlowers(w, r)
			case http.MethodPost:
				createFlower(w, r)
			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
			return
		}

		switch r.Method {
		case http.MethodGet:
			getFlowerByID(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server started on :9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
