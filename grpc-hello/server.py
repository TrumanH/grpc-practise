# -*- coding: utf-8 -*-
import grpc
import data_pb2,data_pb2_grpc

from concurrent import futures
import logging


class Greeter(data_pb2_grpc.gRPCServicer):
    def SayHello(self, request, context):
        return data_pb2.HelloReply(message='Hello, %s!' % request.name)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    data_pb2_grpc.add_gRPCServicer_to_server(Greeter(), server)
    server.add_insecure_port('[::]:8080')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    logging.basicConfig()
    serve()
