package main

import (
	"log"
	"src/mock-unit-test/repository"
	service2 "src/mock-unit-test/service"
)

func main() {
	service := service2.CategoryService{Repository: repository.CategoryRepositoryReal{}}
	category, errs := service.Get("1")
	if errs != nil {
		log.Printf("error -> %s", errs.Error())
	} else {
		log.Printf("data -> %v", category)
	}
}
