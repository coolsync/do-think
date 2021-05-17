gen:
	# protoc --go_out=pb proto/*.proto
	protoc --proto_path=proto --go_out=pb proto/*.proto
clean:
	rm pb/*.go

run:
	go run main.go