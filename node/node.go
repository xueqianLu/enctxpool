package node

import "github.com/xueqianLu/enctxpool/core/mempool"

type Node struct {
	txpool *mempool.TxPool
}

func NewNode() *Node {
	n := new(Node)
	n.txpool = mempool.NewTxPool(mempool.DefaultTxPoolConfig, nil)

	return n
}

func (n *Node) TxPool() *mempool.TxPool {
	return n.txpool
}