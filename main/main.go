package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
}

type Order struct {
	ID          int
	Delivery    bool
	Products    []Product
	TotalAmount float64
}

type System struct {
	Products  []Product
	Orders    []Order
	FoodTruck struct {
		TotalProducts int
		TotalOrders   int
		TotalRevenue  float64
	}
}

func main() {
	var system System
	for {
		fmt.Println("McRonald's - Sistema de Pedidos Eletrônicos")
		fmt.Println("1. Cadastrar Produto")
		fmt.Println("2. Remover Produto")
		fmt.Println("3. Exibir Produtos")
		fmt.Println("4. Fazer Pedido")
		fmt.Println("5. Expedir Pedido")
		fmt.Println("6. Exibir Métricas do Sistema")
		fmt.Println("7. Sair")

		var choice int
		fmt.Print("Escolha uma opção: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			system.AddProduct()
		case 2:
			system.RemoveProduct()
		case 3:
			system.ShowProducts()
		case 4:
			system.PlaceOrder()
		case 5:
			system.ExpediteOrder()
		case 6:
			system.ShowMetrics()
		case 7:
			fmt.Println("Saindo do sistema.")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}

func (s *System) AddProduct() {
	var product Product
	product.ID = s.FoodTruck.TotalProducts + 1

	fmt.Print("Nome do produto: ")
	name, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	product.Name = strings.TrimSpace(name)

	fmt.Print("Descrição do produto: ")
	description, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	product.Description = strings.TrimSpace(description)

	fmt.Print("Preço do produto (R$): ")
	fmt.Scanln(&product.Price)

	s.Products = append(s.Products, product)
	s.FoodTruck.TotalProducts++
	fmt.Println("Produto cadastrado com sucesso.")
}

func (s *System) RemoveProduct() {
	fmt.Print("Digite o ID do produto a ser removido: ")
	var productID int
	fmt.Scanln(&productID)

	for i, product := range s.Products {
		if product.ID == productID {
			s.Products = append(s.Products[:i], s.Products[i+1:]...)
			fmt.Println("Produto removido com sucesso.")
			return
		}
	}

	fmt.Println("Produto não encontrado.")
}

func (s *System) ShowProducts() {
	fmt.Println("Lista de Produtos:")
	for _, product := range s.Products {
		fmt.Printf("ID: %d, Nome: %s, Descrição: %s, Preço: R$%.2f\n", product.ID, product.Name, product.Description, product.Price)
	}
}

func (s *System) PlaceOrder() {
	var order Order
	order.ID = s.FoodTruck.TotalOrders + 1

	fmt.Print("Pedido para entrega (true/false): ")
	fmt.Scanln(&order.Delivery)

	fmt.Println("Escolha até 10 produtos para o pedido:")
	for {
		if len(order.Products) >= 10 {
			break
		}

		fmt.Print("ID do produto (0 para parar): ")
		var productID int
		fmt.Scanln(&productID)
		if productID == 0 {
			break
		}

		// Encontre o produto pelo ID e adicione ao pedido
		for _, product := range s.Products {
			if product.ID == productID {
				order.Products = append(order.Products, product)
				order.TotalAmount += product.Price
			}
		}
	}

	s.Orders = append(s.Orders, order)
	s.FoodTruck.TotalOrders++
	fmt.Println("Pedido feito com sucesso.")
}

func (s *System) ExpediteOrder() {
	if len(s.Orders) == 0 {
		fmt.Println("Nenhum pedido pendente para expedir.")
		return
	}

	order := s.Orders[0]
	s.Orders = s.Orders[1:]
	s.FoodTruck.TotalRevenue += order.TotalAmount

	fmt.Printf("Pedido #%d expedido:\n", order.ID)
	fmt.Printf("Entrega: %t, Total: R$%.2f\n", order.Delivery, order.TotalAmount)
}

func (s *System) ShowMetrics() {
	fmt.Printf("Número total de produtos cadastrados: %d\n", s.FoodTruck.TotalProducts)
	fmt.Printf("Número de pedidos encerrados: %d\n", s.FoodTruck.TotalOrders)
	fmt.Printf("Faturamento total até o momento: R$%.2f\n", s.FoodTruck.TotalRevenue)
}
