package r2o

// ProductService handles communication with the issue related
// methods of the ready2order API.
//
// ready2order API docs: https://app.swaggerhub.com/apis-docs/ready2ordergmbh/ready2order-api-production/1.0.329#/
type ProductService service

type ProductResponse struct {
	ProductgroupActive      *bool   `json:"productgroup_active"`
	ProductgroupCreatedAt   *string `json:"productgroup_created_at"`
	ProductgroupDescription *string `json:"productgroup_description"`
	ProductgroupID          *int64  `json:"productgroup_id"`
	ProductgroupName        *string `json:"productgroup_name"`
	ProductgroupShortcut    *string `json:"productgroup_shortcut"`
	ProductgroupSortIndex   *int64  `json:"productgroup_sortIndex"`
	ProductgroupTypeID      *int64  `json:"productgroup_type_id"`
	ProductgroupUpdatedAt   *string `json:"productgroup_updated_at"`
}
