package main

import "testing"

func Test_calculate(t *testing.T) {
	tests := []struct {
		name string
		mass int
		fuel int
	}{
		{"Fuel: 12", 12, 2},
		{"Fuel: 14", 14, 2},
		{"Fuel: 1969", 1969, 654},
		{"Fuel: 100756", 100756, 33583},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFuel := calculateFuel(tt.mass); gotFuel != tt.fuel {
				t.Errorf("calculateFuel() = %v, want %v", gotFuel, tt.fuel)
			}
		})
	}
}
