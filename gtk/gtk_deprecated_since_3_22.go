//+build gtk_3_6 gtk_3_8 gtk_3_10 gtk_3_12 gtk_3_14 gtk_3_16 gtk_3_18 gtk_3_20 gtk_deprecated

package gtk

// #include <gtk/gtk.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"

	"github.com/einride-labs/gotk3/gdk"
)

/*
 * GtkWidget
 */

// TODO:
// gtk_widget_send_expose().
// gtk_widget_is_composited().

// PopupAtMouse() is a wrapper for gtk_menu_popup(), without the option for a custom positioning function.
func (v *Menu) PopupAtMouseCursor(parentMenuShell IMenu, parentMenuItem IMenuItem, button gdk.Button, activateTime uint32) {
	wshell := nullableWidget(parentMenuShell)
	witem := nullableWidget(parentMenuItem)

	C.gtk_menu_popup(v.native(),
		wshell,
		witem,
		nil,
		nil,
		C.guint(button),
		C.guint32(activateTime))
}

func (v *SizeGroup) GetIgnoreHidden() bool {
	c := C.gtk_size_group_get_ignore_hidden(v.native())
	return gobool(c)
}

// SetWMClass is a wrapper around gtk_window_set_wmclass().
func (v *Window) SetWMClass(name, class string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cClass := C.CString(class)
	defer C.free(unsafe.Pointer(cClass))
	C.gtk_window_set_wmclass(v.native(), (*C.gchar)(cName), (*C.gchar)(cClass))
}

func (v *SizeGroup) SetIgnoreHidden(ignoreHidden bool) {
	C.gtk_size_group_set_ignore_hidden(v.native(), gbool(ignoreHidden))
}

// GetFontName is a wrapper around gtk_font_button_get_font_name().
func (v *FontButton) GetFontName() string {
	c := C.gtk_font_button_get_font_name(v.native())
	return goString(c)
}

// SetFontName is a wrapper around gtk_font_button_set_font_name().
func (v *FontButton) SetFontName(fontname string) bool {
	cstr := C.CString(fontname)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_font_button_set_font_name(v.native(), (*C.gchar)(cstr))
	return gobool(c)
}
