package employee

import (
	"context"
	"fmt"
	"strings"
	"time"

	"gitlab.dataon.com/gophers/sf7-kit/pkg/employee/model"
	"gitlab.dataon.com/gophers/sf7-kit/shared/utils/logger"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetEmployeeInformation(ctx context.Context, sc model.GetEmployeeInformationRequest) (error, []model.GetEmployeeInformationResponse)
	GetEmployeeEditInformation(ctx context.Context, sc model.GetEmployeeByIdRequest) (error, []model.GetEmployeeByIdResponse)
	GetEmployeeMasterAddress(ctx context.Context, sc model.GetEmployeeMasterAddressRequest) (error, []model.GetEmployeeMasterAddressResponse)
	UpdateEmployeeMasterAddress(ctx context.Context, sc model.UpdateEmployeeMasterAddressRequest) (error, string)
	CreateEmployeeMasterAddress(ctx context.Context, sc model.CreateEmployeeMasterAddressRequest) (error, string)
	GetCity(ctx context.Context, sc model.GetCityRequest) (error, []model.GetCityResponse)
	GetAddressType(ctx context.Context, sc model.GetAddressTypeRequest) (error, []model.GetAddressTypeResponse)
	GetOwnerStatus(ctx context.Context, sc model.GetOwnerStatusRequest) (error, []model.GetOwnerStatusResponse)
	GetStayStatus(ctx context.Context, sc model.GetStayStatusRequest) (error, []model.GetStayStatusResponse)
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

func (repo *repo) CreateEmployeeMasterAddress(ctx context.Context, req model.CreateEmployeeMasterAddressRequest) (error, string) {
	var (
		result      string
		errData     error
		queryCreate string
	)
	result = "OK"

	if strings.Trim(req.EmployeeId, " ") == "" {
		result = "employee_id is mandatory"
		return errData, result
	}

	if strings.Trim(req.EmployeeAddressType, " ") == "" {
		result = "employee_address_type is mandatory"
		return errData, result
	}

	queryCreate = `INSERT INTO dbSF6_QA.dbo.TEODEMPADDRESS
	(emp_id, addresstype_code, address, rt, rw, 
	subdistrict, district, city_id, state_id, country_id, 
	zipcode, phone, living_status, owner_status, address_distance, 
	created_date, created_by, lat_lng, local_address, modified_date,
	modified_by)
	VALUES( ?, ?, ?, ?, ?, 
			?, ?, ?, ?, ?, 
			?, ?, ?, ?, ?, 
			getdate(), ?, ?, ?, getdate(), 
			?);`
	queryCreate = repo.dbMaster.Rebind(queryCreate)
	_, errData = repo.dbMaster.Exec(queryCreate,
		req.EmployeeId, req.EmployeeAddressType, req.EmployeeAddress, req.EmployeeAddressRt, req.EmployeeAddressRw,
		req.EmployeeAddressSubdistrict, req.EmployeeAddressDistrict, req.EmployeeAddressCityId, req.EmployeeAddressStateId, req.EmployeeAddressCountryId,
		req.EmployeeAddressZipcode, req.EmployeeAddressPhone, req.EmployeeAddressLivingStatus, req.EmployeeAddressOwnerStatus, req.EmployeeAddressDistance,
		req.Username, req.EmployeeAddressLatLong, req.EmployeeAddressLocal,
		req.Username)

	if errData != nil {
		// fmt.Println(queryUpdate)
		return errData, errData.Error()
	}
	return errData, result
}

func (repo *repo) UpdateEmployeeMasterAddress(ctx context.Context, req model.UpdateEmployeeMasterAddressRequest) (error, string) {
	var (
		result      string
		errData     error
		queryUpdate string
	)
	result = "OK"

	if strings.Trim(req.EmployeeId, " ") == "" {
		result = "employee_id is mandatory"
		return errData, result
	}

	if strings.Trim(req.EmployeeAddressType, " ") == "" {
		result = "employee_address_type is mandatory"
		return errData, result
	}
	//
	//
	//
	queryUpdate = `UPDATE dbSF6_QA.dbo.TEODEMPADDRESS
					SET address = ?, rt = ?, rw = ?, subdistrict = ?, district = ?, 
					city_id = ?, state_id = ?, country_id = ?, zipcode = ?, phone=?,
					living_status = ?, owner_status = ?, address_distance = ?, modified_date = getdate(), modified_by = ?, 
					lat_lng = ?, local_address = ?
					WHERE emp_id = ? AND addresstype_code = ? ;`

	queryUpdate = repo.dbMaster.Rebind(queryUpdate)
	_, errData = repo.dbMaster.Exec(queryUpdate, req.EmployeeAddress, req.EmployeeAddressRt, req.EmployeeAddressRw, req.EmployeeAddressSubdistrict, req.EmployeeAddressDistrict,
		req.EmployeeAddressCityId, req.EmployeeAddressStateId, req.EmployeeAddressCountryId, req.EmployeeAddressZipcode, req.EmployeeAddressPhone,
		req.EmployeeAddressLivingStatus, req.EmployeeAddressOwnerStatus, req.EmployeeAddressDistance, req.Username,
		req.EmployeeAddressLatLong, req.EmployeeAddressLocal,
		req.EmployeeId, req.EmployeeAddressType)

	if errData != nil {
		// fmt.Println(queryUpdate)
		return errData, errData.Error()
	}
	return errData, result
}

func (repo *repo) GetEmployeeMasterAddress(ctx context.Context, req model.GetEmployeeMasterAddressRequest) (error, []model.GetEmployeeMasterAddressResponse) {
	var (
		dataEmployeeAddress []model.GetEmployeeMasterAddressResponse
		errData             error
		queryGetData        string
		paramData           []interface{}
	)

	paramData = append(paramData, req.EmployeeId)

	queryGetData = `SELECT adr.emp_id, adr.addresstype_code, adr.address, adr.rt, adr.rw, 
					adr.subdistrict, adr.district, adr.city_id, adr.state_id, adr.country_id,
					adr.zipcode, adr.phone, adr.living_status, adr.owner_status, adr.address_distance, 
					adr.lat_lng, adr.local_address, ct.city_name, st.state_name
					FROM dbSF6_QA.dbo.TEODEMPADDRESS adr
					LEFT JOIN TGEMCITY ct ON  ct.city_id = adr.city_id
					LEFT JOIN TGEMSTATE st ON st.state_id = adr.state_id
					WHERE adr.emp_id = ?;`

	queryGetData = repo.dbSlave.Rebind(queryGetData)
	res1, errData := repo.dbSlave.Queryx(queryGetData, paramData...)

	if errData != nil {
		logger.Error(nil, errData)
		// logger.Println(queryListing)
		return errData, dataEmployeeAddress
	}
	defer res1.Close()

	if res1.Next() {
		var temp model.GetEmployeeMasterAddressResponse

		errData := res1.Scan(&temp.EmployeeId, &temp.EmployeeAddressType, &temp.EmployeeAddress, &temp.EmployeeAddressRt, &temp.EmployeeAddressRw,
			&temp.EmployeeAddressSubdistrict, &temp.EmployeeAddressDistrict, &temp.EmployeeAddressCityId, &temp.EmployeeAddressStateId, &temp.EmployeeAddressCountryId,
			&temp.EmployeeAddressZipcode, &temp.EmployeeAddressPhone, &temp.EmployeeAddressLivingStatus, &temp.EmployeeAddressOwnerStatus, &temp.EmployeeAddressDistance,
			&temp.EmployeeAddressLatLong, &temp.EmployeeAddressLocal, &temp.EmployeeAddressCityName, &temp.EmployeeAddressStateName)
		if errData != nil {
			logger.Error(nil, errData)
			// logger.Println(queryListing)
			res1.Close()
			return errData, dataEmployeeAddress
		}

		dataEmployeeAddress = append(dataEmployeeAddress, temp)
	}
	return errData, dataEmployeeAddress
}

func (repo *repo) GetEmployeeEditInformation(ctx context.Context, req model.GetEmployeeByIdRequest) (error, []model.GetEmployeeByIdResponse) {
	var (
		dataEmployeeInfo []model.GetEmployeeByIdResponse
		errData          error
		queryGetData     string
		paramData        []interface{}
	)
	if req.Language == "" {
		req.Language = "en"
	}

	paramData = append(paramData, req.CompanyId)
	paramData = append(paramData, req.EmployeeId)

	queryGetData = `SELECT E.emp_id, E.taxno, E.first_name, E.middle_name, E.last_name,
					E.gender, D.birthplace, D.birthdate, E.photo, D.phone,
					E.email, D.maritalstatus, EC.start_date, EC.emp_no, EC.status,
					EC.is_main, EC.company_id, EC.position_id, EC.employ_code emp_status, EC.grade_code job_grade,
					full_name emp_name, P.pos_name_` + req.Language + ` pos_name, year(birthdate)year_birth, month(birthdate)month_birth, DEPT.pos_name_` + req.Language + ` dept_name,
					phone_ext FROM TEOMEmpPersonal E INNER JOIN TEODEMPCOMPANY EC ON E.emp_id = EC.emp_id AND EC.company_id = ? AND is_main = 1
					LEFT JOIN TEODEMPPERSONAL D ON E.emp_id = D.emp_id 
					LEFT JOIN TEOMPosition P ON EC.position_id = P.position_id AND P.company_id = EC.company_id 
					LEFT OUTER JOIN TEOMPosition DEPT ON P.dept_id = DEPT.position_id AND DEPT.company_id = P.company_id 
					WHERE E.emp_id = ?;`

	queryGetData = repo.dbSlave.Rebind(queryGetData)
	res1, errData := repo.dbSlave.Queryx(queryGetData, paramData...)

	if errData != nil {
		logger.Error(nil, errData)
		// logger.Println(queryListing)
		return errData, dataEmployeeInfo
	}

	defer res1.Close()

	if res1.Next() {
		var temp model.GetEmployeeByIdResponse
		errData := res1.Scan(&temp.EmployeeId, &temp.EmployeeTaxNo, &temp.EmployeeFirstName, &temp.EmployeeMiddleName, &temp.EmployeeLastName,
			&temp.EmployeeGender, &temp.EmployeeBirthplace, &temp.EmployeeBirthdate, &temp.EmployeePhoto, &temp.EmployeePhone,
			&temp.EmployeeEmail, &temp.EmployeeMaritalStatus, &temp.EmployeeStartDate, &temp.EmployeeNo, &temp.EmployeeStatus,
			&temp.EmployeeIsMain, &temp.EmployeeCompanyId, &temp.EmployeePositionId, &temp.EmployeeCodeStatus, &temp.EmployeeGrade,
			&temp.EmployeeFullName, &temp.EmployeePosName, &temp.EmployeeYearBirth, &temp.EmployeeMonthBirth, &temp.EmployeeDeptName,
			&temp.EmployeePhoneExt)
		if errData != nil {
			logger.Error(nil, errData)
			// logger.Println(queryListing)
			res1.Close()
			return errData, dataEmployeeInfo
		}

		dataEmployeeInfo = append(dataEmployeeInfo, temp)
	}
	return errData, dataEmployeeInfo
}

func (repo *repo) GetEmployeeInformation(ctx context.Context, req model.GetEmployeeInformationRequest) (error, []model.GetEmployeeInformationResponse) {
	var (
		dataEmployeeInfo  []model.GetEmployeeInformationResponse
		errData           error
		queryListing      string
		jumlahSudahTampil int64
		paramData         []interface{}
	)

	TglSekarang := time.Now()
	TglSekarangStr := TglSekarang.Format("2006-01-02 15:04:05")

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
	paramData = append(paramData, req.UserId)
	paramData = append(paramData, req.EmployeeId)
	paramData = append(paramData, fmt.Sprintf(`%%,%s,%%`, req.UserId))
	paramData = append(paramData, fmt.Sprintf(`%%,%s,%%`, req.UserId))
	paramData = append(paramData, req.EmployeeId)
	paramData = append(paramData, req.CompanyId)
	paramData = append(paramData, req.CompanyId)

	paramData = append(paramData, req.CompanyId)
	paramData = append(paramData, req.CompanyId)
	paramData = append(paramData, req.CompanyId)
	paramData = append(paramData, req.CompanyId)
	paramData = append(paramData, req.CompanyId)
	paramData = append(paramData, fmt.Sprintf(`%%,%s,%%`, req.EmployeeId))
	paramData = append(paramData, req.CompanyId)
	paramData = append(paramData, fmt.Sprintf(`%%,%s,%%`, req.EmployeeId))
	paramData = append(paramData, req.CompanyId)
	paramData = append(paramData, TglSekarangStr)
	paramData = append(paramData, req.UserId)
	paramData = append(paramData, fmt.Sprintf(`%%,%s,%%`, req.UserId))
	paramData = append(paramData, fmt.Sprintf(`%%,%s,%%`, req.UserId))

	queryListing = `SELECT TOP(CAST(? AS INT)) E.full_name emp_name, E.emp_id, EC.emp_no emp_no, P.pos_name_` + req.Language + ` emp_pos, phone_ext,
						DEPT.pos_name_` + req.Language + ` dept, EC.start_date, Grade.grade_name grade, STATUS.employmentstatus_name_` + req.Language + ` status, E.email, 
						E.photo emp_photo, EP.phone, MAR.name_` + req.Language + ` marital_status, EC.end_date, gender_name = CASE WHEN gender = '0' THEN 'Female' WHEN gender = '1' THEN 'Male' END,
						CASE WHEN EXISTS( SELECT TOP 1 seq_id FROM TCLTRequest
							WHERE reqemp = E.emp_id 
							AND sfobject != ''
							AND TCLTRequest.company_id = ?
							AND ( ( ( requester = ? OR reqemp = ? ) AND status NOT IN (0, 3, 8, 9) )
							OR ( ',' + approval_list + ',' LIKE ? AND ',' + approved_list + ',' NOT LIKE ? AND status IN ( 1, 2) ) ) ) 
							THEN 2
							ELSE 1
						END req_flag,
						E.gender
					FROM
						TEOMEmpPersonal E
					LEFT OUTER JOIN TEODEmpPersonal EP ON
						EP.emp_id = E.emp_id
					LEFT OUTER JOIN TEODEmpCompany EC ON
						EC.emp_id = E.emp_id
					LEFT OUTER JOIN TEOMMarital MAR ON
						MAR.code = EP.maritalstatus
					LEFT OUTER JOIN TEOMPosition P ON
						P.position_id = EC.position_id
					LEFT OUTER JOIN TEOMPosition Dept ON
						Dept.position_id = P.dept_id
					LEFT OUTER JOIN TEOMJobGrade Grade ON
						Grade.company_id = EC.company_id
						AND Grade.grade_code = EC.grade_code
					LEFT OUTER JOIN TEOMEmploymentStatus STATUS ON
						STATUS.employmentstatus_code = EC.employ_code
					INNER JOIN( SELECT emp_id empidpkdauthkey FROM VIEW_DATAAUTH
						WHERE (emp_id = ? AND company_id = ? )
							OR ( company_id = ? AND emp_id IN ( SELECT RGM.emp_id FROM TCLRGroupMember RGM
								WHERE RGM.datagroup_id IN (36, 99, 102, 115, 121, 122, 124, 127, 134, 138, 140, 142, 143, 158, 159, 163, 172, 173, 179, 184, 186, 190, 192, 193, 199, 205, 208, 219, 221, 225, 228, 229, 232, 234, 351, 359, 362, 366, 369, 382, 403, 413, 414, 426, 427, 428, 432, 434, 440) ) )
								OR ((company_id = ? AND (grade_category = 'SPV' AND 1 = 1 AND 1 = 1))
								OR (company_id = ? AND (grade_category = 'PRES'))
								OR (company_id = ? AND (1 = 0))
								OR (company_id = ? AND (worklocation_code = 'HEADOFFICE'))
								OR (company_id = ? AND ',' + spv_path + ',' LIKE ?)
								OR (company_id = ? AND ',' + spv_path + ',' LIKE ?)) ) SubAuth ON
						SubAuth.empidpkdauthkey = E.emp_id
					WHERE
						EC.company_id = ?
						AND (EC.end_date >= ? 
							OR EC.end_date IS NULL)
						AND EC.status = 1
						AND ( 1 = 1
							OR E.emp_id IN (SELECT reqemp FROM TCLTRequest R
							WHERE R.status IN (1, 2, 4) 
							AND ( requester = ? OR ( ',' + approval_list + ',' LIKE ? AND ',' + approved_list + ',' NOT LIKE ? AND 1 = 0 )))) `

	// ORDER BY
	// req_flag DESC,
	// emp_name,
	// emp_no

	if jumlahSudahTampil > 0 {
		paramData = append(paramData, jumlahSudahTampil)
		paramData = append(paramData, req.CompanyId)
		paramData = append(paramData, req.UserId)
		paramData = append(paramData, req.EmployeeId)
		paramData = append(paramData, fmt.Sprintf(`%%,%s,%%`, req.UserId))
		paramData = append(paramData, fmt.Sprintf(`%%,%s,%%`, req.UserId))
		paramData = append(paramData, req.EmployeeId)
		paramData = append(paramData, req.CompanyId)
		paramData = append(paramData, req.CompanyId)

		paramData = append(paramData, req.CompanyId)
		paramData = append(paramData, req.CompanyId)
		paramData = append(paramData, req.CompanyId)
		paramData = append(paramData, req.CompanyId)
		paramData = append(paramData, req.CompanyId)
		paramData = append(paramData, fmt.Sprintf(`%%,%s,%%`, req.EmployeeId))
		paramData = append(paramData, req.CompanyId)
		paramData = append(paramData, fmt.Sprintf(`%%,%s,%%`, req.EmployeeId))
		paramData = append(paramData, req.CompanyId)
		paramData = append(paramData, TglSekarangStr)
		paramData = append(paramData, req.UserId)
		paramData = append(paramData, fmt.Sprintf(`%%,%s,%%`, req.UserId))
		paramData = append(paramData, fmt.Sprintf(`%%,%s,%%`, req.UserId))
		queryListing = queryListing + ` AND E.emp_id NOT IN ( SELECT emp_id FROM (SELECT TOP(CAST(? AS INT)) E.emp_id,
																E.full_name,
																EC.emp_no,
																CASE
																	WHEN EXISTS(
																	SELECT
																		TOP 1 seq_id
																	FROM
																		TCLTRequest
																	WHERE
																		reqemp = E.emp_id
																		AND sfobject != ''
																		AND TCLTRequest.company_id = ?
																		AND ( ( ( requester = ?
																			OR reqemp = ? )
																		AND status NOT IN (0, 3, 8, 9) )
																			OR ( ',' + approval_list + ',' LIKE ?
																				AND ',' + approved_list + ',' NOT LIKE ?
																				AND status IN ( 1, 2) ) ) ) THEN 2
																	ELSE 1
																END req_flag
															FROM
																TEOMEmpPersonal E
															LEFT OUTER JOIN TEODEmpPersonal EP ON
																EP.emp_id = E.emp_id
															LEFT OUTER JOIN TEODEmpCompany EC ON
																EC.emp_id = E.emp_id
															LEFT OUTER JOIN TEOMMarital MAR ON
																MAR.code = EP.maritalstatus
															LEFT OUTER JOIN TEOMPosition P ON
																P.position_id = EC.position_id
															LEFT OUTER JOIN TEOMPosition Dept ON
																Dept.position_id = P.dept_id
															LEFT OUTER JOIN TEOMJobGrade Grade ON
																Grade.company_id = EC.company_id
																AND Grade.grade_code = EC.grade_code
															LEFT OUTER JOIN TEOMEmploymentStatus STATUS ON
																STATUS.employmentstatus_code = EC.employ_code
															INNER JOIN( SELECT emp_id empidpkdauthkey FROM VIEW_DATAAUTH
																WHERE (emp_id = ? AND company_id = ? )
																	OR ( company_id = ? AND emp_id IN ( SELECT RGM.emp_id FROM TCLRGroupMember RGM
																		WHERE RGM.datagroup_id IN (36, 99, 102, 115, 121, 122, 124, 127, 134, 138, 140, 142, 143, 158, 159, 163, 172, 173, 179, 184, 186, 190, 192, 193, 199, 205, 208, 219, 221, 225, 228, 229, 232, 234, 351, 359, 362, 366, 369, 382, 403, 413, 414, 426, 427, 428, 432, 434, 440) ) )
																		OR ((company_id = ? AND (grade_category = 'SPV' AND 1 = 1 AND 1 = 1))
																		OR (company_id = ? AND (grade_category = 'PRES'))
																		OR (company_id = ? AND (1 = 0))
																		OR (company_id = ? AND (worklocation_code = 'HEADOFFICE'))
																		OR (company_id = ? AND ',' + spv_path + ',' LIKE ?)
																		OR (company_id = ? AND ',' + spv_path + ',' LIKE ?)) ) SubAuth ON
																SubAuth.empidpkdauthkey = E.emp_id
															WHERE
																EC.company_id = ?
																AND (EC.end_date >= ? 
																	OR EC.end_date IS NULL)
																AND EC.status = 1
																AND ( 1 = 1
																	OR E.emp_id IN (SELECT reqemp FROM TCLTRequest R
																	WHERE R.status IN (1, 2, 4) 
																	AND ( requester = ? OR ( ',' + approval_list + ',' LIKE ? AND ',' + approved_list + ',' NOT LIKE ? AND 1 = 0 )))) 
																ORDER BY
																	req_flag DESC,
																	E.full_name,
																	EC.emp_no) AS D)`
	}
	queryListing = queryListing + `ORDER BY req_flag DESC, emp_name, emp_no`
	queryListing = repo.dbSlave.Rebind(queryListing)
	res1, errData := repo.dbSlave.Queryx(queryListing, paramData...)
	// fmt.Print(queryListing)
	if errData != nil {
		logger.Error(nil, errData)
		// logger.Println(queryListing)
		return errData, dataEmployeeInfo
	}

	defer res1.Close()

	if res1.Next() {
		var temp model.GetEmployeeInformationResponse

		errData := res1.Scan(&temp.EmployeeName, &temp.EmployeeId, &temp.EmployeeNo, &temp.EmployeePos, &temp.EmployeePhoneExt,
			&temp.EmployeeDept, &temp.EmployeeStartDate, &temp.EmployeeGrade, &temp.EmployeeStatus, &temp.EmployeeEmail,
			&temp.EmployeePhoto, &temp.EmployeePhone, &temp.EmployeeMaritalStatus, &temp.EmployeeEndDate, &temp.EmployeeGenderName,
			&temp.EmployeeReqFlag, &temp.EmployeeGenderCode)
		if errData != nil {
			logger.Error(nil, errData)
			// logger.Println(queryListing)
			res1.Close()
			return errData, dataEmployeeInfo
		}

		dataEmployeeInfo = append(dataEmployeeInfo, temp)
		for res1.Next() {
			errData := res1.Scan(&temp.EmployeeName, &temp.EmployeeId, &temp.EmployeeNo, &temp.EmployeePos, &temp.EmployeePhoneExt,
				&temp.EmployeeDept, &temp.EmployeeStartDate, &temp.EmployeeGrade, &temp.EmployeeStatus, &temp.EmployeeEmail,
				&temp.EmployeePhoto, &temp.EmployeePhone, &temp.EmployeeMaritalStatus, &temp.EmployeeEndDate, &temp.EmployeeGenderName,
				&temp.EmployeeReqFlag, &temp.EmployeeGenderCode)
			if errData != nil {
				logger.Error(nil, errData)
				// logger.Println(queryListing)
				res1.Close()
				return errData, dataEmployeeInfo
			}

			dataEmployeeInfo = append(dataEmployeeInfo, temp)
		}
	}
	return errData, dataEmployeeInfo
}

func (repo *repo) GetCity(ctx context.Context, req model.GetCityRequest) (error, []model.GetCityResponse) {
	var (
		recordset   []model.GetCityResponse
		errData     error
		queryString string
		paramData   []interface{}
	)

	queryString = `
		SELECT TOP 50 
			CTY.city_id
			, CTY.city_name + ' (' + STE.state_name + ', ' + CTR.country_name +')' as city_name
			, CTY.state_id
			, STE.state_name
			, STE.country_id
			, CTR.country_name
		FROM TGEMCITY CTY
		INNER JOIN TGEMSTATE STE
			ON CTY.state_id = STE.state_id
		INNER JOIN TGEMCOUNTRY CTR
			ON STE.country_id = CTR.country_id
		WHERE 1=1 `

	if req.Id != 0 {
		paramData = append(paramData, req.Id)
		queryString = queryString + ` AND CTY.city_id = ? `
	}

	queryString = queryString + ` ORDER BY city_name `

	queryString = repo.dbSlave.Rebind(queryString)
	res, errData := repo.dbSlave.Queryx(queryString, paramData...)

	if errData != nil {
		logger.Error(nil, errData)
		return errData, recordset
	}

	defer res.Close()

	if res.Next() {
		var temp model.GetCityResponse

		errData := res.Scan(&temp.Value, &temp.Label, &temp.StateId, &temp.StateName, &temp.CountryId, &temp.CountryName)
		if errData != nil {
			logger.Error(nil, errData)
			res.Close()
			return errData, recordset
		}

		recordset = append(recordset, temp)
		for res.Next() {
			errData := res.Scan(&temp.Value, &temp.Label, &temp.StateId, &temp.StateName, &temp.CountryId, &temp.CountryName)
			if errData != nil {
				logger.Error(nil, errData)
				res.Close()
				return errData, recordset
			}

			recordset = append(recordset, temp)
		}
	}
	return errData, recordset
}

func (repo *repo) GetAddressType(ctx context.Context, req model.GetAddressTypeRequest) (error, []model.GetAddressTypeResponse) {
	var (
		recordset   []model.GetAddressTypeResponse
		errData     error
		queryString string
		paramData   []interface{}
	)

	if req.Language == "" {
		req.Language = "en"
	}

	queryString = `SELECT code, name_` + req.Language + `, order_no FROM TEOMADDRESSTYPE WHERE 1=1 `

	// paramData = append(paramData, req.CompanyId)
	// paramData = append(paramData, req.UserId)
	// paramData = append(paramData, req.EmployeeId)
	if req.Code != "" {
		paramData = append(paramData, req.Code)
		queryString = queryString + ` AND code = ? `
	}

	queryString = queryString + ` ORDER BY order_no `

	queryString = repo.dbSlave.Rebind(queryString)
	res, errData := repo.dbSlave.Queryx(queryString, paramData...)

	if errData != nil {
		logger.Error(nil, errData)
		return errData, recordset
	}

	defer res.Close()

	if res.Next() {
		var temp model.GetAddressTypeResponse

		errData := res.Scan(&temp.Value, &temp.Label, &temp.OrderNo)
		if errData != nil {
			logger.Error(nil, errData)
			res.Close()
			return errData, recordset
		}

		recordset = append(recordset, temp)
		for res.Next() {
			errData := res.Scan(&temp.Value, &temp.Label, &temp.OrderNo)
			if errData != nil {
				logger.Error(nil, errData)
				res.Close()
				return errData, recordset
			}

			recordset = append(recordset, temp)
		}
	}
	return errData, recordset
}

func (repo *repo) GetOwnerStatus(ctx context.Context, req model.GetOwnerStatusRequest) (error, []model.GetOwnerStatusResponse) {
	var (
		recordset   []model.GetOwnerStatusResponse
		errData     error
		queryString string
		paramData   []interface{}
	)

	if req.Language == "" {
		req.Language = "en"
	}

	queryString = `SELECT code, name_` + req.Language + `, order_no FROM TEOMOWNERSTATUS WHERE 1=1 `

	// paramData = append(paramData, req.CompanyId)
	// paramData = append(paramData, req.UserId)
	// paramData = append(paramData, req.EmployeeId)
	if req.Code != "" {
		paramData = append(paramData, req.Code)
		queryString = queryString + ` AND code = ? `
	}

	queryString = queryString + ` ORDER BY order_no `

	queryString = repo.dbSlave.Rebind(queryString)
	res, errData := repo.dbSlave.Queryx(queryString, paramData...)

	if errData != nil {
		logger.Error(nil, errData)
		return errData, recordset
	}

	defer res.Close()

	if res.Next() {
		var temp model.GetOwnerStatusResponse

		errData := res.Scan(&temp.Value, &temp.Label, &temp.OrderNo)
		if errData != nil {
			logger.Error(nil, errData)
			res.Close()
			return errData, recordset
		}

		recordset = append(recordset, temp)
		for res.Next() {
			errData := res.Scan(&temp.Value, &temp.Label, &temp.OrderNo)
			if errData != nil {
				logger.Error(nil, errData)
				res.Close()
				return errData, recordset
			}

			recordset = append(recordset, temp)
		}
	}
	return errData, recordset
}

func (repo *repo) GetStayStatus(ctx context.Context, req model.GetStayStatusRequest) (error, []model.GetStayStatusResponse) {
	var (
		recordset   []model.GetStayStatusResponse
		errData     error
		queryString string
		paramData   []interface{}
	)

	if req.Language == "" {
		req.Language = "en"
	}

	queryString = `SELECT code, name_` + req.Language + `, order_no FROM TEOMSTAYSTATUS WHERE 1=1 `

	// paramData = append(paramData, req.CompanyId)
	// paramData = append(paramData, req.UserId)
	// paramData = append(paramData, req.EmployeeId)
	if req.Code != "" {
		paramData = append(paramData, req.Code)
		queryString = queryString + ` AND code = ? `
	}

	queryString = queryString + ` ORDER BY order_no `

	queryString = repo.dbSlave.Rebind(queryString)
	res, errData := repo.dbSlave.Queryx(queryString, paramData...)

	if errData != nil {
		logger.Error(nil, errData)
		return errData, recordset
	}

	defer res.Close()

	if res.Next() {
		var temp model.GetStayStatusResponse

		errData := res.Scan(&temp.Value, &temp.Label, &temp.OrderNo)
		if errData != nil {
			logger.Error(nil, errData)
			res.Close()
			return errData, recordset
		}

		recordset = append(recordset, temp)
		for res.Next() {
			errData := res.Scan(&temp.Value, &temp.Label, &temp.OrderNo)
			if errData != nil {
				logger.Error(nil, errData)
				res.Close()
				return errData, recordset
			}

			recordset = append(recordset, temp)
		}
	}
	return errData, recordset
}
