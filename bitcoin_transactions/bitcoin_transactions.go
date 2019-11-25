package main

import (
	"log"
	"runtime"
	"time"
)

type Transaction struct {
	TxID      string
	Address   string
	InputTxID string
	Value     uint
}
type TestData struct {
	Txs []Transaction
}

// Analyze a block of transactions from a Bitcoin family blockchain and, given 2 addresses,
// find all transaction chains that lead from our address to the banned address.
// Implement function findChains that returns an array fo chains:
// linked list of TxIDs in as array where every element is the TxID of the tx originating from the previous,
// starting from transaction that sends funds to our address.
// There is only one transaction that sends funds to our address.
// There are no transactions from the banned address.
func main() {
	startTime := time.Now()
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)

	testData := [1]TestData{}

	testData[0] = TestData{Txs: []Transaction{
		Transaction{TxID: "tx1", Address: "bgl", Value: 10, InputTxID: "none"},

		Transaction{TxID: "tx2", Address: "address1", Value: 1, InputTxID: "tx1"},
		Transaction{TxID: "tx3", Address: "address2", Value: 2, InputTxID: "tx1"},
		Transaction{TxID: "tx4", Address: "address3", Value: 3, InputTxID: "tx1"},
		Transaction{TxID: "tx5", Address: "address4", Value: 4, InputTxID: "tx1"},

		Transaction{TxID: "tx6", Address: "address5", Value: 1, InputTxID: "tx2"},
		Transaction{TxID: "tx7", Address: "address6", Value: 1, InputTxID: "tx3"},
		Transaction{TxID: "tx8", Address: "bgl", Value: 1, InputTxID: "tx3"},
		Transaction{TxID: "tx9", Address: "address6", Value: 1, InputTxID: "tx4"},
		Transaction{TxID: "tx10", Address: "address7", Value: 2, InputTxID: "tx5"},
		Transaction{TxID: "tx11", Address: "banned", Value: 2, InputTxID: "tx100"},
		Transaction{TxID: "tx12", Address: "banned", Value: 2, InputTxID: "tx5"},
		Transaction{TxID: "tx13", Address: "banned", Value: 1, InputTxID: "tx7"},
	}}

	for _, td := range testData {
		txsChain := findChains("bgl", "banned", td.Txs)
		log.Println(txsChain)
	}
	endTime := time.Now()
	log.Println("Time:   ", endTime.Sub(startTime))

	runtime.ReadMemStats(&m2)
	log.Println("Memory: ", (m2.TotalAlloc-m1.TotalAlloc)/1024)
}

func findChains(beginAddress string, endAddress string, txs []Transaction) [][]string {

	chain := map[string]bool{}
	ans := [][]string{}
	var startTx Transaction
	for _, tx := range txs {
		if tx.Address == beginAddress {
			startTx = tx
			break
		}
	}
	traverse(endAddress, startTx, chain, txs, []string{}, &ans)
	return ans
}

func traverse(endAddress string, startTx Transaction, visited map[string]bool, txs []Transaction, chain []string, ans *[][]string) {
	visited[startTx.TxID] = true
	subChain := make([]string, len(chain)+1)
	copy(subChain, append(chain, startTx.TxID))
	log.Println(subChain)
	if startTx.Address == endAddress {
		*ans = append(*ans, subChain)
	} else {

		for _, tx := range txs {
			if !visited[tx.TxID] && tx.InputTxID == startTx.TxID {
				traverse(endAddress, tx, visited, txs, subChain, ans)
			}
		}
	}
}
