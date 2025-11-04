package main

import (
	"github.com/Mathias002/TP-fil-rouge-GO-efrei/internal/app"
	"github.com/Mathias002/TP-fil-rouge-GO-efrei/internal/store"
)

func main() {
	// init storer
	var store = storage.NewMemoryStore()

    app.Run(store)
}
