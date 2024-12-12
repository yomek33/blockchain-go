package main

type BlockChain struct {
	blocks []*Block
}

func (bs *BlockChain) AddBlock(data string) {
	prevBlock := bs.blocks[len(bs.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bs.blocks = append(bs.blocks, newBlock)
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}
