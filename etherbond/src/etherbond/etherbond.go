package etherbond

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
)

const (
	ETHERTYPE = 0x8888
)

const (
	ADV_CLU_FR = 0 + iota
	ADV_CH_FR
	ADV_IN_FR
	ADV_OUT_FR
	IO_TR_FR
	ACK_FR
)

const (
	ADV_CLU_CM = 0x01
	ADV_CH_CM  = 0x02
	ADV_IN_CM  = 0x03
	ADV_OUT_CM = 0x04
	IO_TR_CM   = 0x05
	ACK_CM     = 0xff
)

const (
	TRANSNEW = uint8(0) + iota
	TRANSDONE
)

// Config struct

type Config struct {
	Rsize            uint8
	ifi              *net.Interface
	Debug            bool
	Done             chan bool
	kill_sender      chan bool
	kill_receiver    chan bool
	kill_advertiser  chan bool
	frame_send_chan  chan string
	tag_chan         chan uint32
	transaction_chan chan Transaction
}

// Peers description

type Peer struct {
	PeerId   uint32
	Channels []uint32
	Inputs   []uint32
	Outputs  []uint32
}

type Cluster struct {
	ClusterId uint32
	Peers     []Peer
}

//

type peerlist map[int]int

// Peers status

type Peer_runinfo struct {
	HAddr    []byte
	Channels map[uint32]bool
	Inputs   map[uint32]bool
	Outputs  map[uint32]bool
}

type Cluster_runinfo struct {
	ClusterId uint32
	Peers     map[uint32]Peer_runinfo
	Quorate   bool
	Degraded  bool
}

//

type Macs struct {
	Assoc map[string]string
}

//

type ChOp interface {
}

type InOp interface {
	SetId(id uint32)
	GetId() uint32
}

type OutOp interface {
	SetId(id uint32)
	GetId() uint32
}

type Peer_op struct {
	PeerId   uint32
	Channels []ChOp
	Inputs   []InOp
	Outputs  []OutOp
}

func (c *Cluster) String() string {
	result, _ := json.MarshalIndent(&c, "", "\t")
	return string(result)
}

func (c *Cluster_runinfo) String() string {
	result := "ClusterId: " + strconv.Itoa(int(c.ClusterId)) + "\n"
	if c.Quorate {
		result += "Quorum: yes\n"
	} else {
		result += "Quorum: no\n"
	}
	if c.Degraded {
		result += "Degraded: yes\n"
	} else {
		result += "Degraded: no\n"
	}
	for pid, peer := range c.Peers {
		result += "\tPeer: " + strconv.Itoa(int(pid)) + "\n"
		result += fmt.Sprintf("%02x %02x %02x %02x %02x %02x", peer.HAddr[0], peer.HAddr[1], peer.HAddr[2], peer.HAddr[3], peer.HAddr[4], peer.HAddr[5])
	}
	return string(result)
}

type Transaction struct {
	Ttype uint8
	Tag   uint32
	Data  string
}


