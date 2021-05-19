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
	HealthCheck(ctx context.Context, req *modelProtoc.HealthCheckRequest) (modelProtoc.HealthCheckResponse, error)
	AddFieldCheck(ctx context.Context, req *modelProtoc.AddFieldCheckRequest) (modelProtoc.AddFieldCheckResponse, error)
}

func NewRepo(dbSlave, dbMaster *sqlx.DB) Repository {
	return &repo{
		dbSlave:  dbSlave,
		dbMaster: dbMaster,
	}
}

func (r repo) HealthCheck(ctx context.Context, req *modelProtoc.HealthCheckRequest) (modelProtoc.HealthCheckResponse, error) {
	// @ select table
	//err := r.db.QueryRow("SELECT email FROM mt_user WHERE id=$1", wording).Scan(&wording)
	//if err != nil {
	//	level.Error(logCreate).Log("err", err)
	//	return "", RepoErr
	//}

	// @ insert
	//var query = `
	//	INSERT INTO mt_user (id, email, password)
	//	VALUES ($1, $2, $3)`
	//
	//_, err := repo.db.ExecContext(ctx, query, user.ID, user.Email, user.Password)
	//if err != nil {
	//	level.Error(logCreate).Log("err", err)
	//	return err
	//}

	var res modelProtoc.HealthCheckResponse
	res.Wording = "Are You Say, " + req.Wording

	return res, nil
}

func (r repo) AddFieldCheck(ctx context.Context, req *modelProtoc.AddFieldCheckRequest) (modelProtoc.AddFieldCheckResponse, error) {
	var (
		paramData []interface{}
	)

	paramData = append(paramData, req.CompanyId)
	paramData = append(paramData, req.PageId)

	queryString := `select company_id, page_id, table_name, additional_fields from sf7_core_manage_field where company_id = ? AND page_id = ?`
	queryString = r.dbSlave.Rebind(queryString)
	resData, errData := r.dbSlave.Queryx(queryString, paramData...)
	if errData != nil {
		logger.Error(nil, errData)
	}
	var p model.CoreManageField
	err := resData.Scan(&p.CompanyId,&p.PageId,&p.TableName, &p.AdditionalFields)
	if err != nil {
		logger.Error(nil, err)
	}
	res := modelProtoc.AddFieldCheckResponse{
		Name: p.AdditionalFields,
	}
	return res, RepoErr
}