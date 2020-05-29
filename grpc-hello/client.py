# -*- coding: utf-8 -*-
import logging
import grpc
import data_pb2,data_pb2_grpc

"""
 _HOST = 'localhost'
_PORT = '8080'

def run():
    conn = grpc.insecure_channel(_HOST + ':' + _PORT)
    client = data_pb2_grpc.gRPCStub(channel=conn)
    response = client.SayHello(data_pb2.HelloRequest(name='hello,world!'))
    print("received: " + response.text)

if __name__ == '__main__':
    run()
"""

def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel('localhost:8080') as channel:
        stub = data_pb2_grpc.gRPCStub(channel)
        response = stub.SayHello(data_pb2.HelloRequest(name='you'))
    print("Greeter client received: " + response.message)


if __name__ == '__main__':
    logging.basicConfig()
    run()
