// mypackage.go
package mypackage

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)

func Test() string {
	var app = qml.NewQQmlApplicationEngine(nil)
	app.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	return "test"
}
