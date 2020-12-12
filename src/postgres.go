package challengeReport

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	*sqlx.DB
}

func NewClient() (Database, error) {
	// tenta a conexão e retorna diretamente um error em caso de falha
	db := sqlx.MustConnect("postgres", "user=postgres dbname=postgres sslmode=disable")

	schema := `CREATE TABLE IF NOT EXISTS relatorios (
		cpf VARCHAR(14) PRIMARY KEY,
		private BOOLEAN,
		incompleto BOOLEAN,
		data_ultima_compra DATE,
		ticket_medio DECIMAL(10,2),
		ticket_ultima_compra DECIMAL(10,2),
		loja_mais_frequente VARCHAR(14),
		loja_ultima_compra VARCHAR(14)
	);`

	// creates the relatorios table
	output, err := db.Exec(schema)
	if err != nil {
		return Database{}, fmt.Errorf("db.Exec() create table: %v", err)
	}
	log.Println("return: ", output)

	output, err = db.Exec("TRUNCATE TABLE relatorios")
	if err != nil {
		return Database{}, fmt.Errorf("db.Exec() truncate table: %v", err)
	}
	log.Println("return: ", output)

	return Database{db}, nil
}

//InsertReport
func (db *Database) InsertReport(r Row) {
	q := `INSERT INTO relatorios VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	db.MustExec(
		q,
		r.CPF,
		r.Private,
		r.Incompleto,
		r.DataDaUltimaCompra,
		r.TicketMedio,
		r.TicketDaUltimaCompra,
		r.LojaMaisFrequente,
		r.LojaDaUltimaCompra,
	)
}
