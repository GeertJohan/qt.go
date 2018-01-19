//  header block begin
// /usr/include/qt/QtCore/qsharedmemory.h
// #include <qsharedmemory.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 42
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QSharedMemory struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qsharedmemory.h:54
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QSharedMemory) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSharedMemory10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:77
// index:0
// void QSharedMemory(class QObject *)
func NewQSharedMemory(parent unsafe.Pointer) *QSharedMemory {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSharedMemoryC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QSharedMemory{cthis}
}

// /usr/include/qt/QtCore/qsharedmemory.h:78
// index:1
// void QSharedMemory(const class QString &, class QObject *)
func NewQSharedMemory_1(key unsafe.Pointer, parent unsafe.Pointer) *QSharedMemory {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSharedMemoryC2ERK7QStringP7QObject", ffiqt.FFI_TYPE_VOID, cthis, key, parent)
	gopp.ErrPrint(err, rv)
	return &QSharedMemory{cthis}
}

// /usr/include/qt/QtCore/qsharedmemory.h:79
// index:0
// virtual
// void ~QSharedMemory()
func DeleteQSharedMemory(*QSharedMemory) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSharedMemoryD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:81
// index:0
// void setKey(const class QString &)
func (this *QSharedMemory) SetKey(key unsafe.Pointer) {
	// 0: (, const QString & key), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSharedMemory6setKeyERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:82
// index:0
// QString key()
func (this *QSharedMemory) Key() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSharedMemory3keyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:83
// index:0
// void setNativeKey(const class QString &)
func (this *QSharedMemory) SetNativeKey(key unsafe.Pointer) {
	// 0: (, const QString & key), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSharedMemory12setNativeKeyERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:84
// index:0
// QString nativeKey()
func (this *QSharedMemory) NativeKey() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSharedMemory9nativeKeyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:86
// index:0
// bool create(int, enum QSharedMemory::AccessMode)
func (this *QSharedMemory) Create(size int, mode int) {
	// 0: (, int size, QSharedMemory::AccessMode mode), (&size, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSharedMemory6createEiNS_10AccessModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &size, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:87
// index:0
// int size()
func (this *QSharedMemory) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSharedMemory4sizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:89
// index:0
// bool attach(enum QSharedMemory::AccessMode)
func (this *QSharedMemory) Attach(mode int) {
	// 0: (, QSharedMemory::AccessMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSharedMemory6attachENS_10AccessModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:90
// index:0
// bool isAttached()
func (this *QSharedMemory) IsAttached() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSharedMemory10isAttachedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:91
// index:0
// bool detach()
func (this *QSharedMemory) Detach() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSharedMemory6detachEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:93
// index:0
// void * data()
func (this *QSharedMemory) Data() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSharedMemory4dataEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:95
// index:1
// const void * data()
func (this *QSharedMemory) Data_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSharedMemory4dataEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:94
// index:0
// const void * constData()
func (this *QSharedMemory) ConstData() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSharedMemory9constDataEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:98
// index:0
// bool lock()
func (this *QSharedMemory) Lock() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSharedMemory4lockEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:99
// index:0
// bool unlock()
func (this *QSharedMemory) Unlock() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSharedMemory6unlockEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:102
// index:0
// QSharedMemory::SharedMemoryError error()
func (this *QSharedMemory) Error() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSharedMemory5errorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedmemory.h:103
// index:0
// QString errorString()
func (this *QSharedMemory) ErrorString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSharedMemory11errorStringEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
