package block

import (
	"testing"
)

func TestHash(t *testing.T) {
	block := Block{
		ID:    0,
		Nonce: 12345,
		PHash: "0000000000000000000000000000000000000000000000000000000000000000",
	}

	_, err := block.Hash()

	if err != nil {
		t.Error(err)
	}
}

func TestHashFail(t *testing.T) {
	block := Block{
		ID:    0,
		Nonce: 12345,
		PHash: "000000000000000000000000000000000000000000000000000000000000000",
	}

	_, err := block.Hash()

	if err == nil {
		t.Error(err)
	}
}

func TestHashData(t *testing.T) {
	block := Block{
		ID:    0,
		Nonce: 12345,
		Data:  "muh data",
		PHash: "0000000000000000000000000000000000000000000000000000000000000000",
	}

	_, err := block.Hash()

	if err != nil {
		t.Error(err)
	}
}

func TestHashPHash(t *testing.T) {
	block := Block{
		ID:    0,
		Nonce: 12345,
		Data:  "muh data",
		PHash: "0000000000000000000000000000000000000000000000000000000000000000",
	}

	hash1, err := block.Hash()

	if err != nil {
		t.Error(err)
	}

	block = Block{
		ID:    0,
		Nonce: 12345,
		Data:  "muh data",
		PHash: string(hash1),
	}

	hash2, err := block.Hash()

	if err != nil {
		t.Error(err)
	}

	if string(hash1) == string(hash2) {
		t.Errorf("Hash 1 = hash 2\n(%s == %s)\n", string(hash1), string(hash2))
	}
}
