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

	GetDataRemainingLeaveReq struct {
		CompanyId  	*int64 `json:"company_id" binding:"required"`
		EmployeeId 	string `json:"employee_id"`
		LeaveCode 	string `json:"leave_code"`
	}

	GetDataRemainingLeaveResponse struct {
		StartValidDate  time.Time 	`json:"startvaliddate"`
		EndValidDate  	time.Time 	`json:"endvaliddate"`
		Remaining  		float64		`json:"remaining"`
		Total  			float64		`json:"total"`
		RequiredRefDoc 	string 		`json:"required_refdoc"`
		RequiredRemark 	string 		`json:"required_remark"`
	}

	GetDataRequiredRefDocResponse struct {
		RequiredRefDoc 	string 		`json:"required_refdoc"`
	}

	GetDataRequiredRemarkResponse struct {
		RequiredRemark 	string 		`json:"required_remark"`
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
		RequestNo      string  `schema:"request_no"`
		CompanyId      string  `schema:"company_id"`
		Requestedby    string  `schema:"requestedby"`
		Requestfor     string  `schema:"requestfor"`
		LeaveCode      string  `schema:"leave_code"`
		LeaveType      string  `schema:"leave_type"`
		LeaveStartdate string  `schema:"leave_startdate"`
		LeaveEnddate   string  `schema:"leave_enddate"`
		LeaveStarttime string  `schema:"leave_starttime"`
		LeaveEndtime   string  `schema:"leave_endtime"`
		Remark         *string `schema:"remark"`
		Username       *string `schema:"username"`
		Refdoc         string  `json:"refdoc"`
		Reqfullday     string  `json:"reqfullday"`
		HdtypeStarttime string  `json:"hdtypestarttime"`
		HdtypeEndtime  string  `json:"hdtypeendtime"`
	}
)
