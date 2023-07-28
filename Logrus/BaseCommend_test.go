package Logrus

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestBase(t *testing.T) {
	Base()
}

func TestSetFields(t *testing.T) {
	SetFields()
}

func TestSetFields2(t *testing.T) {
	SetFields2()
}

func TestCustomFormat(t *testing.T) {
	CustomFormat()
}

// Levels 作用于那个日志生效的
func (f MyHook) Levels() []logrus.Level {
	//return []logrus.Level{logrus.ErrorLevel}
	return logrus.AllLevels
}

// 每一次做的事情
func (f MyHook) Fire(entry *logrus.Entry) error {
	//file, _ := os.OpenFile("./error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//line, _ := entry.String()
	//file.Write([]byte(line))
	//logrus.WithField("name", "Sakura")
	entry.Data["name"] = "sakura"
	return nil
}

func CustomFormat() {
	logrus.AddHook(&MyHook{})

	logrus.Warnln("warn")
	logrus.Errorln("Error")

}
