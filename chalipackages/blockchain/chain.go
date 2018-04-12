package blockchain

import (
	"github.com/DesmondANIMUS/chainchali/chalipackages/chalihelp"
	"github.com/DesmondANIMUS/chainchali/chalipackages/chalimodel"
)

type Block struct {
	Hash          string
	PrevBlockHash string
	Data          string
}

type Blockchain struct {
	Blocks []*Block
}

func NewBlock(data, prevBlockHash string) *Block {
	block := &Block{
		Data:          data,
		PrevBlockHash: prevBlockHash,
		Hash:          chalihelp.GetHash(prevBlockHash + data + chalimodel.Random),
	}

	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Gensis Block", "")
}

func NewBlockChain() *Blockchain {
	return &Blockchain{
		[]*Block{
			NewGenesisBlock(),
		},
	}
}

func (bc *Blockchain) AddBlock(data string) *Block {
	prevBlockHash := bc.Blocks[len(bc.Blocks)-1].Hash
	nayaBlock := NewBlock(data, prevBlockHash)
	bc.Blocks = append(bc.Blocks, nayaBlock)

	return nayaBlock
}
