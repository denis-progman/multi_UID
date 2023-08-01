package errs

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func getCurrentTime() string {
	return time.Now().Format(time.ANSIC)
}

const (
	death = 0
	stop = 1
	warning = 2
	note = 3
)

type DefaultError struct {
	Time string `json:"time"`
	Message string `json:"message"`
	Code int `json:"code"`
	File string `json:"file"`
	Line int `json:"line"`
	Method string `json:"method"`
	Behavior int `json:"behavior"`
	Trace string `json:"trace"`
}

type Error interface {
	setActionPlace()
	getTraceRecord() string
	ToString(name string) string
}



func (err *DefaultError) setActionPlace() {
	pc, file, line, _ := runtime.Caller(1)
	err.Method = fmt.Sprintf("%T", runtime.FuncForPC(pc).Name())
	err.File = file
	err.Line = line
}



func (err *DefaultError) getTraceRecord() string {
	return fmt.Sprintf("%s:%d <%s>", err.Method, err.Line, err.File)
}

func newOne(behavior int, message string, tracing bool, code int) Error {
	errStruct := DefaultError{
		Time: getCurrentTime(),
		Message: message,
		Code: code,
		Behavior: behavior,
	}
	if (tracing) {
		errStruct.Trace = errStruct.getTraceRecord()
	}
	var err Error = &errStruct
	err.setActionPlace()
	return err
}

func (err *DefaultError) ToString(name string) string {
	trace := ""
	if (err.Trace != "") {
		trace = fmt.Sprintf("trace: \n%s\n", err.Trace)
	}
	return fmt.Sprintf("%s: %s code: %d\n%s\n%s\n", name, err.Message, err.Code, err.getTraceRecord(), trace)
}

func DeathWill(message string, code int) {
	name := "DeathWill"
	err := newOne(death, fmt.Sprintf("%s: %s", name, message), true, code)
	log.Fatal(err.ToString(name))
}

func StopException(message string, code int) Error {
	name := "StopException"
	err := newOne(stop, fmt.Sprintf("%s: %s", name, message), true, code)
	return err
}

func Warning(message string, code int) Error {
	name := "Warning"
	err := newOne(warning, fmt.Sprintf("%s: %s", name, message), true, code)
	return err
}

func Note(message string, code int) {
	name := "Warning"
	err := newOne(note, fmt.Sprintf("%s: %s", name, message), true, code)
	log.Println(err.ToString(name))
}



// type DeathWill struct {
	
// }


type Collection struct {

}