package bondmachine

import (
	"encoding/hex"
	"strconv"
	"strings"
	"udpbond"
)

type NetParameters map[string]string

type Udpbond_extra struct {
	Config    *udpbond.Config
	Cluster   *udpbond.Cluster
	Ips       *udpbond.Ips
	PeerID    uint32
	Maps      *IOmap
	Flavor    string
	Ip        string
	Broadcast string
	Netmask   string
	Port      string
	NetParams *NetParameters
}

type FirmwareCommand struct {
	PrimaryState   string // SM state for the command
	SecondaryState string // state for eventually internal State machines
	Command        string // ASCII command
	Description    string
	VHDLRapp       string // VHDL hex
	Signal         string // Eventually associated signal
	Starting       int    // Position of start within the memory
	OmitReturn     bool
	Payload        []string // Payload to transmit
	Payload_relpos []int    // Pauload relative position
}

func Ascii2Hex(in string) string {
	encoded := ""
	remaining_string := in

	for remaining_string != "" {

		lidx := strings.Index(remaining_string, "<<<")

		if lidx == -1 {
			tenc := []byte(remaining_string)
			encoded += hex.EncodeToString(tenc)
			remaining_string = ""
		} else {
			if lidx != 0 {
				tenc := []byte(remaining_string[0:lidx])
				encoded += hex.EncodeToString(tenc)
			}

			ridx := strings.Index(remaining_string, ">>>")
			if lidx != -1 {
				encoded += remaining_string[lidx+3 : ridx]
				remaining_string = remaining_string[ridx+3:]
			}
		}
	}
	return encoded
}

func CompleteCommands(fc []FirmwareCommand) ([]FirmwareCommand, int) {
	result := make([]FirmwareCommand, len(fc))
	poscounter := 0

	statelen := make(map[string]int)

	for i, com := range fc {
		result[i].PrimaryState = com.PrimaryState
		result[i].SecondaryState = com.SecondaryState
		result[i].Command = com.Command
		result[i].Description = com.Description
		result[i].Signal = com.Signal
		result[i].OmitReturn = com.OmitReturn
		result[i].Starting = poscounter
		result[i].Payload = append([]string(nil), com.Payload...)
		result[i].Payload_relpos = append([]int(nil), com.Payload_relpos...)

		result[i].VHDLRapp = ""

		encoded := ""
		remaining_string := com.Command

		for remaining_string != "" {

			lidx := strings.Index(remaining_string, "<<<")

			if lidx == -1 {
				tenc := []byte(remaining_string)
				encoded += hex.EncodeToString(tenc)
				remaining_string = ""
			} else {
				if lidx != 0 {
					tenc := []byte(remaining_string[0:lidx])
					encoded += hex.EncodeToString(tenc)
				}

				ridx := strings.Index(remaining_string, ">>>")
				if lidx != -1 {
					encoded += remaining_string[lidx+3 : ridx]
					remaining_string = remaining_string[ridx+3:]
				}
			}
		}

		if !com.OmitReturn {
			encoded += "0d0a"
		}

		encoded += "ff"

		for j := 0; j < len(encoded); j = j + 2 {
			result[i].VHDLRapp += "x\"" + encoded[j:j+2] + "\","
		}

		poscounter += len(encoded) / 2

		pmstate := com.PrimaryState
		if sl, ok := statelen[pmstate]; ok {
			statelen[pmstate] = sl + 1
		} else {
			statelen[pmstate] = 1
		}

	}

	lencounter := make(map[string]int)

	for i, com := range fc {

		pmstate := com.PrimaryState
		nbits := Needed_bits(statelen[pmstate])

		if st, ok := lencounter[pmstate]; !ok {
			result[i].SecondaryState = zeros_prefix(nbits, get_binary(0))
			lencounter[pmstate] = 1
		} else {
			result[i].SecondaryState = zeros_prefix(nbits, get_binary(st))
			lencounter[pmstate] = st + 1
		}
	}
	return result, poscounter
}

func LocateCommand(fc []FirmwareCommand, ps string, ss string) int {
	for _, com := range fc {
		if com.PrimaryState == ps {
			if com.SecondaryState == ss {
				return com.Starting
			}
		}
	}
	return -1
}

func LocateCommandbyIndex(fc []FirmwareCommand, ps string, index int) (FirmwareCommand, bool) {
	i := 0
	for _, com := range fc {
		if com.PrimaryState == ps {
			if i == index {
				return com, true
			}
			i++
		}
	}
	return FirmwareCommand{}, false
}

func (sl *Udpbond_extra) Get_Name() string {
	return "udpbond"
}

