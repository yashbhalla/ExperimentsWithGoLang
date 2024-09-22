package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Node struct {
	ID    int
	Data  map[string]interface{}
	Peers []*Node
	mutex sync.Mutex
}

func (n *Node) gossip() {
	for {
		if len(n.Peers) > 0 {
			peer := n.Peers[rand.Intn(len(n.Peers))]
			n.sendData(peer)
		}
		time.Sleep(time.Second)
	}
}

func (n *Node) sendData(peer *Node) {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	for key, value := range n.Data {
		peer.updateData(key, value)
	}
}

func (n *Node) updateData(key string, value interface{}) {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	n.Data[key] = value
	fmt.Printf("Node %d updated: %s = %v\n", n.ID, key, value)
}

func main() {
	numNodes := 5
	nodes := make([]*Node, numNodes)

	for i := 0; i < numNodes; i++ {
		nodes[i] = &Node{
			ID:    i,
			Data:  make(map[string]interface{}),
			Peers: make([]*Node, 0),
		}
	}

	for i := 0; i < numNodes; i++ {
		for j := 0; j < numNodes; j++ {
			if i != j {
				nodes[i].Peers = append(nodes[i].Peers, nodes[j])
			}
		}
	}

	for _, node := range nodes {
		go node.gossip()
	}

	nodes[0].updateData("temperature", 25.5)
	time.Sleep(time.Second * 2)
	nodes[2].updateData("humidity", 60)

	time.Sleep(time.Second * 10)

	fmt.Println("\nFinal state:")
	for _, node := range nodes {
		fmt.Printf("Node %d data: %v\n", node.ID, node.Data)
	}
}
