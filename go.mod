module github.com/gilbertchen/duplicacy

go 1.16

require (
	cloud.google.com/go v0.38.0
	github.com/aryann/difflib v0.0.0-20170710044230-e206f873d14a
	github.com/aws/aws-sdk-go v1.30.7
	github.com/bkaradzic/go-lz4 v1.0.0
	github.com/gilbertchen/azure-sdk-for-go v14.1.2-0.20180323033227-8fd4663cab7c+incompatible
	github.com/gilbertchen/cli v1.2.1-0.20160223210219-1de0a1836ce9
	github.com/gilbertchen/go-dropbox v0.0.0-20221207034530-08c0c180a4f9
	github.com/gilbertchen/go-ole v1.2.0
	github.com/gilbertchen/goamz v0.0.0-20170712012135-eada9f4e8cc2
	github.com/gilbertchen/gopass v0.0.0-20170109162249-bf9dde6d0d2c
	github.com/gilbertchen/keyring v0.0.0-20221004152639-1661cbebc508
	github.com/gilbertchen/xattr v0.0.0-20160926155429-68e7a6806b01
	github.com/klauspost/reedsolomon v1.9.9
	github.com/minio/blake2b-simd v0.0.0-20160723061019-3f5f724cb5b1
	github.com/minio/highwayhash v1.0.2
	github.com/ncw/swift/v2 v2.0.1
	github.com/pkg/sftp v1.11.0
	github.com/pkg/xattr v0.4.1
	github.com/vmihailenco/msgpack v4.0.4+incompatible
	golang.org/x/crypto v0.0.0-20220131195533-30dcbda58838
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	google.golang.org/api v0.21.0
	storj.io/uplink v1.9.0
)

require (
	github.com/Azure/go-autorest v10.15.5+incompatible // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dnaeon/go-vcr v1.2.0 // indirect
	github.com/gilbertchen/highwayhash v0.0.0-20221109044721-eeab1f4799d8
	github.com/goamz/goamz v0.0.0-20180131231218-8b901b531db8 // indirect
	github.com/godbus/dbus v4.1.0+incompatible // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/klauspost/cpuid v1.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/marstr/guid v1.1.0 // indirect
	github.com/mmcloughlin/avo v0.0.0-20200803215136-443f81d77104 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/segmentio/go-env v1.1.0 // indirect
	github.com/skratchdot/open-golang v0.0.0-20200116055534-eef842397966
	github.com/vaughan0/go-ini v0.0.0-20130923145212-a98ad7ee00ec // indirect
	go.opencensus.io v0.22.3 // indirect
	google.golang.org/appengine v1.6.5 // indirect
	google.golang.org/genproto v0.0.0-20200409111301-baae70f3302d // indirect
	google.golang.org/grpc v1.28.1 // indirect
)

replace github.com/gilbertchen/go-ole => github.com/go-ole/go-ole v1.2.6
