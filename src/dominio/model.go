package dominio

type ActualizarProductoInput struct {
	Nombre      *string `json:"nombre,omitempty"`
	SKU         *string `json:"sku,omitempty"`
	Precio      *string `json:"precio,omitempty"`
	Descripcion *string `json:"descripcion,omitempty"`
}

type CrearProductoInput struct {
	Nombre      string `json:"nombre"`
	SKU         string `json:"sku"`
	Precio      string `json:"precio"`
	Descripcion string `json:"descripcion"`
}

type RespuestaEliminacion struct {
	Mensaje string `json:"mensaje"`
}

type Producto struct {
	ID          string `json:"id"`
	Nombre      string `json:"nombre"`
	SKU         string `json:"sku"`
	Precio      string `json:"precio"`
	Descripcion string `json:"descripcion"`
}

func (Producto) IsEntity() {}
