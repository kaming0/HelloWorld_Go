# HelloWorld_Go
Write HelloWorld in Go

# To call HelloWorld from C++
go build -o hello.so -buildmode=c-shared hello_cpp.go
g++ -o main main.cpp ./hello.so
