import socket

# 服务器的地址和端口
server_address = ('127.0.0.1', 10000)

# 创建UDP套接字
client_socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

try:
    # 发送数据
    message = 'This is a message'
    print(f'Sending: {message}')
    client_socket.sendto(message.encode(), server_address)

finally:
    print('Closing socket')
    client_socket.close()
