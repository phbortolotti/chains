package main

// Blockchain mantém uma sequencia de blocos
type Blockchain struct {
	blocks []*Block
}

// AddBlock adiciona um novo bloco
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain cria um novo blockchain com o cloco genesis
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
