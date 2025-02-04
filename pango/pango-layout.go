/*
 * Copyright (c) 2015- terrak <terrak1975@gmail.com>
 *
 * This file originated from: http://www.terrak.net/
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package pango

// #include <pango/pango.h>
// #include "pango.go.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/go-gst/go-glib/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.pango_alignment_get_type()), marshalAlignment},
		{glib.Type(C.pango_ellipsize_mode_get_type()), marshalEllipsizeMode},
		{glib.Type(C.pango_wrap_mode_get_type()), marshalWrapMode},
		{glib.Type(C.pango_tab_align_get_type()), marshalTabAlign},

		// Objects/Interfaces
		//		{glib.Type(C.pango_layout_get_type()), marshalLayout},
	}
	glib.RegisterGValueMarshalers(tm)
}

// Layout is a representation of PangoLayout.
type Layout struct {
	pangoLayout *C.PangoLayout
}

// Native returns a pointer to the underlying PangoLayout.
func (v *Layout) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *Layout) native() *C.PangoLayout {
	if v == nil {
		return nil
	}
	return (*C.PangoLayout)(unsafe.Pointer(v.pangoLayout))
}

func WrapLayout(p uintptr) *Layout {
	layout := new(Layout)
	layout.pangoLayout = (*C.PangoLayout)(unsafe.Pointer(p))
	return layout
}

// LayoutLine is a representation of PangoLayoutLine.
type LayoutLine struct {
	pangoLayoutLine *C.PangoLayout
}

// Native returns a pointer to the underlying PangoLayoutLine.
func (v *LayoutLine) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *LayoutLine) native() *C.PangoLayoutLine {
	if v == nil {
		return nil
	}
	return (*C.PangoLayoutLine)(unsafe.Pointer(v.pangoLayoutLine))
}

/*
 * Constants
 */

// Alignment is a representation of Pango's PangoAlignment.
type Alignment int

const (
	ALIGN_LEFT   Alignment = C.PANGO_ALIGN_LEFT
	ALIGN_CENTER Alignment = C.PANGO_ALIGN_CENTER
	ALIGN_RIGHT  Alignment = C.PANGO_ALIGN_RIGHT
)

func marshalAlignment(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Alignment(c), nil
}

// WrapMode is a representation of Pango's PangoWrapMode.
type WrapMode int

const (
	WRAP_WORD      WrapMode = C.PANGO_WRAP_WORD
	WRAP_CHAR      WrapMode = C.PANGO_WRAP_CHAR
	WRAP_WORD_CHAR WrapMode = C.PANGO_WRAP_WORD_CHAR
)

func marshalWrapMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return WrapMode(c), nil
}

// EllipsizeMode is a representation of Pango's PangoEllipsizeMode.
type EllipsizeMode int

const (
	ELLIPSIZE_NONE   EllipsizeMode = C.PANGO_ELLIPSIZE_NONE
	ELLIPSIZE_START  EllipsizeMode = C.PANGO_ELLIPSIZE_START
	ELLIPSIZE_MIDDLE EllipsizeMode = C.PANGO_ELLIPSIZE_MIDDLE
	ELLIPSIZE_END    EllipsizeMode = C.PANGO_ELLIPSIZE_END
)

func marshalEllipsizeMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return EllipsizeMode(c), nil
}

type TabAlign int

const (
	TAB_LEFT TabAlign = C.PANGO_TAB_LEFT
)

func marshalTabAlign(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return TabAlign(c), nil
}

/*
func marshalLayout(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapLayout(obj), nil
}

func wrapLayout(obj *glib.Object) *Layout {
	return &Layout{obj}
}
*/

// LayoutNew is a wrapper around pango_layout_new().
func LayoutNew(context *Context) *Layout {
	c := C.pango_layout_new(context.native())

	layout := new(Layout)
	layout.pangoLayout = (*C.PangoLayout)(c)
	return layout
}

// Copy is a wrapper around pango_layout_copy().
func (v *Layout) Copy() *Layout {
	c := C.pango_layout_copy(v.native())

	layout := new(Layout)
	layout.pangoLayout = (*C.PangoLayout)(c)
	return layout
}

// GetContext is a wrapper around pango_layout_get_context().
func (v *Layout) GetContext() *Context {
	c := C.pango_layout_get_context(v.native())

	context := new(Context)
	context.pangoContext = (*C.PangoContext)(c)

	return context
}

