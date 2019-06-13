module sayhi

go 1.12

replace (
	cloud.google.com/go v0.26.0 => github.com/googleapis/google-cloud-go v0.26.0
	golang.org/x/lint v0.0.0-20181026193005-c67002cb31c3 => github.com/golang/lint v0.0.0-20181026193005-c67002cb31c3
	golang.org/x/net v0.0.0-20180826012351-8a410e7b638d => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be => github.com/golang/oauth2 v0.0.0-20180821212333-d2e6202438be
	golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f => github.com/golang/sync v0.0.0-20180314180146-1d60e4601c6f
	golang.org/x/sys v0.0.0-20180830151530-49385e6e1522 => github.com/golang/sys v0.0.0-20180830151530-49385e6e1522
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
	golang.org/x/tools v0.0.0-20190114222345-bf090417da8b => github.com/golang/tools v0.0.0-20190114222345-bf090417da8b
	google.golang.org/appengine v1.1.0 => github.com/golang/appengine v1.1.0
	google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8 => github.com/google/go-genproto v0.0.0-20180817151627-c66870c02cf8
	google.golang.org/grpc v1.19.1 => github.com/grpc/grpc-go v1.19.1
)

require (
	github.com/golang/protobuf v1.3.1
	google.golang.org/grpc v1.19.1
)
