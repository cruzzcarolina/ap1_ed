package pedido

import (
	produto "mcronalds/produto"
	"time"
)

type Pedido struct {
	ID         int
	Entrega    bool
	Produtos   []produto.Produto
	ValorTotal float64
	DataHora   time.Time
}
