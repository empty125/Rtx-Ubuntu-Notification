# -*- coding: utf-8 -*-

import socketserver
import notify

class NotificationUDPHandler(socketserver.BaseRequestHandler):

	def handle(self):
		data = self.request[0].decode('utf-8').strip()
		socket = self.request[1]
		print("{0} wrote:{1}".format(self.client_address[0],data))
		notify.show(data)
		socket.sendto("ok \n".encode('utf-8'),self.client_address)


if __name__ =="__main__":
	# Create the server, binding to localhost on port 9999
	server = socketserver.UDPServer(("192.168.122.1", 10086), NotificationUDPHandler)

	# Activate the server; this will keep running until you
	# interrupt the program with Ctrl-C
	server.serve_forever()