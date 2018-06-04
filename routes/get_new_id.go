package routes

import (
	"profzone/libtools/courier/httpx"
	"context"
	"profzone/service-id/modules/algorithm"
	"profzone/service-id/global"
	"profzone/service-id/constants/errors"
	"github.com/sirupsen/logrus"
)

// 获取一个新的全局唯一ID
type GetNewId struct {
	httpx.MethodGet
}

func (req GetNewId) Path() string {
	return "/id"
}

type GetNewIdResponse struct {
	ID uint64 `json:"id"`
}

func (req GetNewId) Output(ctx context.Context) (result interface{}, err error) {
	generator := algorithm.GeneratorContainerInstance.GetAlgorithm(global.Config.GenerateAlgorithm)
	if generator == nil {
		return nil, errors.AlgorithmNotFound
	}
	resp, err := generator.GenerateUniqueID()
	if err != nil {
		logrus.Errorf("generator.GenerateUniqueID err: %v", err)
		return nil, err
	}

	return GetNewIdResponse{
		ID: resp,
	}, nil
}

