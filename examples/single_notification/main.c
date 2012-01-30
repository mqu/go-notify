/*
 * main.c for go-notify
 * by lenorm_f
 */

#include <stdlib.h>
#include <unistd.h>
#include <libnotify/notify.h>

#define DELAY 3000

int main () {
	NotifyNotification *hello;

	notify_init("Hello World!");
	hello = notify_notification_new("Hello World!",
			"This is an example notification.", NULL);
	if (!hello)
		return 1;
	notify_notification_set_timeout(hello, DELAY);

	if (!notify_notification_show(hello, NULL))
		return 3;

	sleep(DELAY / 1000);
	notify_notification_close(hello, NULL);

	notify_uninit();

	return 0;
}
