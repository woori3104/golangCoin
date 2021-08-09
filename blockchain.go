package main

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

// block의 struct
// 대문자로 변수 선언하면 외부파일에서 접근 가능 (Public)
type Block struct {
	Data     string // public Data
	Hash     string // public Hash
	PrevHash string // public PrevHash
}

type blockchain struct {
	blocks []*Block
}

// singleton pattern
var b *blockchain

// 병렬적으로 실행되더라도 단한개만 동작하게 하는 패키지 함수
var once sync.Once

func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	// 16진수로 formmat
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].Hash
}

// string을 ㅏey로서 block을 생성해서 return 하는 함수
func createBlock(data string) *Block {
	newBlock := Block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

// 새로운 블럭을 추가
func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

// blockchain변수 b의 instance를 직접 공유하지않고 b에 접근할 수 있도록 하는 함수
func GetBlockchain() *blockchain {
	// b가 없다면 기존의 값을 업데이트
	if b == nil {
		// 한번만 실행되는 함수
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis")
		})
	}
	// 만약 값이 있다면 새로 생성하지않고 기존 b를 공유
	// singleton pattern
	return b
}

func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}
