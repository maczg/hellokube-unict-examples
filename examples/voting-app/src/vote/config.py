import os
import socket


class Config:
    REDIS_HOST = os.getenv('REDIS_HOST', "localhost")
    REDIS_PORT = os.getenv('REDIS_PORT', 6379)

    OPTION_A = os.getenv('OPTION_A', "Cats")
    OPTION_B = os.getenv('OPTION_B', "Dogs")
    HOSTNAME = socket.gethostname()