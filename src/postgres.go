package challengeReport

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "secret"
	dbname   = "postgres"
)

func NewPostgres() (db *sql.DB) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db

}

func creatTableReport() {
	db := NewPostgres()

	schema := `CREATE TABLE IF NOT EXISTS relatorios (
		id SERIAL PRIMARY KEY,
		cpf VARCHAR(14),
		private char(1),
		incompleto char(1),
		data_ultima_compra VARCHAR(14),
		ticket_medio varchar(10),
		ticket_ultima_compra varchar(10),
		loja_mais_frequente VARCHAR(14),
		loja_ultima_compra VARCHAR(14)
	);`

	_, err := db.Exec(schema)

	if err != nil {
		fmt.Errorf("db.Exec() create table: %v", err)
	}

	fmt.Println("Table created with success!")

	_, err = db.Exec("TRUNCATE TABLE relatorios")
	if err != nil {
		fmt.Errorf("db.Exec() truncate table: %v", err)
	}

	fmt.Println("Table truncate Done!")
	defer db.Close()

}

func InsertReport(r Row, db *sql.DB) {

	q := `INSERT INTO relatorios VALUES (DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := db.Exec(
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
	if err != nil {
		panic(err)
	}
	defer fmt.Printf("row inserted! with success, cpf: %v \n", r.CPF)

}
