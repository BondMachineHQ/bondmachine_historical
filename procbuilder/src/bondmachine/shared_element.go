package bondmachine

type Shared_element interface {
	Shr_get_name() string // The name
	Shr_get_desc() string // A description
	Shortname() string
	GV_config(uint8) string
	Instantiate(string) (Shared_instance, bool)
}

type Shared_instance_list []int

type Shared_instance interface {
	String() string
	Shortname() string
	Shr_get_name() string
	GV_config(uint8) string
	Write_verilog(*Bondmachine, int, string, string) string
	Get_wires_perproc(*Bondmachine, int, int, string) string
	Get_header_perproc(*Bondmachine, int, int, string) string
}
