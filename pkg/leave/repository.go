package leave

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"gitlab.dataon.com/gophers/sf7-kit/pkg/leave/model"
	"gitlab.dataon.com/gophers/sf7-kit/shared/utils/logger"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetLeaveRequestListing(ctx context.Context, sc model.GetLeaveRequestListingRequest) (error, []model.GetLeaveRequestListingResponse)
	GetLeaveRequestFilterListing(ctx context.Context, sc model.GetLeaveRequestListingFilterRequest) (error, []model.GetLeaveRequestListingFilterResponse)
	GetDataTypeOfLeave(ctx context.Context, sc model.GetDataTypeOfLeaveReq) (error, []model.GetDataTypeOfLeaveResponse)
	GetDataRequestFor(ctx context.Context, sc model.GetDataRequestForReq) (error, []model.GetDataRequestForResponse)
	GetDataRemainingLeave(ctx context.Context, sc model.GetDataRemainingLeaveReq) (error, []model.GetDataRemainingLeaveResponse)
	CreateLeaveRequest(ctx context.Context, sc model.CreateLeaveRequestReq) (error, string)
	CreateLeaveRequestForm(ctx context.Context, sc model.CreateLeaveRequestFormReq) (error, string)
}

type repo struct {
	dbSlave  *sqlx.DB
	dbMaster *sqlx.DB
}

func NewRepo(dbSlave, dbMaster *sqlx.DB) Repository {
	return &repo{
		dbSlave:  dbSlave,
		dbMaster: dbMaster,
	}
}

func (repo *repo) GetDataRemainingLeave(ctx context.Context, req model.GetDataRemainingLeaveReq) (error, []model.GetDataRemainingLeaveResponse) {
	var (
		dataRemainingLeave      []model.GetDataRemainingLeaveResponse
		errData             	error
		queryDataRemainingLeave string
		queryDataRequiredRefDoc string
		queryDataRequiredRemark string
		paramData           	[]interface{}
		paramDataRefDoc       	[]interface{}
		paramDataRemark       	[]interface{}
	)

	if req.CompanyId == nil {
		errData = errors.New("company_id is mandatory")
		return errData, dataRemainingLeave
	}
	
	LeaveCodeSplit := strings.Split(req.LeaveCode, "|")

	paramData = append(paramData, req.EmployeeId)
	paramData = append(paramData, LeaveCodeSplit[0])
	paramData = append(paramData, req.CompanyId)
	queryDataRemainingLeave = `SELECT startvaliddate, remaining
							FROM TTADEMPGETLEAVE
							WHERE emp_id = ?
							AND leave_code = ?
							AND active_status = 1
							AND company_id = ?
							AND (
								endvaliddate >= getdate()
								OR
								endvaliddate IS NULL
							)
							order by startvaliddate,endvaliddate asc`

	queryDataRemainingLeave = repo.dbSlave.Rebind(queryDataRemainingLeave)
	res1, errData := repo.dbSlave.Queryx(queryDataRemainingLeave, paramData...)
	if errData != nil {
		logger.Error(nil, errData)
		// logger.Println(queryListing)
		return errData, dataRemainingLeave
	}

	defer res1.Close()
	
	if res1.Next() {
		var temp model.GetDataRemainingLeaveResponse
		errData := res1.Scan(&temp.StartValidDate, &temp.Remaining)
		if errData != nil {
			logger.Error(nil, errData)
			// logger.Println(queryListing)
			res1.Close()
			return errData, dataRemainingLeave
		}

		// Start Process Required RefDoc
		paramDataRefDoc = append(paramDataRefDoc, LeaveCodeSplit[0])
		paramDataRefDoc = append(paramDataRefDoc, req.CompanyId)
		queryDataRequiredRefDoc = `SELECT requiredattachment
								FROM TTAMLEAVETYPE
								WHERE leave_code = ?
								AND company_id = ?`

		queryDataRequiredRefDoc = repo.dbSlave.Rebind(queryDataRequiredRefDoc)
		res2, errData2 := repo.dbSlave.Queryx(queryDataRequiredRefDoc, paramDataRefDoc...)
		if errData2 != nil {
			logger.Error(nil, errData2)
			// logger.Println(queryListing)
			return errData2, dataRemainingLeave
		}

		defer res2.Close()

		if res2.Next() {
			var temp2 model.GetDataRequiredRefDocResponse
			errData2 := res2.Scan(&temp2.RequiredRefDoc)
			if errData2 != nil {
				logger.Error(nil, errData2)
				// logger.Println(queryListing)
				res2.Close()
				return errData2, dataRemainingLeave
			}

			for res2.Next() {
				errData2 := res2.Scan(&temp2.RequiredRefDoc)
				if errData2 != nil {
					logger.Error(nil, errData2)
					// logger.Println(queryListing)
					res2.Close()
					return errData2, dataRemainingLeave
				}
			}

			temp.RequiredRefDoc = temp2.RequiredRefDoc
		}
		// End Process Required RefDoc

		// Start Process Required Remark
		paramDataRemark = append(paramDataRemark, req.CompanyId)
		queryDataRequiredRemark = `SELECT field_value
								FROM TCLCAPPCOMPANY
								WHERE field_code = 'remarkisrequired'
								AND company_id = ?`

		queryDataRequiredRemark = repo.dbSlave.Rebind(queryDataRequiredRemark)
		res3, errData3 := repo.dbSlave.Queryx(queryDataRequiredRemark, paramDataRemark...)
		if errData3 != nil {
			logger.Error(nil, errData3)
			// logger.Println(queryListing)
			return errData3, dataRemainingLeave
		}

		defer res3.Close()

		if res3.Next() {
			var temp2 model.GetDataRequiredRemarkResponse
			errData3 := res3.Scan(&temp2.RequiredRemark)
			if errData3 != nil {
				logger.Error(nil, errData3)
				// logger.Println(queryListing)
				res3.Close()
				return errData3, dataRemainingLeave
			}

			for res3.Next() {
				errData3 := res3.Scan(&temp2.RequiredRemark)
				if errData3 != nil {
					logger.Error(nil, errData3)
					// logger.Println(queryListing)
					res3.Close()
					return errData3, dataRemainingLeave
				}
			}

			temp.RequiredRemark = temp2.RequiredRemark
		}
		// End Process Required Remark

		dataRemainingLeave = append(dataRemainingLeave, temp)
		for res1.Next() {
			errData := res1.Scan(&temp.StartValidDate, &temp.Remaining)
			if errData != nil {
				logger.Error(nil, errData)
				// logger.Println(queryListing)
				res1.Close()
				return errData, dataRemainingLeave
			}

			dataRemainingLeave = append(dataRemainingLeave, temp)
		}
	}

	return errData, dataRemainingLeave
}

