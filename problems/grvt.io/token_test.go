package problems

import (
	"encoding/csv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIngest(t *testing.T) {
	tests := []struct {
		arg string
	}{
		{arg: "1,09:50,GOOG,sell,42.24,100"},
		{arg: "2,10:50,\"GOOG, INC\",sell,42.24,100"},
		{arg: "2,10:50,\"GOOGLE , asdsa ,INC\",sell,42.24,100"},
		{arg: "2,10:50,\"GOOGLE,INC\",\"SELL,ONLY\",42.24,100"},
		{arg: "3,11:50,\"GOOG, \"\"ALPHA, INC\",sell,42.24,100"},
		{arg: "3,11:50,\"\",\"\",sell,42.24,100"},
		{arg: "3,11:50,\"\"\",\"\"\",sell,42.24,100"},
		{arg: "3,11:50,\"a\"\",\"\"b\",sell,42.24,100"},
		{arg: "3,11:50,\"GOOG, \"\"\"\"ALPHA, INC\",sell,\"42.24,\"\"\",100"},
	}

	for _, tt := range tests {
		r := csv.NewReader(strings.NewReader(tt.arg))
		expected, err := r.Read()
		if err != nil {
			t.Errorf("Error not expected")
		}
		assert.Equal(t, expected, ingest2(tt.arg))
	}
}
