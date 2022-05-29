package gRPC

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

//GetProductByID
//Passes gRPC request to controller, changes result to proto message and returns that message
func (g *grpcService) GetProductByID(ctx context.Context, id *productService.ID) (*productService.Product, error) {
	product, err := g.pc.GetProductByID(id.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "Product with this ID not found")
	}

	return product.ToProto(), nil
}

//CreateProduct
//Changes gRPC request into Product model and passes that to controller, returns created Product as proto message
func (g *grpcService) CreateProduct(ctx context.Context, pGrpc *productService.Product) (*productService.Product, error) {
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

	return pGrpc, nil
}

//UpdateProduct
////Changes gRPC request into Product model and passes that to controller, returns updated Product as proto message
func (g *grpcService) UpdateProduct(ctx context.Context, pGrpc *productService.Product) (*productService.Product, error) {
	product := model.Product{
		ID:         pGrpc.GetId(),
		Name:       pGrpc.GetName(),
		Detail:     pGrpc.GetDetail(),
		Price:      pGrpc.GetPrice(),
		IsCampaign: pGrpc.GetIsCampaign(),
	}
	_, err := g.pc.UpdateProduct(pGrpc.GetId(), &product)
	if err != nil {
		return nil, err
	}
	return pGrpc, nil
}

//DeleteProduct
//Gets id from the proto message and deletes the record with that id from table. Returns nil(need to change this)
func (g *grpcService) DeleteProduct(ctx context.Context, id *productService.ID) (*productService.Product, error) {
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
