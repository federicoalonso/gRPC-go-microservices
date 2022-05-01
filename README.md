# gRPC-go-microservices

## Introduction

The main porpoise of this project is to understand how microservices talks to each other, and how to use gRPC to communicate with them.

Also, in this project I intend to practice and understand better the golang programming language.

URL of this course is [link to udemy course](https://www.udemy.com/share/101Zo03@7j5uYCH61FX0Xqkd-vkDY6MNUaAOEqCZjkA_g1UQtvBr62TgxOlRiUAm6TMFBUyv/).

## Memory Helper

```bash
protoc -Igreet/proto --go_out=. --go_opt=module=github.com/federicoalonso/gRPC-go-microservices --go-grpc_out=. --go-grpc_opt=module=github.com/federicoalonso/gRPC-go-microservices greet/proto/dummy.proto

# If make greet isn`t working, you can use this command to repair it

export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin
source $HOME/.bashrc
```