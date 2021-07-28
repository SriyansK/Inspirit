package main

import(
	"fmt"
	"strconv"
	"context"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"Inspirit/Inspirit"
)

type server struct {
	proto.UnimplementedUsersServer
}

type person struct {
	firstName string
	lastName string
	age int64
	dateJoined string
	billingAddress string
}
const (
	port = ":50051"
)

// global variables declaration
var idCounter int64 = 0
var personMap map[int64]person
var currPerson person 

func (s* server) SearchInfo (ctx context.Context ,request *proto.Request) (*proto.Response , error){

	currId := request.GetId()
	_ , ok := personMap[currId]

	fmt.Printf("ID received  = %d \n",currId)

	if !ok {
		return &proto.Response{
			FirstName: "" , 
			LastName: "",
			Age: -1,
			DateJoined: "",
			BillingAddress: "",
			Valid: false,
		} , nil
	}

	currPerson = personMap[currId]
	return &proto.Response{
		FirstName: currPerson.firstName , 
		LastName: currPerson.lastName,
		Age: currPerson.age,
		DateJoined: currPerson.dateJoined,
		BillingAddress: currPerson.billingAddress,
		Valid: true,
	} , nil
}

func(s* server) AddUser (ctx context.Context ,request *proto.NewUserRequest) (*proto.NewUserResponse , error){
	
	currPerson.firstName = request.GetFirstName()
	currPerson.lastName = request.GetLastName()
	currPerson.age = request.GetAge()
	currPerson.dateJoined = request.GetDateJoined()
	currPerson.billingAddress = request.GetBillingAddress()
	idCounter++
	personMap[idCounter] = currPerson

	fmt.Printf("New user information received %v\n",currPerson)

	return &proto.NewUserResponse{
		Id: idCounter,
		FirstName: currPerson.firstName , 
		LastName: currPerson.lastName,
		Age: currPerson.age,
		DateJoined: currPerson.dateJoined,
		BillingAddress: currPerson.billingAddress,
	} , nil
}

func(s* server) UpdateUser(ctx context.Context ,request *proto.UpdateUserRequest) (*proto.UpdateUserResponse , error){
	currId := request.GetId()
	_ , ok := personMap[currId]

	if !ok {
		return &proto.UpdateUserResponse{
		FirstName: "" , 
		LastName: "",
		Age: -1,
		BillingAddress: "",
		Valid: false,
		} , nil
	}

	currPerson = personMap[currId]
	currPerson.firstName = request.FirstName
	currPerson.lastName = request.LastName
	currPerson.age = request.Age
	currPerson.billingAddress = request.BillingAddress
	personMap[currId] = currPerson

	fmt.Printf("User Updated!\n");

	return &proto.UpdateUserResponse{
		FirstName: currPerson.firstName , 
		LastName: currPerson.lastName,
		Age: currPerson.age,
		BillingAddress: currPerson.billingAddress,
		Valid: true,
	} , nil
}

func (s* server) DeleteUser (ctx context.Context ,request *proto.DeleteUserRequest) (*proto.DeleteUserResponse , error){

	currId := request.GetId()
	_ , ok := personMap[currId]

	if !ok {
		return &proto.DeleteUserResponse{Mssg: "No User Found!"} , nil
	}

	delete(personMap, currId);
	mssg := "User with ID = " + strconv.Itoa(int(currId)) + " Deleted"
	return &proto.DeleteUserResponse{Mssg: mssg} , nil
}

func main(){

	personMap = make(map[int64]person)
	lis , err := net.Listen("tcp", port)
	if err!=nil {
		panic(err);
	}

	srv := grpc.NewServer()
	proto.RegisterUsersServer(srv , &server{})
	
	reflection.Register(srv) 

	if e:=srv.Serve(lis); e!=nil{
		panic(e)
	}
}