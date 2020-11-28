package main

import (
	"log"
	"runtime"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/purwandi/platform/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cmd.Execute()
}
