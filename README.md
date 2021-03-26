# Imput unstructured File (Without Tab)

Entrega de desafio proposto em processo seletivo NeoWay

## Setup

Para execução do projeto será necessario á instalação dos seguintes itens:

- Golang (https://golang.org/doc/install)

- Docker (https://docs.docker.com/get-docker/)

Após a instalação do Go & Docker será necessário realizar o pull da imagem referente ao postgres

```bash
docker pull postgres
```

Tendo as devidas instalações concluidas, basta executar os comandos **make** disponibilizados para facilitar a validação no processo seletivo. (Apontamentos realizados no arquivo **Makefile**)

- Navegue até a pasta **/src** e execute os comandos:
```bash
# Executa o servidor postgres no Docker redirecionando com suas portas locais
make postgres

# Executa o adminer localmente para que seja possivel visualizar os registros da tabela
make adminer

# Executa o programa solicitado no desáfio do processo seletivo
make release
```
## Adminer

- O item não havia sido solicitado, porém achei interessante incluir para facilitar a visualização dos dados no **Postgres**
- O **Adminer** fica disponível em: http://localhost:8080/ durante a execução 

Credenciais de Acesso:

```
	host     = "localhost"
	user     = "postgres"
	password = "secret"
	dbname   = "postgres"
```

## Observações

> Desde já, gostaria de agradecer a oportunidade de particiar do processo seletivo!

