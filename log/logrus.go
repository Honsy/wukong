package log

import (
	"bytes"
	"fmt"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	log    logrus.Ext1FieldLogger
	prefix string
	fields Fields
}

// Level type
type Level uint32

// These are the different logging levels. You can set the logging level to log
// on your instance of logger, obtained with `logrus.New()`.
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

type LogFormatter struct {
}

func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")

	b.WriteString(fmt.Sprintf("[%s] [%s]", timestamp, entry.Level))
	f.writeFields(b, entry)
	b.WriteString(fmt.Sprintf(" msg=%s \n", entry.Message))
	return b.Bytes(), nil
}

func (f *LogFormatter) writeFields(b *bytes.Buffer, entry *logrus.Entry) {
	if len(entry.Data) != 0 {
		fields := make([]string, 0, len(entry.Data))
		for field := range entry.Data {
			fields = append(fields, field)
		}

		for _, field := range fields {
			f.writeField(b, entry, field)
		}
	}
}

func (f *LogFormatter) writeField(b *bytes.Buffer, entry *logrus.Entry, field string) {
	fmt.Fprintf(b, "[%s:%v]", field, entry.Data[field])
}

func NewLogrusLogger(logrus logrus.Ext1FieldLogger, prefix string, fields Fields) *LogrusLogger {
	return &LogrusLogger{
		log:    logrus,
		prefix: prefix,
		fields: fields,
	}
}

func NewDefaultLogrusLogger() *LogrusLogger {
	logger := logrus.New()
	logger.SetFormatter(&LogFormatter{})
	// logger.Formatter = &prefixed.TextFormatter{
	// 	FullTimestamp:   true,
	// 	TimestampFormat: "2006-01-02 15:04:05.000",
	// }

	return NewLogrusLogger(logger, "main", nil)
}

func (l *LogrusLogger) Print(args ...interface{}) {
	l.prepareEntry().Print(args...)
}

func (l *LogrusLogger) Printf(format string, args ...interface{}) {
	l.prepareEntry().Printf(format, args...)
}

func (l *LogrusLogger) Trace(args ...interface{}) {
	l.prepareEntry().Trace(args...)
}

func (l *LogrusLogger) Tracef(format string, args ...interface{}) {
	l.prepareEntry().Tracef(format, args...)
}

func (l *LogrusLogger) Debug(args ...interface{}) {
	l.prepareEntry().Debug(args...)
}

func (l *LogrusLogger) Debugf(format string, args ...interface{}) {
	l.prepareEntry().Debugf(format, args...)
}

func (l *LogrusLogger) Info(args ...interface{}) {
	l.prepareEntry().Info(args...)
}

func (l *LogrusLogger) Infof(format string, args ...interface{}) {
	l.prepareEntry().Infof(format, args...)
}

func (l *LogrusLogger) Warn(args ...interface{}) {
	l.prepareEntry().Warn(args...)
}

func (l *LogrusLogger) Warnf(format string, args ...interface{}) {
	l.prepareEntry().Warnf(format, args...)
}

func (l *LogrusLogger) Error(args ...interface{}) {
	l.prepareEntry().Error(args...)
}

func (l *LogrusLogger) Errorf(format string, args ...interface{}) {
	l.prepareEntry().Errorf(format, args...)
}

func (l *LogrusLogger) Fatal(args ...interface{}) {
	l.prepareEntry().Fatal(args...)
}

func (l *LogrusLogger) Fatalf(format string, args ...interface{}) {
	l.prepareEntry().Fatalf(format, args...)
}

func (l *LogrusLogger) Panic(args ...interface{}) {
	l.prepareEntry().Panic(args...)
}

func (l *LogrusLogger) Panicf(format string, args ...interface{}) {
	l.prepareEntry().Panicf(format, args...)
}

func (l *LogrusLogger) WithPrefix(prefix string) Logger {
	return NewLogrusLogger(l.log, prefix, l.Fields())
}

func (l *LogrusLogger) Prefix() string {
	return l.prefix
}

func (l *LogrusLogger) WithFields(fields Fields) Logger {
	return NewLogrusLogger(l.log, l.Prefix(), l.Fields().WithFields(fields))
}

func (l *LogrusLogger) Fields() Fields {
	return l.fields
}

func (l *LogrusLogger) prepareEntry() *logrus.Entry {
	return l.log.
		WithFields(logrus.Fields(l.Fields())).
		WithField("prefix", l.Prefix())
}

func (l *LogrusLogger) SetLevel(level Level) {
	if ll, ok := l.log.(*logrus.Logger); ok {
		ll.SetLevel(logrus.Level(level))
	}
}
