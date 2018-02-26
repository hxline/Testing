package main

import (
	repository "hxline/repository"

	_ "github.com/lib/pq"
)

func main() {
	repository.Insert()
}
