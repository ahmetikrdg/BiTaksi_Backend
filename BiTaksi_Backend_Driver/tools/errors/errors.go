package errors

import "BiTaksi_Backend_Driver/tools/zap_logger"

var log = zap_logger.ZapMethod()

func StandartErrorWithErrorLog(message error, args interface{}) {
	if message != nil || args != nil {
		log.Error(message.Error())
	}
}

func ServerErrorWithErrorLog(message error) {
	if message != nil {
		log.Error(message.Error())
	}
}
