module github.com/noovertime7/gin-mysqlbak

go 1.14

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/aliyun/aliyun-oss-go-sdk v2.2.4+incompatible
	github.com/e421083458/gin_scaffold v0.0.0-20211226015300-616fc31c0400
	github.com/e421083458/golang_common v1.2.1
	github.com/gin-gonic/contrib v0.0.0-20190526021735-7fb7810ed2a0
	github.com/gin-gonic/gin v1.7.7
	github.com/go-openapi/spec v0.20.6 // indirect
	github.com/go-openapi/swag v0.21.1 // indirect
	github.com/go-playground/locales v0.13.0
	github.com/go-playground/universal-translator v0.17.0
	github.com/go-xorm/xorm v0.7.9
	github.com/golang/protobuf v1.5.2
	github.com/gorhill/cronexpr v0.0.0-20180427100037-88b0669f7d75
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/minio/minio-go/v7 v7.0.34
	github.com/noovertime7/mysqlbak v0.0.0-20220612083217-fdb12cd90242
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.8.2
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	golang.org/x/tools v0.1.11 // indirect
	google.golang.org/protobuf v1.28.0
	gopkg.in/go-playground/validator.v9 v9.29.0
	gopkg.in/ini.v1 v1.67.0
	gorm.io/gorm v1.22.4
)

replace github.com/gin-contrib/sse v0.1.0 => github.com/e421083458/sse v0.1.1

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
