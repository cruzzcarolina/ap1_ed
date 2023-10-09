package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Produto struct {
	ID        int
	Nome      string
	Descricao string
	Preco     float64
}

type Pedido struct {
	ID         int
	Entrega    bool
	Produtos   []Produto
	ValorTotal float64
}

type Sistema struct {
	Produtos []Produto
	Pedidos  []Pedido
	Carrinho struct {
		TotalProdutos int
		TotalPedidos  int
		TotalReceita  float64
	}
}

func main() {
	var sistema Sistema
	for {
		fmt.Println("McRonald's - Sistema de Pedidos Eletrônicos")
		fmt.Println("1. Cadastrar Produto")
		fmt.Println("2. Remover Produto")
		fmt.Println("3. Exibir Produtos")
		fmt.Println("4. Fazer Pedido")
		fmt.Println("5. Expedir Pedido")
		fmt.Println("6. Exibir Métricas do Sistema")
		fmt.Println("7. Sair")

		var escolha int
		fmt.Print("Escolha uma opção: ")
		fmt.Scanln(&escolha)

		switch escolha {
		case 1:
			sistema.AdicionarProduto()
		case 2:
			sistema.RemoverProduto()
		case 3:
			sistema.ExibirProdutos()
		case 4:
			sistema.FazerPedido()
		case 5:
			sistema.ExpedirPedido()
		case 6:
			sistema.ExibirMetricas()
		case 7:
			fmt.Println("Saindo do sistema.")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}

func (s *Sistema) AdicionarProduto() {
	var produto Produto
	produto.ID = s.Carrinho.TotalProdutos + 1

	fmt.Print("Nome do produto: ")
	nome, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	produto.Nome = strings.TrimSpace(nome)

	fmt.Print("Descrição do produto: ")
	descricao, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	produto.Descricao = strings.TrimSpace(descricao)

	fmt.Print("Preço do produto (R$): ")
	fmt.Scanln(&produto.Preco)

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
	var pedido Pedido
	pedido.ID = s.Carrinho.TotalPedidos + 1

	fmt.Print("Pedido para delievery (true/false): ")
	fmt.Scanln(&pedido.Entrega)

	fmt.Println("Escolha até 10 produtos para o pedido:")
	for {
		if len(pedido.Produtos) >= 10 {
			break
		}

		fmt.Print("ID do produto (0 para parar): ")
		var idProduto int
		fmt.Scanln(&idProduto)
		if idProduto == 0 {
			break
		}

		// Encontre o produto pelo ID e adicione ao pedido
		for _, produto := range s.Produtos {
			if produto.ID == idProduto {
				pedido.Produtos = append(pedido.Produtos, produto)
				pedido.ValorTotal += produto.Preco
			}
		}
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
