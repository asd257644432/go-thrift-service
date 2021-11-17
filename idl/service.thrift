namespace go github.go.go_thrift_service

struct ReqOfHello {
}

struct RespOfHello {
    1: i32 Errno,
}

service HelloService {
    RespOfHello Hello(ReqOfHello request)
}