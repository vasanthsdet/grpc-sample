package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "grpc-sample/proto"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewUserServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // Call CreateUser
    createRes, err := client.CreateUser(ctx, &pb.CreateUserRequest{Name: "Alice", Age: 25})
    if err != nil {
        log.Fatalf("could not create user: %v", err)
    }
    log.Printf("User created with ID: %s", createRes.Id)

    // Call GetUser
    getRes, err := client.GetUser(ctx, &pb.GetUserRequest{Id: createRes.Id})
    if err != nil {
        log.Fatalf("could not get user: %v", err)
    }
    log.Printf("User: %s, Age: %d", getRes.Name, getRes.Age)
}
