package memory_test

import (
	core2 "ecommerce/api/core"
	memory2 "ecommerce/api/in-memory"
	"testing"
	"time"
)

//Checking interface implementation
var _ core2.Service = &memory2.InMemoryService{}

func TestInMemoryService_GetLastSeenProducts(t *testing.T) {
	//Arrange
	yesterday := time.Now().Add(-time.Hour * 24)
	today := time.Now()
	tomorrow := time.Now().Add(time.Hour * 24)

	y1 := &core2.Product{
		LastSeen: yesterday,
	}

	y2 := &core2.Product{
		LastSeen: yesterday,
	}

	y3 := &core2.Product{
		LastSeen: yesterday,
	}

	y4 := &core2.Product{
		LastSeen: yesterday,
	}

	t1 := &core2.Product{
		LastSeen: today,
	}

	t2 := &core2.Product{
		LastSeen: today,
	}

	tw1 := &core2.Product{
		LastSeen: tomorrow,
	}

	tw2 := &core2.Product{
		LastSeen: tomorrow,
	}

	products := []*core2.Product{
		y1,
		y2,
		y3,
		y4,
		t1,
		t2,
		tw1,
		tw2,
	}

	s := &memory2.InMemoryService{
		Products: products,
	}

	//Act
	res := s.GetLastSeenProducts()

	//Assert
	expected := []*core2.Product{
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
	var products []*core2.Product

	for i := 0; i < 20; i++ {
		ui := uint(i)
		products = append(products, &core2.Product{TimesSeen: &ui})
	}

	expectedItems := 8
	var expected []*core2.Product
	for i := len(products) - 1; i > len(products)-expectedItems-1; i-- {
		expected = append(expected, products[i])
	}

	s := &memory2.InMemoryService{Products: products}

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
