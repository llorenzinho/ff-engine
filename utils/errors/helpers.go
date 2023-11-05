package helpers

import "log"

func Must[T any](r T, e error) T {
	if e != nil {
		log.Fatalf("Error: %s", e.Error())
	}
	return r
}

func Should[T any](r T, e error) T {
	if e != nil {
		log.Printf("Error: %s", e.Error())
	}
	return r
}
