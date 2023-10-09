# AP1 Estrutura de dados 
Projeto AP1 - Carolina Cruz, Juliana Hall e Maria Mello

Requisitos
-

Identificando essas falhas no processo, sua equipe foi contratada para elaborar um MVP para uma solução de pedidos eletrônicos para o McRonald’s. O sistema deve rodar através de uma CLI, que deve ter as seguintes opções:

Cadastrar produto: deve ser possível armazenar até 50 produtos diferentes ativos. Cada produto deve ser armazenado com um identificador único, representado por um número inteiro, começando no 1. Caso um produto seja removido, o seu identificador não deve ser substituído. Ou seja, se o usuário cadastrou 20 produtos (1, 2, … 20), e excluiu o produto 20, o próximo produto a ser cadastrado terá o identificador 21. Cada produto deve conter as seguintes informações (além do identificador):
Nome: nome curto;
Descrição: descrição completa do item;
Valor: o preço do produto, em reais;
Cadastro de produtos em lote: os sócios gostariam de, com um único comando na CLI, poder passar os dados de vários produtos ao mesmo tempo, para que o sistema armazenasse todas as informações sem precisar voltar para o menu inicial múltiplas vezes;
Remover produto: o sistema permite excluir um produto, inserindo o seu identificador único;
Busca dos produtos: o sistema precisa permitir a busca dos produtos através do seu identificador, exibindo todas as informações na tela;
Exibir produtos: exibe todos os produtos cadastrados no sistema.
Adicionar pedido: o programa deve suportar até 1000 pedidos enquanto estiver sendo executado. O pedido deve incluir os seguintes campos:
Identificador: número único, iniciado com 1;
Delivery: campo booleano, valendo false caso o pedido seja para consumo ou retirada no local, e true caso seja para entrega;
Produtos: lista com até 10 produtos disponíveis no cardápio. Cada produto pode ter uma quantidade de itens diferente;
Valor total: valor do pedido, incluindo os produtos e taxa de entrega (se houver);
Expedir pedido: o programa libera o primeiro pedido cadastrado que esteja ainda aberto, exibindo as suas informações na tela, e adiciona o valor total do pedido ao faturamento do sistema;
Exibir métricas do sistema: o sistema exibe na tela o número total de produtos cadastrados, o número de pedidos já encerrados, o número de pedidos em andamento e o faturamento total até o momento.
Há alguns requisitos que os sócios do McRonald’s gostariam muito que fossem implementados já no MVP, mas que não estão no escopo do serviço:

Leitura dos produtos em arquivo de texto: caso os produtos possam ser lidos a partir de um arquivo .csv no momento em que o programa é iniciado, isso agilizaria muito o processo de testes e validação do sistema;
Exibir os pedidos ainda em aberto: exibe na tela todos os pedidos que estão em andamento, na ordem em que foram abertos;
Buscar os produtos cadastrados pelo nome: lista todos os produtos cujo nome inicie com o texto buscado pelo usuário;
Registro de data e hora dos pedidos: o sistema deve incluir, em cada pedido, a data e a hora em que o pedido foi realizado, bem como apresentar na tela a data e a hora no momento em que o pedido for expedido;
Registro do tempo médio de expedição dos pedidos: além das métricas já apresentadas, os sócios gostariam de acompanhar o tempo (em minutos) que cada pedido leva, do momento em que ele é registrado ao momento em que ele é expedido. A funcionalidade de exibir métricas do sistema deve incluir um tempo médio, que é a soma dos tempos de cada pedido expedido, dividido pelo número de pedidos já finalizados.
