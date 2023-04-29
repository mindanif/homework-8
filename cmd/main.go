package main

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"google.golang.org/grpc"
	"homework-5/internal/pb"
	"homework-5/internal/pkg/db"
	"homework-5/internal/pkg/repository/postgresql/products"
	"homework-5/internal/pkg/repository/postgresql/warehouses"
	httpServer "homework-5/internal/pkg/server"
	"log"
	"net"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "test"
	dbname   = "test"
)

const (
	service     = "api"
	environment = "development"
)

func main() {
	tp, err := tracerProvider("http://localhost:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}
	otel.SetTracerProvider(tp)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	database, err := db.NewDB(ctx, dsn)
	if err != nil {
		fmt.Println(err)
		return
	}

	productRepo := products.NewProducts(database)
	warehouseRepo := warehouses.NewWarehouses(database)

	defer database.GetPool(ctx).Close()

	consoleServer := httpServer.New(productRepo, warehouseRepo)

	consoleServer.RunServer(ctx)

	go http.ListenAndServe(":9091", promhttp.Handler())

	server := grpc.NewServer()

	productsServer := pb.NewGrpcProductServer(productRepo)
	warehousesServer := pb.NewGrpcWarehouseServer(warehouseRepo)

	pb.RegisterProductServiceServer(server, productsServer)
	pb.RegisterWarehouseServiceServer(server, warehousesServer)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	if err = server.Serve(listener); err != nil {
		panic(err)
	}
}
func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(service),
			attribute.String("environment", environment),
		)),
	)
	return tp, nil
}
