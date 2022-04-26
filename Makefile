test:
	protoc --proto_path=greet/proto \
	--go_out=greet/proto \
	--go_opt=paths=source_relative \
	greet/proto/dummy.proto