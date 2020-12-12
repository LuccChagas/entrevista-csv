package challengeReport

import (
	"testing"
)

func TestChallenge(t *testing.T) {
	creatTableReport()

	_, err := ParseData("files/base_teste.txt")
	if err != nil {
		t.Fatal(err)
	}

}
