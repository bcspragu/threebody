package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/bcspragu/threebody/gif"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	if err := gif.GIF(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
