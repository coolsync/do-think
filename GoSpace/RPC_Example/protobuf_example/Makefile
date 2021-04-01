gen:
	# protoc --proto_path=proto --go_out=pb *.proto
	protoc --proto_path=micro_server/proto --go_out=micro_server/pb micro_server/proto/*.proto

clean:
	rm pb/*.go

run:
	go run main.go