package strategy

import "testing"

func Test(t *testing.T) {
	var myLog = Log{}
	myLog.ILogging = DBLog{}
	myLog.ILogging.Info()
	myLog.ILogging.Error()

	myLog.ILogging = FileLog{}
	myLog.ILogging.Info()
	myLog.ILogging.Error()
}
