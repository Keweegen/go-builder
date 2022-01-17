package services

import (
	"context"
	"errors"
	"go-builder/internal/builders"
)

type Company struct{}

func NewCompany() *Company {
	return &Company{}
}

func (c *Company) Create(ctx context.Context, companyType, name, webSite, email, bin string, active bool) (builders.CompanyBuilder, error) {
	builder, err := builders.GetCompanyBuilder(companyType)
	if err != nil {
		return nil, err
	}

	company := builder.
		SetName(name).
		SetWebSite(webSite).
		SetEmail(email).
		SetBIN(bin).
		SetActive(active).
		Build()

	// TODO: Insert into database

	return company, nil
}

func (c *Company) CreateBookmaker(ctx context.Context, name, webSite, email, bin string, active bool) (*builders.BookmakerCompany, error) {
	builder, err := c.Create(ctx, builders.CompanyTypeBookmaker, name, webSite, email, bin, active)
	if err != nil {
		return nil, err
	}

	company, ok := builder.(*builders.BookmakerCompany)
	if !ok {
		return nil, errors.New("failed convert company to bookmaker")
	}

	return company, nil
}

func (c *Company) CreatePaymentSystem(ctx context.Context, name, webSite, email, bin string, active bool) (*builders.PaymentSystemCompany, error) {
	builder, err := c.Create(ctx, builders.CompanyTypePaymentSystem, name, webSite, email, bin, active)
	if err != nil {
		return nil, err
	}

	company, ok := builder.(*builders.PaymentSystemCompany)
	if !ok {
		return nil, errors.New("failed convert company to payment system")
	}

	return company, nil
}
