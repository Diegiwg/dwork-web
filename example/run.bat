@echo off

set project=%1

cd %project%
go build .
%project%.exe
cd ..