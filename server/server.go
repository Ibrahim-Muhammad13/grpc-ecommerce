package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/Ibrahim-Muhammad13/e-comm/pb"
	"github.com/golang-jwt/jwt/v4"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var db *gorm.DB

type server struct {
	pb.UnimplementedUserServiceServer
	pb.UnimplementedLoginServiceServer
	pb.UnimplementedProductServiceServer
	//pb.ProductServiceServer
}
type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Role     string
}

func (*server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := User{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Role:     "user",
	}
	result := db.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("internal error %v", result.Error),
		)
	}

	id := user.ID

	return &pb.CreateUserResponse{
		User: &pb.User{
			Id:        uint64(id),
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		},
	}, nil
}

func (*server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	var user User
	db.Where("name = ?", req.GetName()).First(&user)
	if user.ID == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("user not found"),
		)
	}
	if user.Password != req.GetPassword() {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("password not match"),
		)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user.ID,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("hmacSampleSecret")))
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("internal error %v", err),
		)
	}
	return &pb.LoginResponse{
		Token: tokenString,
		User: &pb.User{
			Id:        uint64(user.ID),
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		},
	}, nil
}

// product servie

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
}

func (*server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product := Product{
		Name:        req.GetProductName(),
		Description: req.GetProductDescription(),
		Price:       req.GetProductPrice(),
	}
	result := db.Create(&product)
	if result.Error != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("internal error %v", result.Error),
		)
	}
	id := product.ID
	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:                 uint64(id),
			ProductName:        product.Name,
			ProductDescription: product.Description,
			ProductPrice:       product.Price,
			CreatedAt:          timestamppb.New(product.CreatedAt),
			UpdatedAt:          timestamppb.New(product.UpdatedAt),
		},
	}, nil
}

func (*server) GetProducts(ctx context.Context, req *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	var products []Product
	db.Find(&products)
	var productspb []*pb.Product
	for _, product := range products {
		productspb = append(productspb, ProductToProductpb(&product))
	}
	return &pb.GetProductsResponse{
		Products: productspb,
	}, nil
}
func ProductToProductpb(product *Product) *pb.Product {
	return &pb.Product{
		Id:                 uint64(product.ID),
		ProductName:        product.Name,
		ProductDescription: product.Description,
		ProductPrice:       product.Price,
		CreatedAt:          timestamppb.New(product.CreatedAt),
		UpdatedAt:          timestamppb.New(product.UpdatedAt),
	}
}

func (*server) SearchProduct(ctx context.Context, req *pb.SearchProductRequest) (*pb.SearchProductResponse, error) {
	var products []Product
	db.Where("name LIKE ?", "%"+req.GetProductName()+"%").Find(&products)
	if len(products) == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("product not found"),
		)
	}
	var productspb []*pb.Product
	for _, product := range products {
		productspb = append(productspb, ProductToProductpb(&product))
	}
	return &pb.SearchProductResponse{
		Product: productspb,
	}, nil
}
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	d, err := gorm.Open("mysql", "root:@/grpc-ecom?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	db = d
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Product{})
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Faild to listen %v", err)
	}
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(s, &server{})
	pb.RegisterLoginServiceServer(s, &server{})
	pb.RegisterProductServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("faild to serve %v", err)
	}

}
