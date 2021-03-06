module github.com/PharbersDeveloper/ipaas-job-reg

go 1.12

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.44.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190731235908-ec7cb31e5a56
	golang.org/x/image => github.com/golang/image v0.0.0-20190802002840-cff245a6509b
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190806162312-597adff16ade
	golang.org/x/mod => github.com/golang/mod v0.1.0
	golang.org/x/net => github.com/golang/net v0.0.0-20190724013045-ca1201d0de80
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190812073006-9eafafc0a87e
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190809145639-6d4652c779c4
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20190717185122-a985d3407aa7
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.8.0
	google.golang.org/appengine => github.com/golang/appengine v1.6.1
	google.golang.org/crypto => github.com/golang/crypto v0.0.0-20190701094942-4def268fd1a4
	google.golang.org/genproto => github.com/googleapis/go-genproto v0.0.0-20190801165951-fa694d86fc64
	google.golang.org/grpc => github.com/grpc/grpc-go v1.22.1
)

require (
	github.com/PharbersDeveloper/bp-go-lib v0.0.2
	github.com/alfredyang1986/blackmirror v0.1.1
	github.com/aliyun/aliyun-oss-go-sdk v2.0.2+incompatible
	github.com/elodina/go-avro v0.0.0-20160406082632-0c8185d9a3ba
	github.com/elodina/go-kafka-avro v0.0.0-20160422130714-ab6b1d1c9a23
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/hashicorp/go-uuid v1.0.1
	github.com/mitchellh/mapstructure v1.1.2
	github.com/smartystreets/goconvey v0.0.0-20190731233626-505e41936337
)
