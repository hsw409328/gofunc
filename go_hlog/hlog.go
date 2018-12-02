package go_hlog

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	color_black = uint8(iota + 90)
	color_red
	color_green
	color_yellow
	color_blue
	color_magenta
	color_cyan
	color_white
)

const (
	verbose = "[VERB]"
	trace   = "[TRACE]"
	errors  = "[ERROR]"
	warn    = "[WARN]"
	info    = "[INFO]"
	debug   = "[DEBUG]"
	assert  = "[ASST]"
)

type Logger struct {
	lg *log.Logger
	w  io.Writer
}

func NewLogger(w io.Writer) *Logger {
	hlogObject := &Logger{w: w}
	hlogObject.SetLogFile()
	return hlogObject
}

func GetInstance(fileName string) (logger *Logger) {
	if fileName != "" {
		f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			logger = NewLogger(os.Stdout)
			logger.Error(err.Error())
		} else {
			logger = NewLogger(f)
		}
	} else {
		logger = NewLogger(os.Stdout)
	}

	return
}

func (ctx *Logger) SetLogFile() {
	ctx.lg = log.New(ctx.w, "", 0)
	ctx.lg.SetOutput(ctx.w)
	ctx.lg.SetFlags(log.Lshortfile | log.LstdFlags)
}

func (ctx *Logger) Verbose(message interface{}) {
	level := ctx.formatLevel(verbose)
	ctx.lg.Output(2, fmt.Sprintln(ctx.formatPrefix(level), message))
}

func (ctx *Logger) VerboseF(format string, a ...interface{}) {
	level := ctx.formatLevel(verbose)
	ctx.lg.Output(2, fmt.Sprintln(ctx.formatPrefix(level), fmt.Sprintf(format, a...)))
}

func (ctx *Logger) Trace(message interface{}) {
	level := ctx.formatLevel(trace)
	ctx.lg.Output(2, fmt.Sprintln(ctx.formatPrefix(level), message))
}

func (ctx *Logger) TraceF(format string, a ...interface{}) {
	level := ctx.formatLevel(trace)
	ctx.lg.Output(2, fmt.Sprintln(ctx.formatPrefix(level), fmt.Sprintf(format, a...)))
}
func (ctx *Logger) Error(message interface{}) {
	level := ctx.formatLevel(errors)
	ctx.lg.Output(2, fmt.Sprintln(ctx.formatPrefix(level), message))
}

func (ctx *Logger) ErrorF(format string, a ...interface{}) {
	level := ctx.formatLevel(errors)
	ctx.lg.Output(2, fmt.Sprintln(ctx.formatPrefix(level), fmt.Sprintf(format, a...)))
}

func (ctx *Logger) Warn(message interface{}) {
	level := ctx.formatLevel(warn)
	ctx.lg.Output(2, fmt.Sprintln(ctx.formatPrefix(level), message))
}

func (ctx *Logger) WarnF(format string, a ...interface{}) {
	level := ctx.formatLevel(warn)
	ctx.lg.Output(2, fmt.Sprintln(ctx.formatPrefix(level), fmt.Sprintf(format, a...)))
}

func (ctx *Logger) Info(message interface{}) {
	level := ctx.formatLevel(info)
	ctx.lg.Output(2, fmt.Sprintln(ctx.formatPrefix(level), message))
}

func (ctx *Logger) InfoF(format string, a ...interface{}) {
	level := ctx.formatLevel(info)
	ctx.lg.Output(2, fmt.Sprintln(ctx.formatPrefix(level), fmt.Sprintf(format, a...)))
}

func (ctx *Logger) Debug(message interface{}) {
	level := ctx.formatLevel(debug)
	ctx.lg.Output(2, fmt.Sprintln(ctx.formatPrefix(level), message))
}

func (ctx *Logger) DebugF(format string, a ...interface{}) {
	level := ctx.formatLevel(debug)
	ctx.lg.Output(2, fmt.Sprintln(ctx.formatPrefix(level), fmt.Sprintf(format, a...)))
}

func (ctx *Logger) Asset(message interface{}) {
	level := ctx.formatLevel(assert)
	ctx.lg.Output(2, fmt.Sprintln(ctx.formatPrefix(level), message))
}

func (ctx *Logger) AssetF(format string, a ...interface{}) {
	level := ctx.formatLevel(assert)
	ctx.lg.Output(2, fmt.Sprintln(ctx.formatPrefix(level), fmt.Sprintf(format, a...)))
}

func (ctx *Logger) formatLevel(level string) string {
	switch level {
	case verbose:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_black, level)
	case trace:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_cyan, level)
	case errors:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_red, level)
	case warn:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_yellow, level)
	case info:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_green, level)
	case debug:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_blue, level)
	case assert:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_magenta, level)
	default:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_black, level)
	}
}

func (ctx *Logger) formatPrefix(level string) string {
	return fmt.Sprintf("%s :", level)
}
