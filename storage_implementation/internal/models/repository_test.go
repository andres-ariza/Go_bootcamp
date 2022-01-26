package models

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	product := Product{
		Name: "test",
	}
	myRepo := NewRepo()
	productResult, err := myRepo.Store(product)
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, product.Name, productResult.Name)
}

func TestGetByName(t *testing.T) {
	productResult := NewRepo().GetByName("test")

	assert.Equal(t, "test", productResult.Name)
}

func TestGetOneWithContext(t *testing.T) {
	// usamos un Id que exista en la DB
	id := 4
	// definimos un Product cuyo nombre sea igual al registro de la DB
	product := Product{
		Name: "test",
	}
	myRepo := NewRepo()
	// se define un context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	productResult, err := myRepo.GetOneWithcontext(ctx, id)
	fmt.Println(err)
	assert.Equal(t, product.Name, productResult.Name)
}
