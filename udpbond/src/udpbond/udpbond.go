package udpbond

import ()

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
	Rsize uint8
	//	ifi              *net.Interface
	Debug           bool
	Done            chan bool
	kill_sender     chan bool
	kill_receiver   chan bool
	kill_advertiser chan bool
	frame_send_chan chan string
	tag_chan        chan uint32
	//	transaction_chan chan Transaction
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

type Ips struct {
	Assoc map[string]string
}
