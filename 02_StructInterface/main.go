package main

import (
	"fmt"
)

type Vehicle struct {
	Brand     string
	Type      string
	RentPrice float64
}

type Customer struct {
	Name           string
	Address        string
	RentedVehicles []Vehicle
}

type RentalSystem struct {
	Customers []Customer
	Vehicles  []Vehicle
}

func (rs *RentalSystem) AddCustomer(name, address string) {
	customer := Customer{Name: name, Address: address}
	rs.Customers = append(rs.Customers, customer)
}

func (rs *RentalSystem) AddVehicle(brand, typeVehicle string, rentPrice float64) {
	vehicle := Vehicle{Brand: brand, Type: typeVehicle, RentPrice: rentPrice}
	rs.Vehicles = append(rs.Vehicles, vehicle)
}

func (rs *RentalSystem) RentVehicle(customerName, brand, typeVehicle string) {
	for i, customer := range rs.Customers {
		if customer.Name == customerName {
			for j, vehicle := range rs.Vehicles {
				if vehicle.Brand == brand && vehicle.Type == typeVehicle {
					rs.Customers[i].RentedVehicles = append(rs.Customers[i].RentedVehicles, vehicle)

					rs.Vehicles = append(rs.Vehicles[:j], rs.Vehicles[j+1:]...)
					return
				}
			}
		}
	}
	fmt.Println("Vehicle not available or Customer not found")
}

func (rs *RentalSystem) ReturnVehicle(customerName, brand, typeVehicle string) {
	for i, customer := range rs.Customers {
		if customer.Name == customerName {
			for j, vehicle := range customer.RentedVehicles {
				if vehicle.Brand == brand && vehicle.Type == typeVehicle {

					rs.Vehicles = append(rs.Vehicles, vehicle)

					rs.Customers[i].RentedVehicles = append(rs.Customers[i].RentedVehicles[:j], rs.Customers[i].RentedVehicles[j+1:]...)
					return
				}
			}
		}
	}
	fmt.Println("Vehicle not found in customer's rented vehicles")
}

func (rs *RentalSystem) ListCustomers() {
	fmt.Println("Customer List:")
	for _, customer := range rs.Customers {
		fmt.Printf("Name: %s, Address: %s, Rented Vehicles: %d\n", customer.Name, customer.Address, len(customer.RentedVehicles))
	}
}

func (rs *RentalSystem) ListRentedVehicles() {
	fmt.Println("Rented Vehicles:")
	for _, customer := range rs.Customers {
		for _, vehicle := range customer.RentedVehicles {
			fmt.Printf("Customer: %s, Vehicle: %s %s, Rent Price: %.2f\n", customer.Name, vehicle.Brand, vehicle.Type, vehicle.RentPrice)
		}
	}
}

func main() {
	rentalSystem := &RentalSystem{}

	rentalSystem.AddCustomer("John Doe", "123 Elm St")
	rentalSystem.AddCustomer("Jane Smith", "456 Oak St")

	rentalSystem.AddVehicle("Toyota", "Camry", 50.0)
	rentalSystem.AddVehicle("Honda", "Civic", 45.0)

	rentalSystem.RentVehicle("John Doe", "Toyota", "Camry")

	rentalSystem.ListCustomers()

	rentalSystem.ListRentedVehicles()

	rentalSystem.ReturnVehicle("John Doe", "Toyota", "Camry")

	rentalSystem.ListCustomers()

	rentalSystem.ListRentedVehicles()
}