func (repo *repo) GetDataRequestFor(ctx context.Context, req model.GetDataRequestForReq) (error, []model.GetDataRequestForResponse) {
	var (
		dataRequestFor      []model.GetDataRequestForResponse
		errData             error
		queryDataRequestfor string
		paramData           []interface{}
	)

	if req.CompanyId == nil {
		errData = errors.New("company_id is mandatory")
		return errData, dataRequestFor
	}
	if strings.Trim(req.Status, " ") != "" {
		req.Status = "ACTIVE"
	}
	paramData = append(paramData, strings.ToUpper(req.Status))
	paramData = append(paramData, req.CompanyId)

	queryDataRequestfor = `SELECT distinct A.emp_id emp_id, A.full_name , A.full_name+' ['+b.emp_no+']' ntitle
							FROM TEOMEMPPERSONAL a
							LEFT JOIN VIEW_EMPLOYEE_STS b ON a.emp_id = b.emp_id
							WHERE b.active = ? AND b.company_id = ?`

	if strings.Trim(req.Search, " ") != "" {
		paramData = append(paramData, fmt.Sprintf(`%%%s%%`, req.Search))
		queryDataRequestfor = queryDataRequestfor + ` AND A.full_name+' ['+b.emp_no+']' LIKE ? `
	}
	queryDataRequestfor = queryDataRequestfor + ` ORDER BY ntitle`

	queryDataRequestfor = repo.dbSlave.Rebind(queryDataRequestfor)
	res1, errData := repo.dbSlave.Queryx(queryDataRequestfor, paramData...)
	// fmt.Print(queryListing)
	if errData != nil {
		logger.Error(nil, errData)
		// logger.Println(queryListing)
		return errData, dataRequestFor
	}

	defer res1.Close()

	if res1.Next() {
		var temp model.GetDataRequestForResponse
		errData := res1.Scan(&temp.EmployeeId, &temp.EmployeeName, &temp.EmployeeTitle)
		if errData != nil {
			logger.Error(nil, errData)
			// logger.Println(queryListing)
			res1.Close()
			return errData, dataRequestFor
		}

		dataRequestFor = append(dataRequestFor, temp)
		for res1.Next() {
			errData := res1.Scan(&temp.EmployeeId, &temp.EmployeeName, &temp.EmployeeTitle)
			if errData != nil {
				logger.Error(nil, errData)
				// logger.Println(queryListing)
				res1.Close()
				return errData, dataRequestFor
			}

			dataRequestFor = append(dataRequestFor, temp)
		}
	}
	return errData, dataRequestFor
}

