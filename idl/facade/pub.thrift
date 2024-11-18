namespace go system.pub
include "../base/base.thrift"



struct UploadReq{
    1:  binary file
}
// pub service
service SystemService {
    base.NilResponse Upload(1: UploadReq req) (api.post = "/api/pub/upload/")
}
