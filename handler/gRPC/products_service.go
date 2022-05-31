package gRPC

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"log"
	"net"
	"test-api/controller"
	"test-api/model"
	"test-api/proto"
)

type grpcService struct {
	pc controller.ProductController
	productService.UnimplementedProductServiceServer
	db *gorm.DB
}

func NewGrpcService(pc controller.ProductController) *grpcService {
	return &grpcService{
		pc: pc,
	}

}

//GetAllProducts
//List all products on the database table,takes empty message
func (g *grpcService) GetAllProducts(ctx context.Context, in *emptypb.Empty) (*productService.GetAllProductRes, error) {
	products, err := g.pc.GetAllProducts()

	if err != nil {
		return nil, status.Error(codes.NotFound, "Listing all products have failed")
	}
	var productRes []*productService.ProductReq
	for _, product := range products {
		productRes = append(productRes, product.ToProto2())
	}
	fmt.Println(productRes)
	return &productService.GetAllProductRes{Products: productRes}, nil
}

//GetProductByID
//Passes gRPC request to controller, changes result to proto message and returns that message
func (g *grpcService) GetProductByID(ctx context.Context, id *productService.IdReq) (*productService.ProductRes, error) {
	product, err := g.pc.GetProductByID(id.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "Product with this ID not found")
	}

	return product.ToProto(), nil
}

//CreateProduct
//Changes gRPC request into Product model and passes that to controller, returns created Product as proto message
func (g *grpcService) CreateProduct(ctx context.Context, pGrpc *productService.ProductReq) (*productService.ProductRes, error) {
	product := model.Product{
		ID:         pGrpc.GetId(),
		Name:       pGrpc.GetName(),
		Detail:     pGrpc.GetDetail(),
		Price:      pGrpc.GetPrice(),
		IsCampaign: pGrpc.GetIsCampaign(),
	}
	err := g.pc.CreateProduct(&product)
	if err != nil {
		return nil, err
	}

	return product.ToProto(), nil
}

//UpdateProduct
////Changes gRPC request into Product model and passes that to controller, returns updated Product as proto message
func (g *grpcService) UpdateProduct(ctx context.Context, pGrpc *productService.ProductReq) (*productService.ProductRes, error) {
	product := model.Product{
		ID:         pGrpc.GetId(),
		Name:       pGrpc.GetName(),
		Detail:     pGrpc.GetDetail(),
		Price:      pGrpc.GetPrice(),
		IsCampaign: pGrpc.GetIsCampaign(),
	}
	res, err := g.pc.UpdateProduct(pGrpc.GetId(), &product)
	if err != nil {
		return nil, err
	}
	return res.ToProto(), nil
}

//DeleteProduct
//Gets id from the proto message and deletes the record with that id from table. Returns nil(need to change this)
func (g *grpcService) DeleteProduct(ctx context.Context, id *productService.IdReq) (*emptypb.Empty, error) {
	err := g.pc.DeleteProduct(id.GetId())
	if err != nil {
		return nil, err
	}

	return nil, nil
}

//NewGrpc
//Starts the gRPC server
func (g *grpcService) NewGrpc() {
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Printf("Failed on listener: %v", err)
	}

	s := grpc.NewServer()
	productService.RegisterProductServiceServer(s, g)

	log.Printf("Listener on %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed on server %v", err)
	}

}
