all: go-pb-example example

%.pb.go:	%.proto
	protoc --go_out=. $<  --descriptor_set_out=$?.pbuf

go-pb-example: go-pb-example/main.go example
	go get ./go-pb-example

example: example/test.pb.go
	go get ./example

#test.pb.go: example/test.proto
#	protoc --go_out=. ./example/test.proto

clean:
	rm -rf example/*.pb.go example/*.pbuf

