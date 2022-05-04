package r2o

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// BillService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type BillService service

type BillDocumentGetParams struct {
	Pagination

	//Search query (For example: RG2018/1)
	Query *string `url:"query"`

	//Table-Id which should be filtered (For example: 1234)
	TableId *int `url:"tableId"`

	//Customer-Id which should be filtered (For example: 1234)
	CustomerId *int `url:"customerId"`

	//Date field you want to query (Possible values: dr_startDate or b_dateTime)
	DateField *DateField `url:"dateField"`

	//Date from (For example: 2019-01-01)
	DateFrom *string `url:"dateFrom"`

	//Date to (For example: 2019-12-31
	DateTo *string `url:"dateTo"`

	//Trainingmode on/off
	TestMode *bool `url:"testMode"`

	//Include list of items (Default: false)
	Items *bool `url:"items"`

	//Include list of discounts (Default: false)
	Discounts *bool `url:"discounts"`

	//Include list of payments (Default: false)
	Payments *bool `url:"payments"`
}

type InvoiceAddress struct {
	City       *string `json:"city"`
	Company    *string `json:"company"`
	Country    *string `json:"country"`
	Email      *string `json:"email"`
	FirstName  *string `json:"firstName"`
	LastName   *string `json:"lastName"`
	Phone      *string `json:"phone"`
	Salutation *string `json:"salutation"`
	Street     *string `json:"street"`
	Title      *string `json:"title"`
	VatID      *string `json:"vatId"`
	Zip        *string `json:"zip"`
}

type InvoiceDiscountResponse struct {
	BillDiscountCreatedAt       *string `json:"billDiscount_created_at"`
	BillDiscountDiscountGroupID *int64  `json:"billDiscount_discountGroup_id"`
	BillDiscountDiscountID      *int64  `json:"billDiscount_discount_id"`
	BillDiscountID              *int64  `json:"billDiscount_id"`
	BillDiscountName            *string `json:"billDiscount_name"`
	BillDiscountPercent         *string `json:"billDiscount_percent"`
	BillDiscountUpdatedAt       *string `json:"billDiscount_updated_at"`
	BillDiscountValue           *string `json:"billDiscount_value"`
}

