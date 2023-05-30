package logx

import "errors"

const (
	DebugLevel uint32 = iota
	InfoLevel
	WarnLevel
	// ErrorLevel includes errors, slows, stacks
	ErrorLevel
	// SevereLevel only log severe messages
	SevereLevel
)

const (
	jsonEncodingType = iota
	plainEncodingType

	plainEncoding    = "plain"
	plainEncodingSep = '\t'
	sizeRotationRule = "size"
)

const (
	accessFilename = "access.log"
	debugFilename  = "debug.log"
	warnFilename   = "warn.log"
	errorFilename  = "error.log"
	severeFilename = "severe.log"
	slowFilename   = "slow.log"
	statFilename   = "stat.log"

	fileMode   = "file"
	volumeMode = "volume"

	levelAlert  = "alert"
	levelDebug  = "debug"
	levelWarn   = "warn"
	levelInfo   = "info"
	levelError  = "error"
	levelSevere = "severe"
	levelFatal  = "fatal"
	levelSlow   = "slow"
	levelStat   = "stat"

	backupFileDelimiter = "-"
	flags               = 0x0
)

const (
	callerKey    = "caller"
	contentKey   = "content"
	durationKey  = "duration"
	levelKey     = "level"
	spanKey      = "span"
	timestampKey = "@timestamp"
	traceKey     = "trace"
)

var (
	// ErrLogPathNotSet is an error that indicates the log path is not set.
	ErrLogPathNotSet = errors.New("log path must be set")
	// ErrLogServiceNameNotSet is an error that indicates that the service name is not set.
	ErrLogServiceNameNotSet = errors.New("log service name must be set")
)