func (repo *repo) CreateLeaveRequestForm(ctx context.Context, req model.CreateLeaveRequestFormReq) (error, string) {
	var (
		result      string
		errCreate   error
		queryCreate string
	)
	result = "OK"

	if req.CompanyId == "" {
		result = "company_id is mandatory"
		return errCreate, result
	}

	if req.LeaveCode == "" {
		result = "leave_code is mandatory"
		return errCreate, result
	}

	if req.LeaveStartdate == "" {
		result = "leave_startdate is mandatory"
		return errCreate, result
	}

	if req.LeaveEnddate == "" {
		result = "leave_enddate is mandatory"
		return errCreate, result
	}

	if req.RequestNo == "" {
		result = "request_no is mandatory"
		return errCreate, result
	}

	if req.Requestfor == "" {
		result = "requestfor is mandatory"
		return errCreate, result
	}

	queryCreate = `INSERT INTO dbSF6_QA.dbo.TTADLEAVEREQUEST
		(request_no, company_id, requestedby, requestfor, requestdate, 
		leave_code, leave_startdate, leave_enddate, usecalendar, totaldays, 
		remark, created_by, created_date, refdoc, reqfullday, 
		approval_status, hdtype_starttime, hdtype_endtime, leave_start_halfday, leave_end_halfday)
		VALUES
		(?, ?, ?, ?, getdate(), 
		?, ?, ?, 'N', 1.0000, 
		NULL, N'shiburin1988', '2018-05-14 14:47:43.000', NULL, N'Y', 
		NULL, 0, 0, NULL, NULL);`
	queryCreate = repo.dbMaster.Rebind(queryCreate)

	tx, errCreate := repo.dbMaster.Begin()
	if errCreate != nil {
		return errCreate, result
	}
	createRequest, errCreate := tx.Prepare(queryCreate)
	if errCreate != nil {
		return errCreate, result
	}
	defer createRequest.Close()

	_, errCreate = createRequest.Exec()
	if errCreate != nil {
		return errCreate, result
	}
	tx.Commit()
	return errCreate, result
}

func (repo *repo) CreateLeaveRequest(ctx context.Context, req model.CreateLeaveRequestReq) (error, string) {
	var (
		result      string
		errCreate   error
		queryCreate string
	)
	result = "OK"

	if req.CompanyId == 0 {
		result = "company_id is mandatory"
		return errCreate, result
	}

	if req.LeaveCode == "" {
		result = "leave_code is mandatory"
		return errCreate, result
	}

	if req.LeaveStartdate == "" {
		result = "leave_startdate is mandatory"
		return errCreate, result
	}

	if req.LeaveEnddate == "" {
		result = "leave_enddate is mandatory"
		return errCreate, result
	}

	if req.RequestNo == "" {
		result = "request_no is mandatory"
		return errCreate, result
	}

	if req.Requestfor == "" {
		result = "requestfor is mandatory"
		return errCreate, result
	}

	queryCreate = `INSERT INTO dbSF6_QA.dbo.TTADLEAVEREQUEST
		(request_no, company_id, requestedby, requestfor, requestdate, 
		leave_code, leave_startdate, leave_enddate, usecalendar, totaldays, 
		remark, created_by, created_date, refdoc, reqfullday, 
		approval_status, hdtype_starttime, hdtype_endtime, leave_start_halfday, leave_end_halfday)
		VALUES
		(?, ?, ?, ?, getdate(), 
		?, ?, ?, 'N', 1.0000, 
		NULL, N'shiburin1988', '2018-05-14 14:47:43.000', NULL, N'Y', 
		NULL, 0, 0, NULL, NULL);`
	queryCreate = repo.dbMaster.Rebind(queryCreate)

	tx, errCreate := repo.dbMaster.Begin()
	if errCreate != nil {
		return errCreate, result
	}
	createRequest, errCreate := tx.Prepare(queryCreate)
	if errCreate != nil {
		return errCreate, result
	}
	defer createRequest.Close()

	_, errCreate = createRequest.Exec()
	if errCreate != nil {
		return errCreate, result
	}
	tx.Commit()
	return errCreate, result
}