// SetAttributes is a wrapper around pango_layout_set_attributes().
func (v *Layout) SetAttributes(attrs *AttrList) {
	C.pango_layout_set_attributes(v.native(), attrs.native())
}

// GetAttributes is a wrapper around pango_layout_get_attributes().
func (v *Layout) GetAttributes() *AttrList {
	c := C.pango_layout_get_attributes(v.native())
	return wrapAttrList(c)
}

// SetText is a wrapper around pango_layout_set_text().
func (v *Layout) SetText(text string, length int) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.pango_layout_set_text(v.native(), (*C.char)(cstr), (C.int)(length))
}

// GetText is a wrapper around pango_layout_get_text().
func (v *Layout) GetText() string {
	c := C.pango_layout_get_text(v.native())
	return C.GoString((*C.char)(c))
}

// GetCharacterCount is a wrapper around pango_layout_get_character_count().
func (v *Layout) GetCharacterCount() int {
	c := C.pango_layout_get_character_count(v.native())
	return int(c)
}

// SetMarkup is a wrapper around pango_layout_set_markup().
func (v *Layout) SetMarkup(text string, length int) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.pango_layout_set_markup(v.native(), (*C.char)(cstr), (C.int)(length))
}

//void           pango_layout_set_markup_with_accel (PangoLayout    *layout,
//						   const char     *markup,
//						   int             length,
//						   gunichar        accel_marker,
//						   gunichar       *accel_char);

/*
func (v *Layout) SetMarkupWithAccel (text string, length int, accel_marker, accel_char rune){
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.pango_layout_set_markup_with_accel (v.native(),  (*C.char)(cstr), (C.int)(length), (C.gunichar)(accel_marker), (C.gunichar)(accel_char) )
}
*/

// SetFontDescription is a wrapper around pango_layout_set_font_description().
func (v *Layout) SetFontDescription(desc *FontDescription) {
	C.pango_layout_set_font_description(v.native(), desc.native())
}

// GetFontDescription is a wrapper around pango_layout_get_font_description().
func (v *Layout) GetFontDescription() *FontDescription {
	c := C.pango_layout_get_font_description(v.native())

	desc := new(FontDescription)
	desc.pangoFontDescription = (*C.PangoFontDescription)(c)

	return desc
}

// SetWidth is a wrapper around pango_layout_set_width().
func (v *Layout) SetWidth(width int) {
	C.pango_layout_set_width(v.native(), C.int(width))
}

// GetWidth is a wrapper around pango_layout_get_width().
func (v *Layout) GetWidth() int {
	c := C.pango_layout_get_width(v.native())
	return int(c)
}

// SetHeight is a wrapper around pango_layout_set_height().
func (v *Layout) SetHeight(width int) {
	C.pango_layout_set_height(v.native(), C.int(width))
}

// GetHeight is a wrapper around pango_layout_get_height().
func (v *Layout) GetHeight() int {
	c := C.pango_layout_get_height(v.native())
	return int(c)
}

// SetWrap is a wrapper around pango_layout_set_wrap().
func (v *Layout) SetWrap(wrap WrapMode) {
	C.pango_layout_set_wrap(v.native(), C.PangoWrapMode(wrap))
}

// WrapMode is a wrapper around pango_layout_get_wrap().
func (v *Layout) GetWrap() WrapMode {
	c := C.pango_layout_get_wrap(v.native())
	return WrapMode(c)
}

// IsWrapped is a wrapper around pango_layout_is_wrapped().
func (v *Layout) IsWrapped() bool {
	c := C.pango_layout_is_wrapped(v.native())
	return gobool(c)
}

// SetIndent is a wrapper around pango_layout_set_indent().
func (v *Layout) SetIndent(indent int) {
	C.pango_layout_set_indent(v.native(), C.int(indent))
}

// GetIndent is a wrapper around pango_layout_get_indent().
func (v *Layout) GetIndent() int {
	c := C.pango_layout_get_indent(v.native())
	return int(c)
}

// SetTabs is a wrapper around pango_layout_set_tabs().
func (v *Layout) SetTabs(tabs *TabArray) {
	C.pango_layout_set_tabs(v.native(), tabs.native())
}

