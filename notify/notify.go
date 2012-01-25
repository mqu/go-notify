/*
 * notify.go for go-notify
 * by lenorm_f
 */

package notify

/*
#cgo pkg-config: libnotify

#include <stdlib.h>
#include <libnotify/notify.h>
*/
import "C"
import "unsafe"

/*
 * Exported Types
 */
type NotifyNotification struct {
	_notification *C.NotifyNotification
}

type GError struct {
	domain uint32
	code int
	message string
}

const (
	NOTIFY_URGENCY_LOW = 0;
	NOTIFY_URGENCY_NORMAL = 1;
	NOTIFY_URGENCY_CRITICAL = 2;
)

type NotifyUrgency int;

/*
 * Private Functions
 */
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

/*
 * Exported Functions
 */
// Pure Functions
func Init(app_name string) bool {
	papp_name := C.CString(app_name)
	defer C.free(unsafe.Pointer(papp_name))

	return bool(C.notify_init(papp_name) != 0)
}

func UnInit() {
	C.notify_uninit()
}

func IsInitted() bool {
	return C.notify_is_initted() != 0
}

func GetAppName() string {
	return C.GoString(C.notify_get_app_name())
}

func GetServerCaps() []string {
	var caps []string
	var gcaps *C.GList

	gcaps = C.notify_get_server_caps()
	if gcaps == nil {
		return nil
	}
	for ; gcaps != nil; gcaps = gcaps.next {
		caps = append(caps, C.GoString((*C.char)(gcaps.data)))
	}

	return caps
}

func GetServerInfo(name, vendor, version, spec_version *string) bool {
	var cname *C.char
	var cvendor *C.char
	var cversion *C.char
	var cspec_version *C.char

	ret := C.notify_get_server_info(&cname,
					&cvendor,
					&cversion,
					&cspec_version) != 0
	*name = C.GoString(cname)
	*vendor = C.GoString(cvendor)
	*version = C.GoString(cversion)
	*spec_version = C.GoString(cspec_version)

	return ret
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

func NotificationUpdate(notif *NotifyNotification, summary, body, icon string) bool {
	psummary := C.CString(summary)
	pbody := C.CString(body)
	picon := C.CString(icon)
	defer func() {
		C.free(unsafe.Pointer(psummary))
		C.free(unsafe.Pointer(pbody))
		C.free(unsafe.Pointer(picon))
	}()

	return C.notify_notification_update(notif._notification, psummary, pbody, picon) != 0
}

func NotificationShow(notif *NotifyNotification) *GError {
	var err *C.GError
	C.notify_notification_show(notif._notification, (**C.GError)(&err))

	if err != nil {
		return new_gerror(err)
	}

	return nil
}

func NotificationSetTimeout(notif *NotifyNotification, timeout int) {
	C.notify_notification_set_timeout(notif._notification, C.gint(timeout))
}

func NotificationSetCategory(notif *NotifyNotification, category string) {
	pcategory := C.CString(category)
	defer C.free(unsafe.Pointer(pcategory))

	C.notify_notification_set_category(notif._notification, pcategory)
}

func NotificationSetUrgency(notif *NotifyNotification, urgency NotifyUrgency) {
	C.notify_notification_set_urgency(notif._notification, C.NotifyUrgency(urgency))
}

// Member Functions
func (this *NotifyNotification) Update(summary, body, icon string) bool {
	psummary := C.CString(summary)
	pbody := C.CString(body)
	picon := C.CString(icon)
	defer func() {
		C.free(unsafe.Pointer(psummary))
		C.free(unsafe.Pointer(pbody))
		C.free(unsafe.Pointer(picon))
	}()

	return C.notify_notification_update(this._notification, psummary, pbody, picon) != 0
}

func (this *NotifyNotification) Show() {
	var err *C.GError
	C.notify_notification_show(this._notification, &err)
}

func (this *NotifyNotification) SetTimeout(timeout int) {
	C.notify_notification_set_timeout(this._notification, C.gint(timeout))
}

func (this *NotifyNotification) SetCategory(category string) {
	pcategory := C.CString(category)
	defer C.free(unsafe.Pointer(pcategory))

	C.notify_notification_set_category(this._notification, pcategory)
}

func (this *NotifyNotification) SetUrgency(urgency NotifyUrgency) {
	C.notify_notification_set_urgency(this._notification, C.NotifyUrgency(urgency))
}
