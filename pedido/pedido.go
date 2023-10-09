package pedido

import produto "mcronalds/produto"

type Pedido struct {
	ID         int
	Entrega    bool
	Produtos   []produto.Produto
	ValorTotal float64
}