// GetTabs is a wrapper around pango_layout_get_tabs().
func (v *Layout) GetTabs() (*TabArray, error) {
	c := C.pango_layout_get_tabs(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	ta := wrapTabArray(c)
	runtime.SetFinalizer(ta, func(v *TabArray) { FinalizerStrategy(v.free) })
	return ta, nil
}

// GetSize is a wrapper around pango_layout_get_size().
func (v *Layout) GetSize() (int, int) {
	var w, h C.int
	C.pango_layout_get_size(v.native(), &w, &h)
	return int(w), int(h)
}

/*
 * TabArray
 */

// TabArray is a representation of PangoTabArray.
type TabArray struct {
	pangoTabArray *C.PangoTabArray
}

// Native returns a pointer to the underlying PangoTabArray.
func (v *TabArray) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *TabArray) native() *C.PangoTabArray {
	if v == nil {
		return nil
	}
	return (*C.PangoTabArray)(unsafe.Pointer(v.pangoTabArray))
}

func wrapTabArray(tabArray *C.PangoTabArray) *TabArray {
	return &TabArray{tabArray}
}

func WrapTabArray(p uintptr) *TabArray {
	tabArray := new(TabArray)
	tabArray.pangoTabArray = (*C.PangoTabArray)(unsafe.Pointer(p))
	return tabArray
}

// TabArrayNew is a wrapper around pango_tab_array_new().
func TabArrayNew(initialSize int, positionsInPixels bool) *TabArray {
	c := C.pango_tab_array_new(C.gint(initialSize), gbool(positionsInPixels))

	tabArray := new(TabArray)
	runtime.SetFinalizer(tabArray, func(v *TabArray) { FinalizerStrategy(v.free) })
	tabArray.pangoTabArray = (*C.PangoTabArray)(c)
	return tabArray
}

// TabArrayNewWithPositions is a wrapper around pango_tab_array_new_with_positions().
// func TabArrayNewWithPositions(size int, positionsInPixels bool, ...) *TabArray {
// 	c := C.pango_tab_array_new_with_positions(C.gint(size), gbool(positionsInPixels), ...)

// 	tabArray := new(TabArray)
//	runtime.SetFinalizer(e, func(v *TabArray) { FinalizerStrategy(v.free) })
// 	tabArray.pangoTabArray = (*C.PangoTabArray)(c)
// 	return tabArray
// }

