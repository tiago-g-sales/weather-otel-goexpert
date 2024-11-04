# temp-cep
Sistema retorna  clima atual  baseado em um CEP informado

# Desafio GOLang Consulta Temperatura baseado em um CEP informado - FullCycle 

Aplicação em Go sendo: 
  - Servidor HTTP Rest

&nbsp;
- **Rodando em Cloud Google com CloudRun**
- **Aplicação em Container com Docker - Dockerfile e teste unitário**

## Funcionalidades

- **Consulta de Temperatura com retorno em Celsius, Kelvin e Fahrenheit**
  - O servidor permite consultar a temperatura informando um CEP.
  - Retorno esperado:
```
  {
	"temp_C": 22.3,
	"temp_F": 72.1,
	"temp_K": 295.3
  } 
``` 
  - Sendo temp_F = Fahrenheit
  - Sendo temp_C = Celsius
  - Sendo temp_K = Kelvin  

## Como Utilizar localmente:

1. **Requisitos:** 
   - Certifique-se de ter o Go instalado em sua máquina.
   - Certifique-se de ter o Docker instalado em sua máquina.
   - Foi atulizado a API viaCEP para encontrar a localização que deseja consultar a temperatura: https://viacep.com.br/
   - Foi utilizado a API WeatherAPI para consultar as temperaturas desejadas: https://www.weatherapi.com/


&nbsp;
2. **Clonar o Repositório:**
&nbsp;

```bash
git clone https://github.com/tiago-g-sales/temp-cep.git
```
&nbsp;
3. **Acesse a pasta do app:**
&nbsp;

```bash
cd temp-cep
```
&nbsp;
4. **Rode o docker para buildar a imagem gerando o container:**
&nbsp;

```bash 
 docker build -t nome_que_preferir/temp-cep:latest .
```

&nbsp;
4. **Rode o docker executar ocontainer:**
&nbsp;

```bash 
 docker run --rm -p 8080:8080 nome_que_preferir/temp-cep
```

5. **Acesse a pasta cmd/ e rode o main.go:**
&nbsp;

```bash 
cd cmd/
```

```bash 
go run main.go
```

Observação: Necessario informar a API KEY da plataforma de consulta de temperatura no arquivo config.env na raiz do projeto conforma abaixo:
WEB_SERVER_PORT=:8080
API_KEY=XXXXXXXXXXXXXXXXXXXXX


## Como testar localmente:

### Portas
HTTP server on port :8080 <br />

### HTTP
 - Execute o curl abaixo ou use um aplicação client REST para realizar a requisição.   
 - curl --request GET \
  --url 'http://localhost:8080/?cep=CEP_DESEJADO_8_NUMEROS' \
  --header 'User-Agent: insomnia/10.0.0'

## Como executar os teste unitários e o relatorio de cobertura de testes:
 - Coverage
 ```bash 
 - go test ./... -coverprofile=coverage.out
 ```
 
 ```bash 
 - go tool cover -html=coverage.out
 ```  

## Como executando a aplicação hospedada no Google Cloud (CloudRun):
 - Acesse a url abaixo no navegador ou outa aplicação client REST para realizar a requisição:
 - https://cloudrun-goexpert-challenge-temp-by-cep-cq6sddvtja-uc.a.run.app/?cep=CEP_DESEJADO_8_NUMEROS

Observação: Informar o CEP numerico 8 caracteres como "query parameter" 

