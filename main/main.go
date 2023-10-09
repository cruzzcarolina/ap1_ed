package main

import (
	"fmt"
	s "mcronalds/sistema"
	"os"
)

func main() {
	var sistema s.Sistema
	for {
		fmt.Println("----------------------------")
		fmt.Println("McRonald's - Sistema de Pedidos Eletrônicos")
		fmt.Println("1. Cadastrar Produto")
		fmt.Println("2. Remover Produto")
		fmt.Println("3. Exibir Produtos")
		fmt.Println("4. Fazer Pedido")
		fmt.Println("5. Expedir Pedido")
		fmt.Println("6. Exibir Métricas do Sistema")
		fmt.Println("7. Exibir Pedidos Em Aberto")
		fmt.Println("7. Sair")
		fmt.Println("----------------------------")

		var escolha int
		fmt.Print("Escolha uma opção: ")
		fmt.Scanln(&escolha)

		switch escolha {
		case 1:
			fmt.Println("----------------------------")
			sistema.AdicionarProduto()
		case 2:
			fmt.Println("----------------------------")
			sistema.RemoverProduto()
		case 3:
			fmt.Println("----------------------------")
			sistema.ExibirProdutos()
		case 4:
			fmt.Println("----------------------------")
			sistema.FazerPedido()
		case 5:
			fmt.Println("----------------------------")
			sistema.ExpedirPedido()
		case 6:
			fmt.Println("----------------------------")
			sistema.ExibirMetricas()
		case 7:
			fmt.Println("----------------------------")
			sistema.ExibirPedidosEmAberto()
		case 8:
			fmt.Println("----------------------------")
			fmt.Println("Saindo do sistema.")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}
