package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	rawRows, err := readFile("files/base_teste.txt")
	if err != nil {
		log.Fatal(err)
	}

	rows := []map[string]string{}

	for i := range rawRows[1:] {
		row, err := splitRow(rawRows[i])
		if err != nil {
			log.Printf("splitRow[%05d] err: %v", i, err)
			continue
		}
		rows = append(rows, row)
	}

	log.Println("len(rows) ", len(rows))
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
	for _, item := range strings.Split(row, "  ") { //using two spaces to split the string
		if len(item) == 0 {
			continue
		}
		item = strings.TrimPrefix(item, " ")
		item = strings.TrimSuffix(item, " ")
		splitted = append(splitted, item)
	}
	header := rowHeader()
	if len(splitted) != len(header) {
		return map[string]string{}, fmt.Errorf("invalid row: %s", row)
	}
	output := map[string]string{}

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
