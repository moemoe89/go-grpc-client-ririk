//
//  Practicing gRPC Client
//
//  Copyright © 2020. All rights reserved.
//

package main

import (
	conf "github.com/moemoe89/practicing-grpc-client-golang/config"
	"github.com/moemoe89/practicing-grpc-client-golang/user/proto"

	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial(conf.Configuration.GrpcServer, opts)
	if err != nil {
		panic(err)
	}
	defer cc.Close()

	c := proto.NewUserServiceClient(cc)

	createUserRes, err := c.Create(context.Background(), &proto.UserCreateReq{
		Name:    "John",
		Phone:   "081234567890",
		Email:   "john@doe.com",
		Address: "USA",
	})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("user has been created: %v", createUserRes)

	detailUserRes, err := c.Detail(context.Background(), &proto.UserIDReq{
		Id: createUserRes.GetId(),
	})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("user detail has been retrieved: %v", detailUserRes)

	listUserRes, err := c.List(context.Background(), &proto.UsersReq{
		Page:    "1",
		PerPage: "10",
	})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("user list has been retrieved: %v", listUserRes)

	updateUserRes, err := c.Update(context.Background(), &proto.UserUpdateReq{
		Id:      createUserRes.GetId(),
		Name:    "Johny",
		Phone:   "081234567891",
		Email:   "johny@doe.com",
		Address: "UK",
	})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("user has been updated: %v", updateUserRes)

	deleteUserRes, err := c.Delete(context.Background(), &proto.UserIDReq{
		Id: createUserRes.GetId(),
	})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("user has been deleted: %v", deleteUserRes)
}