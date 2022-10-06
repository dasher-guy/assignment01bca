package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

var for_prv_hash [8]string
var for_verify [10]string

type block struct {
	block_id           int
	transaction_string string
	prev_hash          string
	hash_block         string
}

func CalculateHash(stringToHash string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))
}

func NewBlock(transaction string, nonce int, previousHash string) *block {
	var i int
	i = 0
	b := new(block)
	b.transaction_string = transaction
	b.block_id = nonce
	b.prev_hash = previousHash
	b.hash_block = CalculateHash(strconv.Itoa(nonce) + transaction + previousHash)
	if i < 4 {
		for_prv_hash[nonce] = b.hash_block
		i++
	}

	return b
}

type blockChain struct {
	list []*block
}

func (obj *blockChain) createBlock(nonce int, transaction string, previous_hash string) *block {

	block1 := NewBlock(transaction, nonce, previous_hash)
	obj.list = append(obj.list, block1)
	return block1
}

func ChangeBlock(nonce int, tem *blockChain) {

	var temp_string string
	fmt.Scanln(&temp_string)
	for j := 0; j < nonce; j++ {
		if tem.list[j].block_id == nonce-1 {
			tem.list[nonce-1].transaction_string = temp_string
		}
	}
}

func VerifyChain(obj *blockChain) bool {
	var checking = ""
	for i := 0; i < len(obj.list); i++ {
		checking = strconv.Itoa(obj.list[i].block_id) + obj.list[i].transaction_string + obj.list[i].prev_hash
		sum := CalculateHash(checking)

		if sum != obj.list[i].hash_block {
			fmt.Printf("Block is tempered, Block No. : %d\n", i+1)
			return false
		}
	}
	fmt.Printf("Blockchain is not tempered\n\n")
	return true
}
func listBlock(obj *blockChain, size int) {

	for i := 0; i < size; i++ {
		fmt.Println("Nonce: ", obj.list[i].block_id)
		fmt.Println("Transaction string: ", obj.list[i].transaction_string)
		fmt.Println("Previous Block Hash: ", obj.list[i].prev_hash)
		fmt.Println("Hash of Block: ", obj.list[i].hash_block)
		for_verify[i] = obj.list[i].hash_block
	}
}

func main() {

	for_prv_hash[0] = "XXXX"
	BCh := new(blockChain)

	BCh.createBlock(1, "ALICE TO CHARLIE", for_prv_hash[0])
	BCh.createBlock(2, "CHARLIE TO TRUDY", for_prv_hash[1])
	BCh.createBlock(3, "TRUDY TO MALORY", for_prv_hash[2])
	BCh.createBlock(4, "MALORY TO ANTHONY", for_prv_hash[3])


	fmt.Printf("%s Finch Winch Blockchain Before Change %s\n", strings.Repeat(":(", 25), strings.Repeat(":)", 25))
	listBlock(BCh, 4)
	VerifyChain(BCh)

	fmt.Printf("%s Finch Winch Blockchain After Change %s\n", strings.Repeat(":(", 25), strings.Repeat(":)", 25))
	var t int
	fmt.Scanln(&t)
	ChangeBlock(t, BCh)
	listBlock(BCh, 4)	
	VerifyChain(BCh)

				
			
	

}