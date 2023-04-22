package server

import (
	"github.com/golang/mock/gomock"
	mock_repository "homework-5/internal/pkg/repository/mocks"
	"testing"
)

type usersRepoFixture struct {
	ctrl       *gomock.Controller
	mProduct   *mock_repository.MockProductsRepo
	mWarehouse *mock_repository.MockWarehousesRepo
	Server     *Server
}

func setUp(t *testing.T) usersRepoFixture {
	ctrl := gomock.NewController(t)

	mProduct := mock_repository.NewMockProductsRepo(ctrl)
	mWarehouse := mock_repository.NewMockWarehousesRepo(ctrl)
	Server := New(mProduct, mWarehouse)

	return usersRepoFixture{ctrl: ctrl, mProduct: mProduct, mWarehouse: mWarehouse, Server: Server}

}

func (u *usersRepoFixture) tearDown() {
	u.ctrl.Finish()
}
