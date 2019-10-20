package mock

import (
	"github.com/stretchr/testify/mock"
	"github.com/tpphu/visual-job/model"
)

type ProductRepoImpl struct {
	mock.Mock
}

func (self *ProductRepoImpl) Create(product model.Product) (*model.Product, error) {
	args := self.Called(product)
	out := args.Get(0)
	if out == nil {
		return nil, args.Error(1)
	}
	return out.(*model.Product), args.Error(1)
}
