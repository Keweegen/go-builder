package builders

import (
	"errors"
	"go-builder/internal/tools"
)

const (
	CompanyTypeBookmaker     = "bookmaker"
	CompanyTypePaymentSystem = "payment_system"
)

var ErrCompanyTypeNotFound = errors.New("company: type not found")

type CompanyBuilder interface {
	SetName(name string) CompanyBuilder
	SetActive(value bool) CompanyBuilder
	SetWebSite(webSite string) CompanyBuilder
	SetEmail(email string) CompanyBuilder
	SetBIN(bin string) CompanyBuilder
	Build() CompanyBuilder
}

type Company struct {
	Name    string
	WebSite string
	Email   string
	BIN     string
	Active  bool
}

var companies = map[string]CompanyBuilder{
	CompanyTypeBookmaker:     &BookmakerCompany{},
	CompanyTypePaymentSystem: &PaymentSystemCompany{},
}

func GetCompanyBuilder(companyType string) (CompanyBuilder, error) {
	company, exists := companies[companyType]
	if !exists {
		return nil, ErrCompanyTypeNotFound
	}
	return company, nil
}

//-

type BookmakerCompany struct {
	Company
	Token string
}

func (bc *BookmakerCompany) SetName(name string) CompanyBuilder {
	bc.Name = name
	return bc
}

func (bc *BookmakerCompany) SetWebSite(webSite string) CompanyBuilder {
	bc.WebSite = webSite
	return bc
}

func (bc *BookmakerCompany) SetEmail(email string) CompanyBuilder {
	bc.Email = email
	return bc
}

func (bc *BookmakerCompany) SetBIN(bin string) CompanyBuilder {
	bc.BIN = bin
	return bc
}

func (bc *BookmakerCompany) SetActive(value bool) CompanyBuilder {
	bc.Active = value
	return bc
}

func (bc *BookmakerCompany) Build() CompanyBuilder {
	bc.Token = tools.RandomHash(60)
	return bc
}

//-

type PaymentSystemCompany struct {
	Company
}

func (psc *PaymentSystemCompany) SetName(name string) CompanyBuilder {
	psc.Name = name
	return psc
}

func (psc *PaymentSystemCompany) SetWebSite(webSite string) CompanyBuilder {
	psc.WebSite = webSite
	return psc
}

func (psc *PaymentSystemCompany) SetEmail(email string) CompanyBuilder {
	psc.Email = email
	return psc
}

func (psc *PaymentSystemCompany) SetBIN(bin string) CompanyBuilder {
	psc.BIN = bin
	return psc
}

func (psc *PaymentSystemCompany) SetActive(value bool) CompanyBuilder {
	psc.Active = value
	return psc
}

func (psc *PaymentSystemCompany) Build() CompanyBuilder {
	return psc
}
