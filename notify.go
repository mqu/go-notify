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
import glib "github.com/mattn/go-gtk/glib"

/*
 * Exported Types
 */
type NotifyNotification struct {
	_notification *C.NotifyNotification
}

const (
	NOTIFY_URGENCY_LOW = 0;
	NOTIFY_URGENCY_NORMAL = 1;
	NOTIFY_URGENCY_CRITICAL = 2;
)

type NotifyUrgency int
type NotifyActionCallback func (*NotifyNotification, string, interface {})

/*
 * Private Functions
 */
func new_notify_notification(cnotif *C.NotifyNotification) *NotifyNotification {
	return &NotifyNotification{cnotif}
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

func GetServerCaps() *glib.List {
	var gcaps *C.GList

	gcaps = C.notify_get_server_caps()

	return glib.ListFromNative(unsafe.Pointer(gcaps))
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
	ntext := C.g_utf8_normalize((*C.gchar)(ptext), -1, C.G_NORMALIZE_DEFAULT)
	defer func() {
		C.free(unsafe.Pointer(ptitle))
		C.free(unsafe.Pointer(ptext))
		C.free(unsafe.Pointer(pimage))

		if ntext != nil {
			C.free(unsafe.Pointer(ntext))
		}
	}()

	return new_notify_notification(C.notify_notification_new(ptitle, (*C.char)(ntext), pimage))
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

func NotificationShow(notif *NotifyNotification) *glib.Error {
	var err *C.GError
	C.notify_notification_show(notif._notification, &err)

	return glib.ErrorFromNative(unsafe.Pointer(err))
}

func NotificationSetTimeout(notif *NotifyNotification, timeout int32) {
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

func NotificationSetHintInt32(notif *NotifyNotification, key string, value int32) {
	pkey := C.CString(key)
	defer C.free(unsafe.Pointer(pkey))

	C.notify_notification_set_hint_int32(notif._notification, pkey, C.gint(value))
}

func NotificationSetHintDouble(notif *NotifyNotification, key string, value float64) {
	pkey := C.CString(key)
	defer C.free(unsafe.Pointer(pkey))

	C.notify_notification_set_hint_double(notif._notification, pkey, C.gdouble(value))
}

func NotificationSetHintString(notif *NotifyNotification, key string, value string) {
	pkey := C.CString(key)
	pvalue := C.CString(value)
	defer func() {
		C.free(unsafe.Pointer(pkey))
		C.free(unsafe.Pointer(pvalue))
	}()

	C.notify_notification_set_hint_string(notif._notification, pkey, pvalue)
}

func NotificationSetHintByte(notif *NotifyNotification, key string, value byte) {
	pkey := C.CString(key)
	defer C.free(unsafe.Pointer(pkey))

	C.notify_notification_set_hint_byte(notif._notification, pkey, C.guchar(value))
}

// FIXME: implement
func NotificationSetHintByteArray(notif *NotifyNotification, key string, value []byte, len uint32) {
	pkey := C.CString(key)
	defer C.free(unsafe.Pointer(pkey))

	// C.notify_notification_set_hint_byte_array(notif._notification, pkey, (*C.guchar)(value), C.gsize(len))
}

func NotificationSetHint(notif *NotifyNotification, key string, value interface {}) {
	switch value.(type) {
		case int32: NotificationSetHintInt32(notif, key, value.(int32))
		case float64: NotificationSetHintDouble(notif, key, value.(float64))
		case string: NotificationSetHintString(notif, key, value.(string))
		case byte: NotificationSetHintByte(notif, key, value.(byte))
	}
}

func NotificationClearHints(notif *NotifyNotification) {
	C.notify_notification_clear_hints(notif._notification)
}

// FIXME: the C function is supposed to be allowing the user to pass another function than free
func NotificationAddAction(notif *NotifyNotification, action, label string, callback NotifyActionCallback, user_data interface {}) {
	// C.notify_notification_add_action(notif._notification, paction, plabel, (C.NotifyActionCallback)(callback), user_data, C.free)
}

func NotificationClearActions(notif *NotifyNotification) {
	C.notify_notification_clear_actions(notif._notification)
}

func NotificationClose(notif *NotifyNotification) *glib.Error {
	var err *C.GError

	C.notify_notification_close(notif._notification, &err)

	return glib.ErrorFromNative(unsafe.Pointer(err))
}

// Member Functions
func (this *NotifyNotification) Update(summary, body, icon string) bool {
	return NotificationUpdate(this, summary, body, icon)
}

func (this *NotifyNotification) Show() *glib.Error {
	return NotificationShow(this)
}

func (this *NotifyNotification) SetTimeout(timeout int32) {
	NotificationSetTimeout(this, timeout)
}

func (this *NotifyNotification) SetCategory(category string) {
	NotificationSetCategory(this, category)
}

func (this *NotifyNotification) SetUrgency(urgency NotifyUrgency) {
	NotificationSetUrgency(this, urgency)
}

func (this *NotifyNotification) SetHintInt32(key string, value int32) {
	NotificationSetHintInt32(this, key, value)
}

func (this *NotifyNotification) SetHintDouble(key string, value float64) {
	NotificationSetHintDouble(this, key, value)
}

func (this *NotifyNotification) SetHintString(key string, value string) {
	NotificationSetHintString(this, key, value)
}

func (this *NotifyNotification) SetHintByte(key string, value byte) {
	NotificationSetHintByte(this, key, value)
}

func (this *NotifyNotification) SetHintByteArray(key string, value []byte, len uint32) {
	NotificationSetHintByteArray(this, key, value, len)
}

func (this *NotifyNotification) SetHint(key string, value interface {}) {
	NotificationSetHint(this, key, value)
}

func (this *NotifyNotification) ClearHints() {
	NotificationClearHints(this)
}

func (this *NotifyNotification) AddAction(action, label string, callback NotifyActionCallback, user_data interface {}) {
	NotificationAddAction(this, action, label, callback, user_data)
}

func (this *NotifyNotification) ClearActions() {
	NotificationClearActions(this)
}

func (this *NotifyNotification) Close() *glib.Error {
	return NotificationClose(this)
}
