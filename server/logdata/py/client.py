import socket

# 创建一个UDP套接字
udp_socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

# 定义接收方的IP和端口
target_ip = "127.0.0.1"
target_port = 10000

# 要发送的消息
message = "你好, UDP!"
print(message)

# 发送消息
udp_socket.sendto(message.encode('utf-8'), (target_ip, target_port))

# 关闭套接字
udp_socket.close()
