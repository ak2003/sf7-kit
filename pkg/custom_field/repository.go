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
	AddFieldCheck(ctx context.Context, req *modelProtoc.AddFieldCheckRequest) (*model.CoreManageField, error)
}

func NewRepo(dbSlave, dbMaster *sqlx.DB) Repository {
	return &repo{
		dbSlave:  dbSlave,
		dbMaster: dbMaster,
	}
}

func (r repo) AddFieldCheck(ctx context.Context, req *modelProtoc.AddFieldCheckRequest) (*model.CoreManageField, error) {

	var p model.CoreManageField
	logger.Info(nil, req.CompanyId)

	query := r.dbSlave.Rebind("select company_id, page_id, table_name, additional_fields from dbSF6_QA.dbo.SF7_CORE_MANAGE_FIELD where company_id=? and page_id=?")
	err := r.dbSlave.Get(&p, query, req.CompanyId, req.PageId)
	if err != nil {
		logger.Error(nil, err)
		return nil, RepoErr
	}

	return &p, nil
}