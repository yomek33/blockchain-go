package main

import (
	"log"

	"github.com/boltdb/bolt"
)
	const 
( dbFile = "blockchain.db"
blockBucket = "blocks"
dbOwnerReadWrite = 0600
)

type BlockChain struct {
	tip []byte
	db  *bolt.DB
}


func NewBlockChain() *BlockChain {
	var tip []byte
	db, err := bolt.Open(dbFile, dbOwnerReadWrite, nil)

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Println(err)
				return nil
			}
			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})
	if err != nil{
		log.Println(err)
		return nil
	}

	bc := BlockChain{tip, db}

	return &bc
}

func (bc *BlockChain) AddBlock(data string){
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		lastHash = b.Get([]byte("l"))
		
		return nil
	})
	if err != nil{
		log.Println(err)
	}
	newBlock := NewBlock(data, lastHash)
	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil{
			log.Println(err)
			return nil
		}
		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil{
			log.Println(err)
			return nil
		}
		bc.tip = newBlock.Hash

		return nil
	})

}