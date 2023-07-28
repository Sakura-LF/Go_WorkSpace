package Logrus

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func Base() {
	logrus.Error("出错了")
	logrus.Warnln("警告")
	logrus.Infof("信息")
	logrus.Debugf("debug")
	logrus.Println("打印")

	fmt.Println(logrus.GetLevel())

}

func SetFields() {
	log := logrus.WithFields(logrus.Fields{
		"user_id": "Sakura",
		"ip":      "192.168.57.33",
		"method":  "post",
	})
	log.Error("出错")
}

func SetFields2() {
	log := logrus.WithField("app", "Study").WithField("service", "logrus")
	log.Error("你好")
}

//func CustomFormat() {
//	logrus.SetFormatter(&logrus.TextFormatter{
//		ForceColors: true,
//	})
//	file, _ := os.OpenFile("./logrus.log", os.O_CREATE|os.O_CREATE|os.O_APPEND, 0666)
//
//	logrus.SetOutput(io.MultiWriter(file, os.Stdout))
//	logrus.Info("信息")
//	logrus.Errorln("错误")
//}

type MyHook struct {
}
