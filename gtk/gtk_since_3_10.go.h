/*
 * Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
 *
 * This file originated from: http://opensource.conformal.com/
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

#include "gtk.go.h" // for callbackDelete
#include <stdlib.h>

static GtkHeaderBar *toGtkHeaderBar(void *p) { return (GTK_HEADER_BAR(p)); }

static GtkListBox *toGtkListBox(void *p) { return (GTK_LIST_BOX(p)); }

static GtkListBoxRow *toGtkListBoxRow(void *p) { return (GTK_LIST_BOX_ROW(p)); }

static GtkRevealer *toGtkRevealer(void *p) { return (GTK_REVEALER(p)); }

static GtkSearchBar *toGtkSearchBar(void *p) { return (GTK_SEARCH_BAR(p)); }

static GtkStack *toGtkStack(void *p) { return (GTK_STACK(p)); }

static GtkStackSwitcher *toGtkStackSwitcher(void *p) {
  return (GTK_STACK_SWITCHER(p));
}

extern gboolean goListBoxFilterFuncs(GtkListBoxRow *row, gpointer user_data);

static inline void _gtk_list_box_set_filter_func(GtkListBox *box,
                                                 gpointer user_data) {
  gtk_list_box_set_filter_func(box,
                               (GtkListBoxFilterFunc)(goListBoxFilterFuncs),
                               user_data, (GDestroyNotify)(callbackDelete));
}

extern void goListBoxHeaderFuncs(GtkListBoxRow *row, GtkListBoxRow *before,
                                 gpointer user_data);

static inline void _gtk_list_box_set_header_func(GtkListBox *box,
                                                 gpointer user_data) {
  gtk_list_box_set_header_func(
      box, (GtkListBoxUpdateHeaderFunc)(goListBoxHeaderFuncs), user_data,
      (GDestroyNotify)(callbackDelete));
}

extern gint goListBoxSortFuncs(GtkListBoxRow *row1, GtkListBoxRow *row2,
                               gpointer user_data);

static inline void _gtk_list_box_set_sort_func(GtkListBox *box,
                                               gpointer user_data) {
  gtk_list_box_set_sort_func(box, (GtkListBoxSortFunc)(goListBoxSortFuncs),
                             user_data, (GDestroyNotify)(callbackDelete));
}
