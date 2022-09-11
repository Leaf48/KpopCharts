package main

import (
	"log"

	"main/internal/pkg"
)

func main() {
	r := pkg.Billboard()
	log.Println(r[99].Group)
	log.Println(r[99].Title)
}
