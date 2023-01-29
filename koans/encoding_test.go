package koans

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Encoding struct {
	Type string
}

func (k *KoansTest) EncodingTest(t *testing.T) {
	rows, err := k.db.Query(PragmaEncoding)
	assert.Nil(t, err)
	defer rows.Close()
	n := 0
	for rows.Next() {
		encoding := Encoding{}
		err := rows.Scan(&encoding.Type)
		assert.Nil(t, err)
		assert.Equal(t, UTF8, encoding.Type)
		n = n + 1
	}
	assert.True(t, n > 0)
}
