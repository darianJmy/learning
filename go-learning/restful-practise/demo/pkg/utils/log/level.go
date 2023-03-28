package log

func Debug(msgs ...interface{}) {
	logger.Debug(msgs...)
}

func Debugf(msg string, keysAndValues ...interface{}) {
	logger.Debugf(msg, keysAndValues...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	logger.Debugw(msg, keysAndValues...)
}

func Info(msgs ...interface{}) {
	logger.Info(msgs...)
}

func Infof(msg string, keysAndValues ...interface{}) {
	logger.Infof(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	logger.Infow(msg, keysAndValues...)
}

func Warn(msgs ...interface{}) {
	logger.Warn(msgs...)
}

func Warnf(msg string, keysAndValues ...interface{}) {
	logger.Warnf(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	logger.Warnw(msg, keysAndValues...)
}

func Error(msgs ...interface{}) {
	logger.Error(msgs...)
}

func Errorf(msg string, keysAndValues ...interface{}) {
	logger.Errorf(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	logger.Errorw(msg, keysAndValues...)
}

func Fatal(msgs ...interface{}) {
	logger.Fatal(msgs...)
}

func Fatalf(msg string, keysAndValues ...interface{}) {
	logger.Fatalf(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	logger.Fatalw(msg, keysAndValues...)
}
