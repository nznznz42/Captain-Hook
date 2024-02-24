import socket

def start_server(port):
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as server_socket:
        server_socket.bind(('localhost', port))
        server_socket.listen()

        print(f"Server listening on port {port}")

        while True:
            conn, addr = server_socket.accept()
            with conn:
                print(f"Connection from {addr}")
                data = conn.recv(1024)
                if not data:
                    break
                print(f"Received data: {data.decode('utf-8')}")

                response = data

                conn.sendall(response)
                print("Response Sent")

if __name__ == "__main__":
    port_number = 8080
    start_server(port_number)
