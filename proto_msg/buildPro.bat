protoc.exe --plugin=protoc-gen-go=./protoc-gen-go.exe --go_out=./ *.proto

set SERVER_HOME=D:\work\git\ssServer\lib\src\proto_msg
set CLIENT_HOME=D:\work\git\ssClient

D:
copy %SERVER_HOME%\*.proto %CLIENT_HOME%\protocol

cd %CLIENT_HOME%\protocol
call python protoToJs.py

pause