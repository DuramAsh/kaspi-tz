package main

import (
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
)

type Person struct {
	Name  string `json:"name"`
	IIN   string `json:"iin"`
	Phone string `json:"phone"`
}

func main() {
	root, err := os.Getwd()
	if err != nil {
		return
	}

	envFile := ".env"

	if _, err = os.Stat(envFile); os.IsNotExist(err) {
		log.Fatalf("Environment file %s does not exist", envFile)
	}

	if err = godotenv.Load(filepath.Join(root, envFile)); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		log.Fatal("API_URL is not set")
	}

	reqPath, err := url.Parse(apiURL)
	if err != nil {
		log.Fatalf("Error parsing API_URL: %v", err)
	}

	reqPath = reqPath.JoinPath("people/info")

	threadsStr := os.Getenv("THREADS")
	if threadsStr == "" {
		threadsStr = "10"
	}

	threads, err := strconv.Atoi(threadsStr)
	if err != nil || threads <= 0 {
		log.Fatalf("invalid THREADS: %v", err)
	}

	log.Printf("Starting load test with %d threads...\n", threads)

	wg := errgroup.Group{}

	for _ = range threads {
		wg.Go(func() (err error) {
			for {
				p := generateRandomPerson()
				if err = sendPerson(reqPath.String(), p); err != nil {
					log.Printf("Error sending person to db: %v", err)
				} else {
					log.Printf("Person sent successfully: %s", p.IIN)
				}

				time.Sleep(100 * time.Millisecond)
			}

			// return
		})
	}

	if err := wg.Wait(); err != nil {
		log.Fatalf("Error during load test: %v", err)
	}
}
