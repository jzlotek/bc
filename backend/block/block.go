package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"strconv"
)

type Block struct {
	ID    int    `json:"id"`
	Nonce uint   `json:"nonce"`
	Data  string `json:"data"`
	PHash string `json:"previous_hash"`
}

func (b Block) Bytes() []byte {
	id := []byte(strconv.Itoa(b.ID))
	nonce := []byte(strconv.Itoa(int(b.Nonce)))
	return bytes.Join([][]byte{
		[]byte(b.PHash),
		id,
		nonce,
		[]byte(b.Data),
	}, []byte(""))
}

func (b Block) Hash() ([]byte, error) {
	if len(b.PHash) != 64 {
		errstr := "need to supply previous hash length of 64..."
		err := errors.New(errstr)
		return []byte(""), err
	}
	hash := sha256.New()
	if _, err := io.Copy(hash, bytes.NewReader(b.Bytes())); err != nil {
		return []byte(""), err
	}
	sum := hash.Sum(nil)

	return []byte(hex.EncodeToString(sum)), nil
}
