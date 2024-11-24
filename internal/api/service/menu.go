package service

import (
	"context"

	"github.com/parrot/internal/api/v1/request"
	"github.com/parrot/internal/core"
	"github.com/parrot/internal/pkg/model"
	"github.com/parrot/internal/pkg/utils"
)

type MenuService struct {
}

func (menu MenuService) GetMenus(ctx context.Context, req *request.QueryForm) (ret []*request.MenuResp, err error) {
	m := model.Menu{}
	err = core.DB.WithContext(ctx).Model(model.Menu{}).Where("name = ?", req.Name).First(&m).Error
	if err != nil {
		return
	}

	menus := []*model.Menu{}
	err = core.DB.WithContext(ctx).Model(model.Menu{}).Where("pid = ?", m.ID).Find(&menus).Error
	if err != nil {
		return
	}

	ret = []*request.MenuResp{}
	err = utils.CopyArray(&ret, menus)

	return
}
