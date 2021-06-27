package abstractfactory

import (
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	var a = AbstractFactory()
	a.CreateProxy().Call()
	a.CreateNotify().Exec()
}
