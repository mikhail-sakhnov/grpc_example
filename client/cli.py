import grpc
import proto.service_pb2_grpc
import proto.service_pb2 as dto

if __name__ == "__main__":
    channel = grpc.insecure_channel("127.0.0.1:9999")
    client = proto.service_pb2_grpc.StorageServiceStub(channel)
    put_request = dto.Value()
    put_request.content = "My data to store!"
    response = client.Put(put_request)
    print("Value stored with key {}".format(response.key))
    print("Reading remote value for key {}".format(response.key))
    read_request = (
        dto.Key()
    )  # Or we can just use our response instead of creating new one
    read_request.key = response.key
    read_response = client.Get(read_request)
    print("Value is {}".format(read_response.content))
