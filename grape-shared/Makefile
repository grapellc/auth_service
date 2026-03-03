.PHONY: proto build

PROTO_PATH := proto

proto:
	@mkdir -p proto/authv1 proto/engagementv1 proto/productv1 proto/jobsv1 proto/marketv1 proto/chatv1 proto/feedv1
	protoc --go_out=proto/authv1 --go_opt=paths=source_relative --proto_path=$(PROTO_PATH) proto/auth_service.proto
	protoc --go_out=proto/engagementv1 --go_opt=paths=source_relative --proto_path=$(PROTO_PATH) proto/engagement_service.proto
	protoc --go_out=proto/productv1 --go_opt=paths=source_relative --proto_path=$(PROTO_PATH) proto/product_service.proto
	protoc --go_out=proto/jobsv1 --go_opt=paths=source_relative --proto_path=$(PROTO_PATH) proto/jobs_service.proto
	protoc --go_out=proto/marketv1 --go_opt=paths=source_relative --proto_path=$(PROTO_PATH) proto/market_service.proto
	protoc --go_out=proto/chatv1 --go_opt=paths=source_relative --proto_path=$(PROTO_PATH) proto/chat_service.proto
	protoc --go_out=proto/feedv1 --go_opt=paths=source_relative --proto_path=$(PROTO_PATH) proto/feed_service.proto

build:
	go build ./...
