package produtolote

type ProdutoEmLote struct {
	ProdutoID  int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func NovaProdutoEmLote(produtoID int, nome string, descricao string, preco float64, quantidade int) ProdutoEmLote {
	return ProdutoEmLote{
		ProdutoID:  produtoID,
		Nome:       nome,
		Descricao:  descricao,
		Preco:      preco,
		Quantidade: quantidade,
	}
}
