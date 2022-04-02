package zap_logger

var log = ZapMethod()

func ServerInfoWithInfoLog(info string) {
	if info != "" {
		log.Info(info)
	}

}
