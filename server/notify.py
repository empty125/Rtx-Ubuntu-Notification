# -*- coding: utf-8 -*-

#notify-send "Message Title" "The message body is shown here" -i "notification-message-im" -t 5000

__need_call = False
show = None
notification = None
app_name = 'Rtx-Ubuntu-Notification'

try:
	import notify2 
	notify2.init(app_name)  
	notification  =notify2.Notification(app_name)
except ImportError:
	import subprocess
	__need_call = True


def show_notify2(body,title='Rtx Message Notification',icon='notification-message-im',expire=3):
	notification.update(title,body,icon)
	notification.set_timeout(expire*1000)
	notification.show();

def show_call(body,title='Rtx Message Notification',icon='notification-message-im',expire=3):
	command = ['notify-send',title,body,'-i',icon,'-t',str(expire*1000)]
	subprocess.Popen(command,stderr=subprocess.PIPE)

if __need_call:
	show = show_call
else:
	show = show_notify2

if __name__ == "__main__":
	show('Test Message',expire=1)