func (repo *repo) GetLeaveRequestFilterListing(ctx context.Context, req model.GetLeaveRequestListingFilterRequest) (error, []model.GetLeaveRequestListingFilterResponse) {
	var (
		dataLeaveRequest  []model.GetLeaveRequestListingFilterResponse
		errData           error
		queryListing      string
		jumlahSudahTampil int64
		paramData         []interface{}
	)

	if req.Language == "" {
		req.Language = "en"
	}

	if req.Limit < 1 {
		req.Limit = 50
	}
	jumlahSudahTampil = 0
	if req.Page <= 1 {
		req.Page = 1
	} else {
		jumlahSudahTampil = (req.Page - 1) * req.Limit
	}

	paramData = append(paramData, req.Limit)
	paramData = append(paramData, req.CompanyId)

	queryListing = `SELECT TOP(CAST(? AS INT)) request_no, TTADLEAVEREQUEST.company_id
	, full_name requestfor
	, ISNULL(refdoc,'-') refdoc
	,leave_startdate, leave_enddate, totaldays, leavename_` + req.Language + ` as leave_code,
	TTADLEAVEREQUEST.remark, TGEMREQSTATUS.name_` + req.Language + ` request_status, TTADLEAVEREQUEST.reqfullday
	FROM 	TTADLEAVEREQUEST, TTAMLEAVETYPE, TEOMEMPPERSONAL, TEODEmpCompany EC, TCLTREQUEST RQ, TGEMREQSTATUS
	WHERE	TTADLEAVEREQUEST.leave_code = TTAMLEAVETYPE.leave_code AND TTADLEAVEREQUEST.company_id = TTAMLEAVETYPE.company_id
	AND	TEOMEMPPERSONAL.emp_id = TTADLEAVEREQUEST.requestfor
	AND TEOMEMPPERSONAL.emp_id = EC.emp_id 
	AND TTADLEAVEREQUEST.company_id=EC.company_id 
	AND EC.is_main=1
	AND RQ.status=TGEMREQSTATUS.code
	AND TTADLEAVEREQUEST.request_no = RQ.req_no
	AND TTADLEAVEREQUEST.company_id = ?`

	if jumlahSudahTampil > 0 {
		paramData = append(paramData, jumlahSudahTampil)
		paramData = append(paramData, req.EmployeeId)

		queryListing = queryListing + ` AND TTADLEAVEREQUEST.request_no NOT IN (SELECT TOP(CAST(? AS INT)) TTADLEAVEREQUEST.request_no
		FROM 	TTADLEAVEREQUEST, TTAMLEAVETYPE, TEOMEMPPERSONAL, TEODEmpCompany EC, TCLTREQUEST RQ, TGEMREQSTATUS
		WHERE	TTADLEAVEREQUEST.leave_code = TTAMLEAVETYPE.leave_code 
		AND TTADLEAVEREQUEST.company_id = TTAMLEAVETYPE.company_id
		AND	TEOMEMPPERSONAL.emp_id = TTADLEAVEREQUEST.requestfor
		AND TEOMEMPPERSONAL.emp_id = EC.emp_id 
		AND TTADLEAVEREQUEST.company_id=EC.company_id 
		AND EC.is_main=1
		AND RQ.status=TGEMREQSTATUS.code
		AND TTADLEAVEREQUEST.request_no = RQ.req_no
		AND TTADLEAVEREQUEST.company_id = ?`

		// if req.LeaveCode != "" {
		// 	queryListing = queryListing + ` AND TTAMLEAVETYPE.leavename_` + req.Language + ` LIKE ?`
		// 	paramData = append(paramData, fmt.Sprintf(`%%%s%%`, req.LeaveCode))
		// }
		if len(req.FilterLeaveCode) > 0 {
			query, args, _ := sqlx.In(` AND TTAMLEAVETYPE.leave_code IN (?)`, req.FilterLeaveCode)
			queryListing = queryListing + query

			for _, argss := range args {
				paramData = append(paramData, argss)
			}
		}

		if len(req.FilterRequestStatus) > 0 {
			query, args, _ := sqlx.In(` AND TGEMREQSTATUS.name_`+req.Language+` IN (?)`, req.FilterRequestStatus)
			queryListing = queryListing + query

			for _, argss := range args {
				paramData = append(paramData, argss)
			}
		}

		// if req.RequestStatus != "" {
		// 	queryListing = queryListing + ` AND TGEMREQSTATUS.name_` + req.Language + ` LIKE ?`
		// 	paramData = append(paramData, fmt.Sprintf(`%%%s%%`, req.RequestStatus))
		// }

		if req.FilterLeaveStartdate != "" {
			queryListing = queryListing + ` AND leave_startdate >= ?`
			paramData = append(paramData, req.FilterLeaveStartdate)
		}

		if req.FilterLeaveEnddate != "" {
			queryListing = queryListing + ` AND leave_enddate <= ?`
			paramData = append(paramData, req.FilterLeaveEnddate)
		}

		queryListing = queryListing + ` AND (TTADLEAVEREQUEST.requestfor = ? OR TTADLEAVEREQUEST.requestedby= ? )`
		paramData = append(paramData, req.EmployeeId)
		paramData = append(paramData, req.EmployeeId)
		queryListing = queryListing + ` ORDER BY TTADLEAVEREQUEST.request_no ASC)`
	}

	// if req.LeaveCode != "" {
	// 	queryListing = queryListing + ` AND TTAMLEAVETYPE.leavename_` + req.Language + ` LIKE ?`
	// 	paramData = append(paramData, fmt.Sprintf(`%%%s%%`, req.LeaveCode))
	// }

	if len(req.FilterLeaveCode) > 0 {
		query, args, _ := sqlx.In(` AND TTAMLEAVETYPE.leave_code IN (?)`, req.FilterLeaveCode)
		queryListing = queryListing + query

		for _, argss := range args {
			paramData = append(paramData, argss)
		}
	}

	// if req.RequestStatus != "" {
	// 	queryListing = queryListing + ` AND TGEMREQSTATUS.name_` + req.Language + ` LIKE ?`
	// 	paramData = append(paramData, fmt.Sprintf(`%%%s%%`, req.RequestStatus))
	// }

	if len(req.FilterRequestStatus) > 0 {
		query, args, _ := sqlx.In(` AND TGEMREQSTATUS.name_`+req.Language+` IN (?)`, req.FilterRequestStatus)
		queryListing = queryListing + query

		for _, argss := range args {
			paramData = append(paramData, argss)
		}
	}

	if req.FilterLeaveStartdate != "" {
		queryListing = queryListing + ` AND leave_startdate >= ?`
		paramData = append(paramData, req.FilterLeaveStartdate)
	}

	if req.FilterLeaveEnddate != "" {
		queryListing = queryListing + ` AND leave_enddate <= ?`
		paramData = append(paramData, req.FilterLeaveEnddate)
	}

	queryListing = queryListing + ` AND (TTADLEAVEREQUEST.requestfor = ? OR TTADLEAVEREQUEST.requestedby= ? ) 
	ORDER BY TTADLEAVEREQUEST.request_no ASC`
	paramData = append(paramData, req.EmployeeId)
	paramData = append(paramData, req.EmployeeId)

	queryListing = repo.dbSlave.Rebind(queryListing)
	res1, errData := repo.dbSlave.Queryx(queryListing, paramData...)
	// fmt.Print(queryListing)
	// if configs.FancyHandleError(errData) {
	// 	logger.Println("Query Select GetListingLeaveRequest: Error: ", errData.Error())
	// 	logger.Println(queryListing)
	// 	return errData, dataLeaveRequest
	// }

	if errData != nil {
		logger.Error(nil, errData)
		// logger.Println(queryListing)
		return errData, dataLeaveRequest
	}

	defer res1.Close()

	if res1.Next() {
		var temp model.GetLeaveRequestListingFilterResponse
		errData := res1.Scan(&temp.RequestNo, &temp.CompanyId, &temp.Requestfor, &temp.Refdoc, &temp.LeaveStartdate,
			&temp.LeaveEnddate, &temp.Totaldays, &temp.LeaveCode, &temp.Remark, &temp.RequestStatus,
			&temp.Reqfullday)
		// if configs.FancyHandleError(errData) {
		// 	logger.Println("Scan Query Select GetListingLeaveRequest: Error: " + errData.Error())
		// 	logger.Println(queryListing)
		// 	res1.Close()
		// 	return errData, dataLeaveRequest
		// }

		if errData != nil {
			logger.Error(nil, errData)
			// logger.Println(queryListing)
			return errData, dataLeaveRequest
		}

		dataLeaveRequest = append(dataLeaveRequest, temp)
		for res1.Next() {
			errData := res1.Scan(&temp.RequestNo, &temp.CompanyId, &temp.Requestfor, &temp.Refdoc, &temp.LeaveStartdate,
				&temp.LeaveEnddate, &temp.Totaldays, &temp.LeaveCode, &temp.Remark, &temp.RequestStatus,
				&temp.Reqfullday)
			// if configs.FancyHandleError(errData) {
			// 	logger.Println("Scan Query Select GetListingLeaveRequest: Error: " + errData.Error())
			// 	logger.Println(queryListing)
			// 	res1.Close()
			// 	return errData, dataLeaveRequest
			// }
			if errData != nil {
				logger.Error(nil, errData)
				// logger.Println(queryListing)
				return errData, dataLeaveRequest
			}

			dataLeaveRequest = append(dataLeaveRequest, temp)
		}
	}
	return errData, dataLeaveRequest
}

