package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	pb "go_client/go_client/music_service" // Certifique-se de que o caminho está correto

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	numRequests = 1000 // Número de requisições a serem simuladas
)

// func main() {
// 	// Estabelecer a conexão com o servidor gRPC
// 	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	defer conn.Close()

// 	client := pb.NewMusicServiceClient(conn)

// 	// Simular múltiplas requisições
// 	for i := 0; i < numRequests; i++ {
// 		go func(i int) {
// 			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
// 			defer cancel()

// 			// Criar um novo usuário
// 			user := &pb.User{
// 				Name: fmt.Sprintf("User %d", i),
// 				Age:  int32(20 + (i % 50)),
// 			}

// 			createUserResponse, err := client.CreateUser(ctx, &pb.CreateUserRequest{User: user})
// 			if err != nil {
// 				log.Printf("could not create user: %v", err)
// 				return
// 			}
// 			log.Printf("Created User: %v\n", createUserResponse)

// 			// Obter o usuário criado
// 			getUserResponse, err := client.GetUser(ctx, &pb.GetUserRequest{Id: createUserResponse.Id})
// 			if err != nil {
// 				log.Printf("could not get user: %v", err)
// 				return
// 			}
// 			log.Printf("Got User: %v\n", getUserResponse)

// 			// Atualizar o usuário
// 			updatedUser := &pb.User{
// 				Id:   getUserResponse.Id,
// 				Name: fmt.Sprintf("Updated User %d", i),
// 				Age:  int32(25 + (i % 50)),
// 			}
// 			updateUserResponse, err := client.UpdateUser(ctx, &pb.UpdateUserRequest{User: updatedUser})
// 			if err != nil {
// 				log.Printf("could not update user: %v", err)
// 				return
// 			}
// 			log.Printf("Updated User: %v\n", updateUserResponse)

// 			// Deletar o usuário
// 			deleteUserResponse, err := client.DeleteUser(ctx, &pb.DeleteUserRequest{Id: getUserResponse.Id})
// 			if err != nil {
// 				log.Printf("could not delete user: %v", err)
// 				return
// 			}
// 			log.Printf("Deleted User: %v\n", deleteUserResponse)
// 		}(i)
// 	}

// 	// Aguardar a conclusão das requisições
// 	time.Sleep(time.Second * 30)
// }

// codigo otimizado pra concorrencia
func main() {
	// Estabelecer a conexão com o servidor gRPC
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMusicServiceClient(conn)

	var wg sync.WaitGroup

	// Simular múltiplas requisições
	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			// Criar um novo usuário
			user := &pb.User{
				Name: fmt.Sprintf("User %d", i),
				Age:  int32(20 + (i % 50)),
			}

			createUserResponse, err := client.CreateUser(ctx, &pb.CreateUserRequest{User: user})
			if err != nil {
				log.Printf("could not create user: %v", err)
				return
			}
			log.Printf("Created User: %v\n", createUserResponse)

			// Obter o usuário criado
			getUserResponse, err := client.GetUser(ctx, &pb.GetUserRequest{Id: createUserResponse.Id})
			if err != nil {
				log.Printf("could not get user: %v", err)
				return
			}
			log.Printf("Got User: %v\n", getUserResponse)

			// Atualizar o usuário
			updatedUser := &pb.User{
				Id:   getUserResponse.Id,
				Name: fmt.Sprintf("Updated User %d", i),
				Age:  int32(25 + (i % 50)),
			}
			updateUserResponse, err := client.UpdateUser(ctx, &pb.UpdateUserRequest{User: updatedUser})
			if err != nil {
				log.Printf("could not update user: %v", err)
				return
			}
			log.Printf("Updated User: %v\n", updateUserResponse)

			// Deletar o usuário
			deleteUserResponse, err := client.DeleteUser(ctx, &pb.DeleteUserRequest{Id: getUserResponse.Id})
			if err != nil {
				log.Printf("could not delete user: %v", err)
				return
			}
			log.Printf("Deleted User: %v\n", deleteUserResponse)
		}(i)
	}

	// Aguardar a conclusão das requisições
	wg.Wait()
}
