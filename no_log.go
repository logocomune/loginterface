package loginterface

type nolog struct {
}

// NewNoLog return a no log instance
func NewNoLog() Logger {
	return &nolog{}
}

func (*nolog) Debugf(format string, args ...interface{}) {
	return
}

func (*nolog) Infof(format string, args ...interface{}) {
	return
}

func (*nolog) Warnf(format string, args ...interface{}) {
	return
}

func (*nolog) Errorf(format string, args ...interface{}) {
	return
}

func (*nolog) Debug(args ...interface{}) {
	return
}

func (*nolog) Info(args ...interface{}) {
	return
}

func (*nolog) Warn(args ...interface{}) {
	return
}

func (*nolog) Error(args ...interface{}) {
	return
}

func (*nolog) Debugln(args ...interface{}) {
	return
}

func (*nolog) Infoln(args ...interface{}) {
	return
}

func (*nolog) Warnln(args ...interface{}) {
	return
}

func (*nolog) Errorln(args ...interface{}) {
	return
}
