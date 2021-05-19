package model

type (
	CoreManageField struct {
		Id               string `db:"id"`
		CompanyId        string `db:"company_id"`
		PageId           string `db:"page_id"`
		TableName        string `db:"table_name"`
		AdditionalFields string `db:"additional_fields"`
		StatusFields     string `db:"status_fields"`
	}
)
