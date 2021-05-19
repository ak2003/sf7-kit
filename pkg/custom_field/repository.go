package custom_field

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"gitlab.dataon.com/gophers/sf7-kit/pkg/custom_field/model"
	modelProtoc "gitlab.dataon.com/gophers/sf7-kit/pkg/custom_field/model/protoc/model"
	"gitlab.dataon.com/gophers/sf7-kit/shared/utils/logger"
)

var RepoErr = errors.New("unable to handle repo request")

type repo struct {
	dbSlave  *sqlx.DB
	dbMaster *sqlx.DB
}

type Repository interface {
	AddFieldCheck(ctx context.Context, req *modelProtoc.AddFieldCheckRequest) (modelProtoc.AddFieldCheckResponse, error)
}

func NewRepo(dbSlave, dbMaster *sqlx.DB) Repository {
	return &repo{
		dbSlave:  dbSlave,
		dbMaster: dbMaster,
	}
}

func (r repo) AddFieldCheck(ctx context.Context, req *modelProtoc.AddFieldCheckRequest) (modelProtoc.AddFieldCheckResponse, error) {
	//var (
	//	paramData []interface{}
	//)

	//paramData = append(paramData, req.CompanyId)
	//paramData = append(paramData, req.PageId)

	var p model.CoreManageField
	logger.Info(nil, req.CompanyId)
	//queryString := "select company_id, page_id, table_name, additional_fields from SF7_CORE_MANAGE_FIELD where company_id = $1"
	err := r.dbSlave.Get(&p, "select company_id, page_id, table_name, additional_fields, status_fields from dbSF6_QA.dbo.SF7_CORE_MANAGE_FIELD where company_id='83'")
	if err != nil {
		logger.Error(nil, err)
	}
	//queryString = r.dbSlave.Rebind(queryString)
	//resData, errData := r.dbSlave.Query(queryString, paramData...)
	//if errData != nil {
	//	logger.Error(nil, errData)
	//}

	//err := resData.Scan(&p.CompanyId,&p.PageId,&p.TableName, &p.AdditionalFields)
	//if err != nil {
	//	logger.Error(nil, err)
	//}
	res := modelProtoc.AddFieldCheckResponse{
		Name: p.AdditionalFields,
	}
	return res, RepoErr
}