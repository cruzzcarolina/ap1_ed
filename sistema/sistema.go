package sistema

import (
	"bufio"
	"fmt"
	pedido "mcronalds/pedido"
	produto "mcronalds/produto"
	"os"
	"strings"
)

type Sistema struct {
	Produtos []produto.Produto
	Pedidos  []pedido.Pedido
	Carrinho struct {
		TotalProdutos int
		TotalPedidos  int
		TotalReceita  float64
	}
}

func (s *Sistema) AdicionarProduto() {
	var produto produto.Produto
	produto.ID = s.Carrinho.TotalProdutos + 1

	fmt.Print("Nome do produto: ")
	nome, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	produto.Nome = strings.TrimSpace(nome)

	fmt.Print("Descrição do produto: ")
	descricao, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	produto.Descricao = strings.TrimSpace(descricao)

	fmt.Print("Preço do produto (R$): ")
	fmt.Scanln(&produto.Preco)

	if produto.Preco < 0 {
		fmt.Println("Preço do produto deve ser um valor positivo.")
		return
	}
	// Verificar limite de produtos
	if s.Carrinho.TotalProdutos >= 50 {
		fmt.Println("Limite de produtos atingido. Não é possível adicionar mais produtos.")
		return
	}
	if s.Carrinho.TotalPedidos >= 1000 {
		fmt.Println("Limite de pedidos atingido. Não é possível criar mais pedidos.")
		return
	}

	s.Produtos = append(s.Produtos, produto)
	s.Carrinho.TotalProdutos++
	fmt.Println("Produto cadastrado com sucesso.")
}

func (s *Sistema) RemoverProduto() {
	fmt.Print("Digite o ID do produto a ser removido: ")
	var idProduto int
	fmt.Scanln(&idProduto)

	for i, produto := range s.Produtos {
		if produto.ID == idProduto {
			s.Produtos = append(s.Produtos[:i], s.Produtos[i+1:]...)
			fmt.Println("Produto removido com sucesso.")
			return
		}
	}

	fmt.Println("Produto não encontrado.")
}

func (s *Sistema) ExibirProdutos() {
	fmt.Println("Lista de Produtos:")
	for _, produto := range s.Produtos {
		fmt.Printf("ID: %d, Nome: %s, Descrição: %s, Preço: R$%.2f\n", produto.ID, produto.Nome, produto.Descricao, produto.Preco)
	}
}

func (s *Sistema) FazerPedido() {
	var pedido pedido.Pedido
	pedido.ID = s.Carrinho.TotalPedidos + 1

	fmt.Print("Pedido para delivery (true/false): ")
	fmt.Scanln(&pedido.Entrega)

	fmt.Println("Escolha até 10 produtos para o pedido:")
	for len(pedido.Produtos) < 10 {
		fmt.Println("Escolha como deseja adicionar produtos ao pedido:")
		fmt.Println("1. Por ID")
		fmt.Println("2. Por Nome")
		fmt.Print("Escolha uma opção (1/2): ")
		var opcao int
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			fmt.Print("ID do produto (0 para parar): ")
			var idProduto int
			fmt.Scanln(&idProduto)
			if idProduto == 0 {
				break
			}

			produtoEncontrado := false
			for _, p := range s.Produtos {
				if p.ID == idProduto {
					pedido.Produtos = append(pedido.Produtos, p)
					pedido.ValorTotal += p.Preco
					produtoEncontrado = true
					break
				}
			}
			if !produtoEncontrado {
				fmt.Println("Produto não encontrado.")
			}
		case 2:
			fmt.Print("Nome do produto (0 para parar): ")
			var nomeProduto string
			fmt.Scanln(&nomeProduto)
			if nomeProduto == "" {
				break
			}

			produtoEncontrado := false
			for _, p := range s.Produtos {
				if p.Nome == nomeProduto {
					pedido.Produtos = append(pedido.Produtos, p)
					pedido.ValorTotal += p.Preco
					produtoEncontrado = true
					break
				}
			}
			if !produtoEncontrado {
				fmt.Println("Produto não encontrado.")
			}
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}

		fmt.Print("Deseja adicionar mais produtos? (s/n): ")
		var continuar string
		fmt.Scanln(&continuar)
		if continuar != "s" {
			break
		}
	}

	if len(pedido.Produtos) == 0 {
		fmt.Println("Pedido de delivery deve incluir pelo menos um produto.")
		return
	}

	// Verificar limite de pedidos
	if s.Carrinho.TotalPedidos >= 1000 {
		fmt.Println("Limite de pedidos atingido. Não é possível criar mais pedidos.")
		return
	}

	s.Pedidos = append(s.Pedidos, pedido)
	s.Carrinho.TotalPedidos++
	fmt.Println("Pedido feito com sucesso.")
}

func (s *Sistema) ExpedirPedido() {
	if len(s.Pedidos) == 0 {
		fmt.Println("Nenhum pedido pendente para expedir.")
		return
	}

	pedido := s.Pedidos[0]
	s.Pedidos = s.Pedidos[1:]
	s.Carrinho.TotalReceita += pedido.ValorTotal

	fmt.Printf("Pedido #%d expedido:\n", pedido.ID)
	fmt.Printf("Entrega: %t, Total: R$%.2f\n", pedido.Entrega, pedido.ValorTotal)
}

func (s *Sistema) ExibirMetricas() {
	fmt.Printf("Número total de produtos cadastrados: %d\n", s.Carrinho.TotalProdutos)
	fmt.Printf("Número de pedidos encerrados: %d\n", s.Carrinho.TotalPedidos)
	fmt.Printf("Faturamento total até o momento: R$%.2f\n", s.Carrinho.TotalReceita)
}

func (s *Sistema) CadastrarProdutosEmLote(produtos []produto.Produto) {
	for _, p := range produtos {
		p.ID = s.Carrinho.TotalProdutos + 1
		s.Produtos = append(s.Produtos, p)
		s.Carrinho.TotalProdutos++
	}
}

func (s *Sistema) ExibirPedidosEmAberto() {
	fmt.Println("Pedidos em aberto:")
	for _, pedido := range s.Pedidos {
		fmt.Printf("ID: %d, Entrega: %t, Total: R$%.2f\n", pedido.ID, pedido.Entrega, pedido.ValorTotal)
	}
}
