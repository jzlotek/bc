package block

import (
	"bytes"
	"crypto/sha256"
	"io"
	"strconv"
)

type Block struct {
	ID    int    `json:"id"`
	Nonce int    `json:"nonce"`
	Data  []byte `json:"data"`
	PHash string `json:"previous_hash"`
}

func (b Block) Bytes() []byte {
	id := []byte(strconv.Itoa(b.ID))
	nonce := []byte(strconv.Itoa(b.Nonce))
	return bytes.Join([][]byte{
		[]byte(b.PHash),
		id,
		b.Data,
		nonce,
	}, []byte(""))
}

func (b Block) Hash() ([]byte, error) {
	hash := sha256.New()
	if _, err := io.Copy(hash, bytes.NewReader(b.Bytes())); err != nil {
		return []byte(""), err
	}
	sum := hash.Sum(nil)

	return sum, nil
}
