package logger

import (
	"go.uber.org/zap"
)

func New() {
	zap.NewProductionConfig()
}
