package mempool

import (
	"sync"

	blockchain "github.com/workspace/evoting/ev-blockchain-protocol/core"
)

// =============================
// memory pool for transactions
// =============================
type memPool struct {
	mtx   *sync.RWMutex
	store map[string]blockchain.Transaction
}

var txPerBlock = 10

func NewMemoryPool() *memPool {
	return &memPool{
		mtx:   new(sync.RWMutex),
		store: make(map[string]blockchain.Transaction),
	}
}

func (p *memPool) Add(tx blockchain.Transaction) {
	p.mtx.Lock()

	h := string(tx.Hash()[:])
	if _, ok := p.store[h]; !ok {
		p.store[h] = tx
	}

	p.mtx.Unlock()
}

func (p *memPool) Get(h string) (tx blockchain.Transaction) {
	p.mtx.RLock()
	tx = p.store[h]
	p.mtx.RUnlock()

	return
}

func (p *memPool) Delete(h string) {
	p.mtx.Lock()
	delete(p.store, h)
	p.mtx.Unlock()
}

func (p *memPool) GetVerified() (txs []blockchain.Transaction) {
	n := txPerBlock
	if n == 0 {
		return
	}

	txs = make([]blockchain.Transaction, 0, n)
	for _, tx := range p.store {
		txs = append(txs, tx)

		if n--; n == 0 {
			return
		}
	}

	return
}
