import grpc
import music_service_pb2
import music_service_pb2_grpc

def run():
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = music_service_pb2_grpc.MusicServiceStub(channel)
        
        # Create a new user
        user = music_service_pb2.User(name="John Doe", age=30)
        create_user_response = stub.CreateUser(music_service_pb2.CreateUserRequest(user=user))
        print(f"Created User: {create_user_response}")

        # Get the user
        get_user_response = stub.GetUser(music_service_pb2.GetUserRequest(id=create_user_response.id))
        print(f"Got User: {get_user_response}")

        # Update the user
        updated_user = music_service_pb2.User(id=get_user_response.id, name="John Doe", age=31)
        update_user_response = stub.UpdateUser(music_service_pb2.UpdateUserRequest(user=updated_user))
        print(f"Updated User: {update_user_response}")

        # Delete the user
        delete_user_response = stub.DeleteUser(music_service_pb2.DeleteUserRequest(id=get_user_response.id))
        print(f"Deleted User: {delete_user_response}")

if __name__ == '__main__':
    run()
