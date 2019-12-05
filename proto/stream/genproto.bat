@rem Copyright 2019 ess authors.
@rem Generate the Go code for .proto files
@rem choco install -y protoc
@rem go get -d -u github.com/micro/protoc-gen-micro
@rem go get -d -u github.com/golang/protobuf/protoc-gen-go
@rem visit https://github.com/micro/protoc-gen-micro for this tools


@echo off &TITLE Generation Protobuf Code For Go

mode con cols=100 lines=30
color 0D
cls

setlocal


@rem enter this directory of bat
cd /d %~dp0

echo.
echo ## Current Dir: %cd%
echo.

@rem compile *.proto to go code

protoc -I. --micro_out=. --go_out=. stream.proto

echo.
echo..........work had been done.
echo.
echo..........code had been generated to :  %cd%
echo.
echo..........press any key to exit
pause >nul
