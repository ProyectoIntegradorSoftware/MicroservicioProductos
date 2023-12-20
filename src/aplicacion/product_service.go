package service

import (
	"context"
	"fmt"

	model "github.com/ProyectoIntegradorSoftware/MicroservicioProductos/dominio"
	repository "github.com/ProyectoIntegradorSoftware/MicroservicioProductos/ports"
	pb "github.com/ProyectoIntegradorSoftware/MicroservicioProductos/proto"
)

// este servicio implementa la interfaz ProductServiceServer
// que se genera a partir del archivo proto
type ProductService struct {
	pb.UnimplementedProductServiceServer
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {

	crearProductoInput := model.CrearProductoInput{
		Nombre:      req.GetNombre(),
		SKU:         req.GetSKU(),
		Precio:      req.GetPrecio(),
		Descripcion: req.GetDescripcion(),
	}
	u, err := s.repo.CrearProducto(crearProductoInput)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Producto creado: %v", u)
	response := &pb.CreateProductResponse{
		Id:          u.ID,
		Nombre:      u.Nombre,
		SKU:         u.SKU,
		Precio:      u.Precio,
		Descripcion: u.Descripcion,
	}
	fmt.Printf("Producto creado: %v", response)
	return response, nil
}

func (s *ProductService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	u, err := s.repo.Producto(req.GetId())
	if err != nil {
		return nil, err
	}
	response := &pb.GetProductResponse{
		Id:          u.ID,
		Nombre:      u.Nombre,
		SKU:         u.SKU,
		Precio:      u.Precio,
		Descripcion: u.Descripcion,
	}
	return response, nil
}

func (s *ProductService) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	products, err := s.repo.Productos()
	if err != nil {
		return nil, err
	}
	var response []*pb.Product
	for _, u := range products {
		product := &pb.Product{
			Id:          u.ID,
			Nombre:      u.Nombre,
			SKU:         u.SKU,
			Precio:      u.Precio,
			Descripcion: u.Descripcion,
		}
		response = append(response, product)
	}

	return &pb.ListProductsResponse{Products: response}, nil
}

func (s *ProductService) UpdateUser(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	nombre := req.GetNombre()
	SKU := req.GetSKU()
	precio := req.GetPrecio()
	descripcion := req.GetDescripcion()
	fmt.Printf("Nombre: %v", nombre)
	actualizarProductoInput := &model.ActualizarProductoInput{
		Nombre:      &nombre,
		SKU:         &SKU,
		Precio:      &precio,
		Descripcion: &descripcion,
	}
	fmt.Printf("Producto actualizado input: %v", actualizarProductoInput)
	u, err := s.repo.ActualizarProducto(req.GetId(), actualizarProductoInput)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Producto actualizado: %v", u)
	response := &pb.UpdateProductResponse{
		Id:          u.ID,
		Nombre:      u.Nombre,
		SKU:         u.SKU,
		Precio:      u.Precio,
		Descripcion: u.Descripcion,
	}
	return response, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	respuesta, err := s.repo.EliminarProducto(req.GetId())
	if err != nil {
		return nil, err
	}
	response := &pb.DeleteProductResponse{
		Mensaje: respuesta.Mensaje,
	}
	return response, nil
}
