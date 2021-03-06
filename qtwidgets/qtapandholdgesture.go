package qtwidgets

// /usr/include/qt/QtWidgets/qgesture.h
// #include <qgesture.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*

 */
type QTapAndHoldGesture struct {
	*QGesture
}
type QTapAndHoldGesture_ITF interface {
	QGesture_ITF
	QTapAndHoldGesture_PTR() *QTapAndHoldGesture
}

func (ptr *QTapAndHoldGesture) QTapAndHoldGesture_PTR() *QTapAndHoldGesture { return ptr }

func (this *QTapAndHoldGesture) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGesture.GetCthis()
	}
}
func (this *QTapAndHoldGesture) SetCthis(cthis unsafe.Pointer) {
	this.QGesture = NewQGestureFromPointer(cthis)
}
func NewQTapAndHoldGestureFromPointer(cthis unsafe.Pointer) *QTapAndHoldGesture {
	bcthis0 := NewQGestureFromPointer(cthis)
	return &QTapAndHoldGesture{bcthis0}
}
func (*QTapAndHoldGesture) NewFromPointer(cthis unsafe.Pointer) *QTapAndHoldGesture {
	return NewQTapAndHoldGestureFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgesture.h:254
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QTapAndHoldGesture) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QTapAndHoldGesture10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgesture.h:260
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTapAndHoldGesture(QObject *)

/*

 */
func NewQTapAndHoldGesture(parent qtcore.QObject_ITF /*777 QObject **/) *QTapAndHoldGesture {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QTapAndHoldGestureC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTapAndHoldGestureFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTapAndHoldGesture")
	return gothis
}

// /usr/include/qt/QtWidgets/qgesture.h:260
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTapAndHoldGesture(QObject *)

/*

 */
func NewQTapAndHoldGesture__() *QTapAndHoldGesture {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN18QTapAndHoldGestureC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTapAndHoldGestureFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTapAndHoldGesture")
	return gothis
}

// /usr/include/qt/QtWidgets/qgesture.h:261
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTapAndHoldGesture()

/*

 */
func DeleteQTapAndHoldGesture(this *QTapAndHoldGesture) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QTapAndHoldGestureD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgesture.h:263
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF position() const

/*

 */
func (this *QTapAndHoldGesture) Position() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QTapAndHoldGesture8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgesture.h:264
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPosition(const QPointF &)

/*

 */
func (this *QTapAndHoldGesture) SetPosition(pos qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QTapAndHoldGesture11setPositionERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:266
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setTimeout(int)

/*

 */
func (this *QTapAndHoldGesture) SetTimeout(msecs int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QTapAndHoldGesture10setTimeoutEi", qtrt.FFI_TYPE_POINTER, msecs)
	qtrt.ErrPrint(err, rv)
}
func QTapAndHoldGesture_SetTimeout(msecs int) {
	var nilthis *QTapAndHoldGesture
	nilthis.SetTimeout(msecs)
}

// /usr/include/qt/QtWidgets/qgesture.h:267
// index:0
// Public static Visibility=Default Availability=Available
// [4] int timeout()

/*

 */
func (this *QTapAndHoldGesture) Timeout() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QTapAndHoldGesture7timeoutEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QTapAndHoldGesture_Timeout() int {
	var nilthis *QTapAndHoldGesture
	rv := nilthis.Timeout()
	return rv
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
