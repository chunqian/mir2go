import socket
import threading

def client_handle(client_socket):
    try:
        while True:
            # 接收客户端发送的数据
            request = client_socket.recv(1024)
            
            if not request:
                print("客户端已断开连接。")
                break
            
            decoded_request = request.decode('utf-8')
            print(f"接收到: {decoded_request}")

            # 发送数据
            if decoded_request == "%--$":
                client_socket.sendall("%++$".encode('utf-8'))
    except (socket.error, ConnectionResetError, BrokenPipeError):
        print("出现异常，关闭套接字。")
    finally:
        client_socket.close()

def main():
    server_ip = "0.0.0.0"  # 监听所有可用的网络接口
    server_port = 5500

    # 创建socket对象
    server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)

    # 绑定IP地址和端口号
    server.bind((server_ip, server_port))

    # 开始监听, 最大连接数为5
    server.listen(5)

    print(f"监听 {server_ip}:{server_port}")

    try:
        while True:
            # 接受客户端连接
            client_socket, addr = server.accept()

            print(f"接受来自 {addr} 的连接")

            # 创建一个新的线程来处理该客户端
            client_handler = threading.Thread(target=client_handle, args=(client_socket,))
            client_handler.daemon = True
            client_handler.start()
    except KeyboardInterrupt:
        print("收到 Ctrl+C，正在关闭服务器。")
        server.close()

if __name__ == "__main__":
    main()
