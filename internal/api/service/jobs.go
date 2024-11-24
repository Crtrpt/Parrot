package service

import (
	"context"

	"github.com/parrot/internal/api/v1/request"
	"github.com/parrot/internal/core"
	"github.com/parrot/internal/pkg/model"
	"github.com/parrot/internal/pkg/utils"
)

type JobsService struct {
}

func (srv JobsService) GetJobs(ctx context.Context, req *request.QueryForm) (ret []*request.JobResp, err error) {
	menus := []*model.Job{}
	err = core.DB.WithContext(ctx).Model(model.Menu{}).Where("name = ?", req.Name).Find(&menus).Error
	if err != nil {
		return
	}

	ret = []*request.JobResp{}
	err = utils.CopyArray(&ret, menus)

	return
}

func (srv JobsService) GetJobDetail(ctx context.Context, req *request.QueryForm) (ret []*request.JobResp, err error) {
	menus := []*model.Job{}
	err = core.DB.WithContext(ctx).Model(model.Menu{}).Where(utils.BuildExpr(req)).Find(&menus).Error
	if err != nil {
		return
	}

	ret = []*request.JobResp{}
	err = utils.CopyArray(&ret, menus)

	return
}
