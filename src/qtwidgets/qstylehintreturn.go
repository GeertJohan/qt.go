//  header block begin
// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QStyleHintReturn struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qstyleoption.h:710
// index:0
// void QStyleHintReturn(int, int)
func NewQStyleHintReturn(version int, type_ int) *QStyleHintReturn {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QStyleHintReturnC2Eii", ffiqt.FFI_TYPE_VOID, cthis, &version, &type_)
	gopp.ErrPrint(err, rv)
	return &QStyleHintReturn{cthis}
}

// /usr/include/qt/QtWidgets/qstyleoption.h:711
// index:0
// void ~QStyleHintReturn()
func DeleteQStyleHintReturn(*QStyleHintReturn) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QStyleHintReturnD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end
