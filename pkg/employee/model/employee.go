package model

import "time"

type (
	GetEmployeeInformationRequest struct {
		CompanyId  int64  `json:"company_id" binding:"required"`
		EmployeeId string `json:"employee_id" binding:"required"`
		Language   string `json:"language"`
		UserId     string `json:"user_id"`
		Page       int64  `json:"page"`
		Limit      int64  `json:"limit"`
	}

	GetEmployeeInformationResponse struct {
		EmployeeName          string     `json:"employee_name"`
		EmployeeId            string     `json:"employee_id"`
		EmployeeNo            string     `json:"employee_no"`
		EmployeePos           string     `json:"employee_pos"`
		EmployeePhoneExt      *string    `json:"employee_phone_ext"`
		EmployeeDept          *string    `json:"employee_dept"`
		EmployeeStartDate     time.Time  `json:"employee_start_date"`
		EmployeeGrade         string     `json:"employee_grade"`
		EmployeeStatus        string     `json:"employee_status"`
		EmployeeEmail         *string    `json:"employee_email"`
		EmployeePhoto         string     `json:"employee_photo"`
		EmployeePhone         *string    `json:"employee_phone"`
		EmployeeMaritalStatus string     `json:"employee_marital_status"`
		EmployeeEndDate       *time.Time `json:"employee_end_date"`
		EmployeeGenderName    string     `json:"employee_gender_name"`
		EmployeeReqFlag       string     `json:"employee_req_flag"`
		EmployeeGenderCode    string     `json:"employee_gender_code"`
	}

	GetEmployeeListingRequest struct {
		CompanyId              int64    `json:"company_id" binding:"required"`
		EmployeeId             string   `json:"employee_id" binding:"required"`
		Language               string   `json:"language"`
		UserId                 string   `json:"user_id"`
		FilterName             string   `json:"filter_name"`
		FilterStatus           []string `json:"filter_status"`
		FilterGender           []string `json:"filter_gender"`
		FilterJoinDate         string   `json:"filter_join_date"`
		FilterEmploymentStatus []string `json:"filter_employment_status"`
		FilterGrade            []string `json:"filter_grade"`
		Page                   int64    `json:"page"`
		Limit                  int64    `json:"limit"`
	}

	GetEmployeeResponse struct {
		EmployeeName          string     `json:"employee_name"`
		EmployeeId            string     `json:"employee_id"`
		EmployeeNo            string     `json:"employee_no"`
		EmployeePos           string     `json:"employee_pos"`
		EmployeePhoneExt      *string    `json:"employee_phone_ext"`
		EmployeeDept          string     `json:"employee_dept"`
		EmployeeStartDate     time.Time  `json:"employee_start_date"`
		EmployeeGrade         string     `json:"employee_grade"`
		EmployeeStatus        string     `json:"employee_status"`
		EmployeeEmail         string     `json:"employee_email"`
		EmployeePhoto         string     `json:"employee_photo"`
		EmployeePhone         *string    `json:"employee_phone"`
		EmployeeMaritalStatus string     `json:"employee_marital_status"`
		EmployeeEndDate       *time.Time `json:"employee_end_date"`
		EmployeeGenderName    string     `json:"employee_gender_name"`
		EmployeeReqFlag       string     `json:"employee_req_flag"`
		EmployeeGenderCode    string     `json:"employee_gender_code"`
	}

	GetEmployeeListingResponse struct {
		RecordCount int                   `json:"recordcount"`
		RecordSet   []GetEmployeeResponse `json:"recordset"`
	}

	GetEmployeeByIdRequest struct {
		CompanyId  int64  `json:"company_id" binding:"required"`
		EmployeeId string `json:"employee_id" binding:"required"`
		Language   string `json:"language"`
	}

	GetEmployeeByIdResponse struct {
		EmployeeId            string     `json:"employee_id"`
		EmployeeTaxNo         string     `json:"employee_tax_no"`
		EmployeeFirstName     string     `json:"employee_first_name"`
		EmployeeMiddleName    string     `json:"employee_middle_name"`
		EmployeeLastName      string     `json:"employee_last_name"`
		EmployeeGender        string     `json:"employee_gender"`
		EmployeeBirthplace    string     `json:"employee_birthplace"`
		EmployeeBirthdate     *time.Time `json:"employee_birthdate"`
		EmployeePhoto         string     `json:"employee_photo"`
		EmployeePhone         *string    `json:"employee_phone"`
		EmployeeEmail         string     `json:"employee_email"`
		EmployeeMaritalStatus string     `json:"employee_marital_status"`
		EmployeeStartDate     time.Time  `json:"employee_start_date"`
		EmployeeNo            string     `json:"employee_no"`
		EmployeeStatus        string     `json:"employee_status"`
		EmployeeIsMain        string     `json:"employee_is_main"`
		EmployeeCompanyId     string     `json:"employee_company_id"`
		EmployeePositionId    string     `json:"employee_position_id"`
		EmployeeCodeStatus    string     `json:"employee_code_status"`
		EmployeeGrade         string     `json:"employee_grade"`
		EmployeeFullName      string     `json:"employee_full_name"`
		EmployeePosName       string     `json:"employee_pos_name"`
		EmployeeYearBirth     string     `json:"employee_year_birth"`
		EmployeeMonthBirth    string     `json:"employee_month_birth"`
		EmployeeDeptName      string     `json:"employee_dept_name"`
		EmployeePhoneExt      *string    `json:"employee_phone_ext"`
		EmployeeActiveAddress string     `json:"employee_active_address"`
	}

	GetEmployeeMasterAddressRequest struct {
		EmployeeId      string `json:"employee_id" binding:"required"`
		Language        string `json:"language" binding:"required"`
		AddressTypeCode string `json:"addresstype_code"`
	}

	GetEmployeeMasterAddressResponse struct {
		EmployeeId                  string  `json:"employee_id"`
		EmployeeAddressType         *string `json:"employee_address_type"`
		EmployeeAddressTypeName     *string `json:"employee_address_type_name"`
		EmployeeAddress             *string `json:"employee_address"`
		EmployeeAddressRt           *string `json:"employee_address_rt"`
		EmployeeAddressRw           *string `json:"employee_address_rw"`
		EmployeeAddressSubdistrict  *string `json:"employee_address_subdistrict"`
		EmployeeAddressDistrict     *string `json:"employee_address_district"`
		EmployeeAddressCityId       *string `json:"employee_address_city_id"`
		EmployeeAddressStateId      *string `json:"employee_address_state_id"`
		EmployeeAddressCountryId    *string `json:"employee_address_country_id"`
		EmployeeAddressZipcode      *string `json:"employee_address_zipcode"`
		EmployeeAddressPhone        *string `json:"employee_address_phone"`
		EmployeeAddressLivingStatus *string `json:"employee_address_living_status"`
		EmployeeAddressOwnerStatus  *string `json:"employee_address_owner_status"`
		EmployeeAddressDistance     *string `json:"employee_address_distance"`
		EmployeeAddressLatLong      *string `json:"employee_address_lat_long"`
		EmployeeAddressLocal        *string `json:"employee_address_local"`
		EmployeeAddressCityName     *string `json:"employee_address_city_name"`
		EmployeeAddressStateName    *string `json:"employee_address_state_name"`
	}

	UpdateEmployeeMasterAddressRequest struct {
		EmployeeId                  string  `json:"employee_id"`
		EmployeeAddressType         string  `json:"employee_address_type"`
		EmployeeAddress             string  `json:"employee_address"`
		EmployeeAddressRt           *string `json:"employee_address_rt"`
		EmployeeAddressRw           *string `json:"employee_address_rw"`
		EmployeeAddressSubdistrict  *string `json:"employee_address_subdistrict"`
		EmployeeAddressDistrict     *string `json:"employee_address_district"`
		EmployeeAddressCityId       *string `json:"employee_address_city_id"`
		EmployeeAddressStateId      *string `json:"employee_address_state_id"`
		EmployeeAddressCountryId    string  `json:"employee_address_country_id"`
		EmployeeAddressZipcode      *string `json:"employee_address_zipcode"`
		EmployeeAddressPhone        *string `json:"employee_address_phone"`
		EmployeeAddressLivingStatus *string `json:"employee_address_living_status"`
		EmployeeAddressOwnerStatus  *string `json:"employee_address_owner_status"`
		EmployeeAddressDistance     *string `json:"employee_address_distance"`
		EmployeeAddressLatLong      *string `json:"employee_address_lat_long"`
		EmployeeAddressLocal        *string `json:"employee_address_local"`
		Username                    string  `json:"username"`
	}

	CreateEmployeeMasterAddressRequest struct {
		EmployeeId                  string  `json:"employee_id"`
		EmployeeAddressType         string  `json:"employee_address_type"`
		EmployeeAddress             string  `json:"employee_address"`
		EmployeeAddressRt           *string `json:"employee_address_rt"`
		EmployeeAddressRw           *string `json:"employee_address_rw"`
		EmployeeAddressSubdistrict  *string `json:"employee_address_subdistrict"`
		EmployeeAddressDistrict     *string `json:"employee_address_district"`
		EmployeeAddressCityId       *string `json:"employee_address_city_id"`
		EmployeeAddressStateId      *string `json:"employee_address_state_id"`
		EmployeeAddressCountryId    string  `json:"employee_address_country_id"`
		EmployeeAddressZipcode      *string `json:"employee_address_zipcode"`
		EmployeeAddressPhone        *string `json:"employee_address_phone"`
		EmployeeAddressLivingStatus *string `json:"employee_address_living_status"`
		EmployeeAddressOwnerStatus  *string `json:"employee_address_owner_status"`
		EmployeeAddressDistance     *string `json:"employee_address_distance"`
		EmployeeAddressLatLong      *string `json:"employee_address_lat_long"`
		EmployeeAddressLocal        *string `json:"employee_address_local"`
		Username                    string  `json:"username"`
	}

	/*
		======================= MASTER DATA ========================
	*/

	// Employment Status
	GetEmploymentStatusRequest struct {
		Language string `json:"language"`
		Code     string `json:"code"`
	}

	GetEmploymentStatusResponse struct {
		Value   string `json:"value"`
		Label   string `json:"label"`
		OrderNo string `json:"order_no"`
	}

	// Grade
	GetJobGradeRequest struct {
		CompanyId int64  `json:"company_id"`
		Code      string `json:"code"`
	}

	GetJobGradeResponse struct {
		Value   string `json:"value"`
		Label   string `json:"label"`
		OrderNo string `json:"order_no"`
	}

	// City
	GetCityRequest struct {
		Id string `json:"id"`
	}

	GetCityResponse struct {
		// value untuk unique id, label untuk nama, dibuat seperti ini untuk mengikuti format dari antd
		Value       string `json:"value"`
		Label       string `json:"label"`
		StateId     string `json:"state_id"`
		StateName   string `json:"state_name"`
		CountryId   string `json:"country_id"`
		CountryName string `json:"country_name"`
	}

	// Address Type
	GetAddressTypeRequest struct {
		Language string `json:"language"`
		Code     string `json:"code"`
	}

	GetAddressTypeResponse struct {
		// value untuk unique id, label untuk nama, dibuat seperti ini untuk mengikuti format dari antd
		Value   string `json:"value"`
		Label   string `json:"label"`
		OrderNo string `json:"order_no"`
	}

	// Owner Status
	GetOwnerStatusRequest struct {
		Language string `json:"language"`
		Code     string `json:"code"`
	}

	GetOwnerStatusResponse struct {
		Value   string `json:"value"`
		Label   string `json:"label"`
		OrderNo string `json:"order_no"`
	}

	// Stay Status
	GetStayStatusRequest struct {
		Language string `json:"language"`
		Code     string `json:"code"`
	}

	GetStayStatusResponse struct {
		Value   string `json:"value"`
		Label   string `json:"label"`
		OrderNo string `json:"order_no"`
	}
)
