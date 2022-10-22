package log

import "github.com/rs/zerolog/log"

var Debug = debugLog
var Info = infoLog
var Warn = warnLog
var Error = errorLog

func Reset() {
	Debug = debugLog
	Info = infoLog
	Warn = warnLog
	Error = errorLog
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
