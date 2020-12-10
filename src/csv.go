package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/Nhanderu/brdoc"
)

type row struct {
	CPF                  string
	Private              bool
	Incompleto           bool
	DataDaUltimaCompra   string
	TicketMedio          string
	TicketDaUltimaCompra string
	LojaMaisFrequente    string
	LojaDaUltimaCompra   string
}

func main() {
	data, err := parseData() //use
	if err != nil {
		log.Fatal(err)
	}

}

func parseData() ([]row, error) {
	rawRows, err := readFile("files/base_teste.txt")
	if err != nil {
		return []row{}, err
	}

	rows := []row{}

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
	}

	return rows, nil
}

func parseRow(rawRow map[string]string) (parsedRow row, err error) {

	if !brdoc.IsCPF(rawRow["CPF"]) {
		err = fmt.Errorf("invalid CPF: %s", rawRow["CPF"])
		return
	}

	if !brdoc.IsCNPJ(rawRow["LOJA MAIS FREQUÊNTE"]) {
		err = fmt.Errorf("invalid CNPJ: %s", rawRow["LOJA MAIS FREQUÊNTE"])
		return
	}

	if !brdoc.IsCNPJ(rawRow["LOJA DA ÚLTIMA COMPRA"]) {
		err = fmt.Errorf("invalid CNPJ: %s", rawRow["LOJA DA ÚLTIMA COMPRA"])
		return
	}

	parsedRow.CPF = rawRow["CPF"]
	parsedRow.LojaMaisFrequente = rawRow["LOJA MAIS FREQUÊNTE"]

	// boolean initial state
	parsedRow.Private = false
	parsedRow.Incompleto = false

	if rawRow["PRIVATE"] == "1" {
		parsedRow.Private = true
	}

	if rawRow["Incompleto"] == "1" {
		parsedRow.Private = true
	}

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
		// isolate character if exists whitespaces
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
		"CPF",
		"PRIVATE",
		"INCOMPLETO",
		"DATA DA ÚLTIMA COMPRA",
		"TICKET MÉDIO",
		"TICKET DA ÚLTIMA COMPRA",
		"LOJA MAIS FREQUÊNTE",
		"LOJA DA ÚLTIMA COMPRA",
	}
}
