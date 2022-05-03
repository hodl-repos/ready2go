package r2o

import (
	"context"
	"fmt"
	"net/http"
)

// CountryService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type CustomerService service

type CustomerResponse struct {
	City               *string `json:"city"`
	CustomerCategoryID *int    `json:"customerCategory_id"`
	CustomerBirthday   *string `json:"customer_birthday"`
	CustomerCompany    *string `json:"customer_company"`
	CustomerID         *int    `json:"customer_id"`
	CustomerName       *string `json:"customer_name"`
	CustomerNotes      *string `json:"customer_notes"`
	CustomerNumber     *int    `json:"customer_number"`
	Email              *string `json:"email"`
	FirstName          *string `json:"firstName"`
	LastName           *string `json:"lastName"`
	Phone              *string `json:"phone"`
	Salutation         *string `json:"salutation"`
	Street             *string `json:"street"`
	Title              *string `json:"title"`
	VatID              *string `json:"vatId"`
	Zip                *string `json:"zip"`
}

type CustomerRequest struct {
	Birthday       string `json:"birthday"`
	City           string `json:"city"`
	Company        string `json:"company"`
	CustomerName   string `json:"customer_name"`
	CustomerNumber int64  `json:"customer_number"`
	Email          string `json:"email"`
	FirstName      string `json:"firstName"`
	FkCouID        int64  `json:"fk_cou_id"`
	LastName       string `json:"lastName"`
	Notes          string `json:"notes"`
	Phone          string `json:"phone"`
	Salutation     string `json:"salutation"`
	Street         string `json:"street"`
	Title          string `json:"title"`
	UID            string `json:"uid"`
	Zip            string `json:"zip"`
}

func (as *CustomerService) GetCustomers(ctx context.Context, page *Pagination) (*[]CustomerResponse, error) {
	responseData := make([]CustomerResponse, 0)

	err := as.client.runHttpRequestWithParamsWithContext(ctx, "customers", http.MethodGet, page, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *CustomerService) GetCustomer(ctx context.Context, id *int) (*CustomerResponse, error) {
	responseData := CustomerResponse{}

	u := fmt.Sprintf("customers/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *CustomerService) CreateCustomer(ctx context.Context, data *CustomerRequest) (*CustomerResponse, error) {
	responseData := CustomerResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "customers", http.MethodPut, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *CustomerService) UpdateCustomer(ctx context.Context, id *int, data *CustomerRequest) (*CustomerResponse, error) {
	responseData := CustomerResponse{}

	u := fmt.Sprintf("customers/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPost, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *CustomerService) DeleteCustomer(ctx context.Context, id *int) (*MessageResponse, error) {
	responseData := MessageResponse{}

	u := fmt.Sprintf("customers/%v", *id)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodDelete, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *CustomerService) MergeCustomer(ctx context.Context, masterCustomer *int, customers *[]int) (*MessageResponse, error) {
	responseData := MessageResponse{}

	requestData := struct {
		MasterCustomer *int   `json:"masterCustomer"`
		Customers      *[]int `json:"customerIds"`
	}{
		MasterCustomer: masterCustomer,
		Customers:      customers,
	}

	err := as.client.runHttpRequestWithContext(ctx, "customers/merge", http.MethodPost, &requestData, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
