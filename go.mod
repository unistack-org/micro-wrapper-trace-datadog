module github.com/micro/go-plugins/wrapper/trace/datadog/v2

go 1.13

require (
	github.com/DataDog/datadog-go v3.3.1+incompatible // indirect
	github.com/micro/go-micro/v2 v2.9.1-0.20200716153311-f9bf56239306
	github.com/philhofer/fwd v1.0.0 // indirect
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.5.1
	google.golang.org/grpc v1.35.0
	gopkg.in/DataDog/dd-trace-go.v1 v1.30.0
)

replace github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
