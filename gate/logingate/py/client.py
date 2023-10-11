import socket  # 导入 socket 库

# 创建一个 socket 对象
client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

# 服务器的 IP 地址和端口号
server_address = ('127.0.0.1', 7000)

# 连接到服务器
client_socket.connect(server_address)

try:
    # 发送数据
    message = '%你好, 服务器!$'
    client_socket.sendall(message.encode('utf-8'))

    # 接收数据
    data = client_socket.recv(1024)
    print('收到: ', data.decode('utf-8'))

finally:
    # 关闭连接
    client_socket.close()
