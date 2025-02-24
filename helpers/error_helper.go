package helpers

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

func PanicOnError(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func HandleRecordNotFound(domain interface{}, err error) (interface{}, error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain, errors.New("user not found")
	}
	return domain, err
}
