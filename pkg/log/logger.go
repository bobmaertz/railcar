package log 


type Logger interface {
    Info(string)
    Error(string)
    Debug(string)
    Warn(string)
    Fatal(string)

    Infof(string, ...interface{})
    Errorf(string, ...interface{})
    Debugf(string, ...interface{})
    Warnf(string, ...interface{})
    Fatalf(string, ...interface{})
}
