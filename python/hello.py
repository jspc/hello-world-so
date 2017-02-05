import ctypes
import os

lib_path = os.path.join(os.path.dirname(__file__), '..', 'strings.so')

lib = ctypes.CDLL(lib_path)
hello = lib.C_Hello
hello.argtype = ctypes.c_int
hello.restype = ctypes.c_char_p


print("Running from python")
print(lib.C_Hello(0))
print(lib.C_Hello(1))
