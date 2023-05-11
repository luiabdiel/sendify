# **Sendify** - Serviço de Envio de Emails em Lote

O Sendify é um serviço desenvolvido em **Golang**, utilizando o padrão **REST** e banco de dados **Postgres**, que tem como objetivo facilitar o envio de email em lote para os sistemas clientes, evitando a perda de tempo com o envio de email individualmente.

## **Instalação**

Antes de começar, é necessário ter o Go e o Postgres instalados em sua máquina.

* Clone o repositório para sua máquina
* Configure o arquivo .env com as informações do seu banco de dados Postgres
* Execute o seguinte comando para instalar as dependências do projeto:

```
go get ./...
```

* Execute o seguinte comando para iniciar o servidor:
```
go run main.go
```

## **Utilização**
O serviço possui endpoints que permitem o envio de emails em lote. A API suporta requisições HTTP POST e GET.

## **Envio de email**
Para enviar um email em lote, é necessário fazer uma requisição POST para o endpoint /emails com o seguinte payload:

```
{
  "subject": "Assunto do email",
  "body": "Corpo do email",
  "recipients": [
    "email1@example.com",
    "email2@example.com"
  ]
}

```

O servidor retornará um código 202 Accepted, indicando que o email foi aceito para envio. O envio em si é realizado de forma assíncrona em segundo plano.

## **Consulta de status de envio**
Para consultar o status de envio de um email em lote, é necessário fazer uma requisição GET para o endpoint /emails/{id}, informando o ID do email enviado.

O servidor retornará um JSON com informações sobre o status de envio do email.

## **Contribuição**
Contribuições são bem-vindas! Caso queira contribuir com o projeto, siga os seguintes passos:

* Faça um fork do repositório
* Faça um fork do repositório
* Envie um pull request com suas alterações

## **Licença**
Este projeto está licenciado sob a licença MIT. Veja o arquivo LICENSE para mais informações.