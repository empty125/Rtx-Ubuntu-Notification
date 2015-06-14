# -*- coding: utf-8 -*-

import sys
import getopt
import socketserver
from datetime import datetime
from . import notify


class NotificationUDPHandler(socketserver.BaseRequestHandler):

	def handle(self):
		data = self.request[0].decode('utf-8').strip()
		socket = self.request[1]
		now_time = datetime.now().strftime("%H:%M:%S")
		print("{0} {1} wrote:{2}".format(
			now_time,
			self.client_address[0],			
			data
			))
		notify.show(now_time+' '+data)
		socket.sendto(("%s ok ." % now_time).encode('utf-8'),self.client_address)

host = "192.168.122.1"
port = 10086
opts = None

try:
	opts, args = getopt.getopt(sys.argv[1:], "h:p:", ["host=","port="])
except getopt.GetoptError as e:
	print(e)
	sys.exit(0)

for opt, arg in opts:
	if opt in ("-h", "--host"):
		host = arg.strip()
	if opt in ("-p", "--port"):
		port = int(arg.strip())

print("server listen at %s:%d" % (host,port))
server = socketserver.UDPServer((host,port), NotificationUDPHandler)
server.serve_forever()