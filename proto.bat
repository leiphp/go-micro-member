protoc --proto_path=proto/member/user --micro_out=./ --go_out=./ users.proto
protoc --proto_path=proto/member/user --go_out=./ models.proto
protoc --proto_path=proto/course/course --micro_out=./ --go_out=./ course.proto
protoc-go-inject-tag -input=proto/course/course/course.pb.go