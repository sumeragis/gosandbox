package logger

import (
	"log"

	"go.uber.org/zap"
)

var Log *zap.Logger

func init() {
	l, err := zap.NewProduction()
	if err != nil {
		log.Panicf("Error: failed to init logger err=%s", err.Error())
	}
	Log = l
}
