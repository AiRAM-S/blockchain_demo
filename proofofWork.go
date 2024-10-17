package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const targetBits = 16 // difficulty
const maxNonce = 1000000

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits)) // left move

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevHash,
			pow.block.Data,
			IntToHex(pow.block.Time),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\" \n", pow.block.Data)

	/* 实现 Hashcash 算法：对 nonce 从 0 开始进行遍历，计算每一次哈希是否满足
	条件
		可能会用到的包及函数：big.Int.Cmp(),big.Int.SetBytes()
	*/
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			// 比target小，说明零的个数必然大于target中前置零的个数
			// 因为target是00...01
			break
		} else {
			nonce++
		}
	}

	fmt.Printf("\r%x", hash)
	fmt.Print("\n\n")

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	var isValid bool
	data := pow.prepareData(pow.block.Nonce)

	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid = (bytes.Equal(hash[:],pow.block.Hash)) && (hashInt.Cmp(pow.target) == -1)
	return isValid
}