type InvoiceItem struct {
	CustomerID                               *int64  `json:"customer_id"`
	DailyReportID                            *int64  `json:"dailyReport_id"`
	DeliveryBillID                           *int64  `json:"deliveryBill_id"`
	InvoiceID                                *int64  `json:"invoice_id"`
	ItemAccountingCode                       *string `json:"item_accountingCode"`
	ItemComment                              *string `json:"item_comment"`
	ItemDiscountable                         *bool   `json:"item_discountable"`
	ItemExpirationDate                       *string `json:"item_expirationDate"`
	ItemID                                   *int64  `json:"item_id"`
	ItemInvoiceDiscountGross                 *string `json:"item_invoiceDiscountGross"`
	ItemInvoiceDiscountNet                   *string `json:"item_invoiceDiscountNet"`
	ItemLineDiscountGross                    *string `json:"item_lineDiscountGross"`
	ItemLineDiscountGroupID                  *int64  `json:"item_lineDiscountGroupId"`
	ItemLineDiscountID                       *int64  `json:"item_lineDiscountId"`
	ItemLineDiscountName                     *string `json:"item_lineDiscountName"`
	ItemLineDiscountNet                      *string `json:"item_lineDiscountNet"`
	ItemLineDiscountPercent                  *string `json:"item_lineDiscountPercent"`
	ItemName                                 *string `json:"item_name"`
	ItemNumber                               *string `json:"item_number"`
	ItemPrice                                *string `json:"item_price"`
	ItemPriceBase                            *bool   `json:"item_priceBase"`
	ItemPriceNet                             *string `json:"item_priceNet"`
	ItemProductDiscountable                  *bool   `json:"item_product_discountable"`
	ItemProductPrice                         *string `json:"item_product_price"`
	ItemProductPriceNet                      *string `json:"item_product_priceNet"`
	ItemProductPriceNetPerUnit               *string `json:"item_product_priceNetPerUnit"`
	ItemProductPricePerUnit                  *string `json:"item_product_pricePerUnit"`
	ItemProductVat                           *string `json:"item_product_vat"`
	ItemProductVatCompensation               *string `json:"item_product_vatCompensation"`
	ItemProductVatRate                       *string `json:"item_product_vatRate"`
	ItemQty                                  *string `json:"item_qty"`
	ItemQuantity                             *string `json:"item_quantity"`
	ItemRetour                               *bool   `json:"item_retour"`
	ItemSerialNumber                         *string `json:"item_serialNumber"`
	ItemTestMode                             *bool   `json:"item_testMode"`
	ItemTimestamp                            *string `json:"item_timestamp"`
	ItemTotal                                *string `json:"item_total"`
	ItemTotalNet                             *string `json:"item_totalNet"`
	ItemUndiscountedSumBeforeInvoiceDiscount *string `json:"item_undiscountedSumBeforeInvoiceDiscount"`
	ItemUndiscountedSumBeforeLineDiscount    *string `json:"item_undiscountedSumBeforeLineDiscount"`
	ItemUpdatedStock                         *bool   `json:"item_updatedStock"`
	ItemVat                                  *string `json:"item_vat"`
	ItemVatCompensation                      *string `json:"item_vatCompensation"`
	ItemVatRate                              *string `json:"item_vatRate"`
	PaymentMethodID                          *int64  `json:"paymentMethod_id"`
	PaymentID                                *int64  `json:"payment_id"`
	ProductGroupID                           *int64  `json:"productGroup_id"`
	ProductUnitID                            *int64  `json:"productUnit_id"`
	ProductID                                *int64  `json:"product_id"`
	ProductgroupName                         *string `json:"productgroup_name"`
	RetourInvoiceID                          *int64  `json:"retourInvoice_id"`
	TableID                                  *int64  `json:"table_id"`
	TableName                                *string `json:"table_name"`
	UserID                                   *int64  `json:"user_id"`
	UserName                                 *string `json:"user_name"`
}

type InvoicePayment struct {
	BillPaymentID    *int64  `json:"billPayment_id"`
	BillPaymentName  *string `json:"billPayment_name"`
	BillPaymentValue *string `json:"billPayment_value"`
	PaymentID        *int64  `json:"payment_id"`
	ReferencedBillID *int64  `json:"referenced_bill_id"`
}

type InvoiceTransaction struct {
	ID    *string `json:"id"`
	R2oID *string `json:"r2o_id"`
}

type InvoiceType struct {
	BillTypeID     *int64  `json:"billType_id"`
	BillTypeName   *string `json:"billType_name"`
	BillTypeSymbol *string `json:"billType_symbol"`
}

