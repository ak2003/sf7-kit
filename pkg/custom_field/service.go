package custom_field

import (
	"context"
	"gitlab.dataon.com/gophers/sf7-kit/pkg/custom_field/model/protoc/model"
)

type Service interface {
	CheckAddField(ctx context.Context, req *model.AddFieldCheckRequest) (interface{}, error)
}
