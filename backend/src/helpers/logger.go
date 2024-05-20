package helpers

import (
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/codes"
)

func ZapCodeToLevel(code codes.Code) zapcore.Level {
	// Set OK response to be DEBUG rather than INFO
	// if code == codes.OK {
	// 	return zap.DebugLevel
	// }

	return grpc_zap.DefaultCodeToLevel(code)
}
