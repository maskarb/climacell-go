package main

import (
	"os"

	"github.com/maskarb/climacell-go"
)

func main() {

	c := climacell.NewClient(os.Getenv("CLIMACELL_API_KEY"))

}
