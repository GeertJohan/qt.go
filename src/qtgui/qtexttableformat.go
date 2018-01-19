//  header block begin
// /usr/include/qt/QtGui/qtextformat.h
// #include <qtextformat.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 29
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
type QTextTableFormat struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qtextformat.h:887
// index:0
// void QTextTableFormat()
func NewQTextTableFormat() *QTextTableFormat {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextTableFormatC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QTextTableFormat{cthis}
}

// /usr/include/qt/QtGui/qtextformat.h:889
// index:0
// inline
// bool isValid()
func (this *QTextTableFormat) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextTableFormat7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:891
// index:0
// inline
// int columns()
func (this *QTextTableFormat) Columns() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextTableFormat7columnsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:893
// index:0
// inline
// void setColumns(int)
func (this *QTextTableFormat) SetColumns(columns int) {
	// 0: (, int columns), (&columns)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextTableFormat10setColumnsEi", ffiqt.FFI_TYPE_VOID, this.cthis, &columns)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:898
// index:0
// inline
// QVector<QTextLength> columnWidthConstraints()
func (this *QTextTableFormat) ColumnWidthConstraints() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextTableFormat22columnWidthConstraintsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:901
// index:0
// inline
// void clearColumnWidthConstraints()
func (this *QTextTableFormat) ClearColumnWidthConstraints() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextTableFormat27clearColumnWidthConstraintsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:904
// index:0
// inline
// qreal cellSpacing()
func (this *QTextTableFormat) CellSpacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextTableFormat11cellSpacingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:906
// index:0
// inline
// void setCellSpacing(qreal)
func (this *QTextTableFormat) SetCellSpacing(spacing float64) {
	// 0: (, qreal spacing), (&spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextTableFormat14setCellSpacingEd", ffiqt.FFI_TYPE_VOID, this.cthis, &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:909
// index:0
// inline
// qreal cellPadding()
func (this *QTextTableFormat) CellPadding() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextTableFormat11cellPaddingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:911
// index:0
// inline
// void setCellPadding(qreal)
func (this *QTextTableFormat) SetCellPadding(padding float64) {
	// 0: (, qreal padding), (&padding)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextTableFormat14setCellPaddingEd", ffiqt.FFI_TYPE_VOID, this.cthis, &padding)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:914
// index:0
// inline
// Qt::Alignment alignment()
func (this *QTextTableFormat) Alignment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextTableFormat9alignmentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:917
// index:0
// inline
// void setHeaderRowCount(int)
func (this *QTextTableFormat) SetHeaderRowCount(count int) {
	// 0: (, int count), (&count)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextTableFormat17setHeaderRowCountEi", ffiqt.FFI_TYPE_VOID, this.cthis, &count)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:919
// index:0
// inline
// int headerRowCount()
func (this *QTextTableFormat) HeaderRowCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextTableFormat14headerRowCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
