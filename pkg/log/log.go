package log

import "github.com/rs/zerolog/log"

type Logger func(format string, v ...interface{})

func Debug(format string, v ...interface{}) {
	_debug(format, v...)
}

func Info(format string, v ...interface{}) {
	_info(format, v...)
}

func Warn(format string, v ...interface{}) {
	_warn(format, v...)
}

func Error(format string, v ...interface{}) {
	_error(format, v...)
}

func SetLogger(debug, info, warn, error Logger) {
	if debug == nil {
		debug = debugLog
	}
	_debug = debug
	if info == nil {
		info = infoLog
	}
	_info = info
	if warn == nil {
		warn = warnLog
	}
	_warn = warn
	if error == nil {
		error = errorLog
	}
	_error = error
}

func SetDebugLogger(debug Logger) {
	_debug = debug
}

func SetInfoLogger(info Logger) {
	_info = info
}

func SetWarnLogger(warn Logger) {
	_warn = warn
}

func SetErrorLogger(error Logger) {
	_error = error
}

var _debug = debugLog
var _info = infoLog
var _warn = warnLog
var _error = errorLog

func Reset() {
	_debug = debugLog
	_info = infoLog
	_warn = warnLog
	_error = errorLog
}

func debugLog(format string, v ...interface{}) {
	log.Debug().Msgf(format, v...)
}

func infoLog(format string, v ...interface{}) {
	log.Info().Msgf(format, v...)
}

func warnLog(format string, v ...interface{}) {
	log.Warn().Msgf(format, v...)
}

func errorLog(format string, v ...interface{}) {
	log.Error().Msgf(format, v...)
}
