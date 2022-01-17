package main

import (
	"context"
	"go-builder/internal/services"
	"log"
)

func main() {
	company := services.NewCompany()

	b, err := company.CreateBookmaker(context.Background(), "Bookmaker", "https://bookmaker.com", "support@bookmaker.com", "000033321", true)
	if err != nil {
		log.Fatal(err)
	}

	ps, err := company.CreatePaymentSystem(context.Background(), "Visa", "https://visa.com", "support@visa.com", "674765", true)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Bookmaker", b)
	log.Println("Payment system", ps)
}
