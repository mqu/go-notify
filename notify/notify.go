/*
 * notify.go for go-notify
 * by lenorm_f
 */

/*
	This package provides GO bindings for the C library libnotify.
	Although this package provides full retro-compatibility with the
	regular C library, it also provides OOP-like functions for the
	NotifyNotification object.
*/

package notify

/*
#cgo pkg-config: libnotify

#include <stdlib.h>
#include <libnotify/notify.h>
*/
import "C"
import "unsafe"

// Exported Types
type NotifyNotification struct {
	_notification *C.NotifyNotification
}

type GError struct {
	domain uint32
	code int
	message string
}

// Private Types

// Private Functions
func new_notify_notification(cnotif *C.NotifyNotification) *NotifyNotification {
	notif := new(NotifyNotification)

	notif._notification = cnotif

	return notif
}

func new_gerror(cgerror *C.GError) *GError {
	gerror := new(GError)

	gerror.domain = uint32(cgerror.domain)
	gerror.code = int(cgerror.code)
	gerror.message = C.GoString((*C.char)(cgerror.message))

	return gerror
}

// Exported Functions
func Init(app_name string) bool {
	papp_name := C.CString(app_name)
	defer C.free(unsafe.Pointer(papp_name))

	return bool(C.notify_init(papp_name) != 0)
}

func NotificationNew(title, text, image string) *NotifyNotification {
	ptitle := C.CString(title)
	ptext := C.CString(text)
	pimage := C.CString(image)
	defer func() {
		C.free(unsafe.Pointer(ptitle))
		C.free(unsafe.Pointer(ptext))
		C.free(unsafe.Pointer(pimage))
	}()

	return new_notify_notification(C.notify_notification_new(ptitle, ptext, pimage))
}

func NotificationShow(notif *NotifyNotification) *GError {
	var err *C.GError
	C.notify_notification_show(notif._notification, (**C.GError)(&err))

	if err != nil {
		return new_gerror(err)
	}

	return nil
}

func (this *NotifyNotification) Show() {
	var err *C.GError
	C.notify_notification_show(this._notification, (**C.GError)(&err))
}

func NotificationSetTimeout(notif *NotifyNotification, timeout int) {
	C.notify_notification_set_timeout(notif._notification, C.gint(timeout))
}

func (this *NotifyNotification) SetTimeout(timeout int) {
	C.notify_notification_set_timeout(this._notification, C.gint(timeout))
}
