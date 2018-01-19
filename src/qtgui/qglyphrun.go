//  header block begin
// /usr/include/qt/QtGui/qglyphrun.h
// #include <qglyphrun.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 36
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
}

//  ext block end

//  body block begin
type QGlyphRun struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qglyphrun.h:67
// index:0
// void QGlyphRun()
func NewQGlyphRun() *QGlyphRun {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGlyphRunC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QGlyphRun{cthis}
}

// /usr/include/qt/QtGui/qglyphrun.h:73
// index:0
// void ~QGlyphRun()
func DeleteQGlyphRun(*QGlyphRun) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGlyphRunD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:75
// index:0
// inline
// void swap(class QGlyphRun &)
func (this *QGlyphRun) Swap(other unsafe.Pointer) {
	// 0: (, QGlyphRun & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGlyphRun4swapERS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:77
// index:0
// QRawFont rawFont()
func (this *QGlyphRun) RawFont() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGlyphRun7rawFontEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:78
// index:0
// void setRawFont(const class QRawFont &)
func (this *QGlyphRun) SetRawFont(rawFont unsafe.Pointer) {
	// 0: (, const QRawFont & rawFont), (rawFont)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGlyphRun10setRawFontERK8QRawFont", ffiqt.FFI_TYPE_VOID, this.cthis, rawFont)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:80
// index:0
// void setRawData(const quint32 *, const class QPointF *, int)
func (this *QGlyphRun) SetRawData(glyphIndexArray unsafe.Pointer, glyphPositionArray unsafe.Pointer, size int) {
	// 0: (, const quint32 * glyphIndexArray, const QPointF * glyphPositionArray, int size), (glyphIndexArray, glyphPositionArray, &size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGlyphRun10setRawDataEPKjPK7QPointFi", ffiqt.FFI_TYPE_VOID, this.cthis, glyphIndexArray, glyphPositionArray, &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:84
// index:0
// QVector<quint32> glyphIndexes()
func (this *QGlyphRun) GlyphIndexes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGlyphRun12glyphIndexesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:87
// index:0
// QVector<QPointF> positions()
func (this *QGlyphRun) Positions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGlyphRun9positionsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:90
// index:0
// void clear()
func (this *QGlyphRun) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGlyphRun5clearEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:96
// index:0
// void setOverline(_Bool)
func (this *QGlyphRun) SetOverline(overline bool) {
	// 0: (, bool overline), (&overline)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGlyphRun11setOverlineEb", ffiqt.FFI_TYPE_VOID, this.cthis, &overline)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:97
// index:0
// bool overline()
func (this *QGlyphRun) Overline() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGlyphRun8overlineEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:99
// index:0
// void setUnderline(_Bool)
func (this *QGlyphRun) SetUnderline(underline bool) {
	// 0: (, bool underline), (&underline)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGlyphRun12setUnderlineEb", ffiqt.FFI_TYPE_VOID, this.cthis, &underline)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:100
// index:0
// bool underline()
func (this *QGlyphRun) Underline() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGlyphRun9underlineEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:102
// index:0
// void setStrikeOut(_Bool)
func (this *QGlyphRun) SetStrikeOut(strikeOut bool) {
	// 0: (, bool strikeOut), (&strikeOut)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGlyphRun12setStrikeOutEb", ffiqt.FFI_TYPE_VOID, this.cthis, &strikeOut)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:103
// index:0
// bool strikeOut()
func (this *QGlyphRun) StrikeOut() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGlyphRun9strikeOutEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:105
// index:0
// void setRightToLeft(_Bool)
func (this *QGlyphRun) SetRightToLeft(on bool) {
	// 0: (, bool on), (&on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGlyphRun14setRightToLeftEb", ffiqt.FFI_TYPE_VOID, this.cthis, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:106
// index:0
// bool isRightToLeft()
func (this *QGlyphRun) IsRightToLeft() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGlyphRun13isRightToLeftEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:108
// index:0
// void setFlag(enum QGlyphRun::GlyphRunFlag, _Bool)
func (this *QGlyphRun) SetFlag(flag int, enabled bool) {
	// 0: (, QGlyphRun::GlyphRunFlag flag, bool enabled), (&flag, &enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGlyphRun7setFlagENS_12GlyphRunFlagEb", ffiqt.FFI_TYPE_VOID, this.cthis, &flag, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:110
// index:0
// QGlyphRun::GlyphRunFlags flags()
func (this *QGlyphRun) Flags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGlyphRun5flagsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:112
// index:0
// void setBoundingRect(const class QRectF &)
func (this *QGlyphRun) SetBoundingRect(boundingRect unsafe.Pointer) {
	// 0: (, const QRectF & boundingRect), (boundingRect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGlyphRun15setBoundingRectERK6QRectF", ffiqt.FFI_TYPE_VOID, this.cthis, boundingRect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:113
// index:0
// QRectF boundingRect()
func (this *QGlyphRun) BoundingRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGlyphRun12boundingRectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:115
// index:0
// bool isEmpty()
func (this *QGlyphRun) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGlyphRun7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
