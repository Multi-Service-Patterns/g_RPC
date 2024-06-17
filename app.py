from concurrent import futures
import grpc
import time
import signal
import sys
import music_service_pb2
import music_service_pb2_grpc

# In-memory data storage
users = {}
playlists = {}
songs = {}

class MusicService(music_service_pb2_grpc.MusicServiceServicer):
    def GetUser(self, request, context):
        user = users.get(request.id)
        if not user:
            context.abort(grpc.StatusCode.NOT_FOUND, "User not found")
        return user

    def CreateUser(self, request, context):
        user_id = f"userID{len(users) + 1}"
        user = music_service_pb2.User(
            id=user_id,
            name=request.user.name,
            age=request.user.age,
            playlists=request.user.playlists
        )
        users[user_id] = user
        return user

    def UpdateUser(self, request, context):
        user = users.get(request.user.id)
        if not user:
            context.abort(grpc.StatusCode.NOT_FOUND, "User not found")
        users[request.user.id] = request.user
        return request.user

    def DeleteUser(self, request, context):
        if request.id in users:
            del users[request.id]
        else:
            context.abort(grpc.StatusCode.NOT_FOUND, "User not found")
        return music_service_pb2.Empty()

    def GetPlaylist(self, request, context):
        playlist = playlists.get(request.id)
        if not playlist:
            context.abort(grpc.StatusCode.NOT_FOUND, "Playlist not found")
        return playlist

    def CreatePlaylist(self, request, context):
        playlist_id = f"playlistID{len(playlists) + 1}"
        playlist = music_service_pb2.Playlist(
            id=playlist_id,
            name=request.playlist.name,
            description=request.playlist.description,
            songs=request.playlist.songs
        )
        playlists[playlist_id] = playlist
        return playlist

    def UpdatePlaylist(self, request, context):
        playlist = playlists.get(request.playlist.id)
        if not playlist:
            context.abort(grpc.StatusCode.NOT_FOUND, "Playlist not found")
        playlists[request.playlist.id] = request.playlist
        return request.playlist

    def DeletePlaylist(self, request, context):
        if request.id in playlists:
            del playlists[request.id]
        else:
            context.abort(grpc.StatusCode.NOT_FOUND, "Playlist not found")
        return music_service_pb2.Empty()

    def GetSong(self, request, context):
        song = songs.get(request.id)
        if not song:
            context.abort(grpc.StatusCode.NOT_FOUND, "Song not found")
        return song

    def CreateSong(self, request, context):
        song_id = f"songID{len(songs) + 1}"
        song = music_service_pb2.Song(
            id=song_id,
            artist=request.song.artist,
            name=request.song.name
        )
        songs[song_id] = song
        return song

    def UpdateSong(self, request, context):
        song = songs.get(request.song.id)
        if not song:
            context.abort(grpc.StatusCode.NOT_FOUND, "Song not found")
        songs[request.song.id] = request.song
        return request.song

    def DeleteSong(self, request, context):
        if request.id in songs:
            del songs[request.id]
        else:
            context.abort(grpc.StatusCode.NOT_FOUND, "Song not found")
        return music_service_pb2.Empty()

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    music_service_pb2_grpc.add_MusicServiceServicer_to_server(MusicService(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print("Server started on port 50051")
    
    def handle_signal(signal, frame):
        print("Shutting down server...")
        server.stop(0)
        sys.exit(0)

    signal.signal(signal.SIGINT, handle_signal)
    signal.signal(signal.SIGTERM, handle_signal)
    
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
