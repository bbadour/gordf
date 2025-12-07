package parser

import (
	"fmt"
	"sync/atomic"
)

type NODETYPE string

const (
	LITERAL         NODETYPE = "LITERAL"
	RESOURCELITERAL          = "RESOURCE"
	NODEIDLITERAL            = "NodeIDLiteral"
	BLANK                    = "BNODE"
	IRI                      = "IRI"
)

type Node struct {
	NodeType NODETYPE
	ID       string
}

func (node *Node) String() string {
	return fmt.Sprintf("(%v, %v)", node.NodeType, node.ID)
}

type BlankNodeGetter struct {
	lastid atomic.Int64
}

func (getter *BlankNodeGetter) set(i int) {
	getter.lastid.Store(int64(i))
}

func (getter *BlankNodeGetter) Get() Node {
	lastid := getter.lastid.Add(1)
	return Node{
		NodeType: BLANK,
		ID:       fmt.Sprintf("N%v", lastid),
	}
}

func (getter *BlankNodeGetter) GetFromId(id string) Node {
	return Node{
		NodeType: NODEIDLITERAL,
		ID:       fmt.Sprintf("N%v", id),
	}
}