func (repo *repo) GetLeaveRequestListing(ctx context.Context, req model.GetLeaveRequestListingRequest) (error, []model.GetLeaveRequestListingResponse) {
	var (
		dataLeaveRequest  []model.GetLeaveRequestListingResponse
		errData           error
		queryListing      string
		jumlahSudahTampil int64
		paramData         []interface{}
	)

	if req.Language == "" {
		req.Language = "en"
	}

	if req.Limit < 1 {
		req.Limit = 50
	}
	jumlahSudahTampil = 0
	if req.Page <= 1 {
		req.Page = 1
	} else {
		jumlahSudahTampil = (req.Page - 1) * req.Limit
	}

	if req.Field == "request_no" {
		req.Field = "TTADLEAVEREQUEST.request_no"
	} else if req.Field == "requestfor" {
		req.Field = "TEOMEMPPERSONAL.full_name"
	} else if req.Field == "leave_code" {
		req.Field = "leavename_en"
	} else if req.Field == "leave_startdate" {
		req.Field = "leave_startdate"
	} else if req.Field == "leave_enddate" {
		req.Field = "leave_enddate"
	} else if req.Field == "totaldays" {
		req.Field = "totaldays"
	} else if req.Field == "remark" {
		req.Field = "TTADLEAVEREQUEST.remark"
	} else if req.Field == "request_status" {
		req.Field = "TGEMREQSTATUS.name_en"
	} else {
		req.Field = "TTADLEAVEREQUEST.request_no"
	}

	if strings.ToUpper(req.Order) == "ASC" || strings.ToUpper(req.Order) == "DESC" {
		req.Order = strings.ToUpper(req.Order)
	} else {
		req.Order = "DESC"
	}

	paramData = append(paramData, req.Limit)
	paramData = append(paramData, req.CompanyId)

	queryListing = `SELECT TOP(CAST(? AS INT)) request_no, TTADLEAVEREQUEST.company_id
	, TEOMEMPPERSONAL.full_name requestfor
	, ISNULL(refdoc,'-') refdoc
	,leave_startdate, leave_enddate, totaldays, leavename_` + req.Language + ` as leave_code,
	TTADLEAVEREQUEST.remark, TGEMREQSTATUS.name_` + req.Language + ` request_status, TTADLEAVEREQUEST.reqfullday
	FROM 	TTADLEAVEREQUEST, TTAMLEAVETYPE, TEOMEMPPERSONAL, TEODEmpCompany EC, TCLTREQUEST RQ, TGEMREQSTATUS
	WHERE	TTADLEAVEREQUEST.leave_code = TTAMLEAVETYPE.leave_code AND TTADLEAVEREQUEST.company_id = TTAMLEAVETYPE.company_id
	AND	TEOMEMPPERSONAL.emp_id = TTADLEAVEREQUEST.requestfor
	AND TEOMEMPPERSONAL.emp_id = EC.emp_id 
	AND TTADLEAVEREQUEST.company_id=EC.company_id 
	AND EC.is_main=1
	AND RQ.status=TGEMREQSTATUS.code
	AND TTADLEAVEREQUEST.request_no = RQ.req_no
	AND TTADLEAVEREQUEST.company_id = ?`

	if jumlahSudahTampil > 0 {
		paramData = append(paramData, jumlahSudahTampil)
		paramData = append(paramData, req.CompanyId)

		queryListing = queryListing + ` AND TTADLEAVEREQUEST.request_no NOT IN (SELECT TOP(CAST(? AS INT)) TTADLEAVEREQUEST.request_no
		FROM 	TTADLEAVEREQUEST, TTAMLEAVETYPE, TEOMEMPPERSONAL, TEODEmpCompany EC, TCLTREQUEST RQ, TGEMREQSTATUS
		WHERE	TTADLEAVEREQUEST.leave_code = TTAMLEAVETYPE.leave_code 
		AND TTADLEAVEREQUEST.company_id = TTAMLEAVETYPE.company_id
		AND	TEOMEMPPERSONAL.emp_id = TTADLEAVEREQUEST.requestfor
		AND TEOMEMPPERSONAL.emp_id = EC.emp_id 
		AND TTADLEAVEREQUEST.company_id=EC.company_id 
		AND EC.is_main=1
		AND RQ.status=TGEMREQSTATUS.code
		AND TTADLEAVEREQUEST.request_no = RQ.req_no
		AND TTADLEAVEREQUEST.company_id = ?`

		if len(req.FilterLeaveCode) > 0 {
			query, args, _ := sqlx.In(` AND TTAMLEAVETYPE.leave_code IN (?)`, req.FilterLeaveCode)
			queryListing = queryListing + query

			for _, argss := range args {
				paramData = append(paramData, argss)
			}
		}

		if len(req.FilterRequestStatus) > 0 {
			query, args, _ := sqlx.In(` AND TGEMREQSTATUS.name_`+req.Language+` IN (?)`, req.FilterRequestStatus)
			queryListing = queryListing + query

			for _, argss := range args {
				paramData = append(paramData, argss)
			}
		}

		if req.FilterLeaveStartdate != "" {
			queryListing = queryListing + ` AND leave_startdate >= ?`
			paramData = append(paramData, req.FilterLeaveStartdate)
		}

		if req.FilterLeaveEnddate != "" {
			queryListing = queryListing + ` AND leave_enddate <= ?`
			paramData = append(paramData, req.FilterLeaveEnddate)
		}

		if req.Totaldays != "" {
			queryListing = queryListing + ` AND totaldays = ?`
			paramData = append(paramData, req.Totaldays)
		}

		if req.Remark != "" {
			queryListing = queryListing + ` AND remark LIKE ?`
			paramData = append(paramData, fmt.Sprintf(`%%%s%%`, req.Remark))
		}

		if req.RequestNo != "" {
			queryListing = queryListing + ` AND request_no LIKE ?`
			paramData = append(paramData, fmt.Sprintf(`%%%s%%`, req.RequestNo))
		}

		if req.LeaveCode != "" {
			queryListing = queryListing + ` AND leavename_` + req.Language + ` LIKE ?`
			paramData = append(paramData, fmt.Sprintf(`%%%s%%`, req.LeaveCode))
		}

		if req.Requestfor != "" {
			queryListing = queryListing + ` AND TEOMEMPPERSONAL.full_name LIKE ?`
			paramData = append(paramData, fmt.Sprintf(`%%%s%%`, req.Requestfor))
		}

		if req.RequestStatus != "" {
			queryListing = queryListing + ` AND TGEMREQSTATUS.name_` + req.Language + ` LIKE ?`
			paramData = append(paramData, fmt.Sprintf(`%%%s%%`, req.RequestStatus))
		}

		if req.FilterLeaveStartdate == "" {
			if req.LeaveStartdate != "" {
				queryListing = queryListing + ` AND leave_startdate >= ?`
				paramData = append(paramData, req.LeaveStartdate)
				// , leave_enddate
			}
		}

		if req.FilterLeaveEnddate == "" {
			if req.LeaveEnddate != "" {
				queryListing = queryListing + ` AND leave_enddate <= ?`
				paramData = append(paramData, req.LeaveEnddate)
			}
		}

		queryListing = queryListing + ` AND (TTADLEAVEREQUEST.requestfor = ? OR TTADLEAVEREQUEST.requestedby= ? )`
		paramData = append(paramData, req.EmployeeId)
		paramData = append(paramData, req.EmployeeId)
		queryListing = queryListing + ` ORDER BY ` + req.Field + ` ` + req.Order + `)`
	}

	if len(req.FilterLeaveCode) > 0 {
		query, args, _ := sqlx.In(` AND TTAMLEAVETYPE.leave_code IN (?)`, req.FilterLeaveCode)
		queryListing = queryListing + query

		for _, argss := range args {
			paramData = append(paramData, argss)
		}
	}

	if len(req.FilterRequestStatus) > 0 {
		query, args, _ := sqlx.In(` AND TGEMREQSTATUS.name_`+req.Language+` IN (?)`, req.FilterRequestStatus)
		queryListing = queryListing + query

		for _, argss := range args {
			paramData = append(paramData, argss)
		}
	}

	if req.FilterLeaveStartdate != "" {
		queryListing = queryListing + ` AND leave_startdate >= ?`
		paramData = append(paramData, req.FilterLeaveStartdate)
	}

	if req.FilterLeaveEnddate != "" {
		queryListing = queryListing + ` AND leave_enddate <= ?`
		paramData = append(paramData, req.FilterLeaveEnddate)
	}

	if req.Totaldays != "" {
		queryListing = queryListing + ` AND totaldays = ?`
		paramData = append(paramData, req.Totaldays)
	}

	if req.Remark != "" {
		queryListing = queryListing + ` AND remark LIKE ?`
		paramData = append(paramData, fmt.Sprintf(`%%%s%%`, req.Remark))
	}

	if req.RequestNo != "" {
		queryListing = queryListing + ` AND request_no LIKE ?`
		paramData = append(paramData, fmt.Sprintf(`%%%s%%`, req.RequestNo))
	}

	if req.LeaveCode != "" {
		queryListing = queryListing + ` AND leavename_` + req.Language + ` LIKE ?`
		paramData = append(paramData, fmt.Sprintf(`%%%s%%`, req.LeaveCode))
	}

	if req.Requestfor != "" {
		queryListing = queryListing + ` AND TEOMEMPPERSONAL.full_name LIKE ?`
		paramData = append(paramData, fmt.Sprintf(`%%%s%%`, req.Requestfor))
	}

	if req.RequestStatus != "" {
		queryListing = queryListing + ` AND TGEMREQSTATUS.name_` + req.Language + ` LIKE ?`
		paramData = append(paramData, fmt.Sprintf(`%%%s%%`, req.RequestStatus))
	}

	if req.FilterLeaveStartdate == "" {
		if req.LeaveStartdate != "" {
			queryListing = queryListing + ` AND leave_startdate >= ?`
			paramData = append(paramData, req.LeaveStartdate)
		}
	}

	if req.FilterLeaveEnddate == "" {
		if req.LeaveEnddate != "" {
			queryListing = queryListing + ` AND leave_enddate <= ?`
			paramData = append(paramData, req.LeaveEnddate)
		}
	}

	queryListing = queryListing + ` AND (TTADLEAVEREQUEST.requestfor = ? OR TTADLEAVEREQUEST.requestedby= ? ) 
	ORDER BY ` + req.Field + ` ` + req.Order
	paramData = append(paramData, req.EmployeeId)
	paramData = append(paramData, req.EmployeeId)

	queryListing = repo.dbSlave.Rebind(queryListing)
	res1, errData := repo.dbSlave.Queryx(queryListing, paramData...)
	// fmt.Print(queryListing)
	if errData != nil {
		logger.Error(nil, errData)
		// logger.Println(queryListing)
		return errData, dataLeaveRequest
	}

	defer res1.Close()

	if res1.Next() {
		var temp model.GetLeaveRequestListingResponse
		errData := res1.Scan(&temp.RequestNo, &temp.CompanyId, &temp.Requestfor, &temp.Refdoc, &temp.LeaveStartdate,
			&temp.LeaveEnddate, &temp.Totaldays, &temp.LeaveCode, &temp.Remark, &temp.RequestStatus,
			&temp.Reqfullday)
		if errData != nil {
			logger.Error(nil, errData)
			// logger.Println(queryListing)
			res1.Close()
			return errData, dataLeaveRequest
		}

		dataLeaveRequest = append(dataLeaveRequest, temp)
		for res1.Next() {
			errData := res1.Scan(&temp.RequestNo, &temp.CompanyId, &temp.Requestfor, &temp.Refdoc, &temp.LeaveStartdate,
				&temp.LeaveEnddate, &temp.Totaldays, &temp.LeaveCode, &temp.Remark, &temp.RequestStatus,
				&temp.Reqfullday)
			if errData != nil {
				logger.Error(nil, errData)
				// logger.Println(queryListing)
				res1.Close()
				return errData, dataLeaveRequest
			}

			dataLeaveRequest = append(dataLeaveRequest, temp)
		}
	}
	return errData, dataLeaveRequest
}

