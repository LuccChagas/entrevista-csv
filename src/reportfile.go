package challengeReport

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/Nhanderu/brdoc"
)

//Row structure for each line of the report
type Row struct {
	CPF                  string  `db:"cpf"`
	Private              bool    `db:"private"`
	Incompleto           bool    `db:"incompleto"`
	DataDaUltimaCompra   string  `db:"data_ultima_compra"`
	TicketMedio          float64 `db:"ticket_medio"`
	TicketDaUltimaCompra float64 `db:"ticket_ultima_compra"`
	LojaMaisFrequente    string  `db:"loja_mais_frequente"`
	LojaDaUltimaCompra   string  `db:"loja_ultima_compra"`
}

//ParseData calls all deals made in the file
func ParseData(path string) ([]Row, error) {
	rawRows, err := readFile(path)
	if err != nil {
		return []Row{}, err
	}

	rows := []Row{}

	db, err := NewClient()
	if err != nil {
		return []Row{}, err
	}

	for i := range rawRows[1:] {

		rawRow, err := splitRow(rawRows[i])
		if err != nil {
			log.Printf("splitRow[%05d] err: %v", i, err)
			continue
		}

		parsedRow, err := parseRow(rawRow)
		if err != nil {
			log.Printf("parseRow[%05d] err: %v", i, err)
			continue
		}
		rows = append(rows, parsedRow)

		db.InsertReport(parsedRow)
	}

	return rows, nil
}

//Performs data validation on each type of item
func parseRow(rawRow map[string]string) (parsedRow Row, err error) {

	if !brdoc.IsCPF(rawRow["CPF"]) {
		err = fmt.Errorf("invalid CPF: %s", rawRow["CPF"])
		return
	}
	parsedRow.CPF = strings.Replace(rawRow["CPF"], ".", "", -1)

	if !brdoc.IsCNPJ(rawRow["loja_mais_frequente"]) {
		err = fmt.Errorf("invalid CNPJ: %s", rawRow["loja_mais_frequente"])
		return
	}
	parsedRow.LojaMaisFrequente = strings.Replace(rawRow["loja_mais_frequente"], ".", "", -1)

	if !brdoc.IsCNPJ(rawRow["loja_da_ultima_compra"]) {
		err = fmt.Errorf("invalid CNPJ: %s", rawRow["loja_da_ultima_compra"])
		return
	}
	parsedRow.LojaDaUltimaCompra = strings.Replace(rawRow["loja_mais_frequente"], ".", "", -1)

	//Boolean initial state
	parsedRow.Private = false
	parsedRow.Incompleto = false

	if rawRow["private"] == "1" {
		parsedRow.Private = true
	}

	if rawRow["incompleto"] == "1" {
		parsedRow.Private = true
	}

	var parsedFloat float64
	parsedFloat, err = strconv.ParseFloat(rawRow["ticket_medio"], 64)
	if err != nil {
		return
	}
	parsedRow.TicketMedio = parsedFloat

	parsedFloat, err = strconv.ParseFloat(rawRow["ticket_ultima_compra"], 64)
	if err != nil {
		return
	}
	parsedRow.TicketDaUltimaCompra = parsedFloat

	return
}

func readFile(path string) ([]string, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return []string{}, fmt.Errorf("ioutil.ReadFile err: %v", err)
	}
	return strings.Split(string(file), "\n"), nil
}

func splitRow(row string) (map[string]string, error) {

	splitted := []string{}
	header := rowHeader()
	output := map[string]string{}

	for _, item := range strings.Split(row, "  ") { //using two spaces to split the string
		if len(item) == 0 {
			continue
		}
		// isolating characters if there are blanks before or after
		item = strings.TrimPrefix(item, " ")
		item = strings.TrimSuffix(item, " ")

		splitted = append(splitted, item)
	}

	if len(splitted) != len(header) {
		return map[string]string{}, fmt.Errorf("invalid row: %s", row)
	}

	for i := range splitted {
		output[header[i]] = splitted[i]
	}
	return output, nil
}

func rowHeader() []string {
	return []string{
		"cpf",
		"private",
		"incompleto",
		"data_da_ultima_compra",
		"ticket_medio",
		"ticket_da_ultima_compra",
		"loja_mais_frequente",
		"loja_da_ultima_compra",
	}
}
