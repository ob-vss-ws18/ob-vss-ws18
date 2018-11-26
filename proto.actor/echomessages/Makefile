regenerate:
	go install github.com/gogo/protobuf/protoc-gen-gogoslick
	protoc -I=. -I=${GOPATH}/src --gogoslick_out=grpc\
	Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:. \
	--proto_path=../../../:../../../gogo/protobuf/protobuf/:. echo.proto
