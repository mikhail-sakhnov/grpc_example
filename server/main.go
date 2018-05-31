package main

import (
	"errors"
	"fmt"
	"net"
	"sync"

	"crypto/md5"
	"io"

	proto "github.com/soider/grpc_example/server/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type inMemoryStorage struct {
	sync.Mutex
	storage map[string]string
}

var errorNotFound = errors.New("Value not found")

func (ims *inMemoryStorage) Get(ctx context.Context, key *proto.Key) (*proto.Value, error) {
	ims.Lock()
	defer ims.Unlock()
	if value, found := ims.storage[key.GetKey()]; found {
		return &proto.Value{Content: value}, nil
	}
	return nil, errorNotFound
}

func (ims *inMemoryStorage) Put(ctx context.Context, value *proto.Value) (*proto.Key, error) {

	hash := md5.New()
	io.WriteString(hash, value.Content)
	key := fmt.Sprintf("%x", hash.Sum(nil))
	ims.Lock()
	defer ims.Unlock()
	ims.storage[key] = value.Content
	return &proto.Key{Key: key}, nil
}

func newInMemoryStorage() *inMemoryStorage {
	return &inMemoryStorage{
		storage: make(map[string]string),
	}
}

func main() {
	server := grpc.NewServer()
	grpcLn, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}
	storage := newInMemoryStorage()
	fmt.Println("Start serving")
	reflection.Register(server)
	proto.RegisterStorageServiceServer(server, storage)
	if err := server.Serve(grpcLn); err != nil {
		panic(err)
	}
}
