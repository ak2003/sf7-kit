package model

import "time"

type (
	GetLeaveRequestListingRequest struct {
		CompanyId            int64    `json:"company_id" binding:"required"`
		EmployeeId           string   `json:"employee_id" binding:"required"`
		Language             string   `json:"language"`
		Totaldays            string   `json:"totaldays"`
		Remark               string   `json:"remark"`
		RequestNo            string   `json:"request_no"`
		LeaveCode            string   `json:"leave_code"`
		Requestfor           string   `json:"requestfor"`
		RequestStatus        string   `json:"request_status"`
		LeaveStartdate       string   `json:"leave_startdate"`
		LeaveEnddate         string   `json:"leave_enddate"`
		FilterLeaveStartdate string   `json:"filter_leave_startdate"`
		FilterLeaveEnddate   string   `json:"filter_leave_enddate"`
		FilterLeaveCode      []string `json:"filter_leave_code"`
		FilterRequestStatus  []string `json:"filter_request_status"`
		Page                 int64    `json:"page"`
		Limit                int64    `json:"limit"`
		Field                string   `json:"field"`
		Order                string   `json:"order"`
	}

	GetLeaveRequestListingResponse struct {
		RequestNo      string    `json:"request_no"`
		CompanyId      int64     `json:"company_id"`
		Requestfor     string    `json:"requestfor"`
		Refdoc         string    `json:"refdoc"`
		LeaveStartdate time.Time `json:"leave_startdate"`
		LeaveEnddate   time.Time `json:"leave_enddate"`
		Totaldays      float64   `json:"totaldays"`
		LeaveCode      string    `json:"leave_code"`
		Remark         *string   `json:"remark"`
		RequestStatus  string    `json:"request_status"`
		Reqfullday     string    `json:"reqfullday"`
	}

	GetLeaveRequestListingFilterRequest struct {
		CompanyId            int64    `json:"company_id" binding:"required"`
		EmployeeId           string   `json:"employee_id" binding:"required"`
		Language             string   `json:"language"`
		FilterLeaveStartdate string   `json:"filter_leave_startdate"`
		FilterLeaveEnddate   string   `json:"filter_leave_enddate"`
		FilterLeaveCode      []string `json:"filter_leave_code"`
		FilterRequestStatus  []string `json:"filter_request_status"`
		Page                 int64    `json:"page"`
		Limit                int64    `json:"limit"`
	}

	GetLeaveRequestListingFilterResponse struct {
		RequestNo      string    `json:"request_no"`
		CompanyId      int64     `json:"company_id"`
		Requestfor     string    `json:"requestfor"`
		Refdoc         string    `json:"refdoc"`
		LeaveStartdate time.Time `json:"leave_startdate"`
		LeaveEnddate   time.Time `json:"leave_enddate"`
		Totaldays      float64   `json:"totaldays"`
		LeaveCode      string    `json:"leave_code"`
		Remark         *string   `json:"remark"`
		RequestStatus  string    `json:"request_status"`
		Reqfullday     string    `json:"reqfullday"`
	}

	GetDataTypeOfLeaveReq struct {
		Language   string `json:"language"`
		CompanyId  *int64 `json:"company_id"`
		EmployeeId string `json:"employee_id"`
	}

	GetDataTypeOfLeaveResponse struct {
		Optvalue string `json:"optvalue"`
		Opttext  string `json:"opttext"`
		Optempid string `json:"optempid"`
	}

	GetDataRequestForReq struct {
		Status    string `json:"status"`
		CompanyId *int64 `json:"company_id"`
		Search    string `json:"search"`
	}

	GetDataRequestForResponse struct {
		EmployeeId    string `json:"employee_id"`
		EmployeeName  string `json:"employee_name"`
		EmployeeTitle string `json:"employee_title"`
	}
)
