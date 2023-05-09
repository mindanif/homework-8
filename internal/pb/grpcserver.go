package pb

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	internal "homework-5/internal/model"
	"homework-5/internal/pkg/repository"
	"log"
)

type grpcProductServer struct {
	productRepo repository.ProductsRepo
}

func NewGrpcProductServer(productRepo repository.ProductsRepo) *grpcProductServer {
	return &grpcProductServer{productRepo: productRepo}
}
func (p *grpcProductServer) CreateProduct(ctx context.Context, req *CreateProductRequest) (*CreateProductResponse, error) {
	tr := otel.Tracer("CreateProduct")
	ctx, span := tr.Start(ctx, "received request")
	span.SetAttributes(attribute.Key("params").String(req.String()))
	defer span.End()

	product := repository.Products{
		Name:        req.Name,
		Description: req.Description,
		Price:       int(req.Price),
		WarehouseId: int(req.WarehouseId),
	}

	id, err := p.productRepo.Create(ctx, &product)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	internal.RegProductCounter.Add(1)

	log.Println("New product with id:", id)

	response := CreateProductResponse{Id: uint32(id)}

	return &response, nil
}
func (p *grpcProductServer) UpdateProduct(ctx context.Context, req *UpdateProductRequest) (*UpdateProductResponse, error) {
	tr := otel.Tracer("UpdateProduct")
	ctx, span := tr.Start(ctx, "received request")
	span.SetAttributes(attribute.Key("params").String(req.String()))
	defer span.End()

	product := repository.Products{
		Id:          int(req.Id),
		Name:        req.Name,
		Description: req.Description,
		Price:       int(req.Price),
		WarehouseId: int(req.WarehouseId),
	}

	ok, err := p.productRepo.Update(ctx, &product)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(ok)

	response := UpdateProductResponse{Ok: ok}

	return &response, nil
}
func (p *grpcProductServer) DeleteProduct(ctx context.Context, req *DeleteProductRequest) (*DeleteProductResponse, error) {
	tr := otel.Tracer("DeleteProduct")
	ctx, span := tr.Start(ctx, "DeleteProduct request")
	span.SetAttributes(attribute.Key("params").String(req.String()))
	defer span.End()

	id := int(req.Id)
	ok, err := p.productRepo.Delete(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(ok)

	internal.DeletedProductCounter.Add(1)

	response := DeleteProductResponse{Ok: ok}

	return &response, nil
}
func (p *grpcProductServer) mustEmbedUnimplementedProductServiceServer() {

}

type grpcWarehouseServer struct {
	warehouseRepo repository.WarehousesRepo
}

func NewGrpcWarehouseServer(warehousesRepo repository.WarehousesRepo) *grpcWarehouseServer {
	return &grpcWarehouseServer{warehouseRepo: warehousesRepo}
}

func (w *grpcWarehouseServer) CreateWarehouse(ctx context.Context, req *CreateWarehouseRequest) (*CreateWarehouseResponse, error) {
	tr := otel.Tracer("CreateWarehouse")
	ctx, span := tr.Start(ctx, "CreateWarehouse request")
	span.SetAttributes(attribute.Key("params").String(req.String()))
	defer span.End()

	warehouse := repository.Warehouses{
		Name:   req.Name,
		City:   req.City,
		Square: int(req.Square),
	}

	id, err := w.warehouseRepo.Add(ctx, &warehouse)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("New product with id:", id)

	response := CreateWarehouseResponse{Id: uint32(id)}

	return &response, nil
}
func (w *grpcWarehouseServer) UpdateWarehouse(ctx context.Context, req *UpdateWarehouseRequest) (*UpdateWarehouseResponse, error) {
	tr := otel.Tracer("UpdateWarehouse")
	ctx, span := tr.Start(ctx, "UpdateWarehouse request")
	span.SetAttributes(attribute.Key("params").String(req.String()))
	defer span.End()

	warehouse := repository.Warehouses{
		Id:     int(req.Id),
		Name:   req.Name,
		City:   req.City,
		Square: int(req.Square),
	}

	ok, err := w.warehouseRepo.Update(ctx, &warehouse)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(ok)

	response := UpdateWarehouseResponse{Ok: ok}

	return &response, nil
}
func (w *grpcWarehouseServer) DeleteWarehouse(ctx context.Context, req *DeleteWarehouseRequest) (*DeleteWarehouseResponse, error) {
	tr := otel.Tracer("DeleteWarehouse")
	ctx, span := tr.Start(ctx, "DeleteWarehouse request")
	span.SetAttributes(attribute.Key("params").String(req.String()))
	defer span.End()

	id := int(req.Id)
	ok, err := w.warehouseRepo.Delete(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(ok)

	response := DeleteWarehouseResponse{Ok: ok}

	return &response, nil
}
func (w *grpcWarehouseServer) mustEmbedUnimplementedWarehouseServiceServer() {}
