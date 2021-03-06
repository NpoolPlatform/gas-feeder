package version

import (
	"fmt"

	npool "github.com/NpoolPlatform/message/npool"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/version"
)

func Version() (*npool.VersionResponse, error) {
	info, err := version.GetVersion()
	if err != nil {
		logger.Sugar().Errorf("get service version error: %+w", err)
		return nil, fmt.Errorf("get service version error: %w", err)
	}
	return &npool.VersionResponse{
		Info: info,
	}, nil
}
