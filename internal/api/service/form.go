package service

import (
	"context"

	"github.com/parrot/internal/api/v1/request"
	"github.com/parrot/internal/core"
	"github.com/parrot/internal/pkg/model"
	"github.com/parrot/internal/pkg/utils"
)

type FormService struct {
}

func (srv FormService) GetForms(ctx context.Context, req *request.QueryForm) (ret []*request.FormResp, err error) {
	forms := []*model.Form{}
	err = core.DB.WithContext(ctx).Model(model.Form{}).Find(&forms).Error
	if err != nil {
		return
	}

	ret = []*request.FormResp{}
	err = utils.CopyArray(&ret, forms)

	return
}
