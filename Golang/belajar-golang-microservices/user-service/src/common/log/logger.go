package log

import "go.uber.org/zap"

var Logger *zap.Logger // ← uppercase

func init() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}
