package adapters

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/ProyectoIntegradorSoftware/MicroservicioProductos/database"
	model "github.com/ProyectoIntegradorSoftware/MicroservicioProductos/dominio"
	"github.com/ProyectoIntegradorSoftware/MicroservicioProductos/ports"

	"gorm.io/gorm"
)

/**
* Es un adaptador de salida

 */

type productRepository struct {
	db             *database.DB
	activeSessions map[string]string
}

func NewProductRepository(db *database.DB) ports.ProductRepository {
	return &productRepository{
		db:             db,
		activeSessions: make(map[string]string),
	}
}

func ToJSON(obj interface{}) (string, error) {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonData), err
}

// ObtenerProducto obtiene un producto por su ID.
func (ur *productRepository) Producto(id string) (*model.Producto, error) {
	if id == "" {
		return nil, errors.New("El ID de producto es requerido")
	}

	var productoGORM model.ProductoGORM
	//result := ur.db.GetConn().First(&productoGORM, id)
	result := ur.db.GetConn().First(&productoGORM, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		}
		log.Printf("Error al obtener el producto con ID %s: %v", id, result.Error)
		return nil, result.Error
	}

	return productoGORM.ToGQL()
}

// Productos obtiene todos los productos de la base de datos.
func (ur *productRepository) Productos() ([]*model.Producto, error) {
	var productosGORM []model.ProductoGORM
	result := ur.db.GetConn().Find(&productosGORM)

	if result.Error != nil {
		log.Printf("Error al obtener los productos: %v", result.Error)
		return nil, result.Error
	}

	var productos []*model.Producto
	for _, productoGORM := range productosGORM {
		producto, _ := productoGORM.ToGQL()
		productos = append(productos, producto)
	}

	return productos, nil
}
func (ur *productRepository) CrearProducto(input model.CrearProductoInput) (*model.Producto, error) {

	productoGORM :=
		&model.ProductoGORM{
			Nombre:      input.Nombre,
			SKU:         input.SKU,
			Precio:      input.Precio,
			Descripcion: input.Descripcion,
		}
	result := ur.db.GetConn().Create(&productoGORM)
	if result.Error != nil {
		log.Printf("Error al crear producto: %v", result.Error)
		return nil, result.Error
	}

	response, err := productoGORM.ToGQL()
	return response, err
}

func (ur *productRepository) ActualizarProducto(id string, input *model.ActualizarProductoInput) (*model.Producto, error) {
	var productoGORM model.ProductoGORM
	if id == "" {
		return nil, errors.New("El ID de producto es requerido")
	}

	result := ur.db.GetConn().First(&productoGORM, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Producto con ID %s no encontrado", id)
		}
		return nil, result.Error
	}

	// Solo actualiza los campos proporcionados
	if input.Nombre != nil {
		productoGORM.Nombre = *input.Nombre
	}
	if input.Precio != nil {
		productoGORM.Precio = *input.Precio
	}
	if input.Descripcion != nil {
		productoGORM.Descripcion = *input.Descripcion
	}

	result = ur.db.GetConn().Save(&productoGORM)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Printf("Producto actualizado: %v", productoGORM)
	return productoGORM.ToGQL()
}

// EliminarProducto elimina un producto de la base de datos por su ID.
func (ur *productRepository) EliminarProducto(id string) (*model.RespuestaEliminacion, error) {
	// Intenta buscar el producto por su ID
	var productoGORM model.ProductoGORM
	result := ur.db.GetConn().First(&productoGORM, id)

	if result.Error != nil {
		// Manejo de errores
		if result.Error == gorm.ErrRecordNotFound {
			// El producto no se encontró en la base de datos
			response := &model.RespuestaEliminacion{
				Mensaje: "El producto no existe",
			}
			return response, result.Error

		}
		log.Printf("Error al buscar el producto con ID %s: %v", id, result.Error)
		response := &model.RespuestaEliminacion{
			Mensaje: "Error al buscar el producto",
		}
		return response, result.Error
	}

	// Elimina el producto de la base de datos
	result = ur.db.GetConn().Delete(&productoGORM, id)

	if result.Error != nil {
		log.Printf("Error al eliminar el producto con ID %s: %v", id, result.Error)
		response := &model.RespuestaEliminacion{
			Mensaje: "Error al eliminar el producto",
		}
		return response, result.Error
	}

	// Éxito al eliminar el producto
	response := &model.RespuestaEliminacion{
		Mensaje: "Producto eliminado con éxito",
	}
	return response, result.Error

}
