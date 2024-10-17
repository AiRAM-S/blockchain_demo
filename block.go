package main

import (
	"bytes"
	"time"
	"log"
	"encoding/gob"
)

type Block struct {
	Time		int64
	Data		[]byte
	PrevHash 	[]byte
	Hash		[]byte
	Nonce		int
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevHash, []byte{}, 0}
	block.SetHash()
	return block
}
	
func (b *Block) SetHash() {
	//为Block生成hash，修改为工作量证明算法
	pow := NewProofOfWork(b)
	nonce, hash := pow.Run()
	b.Hash = hash
	b.Nonce = nonce
	return
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}