package database

import (
	"log"

	model "github.com/ProyectoIntegradorSoftware/MicroservicioProductos/dominio"
	"gorm.io/gorm"
)

// EjecutarMigraciones realiza todas las migraciones necesarias en la base de datos.
func EjecutarMigraciones(db *gorm.DB) {

	db.AutoMigrate(&model.ProductoGORM{})

	log.Println("Migraciones completadas")
}
