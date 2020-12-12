package challengeReport

import (
	"testing"
)

func TestChallenge(t *testing.T) {

	data, err := ParseData("files/base_teste.txt")
	if err != nil {
		t.Fatal(err)
	}

	return data

	//len(data)==len(retorno da query)
}
