@echo off

cd %cd%

echo %cd%

protoc  --micro_out=. --go_out=. stream.proto

pause