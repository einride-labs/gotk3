package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import "github.com/einride-labs/gotk3/gdk"

func nativeGdkRectangle(rect gdk.Rectangle) *C.GdkRectangle {
	// Note: Here we can't use rect.GdkRectangle because it would return
	// C type prefixed with gdk package. A ways how to resolve this Go
	// issue with same C structs in different Go packages is documented
	// here https://github.com/golang/go/issues/13467 .
	// This is the easiest way how to resolve the problem.
	return &C.GdkRectangle{
		x:      C.int(rect.GetX()),
		y:      C.int(rect.GetY()),
		width:  C.int(rect.GetWidth()),
		height: C.int(rect.GetHeight()),
	}
}

func nativeGdkGeometry(geom gdk.Geometry) *C.GdkGeometry {
	// Note: Here we can't use geom.GdkGeometry because it would return
	// C type prefixed with gdk package. A ways how to resolve this Go
	// issue with same C structs in different Go packages is documented
	// here https://github.com/golang/go/issues/13467 .
	// This is the easiest way how to resolve the problem.
	return &C.GdkGeometry{
		min_width:   C.gint(geom.GetMinWidth()),
		min_height:  C.gint(geom.GetMinHeight()),
		max_width:   C.gint(geom.GetMaxWidth()),
		max_height:  C.gint(geom.GetMaxHeight()),
		base_width:  C.gint(geom.GetBaseWidth()),
		base_height: C.gint(geom.GetBaseHeight()),
		width_inc:   C.gint(geom.GetWidthInc()),
		height_inc:  C.gint(geom.GetHeightInc()),
		min_aspect:  C.gdouble(geom.GetMinAspect()),
		max_aspect:  C.gdouble(geom.GetMaxAspect()),
		win_gravity: C.GdkGravity(geom.GetWinGravity()),
	}
}