func (sl *Udpbond_extra) Get_Params() *ExtraParams {
	result := new(ExtraParams)
	result.Params = make(map[string]string)
	result.Params["peer_id"] = strconv.Itoa(int(sl.PeerID))
	result.Params["cluster_id"] = strconv.Itoa(int(sl.Cluster.ClusterId))
	result.Params["ip"] = sl.Ip
	result.Params["broadcast"] = sl.Broadcast
	result.Params["port"] = sl.Port

	netparams := *sl.NetParams

	if ssid, ok := netparams["ssid"]; ok {
		result.Params["ssid"] = ssid
	}
	if netmask, ok := netparams["netmask"]; ok {
		result.Params["netmask"] = netmask
	}
	if gateway, ok := netparams["gateway"]; ok {
		result.Params["gateway"] = gateway
	}

	var mypeer udpbond.Peer

	for _, peer := range sl.Cluster.Peers {
		if peer.PeerId == sl.PeerID {
			mypeer = peer
		}

		if sl.Ips != nil {
			peerstr := strconv.Itoa(int(peer.PeerId))
			if ipaddr, ok := sl.Ips.Assoc["peer_"+peerstr]; ok {
				if ipaddr == "auto" {
					result.Params["peer_"+peerstr+"_ip"] = "auto"
				} else if ipaddr == "adv" {
					result.Params["peer_"+peerstr+"_ip"] = "auto"
				} else {
					result.Params["peer_"+peerstr+"_ip"] = ipaddr
				}
			} else {
				result.Params["peer_"+peerstr+"_ip"] = "auto"
			}
		}

	}

	result.Params["input_ids"] = ""
	result.Params["inputs"] = ""
	result.Params["sources"] = ""

	for _, inp := range mypeer.Inputs {
		for iname, ival := range sl.Maps.Assoc {
			if iname[0] == 'i' && ival == strconv.Itoa(int(inp)) {
				result.Params["input_ids"] += "," + ival
				result.Params["inputs"] += "," + iname

				ressource := ""
				for _, opeer := range sl.Cluster.Peers {
					for _, oout := range opeer.Outputs {
						if strconv.Itoa(int(oout)) == ival {
							ressource = strconv.Itoa(int(opeer.PeerId))
							break
						}
					}
				}
				if ressource != "" {
					result.Params["sources"] += "," + ressource
				}

			}
		}
	}

	if result.Params["input_ids"] != "" {
		result.Params["input_ids"] = result.Params["input_ids"][1:len(result.Params["input_ids"])]
		result.Params["inputs"] = result.Params["inputs"][1:len(result.Params["inputs"])]
		result.Params["sources"] = result.Params["sources"][1:len(result.Params["sources"])]
	}

	result.Params["output_ids"] = ""
	result.Params["outputs"] = ""
	// Comma separated and - separated list of peer ids
	result.Params["destinations"] = ""

	for _, outp := range mypeer.Outputs {
		for oname, oval := range sl.Maps.Assoc {
			if oname[0] == 'o' && oval == strconv.Itoa(int(outp)) {
				result.Params["output_ids"] += "," + oval
				result.Params["outputs"] += "," + oname

				resdest := ""
				for _, ipeer := range sl.Cluster.Peers {
					for _, iin := range ipeer.Inputs {
						//fmt.Println(ipeer.PeerId, iin, oval, strconv.Itoa(int(iin)))
						if strconv.Itoa(int(iin)) == oval {
							resdest += "-" + strconv.Itoa(int(ipeer.PeerId))
						}
					}
				}
				//fmt.Println("resdest", resdest)
				if resdest != "" {
					result.Params["destinations"] += "," + resdest[1:len(resdest)]
				}

			}
		}
	}

	if result.Params["output_ids"] != "" {
		result.Params["output_ids"] = result.Params["output_ids"][1:len(result.Params["output_ids"])]
		result.Params["outputs"] = result.Params["outputs"][1:len(result.Params["outputs"])]
		result.Params["destinations"] = result.Params["destinations"][1:len(result.Params["destinations"])]
	}

	return result
}

func (sl *Udpbond_extra) Import(inp string) error {
	return nil
}

func (sl *Udpbond_extra) Export() string {
	return ""
}

func (sl *Udpbond_extra) Check(bmach *Bondmachine) error {
	return nil
}

func (sl *Udpbond_extra) Verilog_headers() string {
	result := "\n"
	return result
}
func (sl *Udpbond_extra) Static_verilog() string {

	result := "\n"
	return result
}

func (sl *Udpbond_extra) ExtraFiles() ([]string, []string) {
	return []string{}, []string{}
}
