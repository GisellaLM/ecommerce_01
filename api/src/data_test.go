package src_test

import (
	"ecommerce/api/src"
	"testing"
	"time"
)

//Checking interface implementation
var _ src.Service = &src.InMemoryService{}

func TestInMemoryService_GetLastSeenProducts(t *testing.T) {
	//Arrange
	yesterday := time.Now().Add(-time.Hour * 24)
	today := time.Now()
	tomorrow := time.Now().Add(time.Hour * 24)

	y1 := &src.Product{
		LastSeen: yesterday,
	}

	y2 := &src.Product{
		LastSeen: yesterday,
	}

	y3 := &src.Product{
		LastSeen: yesterday,
	}

	y4 := &src.Product{
		LastSeen: yesterday,
	}

	t1 := &src.Product{
		LastSeen: today,
	}

	t2 := &src.Product{
		LastSeen: today,
	}

	tw1 := &src.Product{
		LastSeen: tomorrow,
	}

	tw2 := &src.Product{
		LastSeen: tomorrow,
	}

	products := []*src.Product{
		y1,
		y2,
		y3,
		y4,
		t1,
		t2,
		tw1,
		tw2,
	}

	s := &src.InMemoryService{
		Products: products,
	}

	//Act
	res := s.GetLastSeenProducts()

	//Assert
	expected := []*src.Product{
		tw2,
		tw1,
		t2,
		t1,
		y4,
		y3,
	}

	if len(res) != len(expected) {
		t.Errorf("Expected to receive %v items", len(expected))
	}

	for i := 0; i < len(expected); i++ {
		if !res[i].LastSeen.Equal(expected[i].LastSeen) {
			t.Errorf("expected %v at position %v", expected[i].LastSeen, i)
		}
	}
}

func TestInMemoryService_GetTrendingProducts(t *testing.T) {
	//Arrange
	var products []*src.Product

	for i := 0; i < 20; i++ {
		ui := uint(i)
		products = append(products, &src.Product{TimesSeen: &ui})
	}

	expectedItems := 8
	var expected []*src.Product
	for i := len(products) - 1; i > len(products)-expectedItems-1; i-- {
		expected = append(expected, products[i])
	}

	s := &src.InMemoryService{Products: products}

	//Act
	res := s.GetTrendingProducts()

	//Assert
	if len(res) != expectedItems {
		t.Errorf("shoud return %v items", expectedItems)
	}

	//comparing result and expected
	for i := 0; i < len(res); i++ {
		if expected[i].TimesSeen != res[i].TimesSeen {
			t.Errorf("expected %v at position %v", *expected[i], i)
		}
	}
}

func TestInMemoryService_GetFilterCategories(t *testing.T) {

}

func TestInMemoryService_GetPopularCategories(t *testing.T) {

}
