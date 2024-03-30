import socket

def send_request(message, server_address, server_port):
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as client_socket:
        client_socket.connect((server_address, server_port))
        client_socket.sendall(message.encode('utf-8'))

if __name__ == "__main__":
    server_address = 'localhost'
    server_port = 8080

    messages = ["Message 1", "Message 2", "Message 3"]

    for message in messages:
        send_request(message, server_address, server_port)
        print(f"Sent message: {message}")
