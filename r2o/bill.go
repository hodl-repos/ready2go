package r2o

// BillService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type BillService service

type BillDocumentGetParams struct {
	Pagination

	//Search query (For example: RG2018/1)
	Query *string `json:"query"`

	//Table-Id which should be filtered (For example: 1234)
	TableId *int `json:"tableId"`

	//Customer-Id which should be filtered (For example: 1234)
	CustomerId *int `json:"CustomerId"`

	//Date field you want to query (Possible values: dr_startDate or b_dateTime)
	DateField *DateField `json:"dateField"`

	//Date from (For example: 2019-01-01)
	DateFrom *string `json:"dateFrom"`

	//Date to (For example: 2019-12-31
	DateTo *string `json:"dateTo"`

	//Trainingmode on/off
	TestMode *bool `json:"testMode"`

	//Include list of items (Default: false)
	Items *bool `json:"items"`

	//Include list of discounts (Default: false)
	Discounts *bool `json:"discounts"`

	//Include list of payments (Default: false)
	Payments *bool `json:"payments"`
}