func (repo *repo) GetDataTypeOfLeave(ctx context.Context, req model.GetDataTypeOfLeaveReq) (error, []model.GetDataTypeOfLeaveResponse) {
	var (
		dataLeaveRequest     []model.GetDataTypeOfLeaveResponse
		errData              error
		queryDataTypeOfLeave string
		paramData            []interface{}
	)

	if req.Language == "" {
		req.Language = "en"
	}

	if req.CompanyId == nil {
		errData = errors.New("company_id is mandatory")
		return errData, dataLeaveRequest
	}

	if strings.Trim(req.EmployeeId, " ") == "" {
		errData = errors.New("employee_id is mandatory")
		return errData, dataLeaveRequest
	}

	paramData = append(paramData, req.CompanyId)
	paramData = append(paramData, req.EmployeeId)
	paramData = append(paramData, req.CompanyId)
	paramData = append(paramData, req.EmployeeId)

	queryDataTypeOfLeave = `SELECT DISTINCT TTADEMPGETLEAVE.leave_code+'|'+daytype optvalue,leavename_` + req.Language + ` opttext, TTADEMPGETLEAVE.emp_id optempid 
							FROM TTAMLEAVETYPE, TTADEMPGETLEAVE
							WHERE 
								(TTAMLEAVETYPE.graceperiod = 0
								AND TTAMLEAVETYPE.leave_code = TTADEMPGETLEAVE.leave_code AND TTAMLEAVETYPE.company_id = TTADEMPGETLEAVE.company_id
								AND TTADEMPGETLEAVE.company_id = ?
								AND TTADEMPGETLEAVE.active_status = 1
								AND TTADEMPGETLEAVE.emp_id = ?
								AND (
									endvaliddate >= getdate()
									OR
									endvaliddate IS NULL
								))
								OR
								(TTAMLEAVETYPE.graceperiod > 0
								AND TTAMLEAVETYPE.leave_code = TTADEMPGETLEAVE.leave_code AND TTAMLEAVETYPE.company_id = TTADEMPGETLEAVE.company_id
								AND TTADEMPGETLEAVE.company_id = ?
								AND TTADEMPGETLEAVE.emp_id = ?
								AND (
									dateadd(day, graceperiod , endvaliddate) >= getdate()
									OR
									endvaliddate IS NULL
								))
							ORDER BY leavename_` + req.Language + ``

	queryDataTypeOfLeave = repo.dbSlave.Rebind(queryDataTypeOfLeave)
	res1, errData := repo.dbSlave.Queryx(queryDataTypeOfLeave, paramData...)
	// fmt.Print(queryListing)
	if errData != nil {
		logger.Error(nil, errData)
		// logger.Println(queryListing)
		return errData, dataLeaveRequest
	}

	defer res1.Close()

	if res1.Next() {
		var temp model.GetDataTypeOfLeaveResponse
		errData := res1.Scan(&temp.Optvalue, &temp.Opttext, &temp.Optempid)
		if errData != nil {
			logger.Error(nil, errData)
			// logger.Println(queryListing)
			res1.Close()
			return errData, dataLeaveRequest
		}

		dataLeaveRequest = append(dataLeaveRequest, temp)
		for res1.Next() {
			errData := res1.Scan(&temp.Optvalue, &temp.Opttext, &temp.Optempid)
			if errData != nil {
				logger.Error(nil, errData)
				// logger.Println(queryListing)
				res1.Close()
				return errData, dataLeaveRequest
			}

			dataLeaveRequest = append(dataLeaveRequest, temp)
		}
	}
	return errData, dataLeaveRequest
}
