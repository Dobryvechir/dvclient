SET GOPATH=%CD%
echo %GOPATH%
go install main
@del bin\dvclient.exe
@ren bin\main.exe dvclient.exe
rem @copy bin\dvclient.exe C:\Go\bin

