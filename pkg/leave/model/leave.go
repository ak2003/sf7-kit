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

	CreateLeaveRequestReq struct {
		RequestNo      string  `json:"request_no"`
		CompanyId      int64   `json:"company_id"`
		Requestedby    string  `json:"requestedby"`
		Requestfor     string  `json:"requestfor"`
		LeaveCode      string  `json:"leave_code"`
		LeaveStartdate string  `json:"leave_startdate"`
		LeaveEnddate   string  `json:"leave_enddate"`
		Remark         *string `json:"remark"`
		Username       string  `json:"username"`
	}

	CreateLeaveRequestFormReq struct {
		RequestNo      string  `schema:"request_no,required"`
		CompanyId      string  `schema:"company_id,required"`
		Requestedby    string  `schema:"requestedby"`
		Requestfor     string  `schema:"requestfor"`
		LeaveCode      string  `schema:"leave_code"`
		LeaveStartdate string  `schema:"leave_startdate"`
		LeaveEnddate   string  `schema:"leave_enddate"`
		Remark         *string `schema:"remark"`
		Username       *string `schema:"username"`
	}
)
