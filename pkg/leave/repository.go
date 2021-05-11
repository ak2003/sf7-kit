package leave

import (
	"context"
	"fmt"

	"gitlab.dataon.com/gophers/sf7-kit/pkg/leave/model"
	"gitlab.dataon.com/gophers/sf7-kit/shared/utils/logger"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetLeaveRequestListing(ctx context.Context, sc model.GetLeaveRequestListingRequest) (error, []model.GetLeaveRequestListingResponse)
	GetLeaveRequestFilterListing(ctx context.Context, sc model.GetLeaveRequestListingFilterRequest) (error, []model.GetLeaveRequestListingFilterResponse)
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

		if req.FilterLeaveEnddate == "" {
			if req.LeaveEnddate != "" {
				queryListing = queryListing + ` AND leave_startdate >= ?`
				paramData = append(paramData, req.LeaveEnddate)
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
		queryListing = queryListing + ` ORDER BY TTADLEAVEREQUEST.request_no ASC)`
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
	ORDER BY TTADLEAVEREQUEST.request_no ASC`
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
