package demo

import (
	"context"

	"base/gopkg/log"
	"base/gopkg/services"
	"base/internal/model"

	"go.uber.org/zap"
)

func (s *Service) CreateDemo(ctx context.Context, name string, fileType int, projectType int, metadata model.DemoMetadata) (services.Result, error) {
	logPrefix := "/internal/service/demo: Service.CreateDemo()"
	writingOnlineId, err := s.demoDao.Create(ctx, name, fileType, projectType, "", metadata)
	if err != nil {
		log.Sugar().Error(ctx, logPrefix, zap.Any("demo dao Create() error", err))
		return services.Failed(ctx, err)
	}

	log.Sugar().Info(ctx, logPrefix, zap.Any("demo dao Create() success", writingOnlineId))
	return services.Success(ctx, writingOnlineId)
}
