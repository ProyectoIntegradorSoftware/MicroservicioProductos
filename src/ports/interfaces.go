package ports

import (
	model "github.com/ProyectoIntegradorSoftware/MicroservicioProductos/dominio"
)

// puerto de salida
type ProductRepository interface {
	CrearProducto(input model.CrearProductoInput) (*model.Producto, error)
	Producto(id string) (*model.Producto, error)
	ActualizarProducto(id string, input *model.ActualizarProductoInput) (*model.Producto, error)
	EliminarProducto(id string) (*model.RespuestaEliminacion, error)
	Productos() ([]*model.Producto, error)
}
