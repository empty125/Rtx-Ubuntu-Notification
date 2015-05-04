# -*- coding: utf-8 -*-

import socketserver
import notify
from datetime import datetime

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
		socket.sendto("ok \n".encode('utf-8'),self.client_address)


if __name__ =="__main__":
	server = socketserver.UDPServer(("192.168.122.1", 10086), NotificationUDPHandler)
	server.serve_forever()