// Copy is a wrapper around pango_tab_array_copy().
func (v *TabArray) Copy() (*TabArray, error) {
	c := C.pango_tab_array_copy(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	ta := wrapTabArray(c)
	runtime.SetFinalizer(ta, func(v *TabArray) { FinalizerStrategy(v.free) })
	return ta, nil
}

// free is a wrapper around pango_tab_array_free().
func (v *TabArray) free() {
	C.pango_tab_array_free(v.native())
}

// free is a wrapper around pango_tab_array_free().
// This is only to enable other packages within gotk. Should not be used outside the gotk library.
func (v *TabArray) Free() {
	C.pango_tab_array_free(v.native())
}

// GetSize is a wrapper around pango_tab_array_get_size().
func (v *TabArray) GetSize() int {
	return int(C.pango_tab_array_get_size(v.native()))
}

// Resize is a wrapper around pango_tab_array_resize().
func (v *TabArray) Resize(newSize int) {
	C.pango_tab_array_resize(v.native(), C.gint(newSize))
}

// SetTab is a wrapper around pango_tab_array_set_tab().
func (v *TabArray) SetTab(tabIndex int, alignment TabAlign, location int) {
	C.pango_tab_array_set_tab(v.native(), C.gint(tabIndex), C.PangoTabAlign(alignment), C.gint(location))
}

// GetTab is a wrapper around pango_tab_array_get_tab().
func (v *TabArray) GetTab(tabIndex int) (TabAlign, int) {
	var alignment C.PangoTabAlign
	var location C.gint
	C.pango_tab_array_get_tab(v.native(), C.gint(tabIndex), &alignment, &location)
	return TabAlign(alignment), int(location)
}

// GetTabs is a wrapper around pango_tab_array_get_tabs().
// func (v *TabArray) GetTabs() ([]TabAlign, []int) {
// 	var alignment *C.PangoTabAlign
// 	var location *C.gint

// 	C.pango_tab_array_get_tabs(v.native(), &alignment, &location)

// 	size := v.GetSize()

// 	var goAlignments []TabAlign
// 	var goLocations []int

// 	if &alignment != nil {
// 		var ginthelp C.gint
// 		sizeOf := unsafe.Sizeof(ginthelp)
// 		for i := 0; i < int(size); i++ {
// 			goAlignmentElement := TabAlign(*((*C.gint)(unsafe.Pointer(location))))
// 			goAlignments = append(goAlignments, goAlignmentElement)
// 			location += sizeOf
// 		}
// 	}

// 	if &location != nil {
// 		var ginthelp C.gint
// 		sizeOf := unsafe.Sizeof(ginthelp)
// 		for i := 0; i < int(size); i++ {
// 			goLocationElement := int(*((*C.gint)(unsafe.Pointer(location))))
// 			goLocations = append(goLocations, goLocationElement)
// 			location += sizeOf
// 		}

// 		// TODO: free locations
// 	}

// 	return goAlignments, goLocations
// }

// GetPositionsInPixels is a wrapper around pango_tab_array_get_positions_in_pixels().
func (v *TabArray) GetPositionsInPixels() bool {
	return gobool(C.pango_tab_array_get_positions_in_pixels(v.native()))
}

//void           pango_layout_set_spacing          (PangoLayout                *layout,
//						  int                         spacing);
//int            pango_layout_get_spacing          (PangoLayout                *layout);
//void           pango_layout_set_justify          (PangoLayout                *layout,
//						  gboolean                    justify);
//gboolean       pango_layout_get_justify          (PangoLayout                *layout);
//void           pango_layout_set_auto_dir         (PangoLayout                *layout,
//						  gboolean                    auto_dir);
//gboolean       pango_layout_get_auto_dir         (PangoLayout                *layout);
//void           pango_layout_set_alignment        (PangoLayout                *layout,
//						  PangoAlignment              alignment);
//PangoAlignment pango_layout_get_alignment        (PangoLayout                *layout);
//
//void           pango_layout_set_single_paragraph_mode (PangoLayout                *layout,
//						       gboolean                    setting);
//gboolean       pango_layout_get_single_paragraph_mode (PangoLayout                *layout);
//
//void               pango_layout_set_ellipsize (PangoLayout        *layout,
//					       PangoEllipsizeMode  ellipsize);
//PangoEllipsizeMode pango_layout_get_ellipsize (PangoLayout        *layout);
//gboolean           pango_layout_is_ellipsized (PangoLayout        *layout);
//
//int      pango_layout_get_unknown_glyphs_count (PangoLayout    *layout);
//
//void     pango_layout_context_changed (PangoLayout    *layout);
//guint    pango_layout_get_serial      (PangoLayout    *layout);
//
//void     pango_layout_get_log_attrs (PangoLayout    *layout,
//				     PangoLogAttr  **attrs,
//				     gint           *n_attrs);
//
//const PangoLogAttr *pango_layout_get_log_attrs_readonly (PangoLayout *layout,
//							 gint        *n_attrs);
//
//void     pango_layout_index_to_pos         (PangoLayout    *layout,
//					    int             index_,
//					    PangoRectangle *pos);
//void     pango_layout_index_to_line_x      (PangoLayout    *layout,
//					    int             index_,
//					    gboolean        trailing,
//					    int            *line,
//					    int            *x_pos);
//void     pango_layout_get_cursor_pos       (PangoLayout    *layout,
//					    int             index_,
//					    PangoRectangle *strong_pos,
//					    PangoRectangle *weak_pos);
//void     pango_layout_move_cursor_visually (PangoLayout    *layout,
//					    gboolean        strong,
//					    int             old_index,
//					    int             old_trailing,
//					    int             direction,
//					    int            *new_index,
//					    int            *new_trailing);
//gboolean pango_layout_xy_to_index          (PangoLayout    *layout,
//					    int             x,
//					    int             y,
//					    int            *index_,
//					    int            *trailing);
//void     pango_layout_get_extents          (PangoLayout    *layout,
//					    PangoRectangle *ink_rect,
//					    PangoRectangle *logical_rect);
//void     pango_layout_get_pixel_extents    (PangoLayout    *layout,
//					    PangoRectangle *ink_rect,
//					    PangoRectangle *logical_rect);
//void     pango_layout_get_pixel_size       (PangoLayout    *layout,
//					    int            *width,
//					    int            *height);
//int      pango_layout_get_baseline         (PangoLayout    *layout);
//
//int              pango_layout_get_line_count       (PangoLayout    *layout);
//PangoLayoutLine *pango_layout_get_line             (PangoLayout    *layout,
//						    int             line);
//PangoLayoutLine *pango_layout_get_line_readonly    (PangoLayout    *layout,
//						    int             line);
//GSList *         pango_layout_get_lines            (PangoLayout    *layout);
//GSList *         pango_layout_get_lines_readonly   (PangoLayout    *layout);
//
//
//#define PANGO_TYPE_LAYOUT_LINE (pango_layout_line_get_type ())
//
//GType    pango_layout_line_get_type     (void) G_GNUC_CONST;
//
//PangoLayoutLine *pango_layout_line_ref   (PangoLayoutLine *line);
//void             pango_layout_line_unref (PangoLayoutLine *line);
//
//gboolean pango_layout_line_x_to_index   (PangoLayoutLine  *line,
//					 int               x_pos,
//					 int              *index_,
//					 int              *trailing);
//void     pango_layout_line_index_to_x   (PangoLayoutLine  *line,
//					 int               index_,
//					 gboolean          trailing,
//					 int              *x_pos);
//void     pango_layout_line_get_x_ranges (PangoLayoutLine  *line,
//					 int               start_index,
//					 int               end_index,
//					 int             **ranges,
//					 int              *n_ranges);
//void     pango_layout_line_get_extents  (PangoLayoutLine  *line,
//					 PangoRectangle   *ink_rect,
//					 PangoRectangle   *logical_rect);
//void     pango_layout_line_get_pixel_extents (PangoLayoutLine *layout_line,
//					      PangoRectangle  *ink_rect,
//					      PangoRectangle  *logical_rect);
//
//typedef struct _PangoLayoutIter PangoLayoutIter;
//
//#define PANGO_TYPE_LAYOUT_ITER         (pango_layout_iter_get_type ())
//
//GType            pango_layout_iter_get_type (void) G_GNUC_CONST;
//
//PangoLayoutIter *pango_layout_get_iter  (PangoLayout     *layout);
//PangoLayoutIter *pango_layout_iter_copy (PangoLayoutIter *iter);
//void             pango_layout_iter_free (PangoLayoutIter *iter);
//
//int              pango_layout_iter_get_index  (PangoLayoutIter *iter);
//PangoLayoutRun  *pango_layout_iter_get_run    (PangoLayoutIter *iter);
//PangoLayoutRun  *pango_layout_iter_get_run_readonly   (PangoLayoutIter *iter);
//PangoLayoutLine *pango_layout_iter_get_line   (PangoLayoutIter *iter);
//PangoLayoutLine *pango_layout_iter_get_line_readonly  (PangoLayoutIter *iter);
//gboolean         pango_layout_iter_at_last_line (PangoLayoutIter *iter);
//PangoLayout     *pango_layout_iter_get_layout (PangoLayoutIter *iter);
//
//gboolean pango_layout_iter_next_char    (PangoLayoutIter *iter);
//gboolean pango_layout_iter_next_cluster (PangoLayoutIter *iter);
//gboolean pango_layout_iter_next_run     (PangoLayoutIter *iter);
//gboolean pango_layout_iter_next_line    (PangoLayoutIter *iter);
//
//void pango_layout_iter_get_char_extents    (PangoLayoutIter *iter,
//					    PangoRectangle  *logical_rect);
//void pango_layout_iter_get_cluster_extents (PangoLayoutIter *iter,
//					    PangoRectangle  *ink_rect,
//					    PangoRectangle  *logical_rect);
//void pango_layout_iter_get_run_extents     (PangoLayoutIter *iter,
//					    PangoRectangle  *ink_rect,
//					    PangoRectangle  *logical_rect);
//void pango_layout_iter_get_line_extents    (PangoLayoutIter *iter,
//					    PangoRectangle  *ink_rect,
//					    PangoRectangle  *logical_rect);
/* All the yranges meet, unlike the logical_rect's (i.e. the yranges
 * assign between-line spacing to the nearest line)
 */
//void pango_layout_iter_get_line_yrange     (PangoLayoutIter *iter,
//					    int             *y0_,
//					    int             *y1_);
//void pango_layout_iter_get_layout_extents  (PangoLayoutIter *iter,
//					    PangoRectangle  *ink_rect,
//					    PangoRectangle  *logical_rect);
//int  pango_layout_iter_get_baseline        (PangoLayoutIter *iter);
//
