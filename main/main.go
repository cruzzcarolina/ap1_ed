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
		fmt.Println("2.Cadastrar Produtos em Lote via CSV")
		fmt.Println("3. Cadastrar Produtos em Lote")
		fmt.Println("4. Remover Produto")
		fmt.Println("5. Exibir Produtos")
		fmt.Println("6. Fazer Pedido")
		fmt.Println("7. Expedir Pedido")
		fmt.Println("8. Exibir Métricas do Sistema")
		fmt.Println("9. Exibir Pedidos Em Aberto")
		fmt.Println("10. Sair")
		fmt.Println("----------------------------")

		var escolha int
		fmt.Print("Escolha uma opção: ")
		fmt.Scanln(&escolha)

		switch escolha {
		case 1:
			fmt.Println("----------------------------")
			sistema.AdicionarProduto()
		case 2:
			fmt.Println("------------------------------")
			sistema.CadastrarProdutosEmLoteCSV()
		case 3:
			fmt.Println("------------------------------")
			sistema.CadastrarProdutosEmLote2()
		case 4:
			fmt.Println("----------------------------")
			sistema.RemoverProduto()
		case 5:
			fmt.Println("----------------------------")
			sistema.ExibirProdutos()
		case 6:
			fmt.Println("----------------------------")
			sistema.FazerPedido()
		case 7:
			fmt.Println("----------------------------")
			sistema.ExpedirPedido()
		case 8:
			fmt.Println("----------------------------")
			sistema.ExibirMetricas()
		case 9:
			fmt.Println("----------------------------")
			sistema.ExibirPedidosEmAberto()
		case 10:
			fmt.Println("----------------------------")
			fmt.Println("Saindo do sistema.")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}