type InvoiceResponse struct {
	Address                               *InvoiceAddress            `json:"address"`
	BillTypeID                            *int64                     `json:"billType_id"`
	CurrencyID                            *int64                     `json:"currency_id"`
	CustomerCategoryID                    *int64                     `json:"customerCategory_id"`
	CustomerID                            *int64                     `json:"customer_id"`
	Discounts                             *[]InvoiceDiscountResponse `json:"discounts"`
	ID                                    *string                    `json:"id"`
	InvoiceAddressCity                    *string                    `json:"invoice_address_city"`
	InvoiceAddressCompany                 *string                    `json:"invoice_address_company"`
	InvoiceAddressCountry                 *string                    `json:"invoice_address_country"`
	InvoiceAddressEmail                   *string                    `json:"invoice_address_email"`
	InvoiceAddressFirstName               *string                    `json:"invoice_address_firstName"`
	InvoiceAddressLastName                *string                    `json:"invoice_address_lastName"`
	InvoiceAddressPhone                   *string                    `json:"invoice_address_phone"`
	InvoiceAddressSalutation              *string                    `json:"invoice_address_salutation"`
	InvoiceAddressStreet                  *string                    `json:"invoice_address_street"`
	InvoiceAddressTitle                   *string                    `json:"invoice_address_title"`
	InvoiceAddressVatID                   *string                    `json:"invoice_address_vatId"`
	InvoiceAddressZip                     *string                    `json:"invoice_address_zip"`
	InvoiceDeletedReason                  *string                    `json:"invoice_deletedReason"`
	InvoiceDeletedAt                      *string                    `json:"invoice_deleted_at"`
	InvoiceDeliveryDate                   *string                    `json:"invoice_deliveryDate"`
	InvoiceDueDate                        *string                    `json:"invoice_dueDate"`
	InvoiceExternalReferenceNumber        *string                    `json:"invoice_externalReferenceNumber"`
	InvoiceFormatVersion                  *int64                     `json:"invoice_formatVersion"`
	InvoiceID                             *int64                     `json:"invoice_id"`
	InvoiceInPrinterQueue                 *bool                      `json:"invoice_inPrinterQueue"`
	InvoiceInternalInvoiceReferenceNumber *string                    `json:"invoice_internalInvoiceReferenceNumber"`
	InvoiceLocked                         *bool                      `json:"invoice_locked"`
	InvoiceNumber                         *int64                     `json:"invoice_number"`
	InvoiceNumberFull                     *string                    `json:"invoice_numberFull"`
	InvoicePaid                           *bool                      `json:"invoice_paid"`
	InvoicePaidDate                       *string                    `json:"invoice_paidDate"`
	InvoicePdf                            *string                    `json:"invoice_pdf"`
	InvoicePriceBase                      *string                    `json:"invoice_priceBase"`
	InvoiceReferenceID                    *int64                     `json:"invoice_reference_id"`
	InvoiceTestMode                       *bool                      `json:"invoice_testMode"`
	InvoiceText                           *string                    `json:"invoice_text"`
	InvoiceTextBeforeItemsTable           *string                    `json:"invoice_textBeforeItemsTable"`
	InvoiceTimestamp                      *string                    `json:"invoice_timestamp"`
	InvoiceTotal                          *string                    `json:"invoice_total"`
	InvoiceTotalNet                       *string                    `json:"invoice_totalNet"`
	InvoiceTotalTip                       *string                    `json:"invoice_totalTip"`
	InvoiceTotalVat                       *string                    `json:"invoice_totalVat"`
	Items                                 *[]InvoiceItem             `json:"items"`
	Payment                               *[]InvoicePayment          `json:"payment"`
	PaymentMethodID                       *int64                     `json:"paymentMethod_id"`
	PrinterID                             *int64                     `json:"printer_id"`
	R2oID                                 *string                    `json:"r2o_id"`
	TableAreaID                           *int64                     `json:"tableArea_id"`
	TableID                               *int64                     `json:"table_id"`
	Transaction                           *InvoiceTransaction        `json:"transaction"`
	Type                                  *InvoiceType               `json:"type"`
	UserID                                *int64                     `json:"user_id"`
}

type BillResponse struct {
	Count    *int64             `json:"count"`
	DateFrom *string            `json:"dateFrom"`
	DateTo   *string            `json:"dateTo"`
	Invoices *[]InvoiceResponse `json:"invoices"`
	Limit    *int64             `json:"limit"`
	Offset   *int64             `json:"offset"`
	Query    *string            `json:"query"`
}

func (as *BillService) GetInvoices(ctx context.Context, query *BillDocumentGetParams) (*BillResponse, error) {
	responseData := BillResponse{}

	err := as.client.runHttpRequestWithParamsWithContext(ctx, "document/invoice", http.MethodGet, query, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

type InvoiceItemRequest struct {
	DiscountID    int64  `json:"discount_id"`
	DiscountUnit  string `json:"discount_unit"`
	DiscountValue string `json:"discount_value"`
	ItemComment   string `json:"item_comment"`
	ItemName      string `json:"item_name"`
	ItemPrice     string `json:"item_price"`
	ItemPriceBase string `json:"item_priceBase"`
	ItemQuantity  string `json:"item_quantity"`
	ItemVatID     int64  `json:"item_vatId"`
	ItemVatRate   string `json:"item_vatRate"`
	ProductID     int64  `json:"product_id"`
}

type InvoiceRequest struct {
	Address                            *[]InvoiceAddress     `json:"address"`
	BillTypeID                         *int64                `json:"billType_id"`
	CreatePDF                          *bool                 `json:"createPDF"`
	InvoiceDueDate                     *string               `json:"invoice_dueDate"`
	InvoiceExternalReferenceNumber     *string               `json:"invoice_externalReferenceNumber"`
	InvoiceInPrinterQueue              *int64                `json:"invoice_inPrinterQueue"`
	InvoicePaid                        *bool                 `json:"invoice_paid"`
	InvoicePaidDate                    *string               `json:"invoice_paidDate"`
	InvoicePriceBase                   *string               `json:"invoice_priceBase"`
	InvoiceRoundToSmallestCurrencyUnit *string               `json:"invoice_roundToSmallestCurrencyUnit"`
	InvoiceShowRecipient               *bool                 `json:"invoice_showRecipient"`
	InvoiceTestMode                    *bool                 `json:"invoice_testMode"`
	InvoiceText                        *string               `json:"invoice_text"`
	InvoiceTextBeforeItemsTable        *string               `json:"invoice_textBeforeItemsTable"`
	Items                              *[]InvoiceItemRequest `json:"items"`
	PaymentMethodID                    *int64                `json:"paymentMethod_id"`
	PdfFormat                          *string               `json:"pdfFormat"`
	PrinterID                          *int64                `json:"printer_id"`
	UserID                             *int64                `json:"user_id"`
}

func (as *BillService) CreateInvoice(ctx context.Context, data *InvoiceRequest) (*InvoiceResponse, error) {
	responseData := InvoiceResponse{}

	err := as.client.runHttpRequestWithContext(ctx, "document/invoice", http.MethodPut, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (as *BillService) CountInvoices(ctx context.Context) (*int, error) {
	responseItem := struct {
		Error *bool `json:"error"`
		Count *int  `json:"count"`
	}{}

	err := as.client.runHttpRequestWithContext(ctx, "document/invoice/count", http.MethodGet, nil, &responseItem)

	if err != nil {
		return nil, err
	}

	if !*responseItem.Error {
		return nil, errors.New("cannot load count of invoices")
	}

	return responseItem.Count, nil
}

func (as *BillService) GetInvoice(ctx context.Context, invoiceId *int) (*InvoiceResponse, error) {
	responseData := InvoiceResponse{}

	u := fmt.Sprintf("document/invoice/%v", *invoiceId)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, nil, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

type InvoiceDeleteRequest struct {
	InvoiceExternalReferenceNumber *string `json:"invoice_externalReferenceNumber"`
	StornoID                       *int64  `json:"storno_id"`
	StornoReason                   *string `json:"storno_reason"`
}

func (as *BillService) DeleteInvoice(ctx context.Context, invoiceId *int, data *InvoiceDeleteRequest) (*InvoiceResponse, error) {
	responseData := InvoiceResponse{}

	u := fmt.Sprintf("document/invoice/%v/delete", *invoiceId)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPost, data, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

type MiniMessageResponse struct {
	Error   *bool   `json:"error"`
	Message *string `json:"msg"`
}

func (as *BillService) PrintInvoice(ctx context.Context, invoiceId, printerId *int) (*MiniMessageResponse, error) {
	responseData := MiniMessageResponse{}

	requestData := struct {
		PrinterId *int `json:"printer_id"`
	}{
		PrinterId: printerId,
	}

	u := fmt.Sprintf("document/invoice/%v/print", *invoiceId)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodPatch, requestData, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

type InvoicePdfResponse struct {
	Error  *bool   `json:"error"`
	Format *string `json:"format"`
	URI    *string `json:"uri"`
}

func (as *BillService) GetPdfInvoice(ctx context.Context, invoiceId *int, regeneratePdf *bool) (*InvoicePdfResponse, error) {
	responseData := InvoicePdfResponse{}

	requestData := struct {
		RegeneratePDF *bool `json:"regeneratePDF"`
	}{
		RegeneratePDF: regeneratePdf,
	}

	u := fmt.Sprintf("document/invoice/%v/pdf", *invoiceId)
	err := as.client.runHttpRequestWithContext(ctx, u, http.MethodGet, requestData, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
