package bondmachine

import (
	"errors"
	"fmt"
	"os"
	"procbuilder"
	"regexp"
	"simbox"
	"strconv"
	"strings"
)

func nth_assoc(assoc string, seq int) string {
	re := regexp.MustCompile("\\[(?P<to>[0-9]+):(?P<from>[0-9]+)\\] +(?P<name>[a-zA-Z0-9]+)")
	if re.MatchString(assoc) {
		to_string := re.ReplaceAllString(assoc, "${to}")
		from_string := re.ReplaceAllString(assoc, "${from}")
		name_string := re.ReplaceAllString(assoc, "${name}")
		to, _ := strconv.Atoi(to_string)
		from, _ := strconv.Atoi(from_string)
		step := 1
		if from > to {
			step = -1
		}
		pos := from + seq*step

		return name_string + "[" + strconv.Itoa(pos) + "]"
	}
	return assoc
}

func (bmach *Bondmachine) Write_verilog(conf *Config, flavor string, iomaps *IOmap, extramods []ExtraModule, sbox *simbox.Simbox) error {
	if len(bmach.Domains) != 0 {

		pconf := conf.ProcbuilderConfig()

		//Instatiation of the Processor
		for i, dom_id := range bmach.Processors {

			ri := new(procbuilder.RuntimeInfo)
			ri.Init()

			pconf.Runinfo = ri
			dom := bmach.Domains[dom_id]

			sharedlist := ""
			solist := bmach.Shared_links[i]
			for j, so_id := range solist {
				sharedlist += bmach.Shared_objects[so_id].String()
				if j != len(solist)-1 {
					sharedlist += ","
				}
			}

			dom.Arch.Shared_constraints = sharedlist

			arch_filename := "arch_" + strconv.Itoa(i)
			arch_mod_name := "a" + strconv.Itoa(i)

			arch_names := map[string]string{"processor": "p" + strconv.Itoa(i), "rom": "p" + strconv.Itoa(i) + "rom", "ram": "p" + strconv.Itoa(i) + "ram"}

			if _, err := os.Stat(arch_filename + ".v"); os.IsNotExist(err) {
				f, err := os.Create(arch_filename + ".v")
				check(err)
				//defer f.Close()
				_, err = f.WriteString(dom.Arch.Write_verilog(arch_mod_name, arch_names, flavor))
				check(err)
				f.Close()
			}

			if _, err := os.Stat(arch_names["processor"] + ".v"); os.IsNotExist(err) {
				f, err := os.Create(arch_names["processor"] + ".v")
				check(err)
				//defer f.Close()
				_, err = f.WriteString(dom.Arch.Conproc.Write_verilog(pconf, &dom.Arch, arch_names["processor"], flavor))
				check(err)
				f.Close()
			}

			if _, err := os.Stat(arch_names["ram"] + ".v"); os.IsNotExist(err) {
				f, err := os.Create(arch_names["ram"] + ".v")
				check(err)
				//defer f.Close()
				_, err = f.WriteString(dom.Arch.Ram.Write_verilog(dom, arch_names["ram"], flavor))
				check(err)
				f.Close()
			}

			if _, err := os.Stat(arch_names["rom"] + ".v"); os.IsNotExist(err) {
				f, err := os.Create(arch_names["rom"] + ".v")
				check(err)
				//defer f.Close()
				_, err = f.WriteString(dom.Arch.Rom.Write_verilog(dom, arch_names["rom"], flavor))
				check(err)
				f.Close()
			}

		}

		if len(bmach.Shared_objects) > 0 {

			seq := make(map[string]int)

			for i, so := range bmach.Shared_objects {
				sname := so.Shortname()
				if _, ok := seq[sname]; !ok {
					seq[sname] = 0
				}

				if _, err := os.Stat(sname + strconv.Itoa(seq[sname]) + ".v"); os.IsNotExist(err) {
					f, err := os.Create(sname + strconv.Itoa(seq[sname]) + ".v")
					check(err)
					defer f.Close()
					_, err = f.WriteString(so.Write_verilog(bmach, i, sname+strconv.Itoa(seq[sname]), flavor))
					check(err)
				}

				seq[sname]++
			}
		}

		if _, err := os.Stat("bondmachine.v"); os.IsNotExist(err) {
			f, err := os.Create("bondmachine.v")
			check(err)
			defer f.Close()
			_, err = f.WriteString(bmach.Write_verilog_main(conf, "bondmachine", flavor))
			check(err)
		}

		for _, mod := range extramods {
			files, filescode := mod.ExtraFiles()
			for i, file := range files {
				f, err := os.Create(file)
				check(err)
				_, err = f.WriteString(filescode[i])
				check(err)
				f.Close()
			}
		}

		switch flavor {
		case "iverilog_simulation", "iverilog":
			if _, err := os.Stat("bondmachine_tb.v"); os.IsNotExist(err) {
				f, err := os.Create("bondmachine_tb.v")
				check(err)
				defer f.Close()
				_, err = f.WriteString(bmach.Write_verilog_testbench("bondmachine", flavor, iomaps, extramods, sbox))
				check(err)
			}
		case "basys3", "kintex7":
			if _, err := os.Stat("bondmachine_main.v"); os.IsNotExist(err) {
				f, err := os.Create("bondmachine_main.v")
				check(err)
				defer f.Close()
				_, err = f.WriteString(bmach.Write_verilog_board("bondmachine", flavor, iomaps, extramods))
				check(err)
			}
		default:
			return Prerror{"Verilog flavor unknown"}
		}
	} else {
		return Prerror{"No defined domains"}
	}
	return nil
}

func (bmach *Bondmachine) Write_verilog_main(conf *Config, module_name string, flavor string) string {

	result := ""
	result += "module " + module_name + "(clk, reset"

	// The External_inputs connected are defined in the module as port
	for i := 0; i < bmach.Inputs; i++ {
		result += ", i" + strconv.Itoa(i)
	}

	// The External_inputs connected are defined in the module as port
	for i := 0; i < bmach.Outputs; i++ {
		result += ", o" + strconv.Itoa(i)
	}

	result += ");\n\n"

	result += "	//--------------Input Ports-----------------------\n"
	result += "	input clk, reset;\n"

	// The External_inputs connected are defined as input port
	for i := 0; i < bmach.Inputs; i++ {
		result += "	input [" + strconv.Itoa(int(bmach.Rsize)-1) + ":0] i" + strconv.Itoa(i) + ";\n"
	}

	result += "	//--------------Output Ports-----------------------\n"

	// The External_inputs connected are defined as input port
	for i := 0; i < bmach.Outputs; i++ {
		result += "	output [" + strconv.Itoa(int(bmach.Rsize)-1) + ":0] o" + strconv.Itoa(i) + ";\n"
	}

	result += "\n\n"

	// The Internal_inputs connected are wire the unconnected are registers, the machine outputs are always wire handled with assign
	for i, linked := range bmach.Links {
		if linked == -1 {
			result += "	wire [" + strconv.Itoa(int(bmach.Rsize)-1) + ":0] " + strings.ToLower(bmach.Internal_inputs[i].String()) + ";\n"
		}
	}

	result += "\n"

	// The Internal_outputs are always wire except the machine inputs, but here we write only the unconnected
	for _, bond := range bmach.Internal_outputs {
		if conf.Commented_verilog {
			result += "\t//Analyzing Internal output " + strings.ToLower(bond.String()) + "\n"
		}
		if bond.Map_to == 0 {
			//TODO Check if every possibility is taken into account
			result += "	wire [" + strconv.Itoa(int(bmach.Rsize)-1) + ":0] " + strings.ToLower(bond.String()) + ";\n"
			result += "	wire " + strings.ToLower(bond.String()) + "_valid;\n"
			result += "	wire " + strings.ToLower(bond.String()) + "_received;\n"
		} else {
			//TODO Check if every possibility is taken into account
			result += "	wire [" + strconv.Itoa(int(bmach.Rsize)-1) + ":0] " + strings.ToLower(bond.String()) + ";\n"
			result += "	wire " + strings.ToLower(bond.String()) + "_valid;\n"
			result += "	wire " + strings.ToLower(bond.String()) + "_received;\n"
		}
	}

	result += "\n"

	// The internal wire for the connection to the shared memory and the channel
	if len(bmach.Shared_objects) > 0 {
		for proc_id, solist := range bmach.Shared_links {
			for _, so_id := range solist {
				//fmt.Println(proc_id, so_id)
				result += bmach.Shared_objects[so_id].Get_wires_perproc(bmach, proc_id, so_id, flavor)
			}
		}
	}

	result += "\n"

	result += "	//Instantiation of the Processors and Shared Objects\n"

	//Instatiation of the Processor
	for i, dom_id := range bmach.Processors {
		dom := bmach.Domains[dom_id]
		arch := dom.Arch

		arch_mod_name := "a" + strconv.Itoa(i)
		arch_instance_name := arch_mod_name + "_inst"

		result += "	" + arch_mod_name + " " + arch_instance_name + "(clk, reset"

		// A Processor input is connected to a register with its name (if there is no bond) or to an internal output
		// Default: the name is the name of the internal input side if the bond
		for j := 0; j < int(arch.N); j++ {
			map_to := uint8(2)
			res_id := i
			ext_id := j

			for k, linked := range bmach.Links {
				bond := bmach.Internal_inputs[k]
				if bond.Map_to == map_to && bond.Res_id == res_id && bond.Ext_id == ext_id {
					if linked == -1 {
						result += ", " + strings.ToLower(bond.String()) + ", " + strings.ToLower(bond.String()) + "_valid, " + strings.ToLower(bond.String()) + "_received"
					} else {
						result += ", " + strings.ToLower(bmach.Internal_outputs[linked].String()) + ", " + strings.ToLower(bmach.Internal_outputs[linked].String()) + "_valid, " + strings.ToLower(bmach.Internal_outputs[linked].String()) + "_received"
					}
					break
				}
			}
		}

		// A Processor output always use its name connected wheater or not
		for j := 0; j < int(arch.M); j++ {
			map_to := uint8(3)
			res_id := i
			ext_id := j

			bond := Bond{map_to, res_id, ext_id}
			result += ", " + strings.ToLower(bond.String()) + ", " + strings.ToLower(bond.String()) + "_valid, " + strings.ToLower(bond.String()) + "_received"
		}

		//memorize the name of the shared object
		// The module for the connection to the shared memory and the channel
		if len(bmach.Shared_objects) > 0 {
			solist := bmach.Shared_links[i]
			for _, so_id := range solist {
				result += bmach.Shared_objects[so_id].Get_header_perproc(bmach, i, so_id, flavor)
			}
		}
		result += ");\n"
	}

	//Instantiation of the Shared object
	if len(bmach.Shared_objects) > 0 {

		seq := make(map[string]int)

		for i, so := range bmach.Shared_objects {
			sname := so.Shortname()
			if _, ok := seq[sname]; !ok {
				seq[sname] = 0
			}

			result += "	" + sname + strconv.Itoa(seq[sname]) + " " + sname + strconv.Itoa(seq[sname]) + "(clk, reset"

			for proc_id, solist := range bmach.Shared_links {
				for _, so_id := range solist {
					if so_id == i {
						result += bmach.Shared_objects[so_id].Get_header_perproc(bmach, proc_id, so_id, flavor)
					}
				}
			}
			result += ");\n"
			seq[sname]++
		}
	}

	result += "\n"

	for i, linked := range bmach.Links {
		if linked != -1 {
			if bmach.Internal_inputs[i].Map_to == uint8(1) {
				result += "	assign " + strings.ToLower(bmach.Internal_inputs[i].String()) + " = " + strings.ToLower(bmach.Internal_outputs[linked].String()) + ";\n"
			}
		}
	}

	result += "\n"

	result += "endmodule\n"

	return result

}

func (bmach *Bondmachine) Write_verilog_testbench(module_name string, flavor string, iomaps *IOmap, extramods []ExtraModule, sbox *simbox.Simbox) string {

	result := ""
	result += "module " + module_name + "_tb;\n"
	result += "\n"
	result += "	reg clk, reset;\n"
	result += "\n"

	// The External_inputs connected are defined as input port
	for i := 0; i < bmach.Inputs; i++ {
		result += "	reg [" + strconv.Itoa(int(bmach.Rsize)-1) + ":0] Input" + strconv.Itoa(i) + ";\n"
	}

	// The External_inputs connected are defined as input port
	for i := 0; i < bmach.Outputs; i++ {
		result += "	wire [" + strconv.Itoa(int(bmach.Rsize)-1) + ":0] Output" + strconv.Itoa(i) + ";\n"
	}

	result += "\n"

	result += "\t" + module_name + " " + module_name + "_inst " + "(clk, reset"

	// The External_inputs connected are defined as input port
	for i := 0; i < bmach.Inputs; i++ {
		result += ", Input" + strconv.Itoa(i)
	}

	// The External_inputs connected are defined as input port
	for i := 0; i < bmach.Outputs; i++ {
		result += ", Output" + strconv.Itoa(i)
	}

	result += ");\n\n"

	result += "\tinitial  begin\n"
	result += "\t\t$dumpfile (\"working_dir/bondmachine.vcd\");\n"
	result += "\t\t$dumpvars;\n"
	result += "\tend\n"

	result += "\n"
	result += "\tinteger tickN;\n"
	result += "\tlocalparam TICK=20;\n"
	result += "\n"
	result += "\talways\n"
	result += "\t\tbegin\n"
	result += "\t\tclk = 1;\n"
	result += "\t\t#(TICK/2);\n"
	result += "\t\tclk = 0;\n"
	result += "\t\t#(TICK/2);\n"
	result += "\n"
	result += "\t\ttickN = tickN + 1;\n"
	result += "\t\t$display(\"--------------Tick %d---------------\", tickN);\n"
	result += "\tend\n"
	result += "\n"
	result += "\tinitial\n"
	result += "\tbegin\n"
	result += "\t\ttickN = 1;\n"
	result += "\t\treset = 1;\n"
	result += "\t\t#1;\n"
	result += "\t\treset = 0;\n"
	result += "\n"
	result += "\t\t#(100 * TICK);\n"
	result += "\n"
	result += "\t\t$finish;\n"
	result += "\tend\n"
	result += "endmodule\n"
	return result
}

func (bmach *Bondmachine) Write_verilog_basys3_7segment(module_name string, flavor string, iomaps *IOmap, extramods []ExtraModule) (string, error) {
	for _, mod := range extramods {
		if mod.Get_Name() == "basys3_7segment" {

			result := "\n"
			result += "module bond2seg(\n"
			result += "	input clk,\n"
			result += "	input reset,\n"
			result += "	input [15:0] value,\n"
			result += "	output [6:0] segment,\n"
			result += "	output enable_D1,\n"
			result += "	output enable_D2,\n"
			result += "	output enable_D3,\n"
			result += "	output enable_D4,\n"
			result += "	output dp\n"
			result += ");\n"
			result += "\n"
			result += "wire clk_point1hz;\n"
			result += "wire refreshClk;\n"
			result += "reg [3:0] hex;\n"
			result += "reg [3:0] reg_d0;\n"
			result += "reg [3:0] reg_d1;\n"
			result += "reg [3:0] reg_d2;\n"
			result += "reg [3:0] reg_d3;\n"
			result += "\n"
			result += "clkgen Uclkgen(\n"
			result += ".clk(clk),\n"
			result += ".refreshClk(refreshClk),\n"
			result += ".clk_point1hz(clk_point1hz)\n"
			result += ");\n"
			result += "\n"
			result += "enable_sr Uenable(\n"
			result += ".refreshClk(refreshClk),\n"
			result += ".enable_D1(enable_D1),\n"
			result += ".enable_D2(enable_D2),\n"
			result += ".enable_D3(enable_D3),\n"
			result += ".enable_D4(enable_D4)\n"
			result += ");\n"
			result += "\n"
			result += "ssd Ussd(\n"
			result += ".hex(hex),\n"
			result += ".segment(segment),\n"
			result += ".dp(dp)\n"
			result += ");\n"
			result += "\n"
			result += "always @(posedge clk) begin\n"
			result += "    reg_d0 <= value[3:0];\n"
			result += "    reg_d1 <= value[7:4];\n"
			result += "    reg_d2 <= value[11:8];\n"
			result += "    reg_d3 <= value[15:12];\n"
			result += "end\n"
			result += "\n"
			result += "always @ (*)\n"
			result += "\n"
			result += "case ({enable_D1,enable_D2,enable_D3,enable_D4})\n"
			result += "    4'b0111: hex = reg_d0;\n"
			result += "    4'b1011: hex = reg_d1;\n"
			result += "    4'b1101: hex = reg_d2;\n"
			result += "    4'b1110: hex = reg_d3;\n"
			result += "    default: hex = 0; \n"
			result += "endcase \n"
			result += "\n"
			result += "endmodule\n"
			result += "\n"
			result += "module clkgen(\n"
			result += "	input     clk, \n"
			result += "	output    refreshClk,\n"
			result += "	output    clk_point1hz\n"
			result += ");\n"
			result += "\n"
			result += "reg [26:0] 	count = 0;\n"
			result += "reg [16:0] 	refresh = 0;\n"
			result += "\n"
			result += "\n"
			result += "reg      	tmp_clk = 0;\n"
			result += "reg 		rclk = 0;\n"
			result += "\n"
			result += "\n"
			result += "assign clk_point1hz = tmp_clk;\n"
			result += "assign refreshClk = rclk;\n"
			result += "\n"
			result += "\n"
			result += "BUFG clock_buf_0(\n"
			result += "  .I(clk),\n"
			result += "  .O(clk_100mhz)\n"
			result += ");\n"
			result += "\n"
			result += "always @(posedge clk_100mhz) begin\n"
			result += "  if (count < 10000000) begin \n"
			result += "    count <= count + 1;\n"
			result += "  end\n"
			result += "  else begin\n"
			result += "    tmp_clk <= ~tmp_clk;\n"
			result += "    count <= 0;\n"
			result += "  end\n"
			result += "end\n"
			result += "\n"
			result += "always @(posedge clk_100mhz) begin\n"
			result += "	if (refresh < 100000) begin\n"
			result += "		refresh <= refresh + 1;\n"
			result += "	end else begin\n"
			result += "		refresh <= 0;\n"
			result += "		rclk <= ~rclk;\n"
			result += "	end\n"
			result += "end\n"
			result += "\n"
			result += "endmodule\n"
			result += "\n"
			result += "module enable_sr(\n"
			result += "	input         refreshClk,\n"
			result += "	output        enable_D1,\n"
			result += "	output        enable_D2,\n"
			result += "	output        enable_D3,\n"
			result += "	output        enable_D4\n"
			result += ");\n"
			result += "\n"
			result += "reg [3:0] pattern = 4'b0111; \n"
			result += "\n"
			result += "assign enable_D1 = pattern[3];\n"
			result += "assign enable_D2 = pattern[2];\n"
			result += "assign enable_D3 = pattern[1];\n"
			result += "assign enable_D4 = pattern[0];\n"
			result += "\n"
			result += "always @(posedge refreshClk) begin\n"
			result += "	pattern <= {pattern[0],pattern[3:1]};	\n"
			result += "end\n"
			result += "\n"
			result += "\n"
			result += "\n"
			result += "endmodule\n"
			result += "\n"
			result += "module ssd(\n"
			result += "	input [3:0] hex,\n"
			result += "	output reg [6:0] segment,\n"
			result += "	output dp\n"
			result += ");\n"
			result += "\n"
			result += "always @ (*)\n"
			result += "	case (hex) \n"
			result += "		0: segment = 7'b0000001;\n"
			result += "		1: segment = 7'b1001111;\n"
			result += "		2: segment = 7'b0010010;\n"
			result += "		3: segment = 7'b0000110;\n"
			result += "		4: segment = 7'b1001100;\n"
			result += "		5: segment = 7'b0100100;\n"
			result += "		6: segment = 7'b0100000;\n"
			result += "		7: segment = 7'b0001101;\n"
			result += "		8: segment = 7'b0000000;\n"
			result += "		9: segment = 7'b0000100;\n"
			result += "		10: segment = 7'b0001000;\n"
			result += "		11: segment = 7'b1100000;\n"
			result += "		12: segment = 7'b0110001;\n"
			result += "		13: segment = 7'b1000010;\n"
			result += "		14: segment = 7'b0110000;\n"
			result += "		15: segment = 7'b0111000;\n"
			result += "		default: segment = 7'b0000001;\n"
			result += "	endcase	\n"
			result += "assign dp = 4'b1111;\n"
			result += "    \n"
			result += "endmodule\n"
			result += "\n"

			return result, nil

		}
	}

	return "", errors.New("No basys3_7segment module found")
}

func (bmach *Bondmachine) Write_verilog_etherbond(module_name string, flavor string, iomaps *IOmap, extramods []ExtraModule) (string, error) {
	for _, mod := range extramods {
		if mod.Get_Name() == "etherbond" {
			etherbond_params := mod.Get_Params().Params
			fmt.Println(etherbond_params)

			cluster_id, _ := strconv.Atoi(etherbond_params["cluster_id"])
			peer_id, _ := strconv.Atoi(etherbond_params["peer_id"])

			result := ""
			result += "module " + module_name + "_main(\n"
			result += "\n"

			clk_name := "clk"
			rst_name := "reset"

			result += "\tinput " + clk_name + ",\n"
			result += "\tinput " + rst_name + ",\n"
			result += "\toutput sck,\n"
			result += "\toutput mosi,\n"
			result += "\toutput cs_n,\n"
			result += "\tinput miso,\n"
			result += "\tinput INT_n,\n"

			inames := strings.Split(etherbond_params["inputs"], ",")
			iids := strings.Split(etherbond_params["input_ids"], ",")

			onames := strings.Split(etherbond_params["outputs"], ",")
			oids := strings.Split(etherbond_params["output_ids"], ",")
			odests := strings.Split(etherbond_params["destinations"], ",")

			for _, iname := range inames {
				if iname != "" {
					result += "\toutput [" + strconv.Itoa(int(bmach.Rsize)-1) + ":0] " + iname + ",\n"
				}
			}

			for _, oname := range onames {
				if oname != "" {
					result += "\tinput [" + strconv.Itoa(int(bmach.Rsize)-1) + ":0] " + oname + ",\n"
				}
			}

			result = result[0:len(result)-2] + "\n);\n\n"

			result += "\n"

			result += "\t//Etherbond type\n"
			result += "\tlocalparam [15:0] ETHERBOND_TYPE = 16'h8888;\n"

			result += "\n"

			result += "\t//Etherbond commands\n"
			result += "\tlocalparam [7:0] ADV_CLU = 8'h01;\n"
			result += "\tlocalparam [7:0] ADV_CH = 8'h02;\n"
			result += "\tlocalparam [7:0] ADV_IN = 8'h03;\n"
			result += "\tlocalparam [7:0] ADV_OUT = 8'h04;\n"
			result += "\tlocalparam [7:0] IO_TR = 8'h05;\n"
			result += "\tlocalparam [7:0] TAGACK = 8'hff;\n"

			result += "\n"

			result += "\t//packet definition for ethernet transmission\n"
			result += "\tlocalparam [15:0] ethertype = 16'h8888;\n"
			result += "\tlocalparam [47:0] mymac = 48'h0288" + fmt.Sprintf("%08d", peer_id) + ";\n"
			result += "\tlocalparam [31:0] mycluster_id = 32'd" + fmt.Sprintf("%08d", cluster_id) + ";\n"
			result += "\tlocalparam [31:0] mypeer_id = 32'd" + fmt.Sprintf("%08d", peer_id) + ";\n"
			result += "\treg [31:0] tag;\n"

			result += "\n"

			// 1 Advertises
			sendwires := 0

			for i, iname := range inames {
				if iname != "" {
					tmpi, _ := strconv.Atoi(iids[i])
					result += "\tlocalparam [31:0] id_resource_" + iname + " = 32'd" + fmt.Sprintf("%08d", tmpi) + ";\n"
					sendwires++
				}
			}

			result += "\n"

			for i, oname := range onames {
				if oname != "" {
					sendwires++
					tmpi, _ := strconv.Atoi(oids[i])
					result += "\tlocalparam [31:0] id_resource_" + oname + " = 32'd" + fmt.Sprintf("%08d", tmpi) + ";\n"
					ods := strings.Split(odests[i], "-")
					for j, oid := range ods {
						oid_n, _ := strconv.Atoi(oid)
						result += "\tlocalparam [31:0] " + oname + "_dest_" + strconv.Itoa(j) + "_id = 32'd" + fmt.Sprintf("%08d", oid_n) + ";\n"
						if mac, ok := etherbond_params["peer_"+oid+"_mac"]; ok {
							if mac == "auto" {
								result += "\tlocalparam [47:0] " + oname + "_dest_" + strconv.Itoa(j) + "_mac = 48'h0288" + fmt.Sprintf("%08d", oid_n) + ";\n"
							} else if mac == "adv" {
								// TODO just the case here eventually implemented logic has to be written
								result += "\treg [31:0] " + oname + "_dest_" + strconv.Itoa(j) + "_mac;\n"
							} else {
								result += "\tlocalparam [47:0] " + oname + "_dest_" + strconv.Itoa(j) + "_mac = 48'h" + mac + ";\n"
							}
						} else {
							result += "\tlocalparam [47:0] " + oname + "_dest_" + strconv.Itoa(j) + "_mac = 48'h0288" + fmt.Sprintf("%08d", oid_n) + ";\n"
						}
						sendwires++
					}
				}
			}

			result += "\n"

			for _, oname := range onames {
				if oname != "" {
					result += "\twire [" + strconv.Itoa(int(bmach.Rsize)-1) + ":0] " + oname + "_copy;\n"
					result += "\tassign " + oname + "_copy = " + oname + ";\n"
				}
			}

			result += "\n"

			result += "\tlocalparam LASTFRAME = " + strconv.Itoa(Needed_bits(sendwires+2)) + "'d" + strconv.Itoa(sendwires+1) + ";\n"

			framesize := 256 + 24 + int(bmach.Rsize)

			result += "\twire [" + strconv.Itoa(framesize-1) + ":0] ethtxs [0:" + strconv.Itoa(sendwires) + "];\n"

			iw := 1
			result += "\tassign ethtxs[0] = {8'hFF, 48'hFFFFFFFFFFFF, mymac, ethertype, ADV_CLU, mycluster_id, mypeer_id, " + strconv.Itoa(framesize-192) + "'b0};\n"
			for _, iname := range inames {
				if iname != "" {
					result += "\tassign ethtxs[" + strconv.Itoa(iw) + "] = {8'hFF, 48'hFFFFFFFFFFFF, mymac, ethertype, ADV_IN, mycluster_id, mypeer_id, id_resource_" + iname + ", " + strconv.Itoa(framesize-224) + "'b0};\n"
					iw++
				}
			}

			for i, oname := range onames {
				if oname != "" {
					result += "\tassign ethtxs[" + strconv.Itoa(iw) + "] = {8'hFF, 48'hFFFFFFFFFFFF, mymac, ethertype, ADV_OUT, mycluster_id, mypeer_id, id_resource_" + oname + ", " + strconv.Itoa(framesize-224) + "'b0};\n"
					iw++
					ods := strings.Split(odests[i], "-")
					for j, _ := range ods {
						dname := oname + "_dest_" + strconv.Itoa(j) + "_mac"
						result += "\tassign ethtxs[" + strconv.Itoa(iw) + "] = {8'hFF, " + dname + ", mymac, ethertype, IO_TR , tag,  mycluster_id, mypeer_id, id_resource_" + oname + ", " + oname + "_copy, 24'b0};\n"
						iw++
					}
				}
			}

			result += "\twire RESET_n = !reset;\n"

			result += "\n"

			result += "\t// Logic to start the configuration phase\n"
			result += "\t(* keep=\"true\" *) wire end_configuration;\n"
			result += "\treg start_conf;\n"
			result += "\treg [9:0] cnt4cnf;\n"
			result += "\n"
			result += "\talways@(posedge clk)\n"
			result += "\tbegin\n"
			result += "\t\tif(RESET_n==1'b0)\n"
			result += "\t\t\tcnt4cnf <= 32'b0;\n"
			result += "\t\telse\n"
			result += "\t\t\tcnt4cnf <= cnt4cnf + 1'b1;\n"
			result += "\tend\n"

			result += "\talways@(posedge clk)\n"
			result += "\tbegin\n"
			result += "\t\tif(RESET_n==1'b0 | end_configuration)\n"
			result += "\t\t\tstart_conf <= 1'b0;\n"
			result += "\t\telse if(cnt4cnf==9'h1FF)\n"
			result += "\t\t\tstart_conf <= 1'b1;\n"
			result += "\tend\n"

			result += "\n"

			result += "\t// Frame explorer\n"
			result += "\t(* keep=\"true\" *) reg [" + strconv.Itoa(Needed_bits(sendwires+1)) + ":0] frame_explorer;\n"
			result += "\t(* keep=\"true\" *) reg [28:0] frame_counter;\n"

			for i := 0; i < sendwires+1; i++ {
				result += "\treg frame_ready_" + strconv.Itoa(i) + ";\n"
			}
			result += "\treg frame_ready_rx_0;\n"

			for i := 0; i < sendwires+1; i++ {
				result += "\t(* keep=\"true\" *) wire frame_ready_reset_" + strconv.Itoa(i) + ";\n"
			}
			result += "\t(* keep=\"true\" *) wire frame_ready_reset_rx_0;\n"

			result += "\n"

			// Write enabling
			result += "\treg start_write, start_rx;\n"
			result += "\twire done_write;\n"
			result += "\treg [ETHERNET_LENGTH-1:0] EthTx;\n"
			result += "\t(* keep=\"true\" *) reg [ETHERNET_LENGTH-1:0] EthRx;\n"

			result += "\n"

			result += "\talways@(posedge clk)\n"
			result += "\tbegin\n"
			result += "\t\tif(RESET_n==1'b0) \n"
			result += "\t\t\tframe_counter <= 0;\n"
			result += "\t\telse\n"
			result += "\t\t\tframe_counter <= frame_counter + 1;\n"
			result += "\tend\n"

			result += "\n"

			subresult := "frame_ready_rx_0, "
			subresult2 := "1'b1, "
			for i := sendwires; i >= 0; i-- {
				subresult += "frame_ready_" + strconv.Itoa(i) + ", "
				subresult2 += "1'b0, "
			}
			result += "\twire [" + strconv.Itoa(sendwires+1) + ":0] frame_ready = {" + subresult[0:len(subresult)-2] + "};\n"
			result += "\twire [" + strconv.Itoa(sendwires+1) + ":0] frame_type = {" + subresult2[0:len(subresult2)-2] + "};\n"

			result += "\n"

			result += "\twire stop_frame_explorer = (frame_explorer==LASTFRAME) ? 1'b1 : 1'b0;\n"

			result += "\n"

			result += "\talways@(posedge clk)\n"
			result += "\tbegin\n"
			result += "\t\tif(RESET_n==1'b0  | end_configuration==1'b0 | stop_frame_explorer)\n"
			zerostring := ""
			for i := 0; i < Needed_bits(sendwires+2); i++ {
				zerostring += "0"
			}
			result += "\t\t\tframe_explorer <= " + strconv.Itoa(Needed_bits(sendwires+2)) + "'b" + zerostring + ";\n"
			result += "\t\telse if(start_write == 1'b0 & start_rx == 1'b0)\n"
			result += "\t\t\tframe_explorer <= frame_explorer + 1;\n"
			result += "\tend\n"

			result += "\n"

			result += "\treg [" + strconv.Itoa(sendwires+1) + ":0] frame_ready_reset;\n"
			for i := 0; i < sendwires+1; i++ {
				result += "\tassign frame_ready_reset_" + strconv.Itoa(i) + " = frame_ready_reset[" + strconv.Itoa(i) + "];\n"
			}
			result += "\tassign frame_ready_reset_rx_0 = frame_ready_reset[" + strconv.Itoa(sendwires+1) + "];\n"

			result += "\n"

			result += "\talways@(posedge clk)\n"
			result += "\tbegin\n"
			result += "\t\tif(RESET_n==1'b0)\n"
			result += "\t\t\ttag <= 32'b0;\n"
			result += "\t\telse if(done_write == 1'b1)\n"
			result += "\t\t\ttag <= tag + 1'b1;\n"
			result += "\tend\n"

			result += "\n"

			result += "\talways@(posedge clk)\n"
			result += "\tbegin\n"
			result += "\t\tif(RESET_n==1'b0) begin\n"
			result += "\t\t\tstart_write <= 1'b0;\n"
			result += "\t\t\tstart_rx <= 1'b0;\n"
			result += "\t\t\tEthTx <= 'b0;\n"
			result += "\t\tend\n"
			result += "\t\telse\n"
			result += "\t\tbegin\n"
			result += "\t\t\tframe_ready_reset <= 'b0;\n"
			result += "\t\t\tif(frame_ready[frame_explorer] & start_write==1'b0 & start_rx==1'b0 & frame_type[frame_explorer]==1'b0)\n"
			result += "\t\t\tbegin\n"
			result += "\t\t\t\tframe_ready_reset[frame_explorer] <= 1'b1;\n"
			result += "\t\t\t\tEthTx <= ethtxs[frame_explorer];\n"
			result += "\t\t\t\tstart_write <= 1'b1;\n"
			result += "\t\t\tend\n"
			result += "\t\t\telse if(done_write)\n"
			result += "\t\t\t\tstart_write <= 1'b0;\n"
			result += "\t\t\tif(frame_ready[frame_explorer] & start_write==1'b0 & start_rx==1'b0 & frame_type[frame_explorer]==1'b1)\n"
			result += "\t\t\tbegin\n"
			result += "\t\t\t\tframe_ready_reset[frame_explorer] <= 1'b1;\n"
			result += "\t\t\t\tstart_rx <= 1'b1;\n"
			result += "\t\t\tend\n"
			result += "\t\t\telse if(done_rx)\n"
			result += "\t\t\t\tstart_rx <= 1'b0;\n"
			result += "\t\tend\n"
			result += "\tend\n"

			result += "\n"

			result += "\treg old_clu_adv;\n"
			result += "\talways@(posedge clk)\n"
			result += "\tbegin\n"
			result += "\t\tif(RESET_n==1'b0) \n"
			result += "\t\t\told_clu_adv <= 0;\n"
			result += "\t\telse\n"
			result += "\t\t\told_clu_adv <= frame_counter[27];\n"
			result += "\tend\n"

			result += "\talways@(posedge clk)\n"
			result += "\tbegin\n"
			result += "\t\tif(RESET_n==1'b0 | frame_ready_reset_0) \n"
			result += "\t\t\tframe_ready_0 <= 1'b0;\n"
			result += "\t\telse if(old_clu_adv!=frame_counter[27]) \n"
			result += "\t\t\tframe_ready_0 <= 1'b1;\n"
			result += "\tend\n"

			result += "\n"

			iw = 1
			for i, iname := range inames {
				if iname != "" {
					adv_name := "old_in_adv_" + strconv.Itoa(i)
					result += "\treg " + adv_name + ";\n"
					result += "\talways@(posedge clk)\n"
					result += "\tbegin\n"
					result += "\t\tif(RESET_n==1'b0) \n"
					result += "\t\t\t" + adv_name + " <= 0;\n"
					result += "\t\telse\n"
					result += "\t\t\t" + adv_name + " <= frame_counter[27];\n"
					result += "\tend\n"

					result += "\talways@(posedge clk)\n"
					result += "\tbegin\n"
					result += "\t\tif(RESET_n==1'b0 | frame_ready_reset_" + strconv.Itoa(iw) + ") \n"
					result += "\t\t\tframe_ready_" + strconv.Itoa(iw) + " <= 1'b0;\n"
					result += "\t\telse if(" + adv_name + "!=frame_counter[27]) \n"
					result += "\t\t\tframe_ready_" + strconv.Itoa(iw) + " <= 1'b1;\n"
					result += "\tend\n"

					result += "\n"
					iw++
				}
			}

			for i, oname := range onames {
				if oname != "" {
					adv_name := "old_out_adv_" + strconv.Itoa(i)
					result += "\treg " + adv_name + ";\n"
					result += "\talways@(posedge clk)\n"
					result += "\tbegin\n"
					result += "\t\tif(RESET_n==1'b0) \n"
					result += "\t\t\t" + adv_name + " <= 'b0;\n"
					result += "\t\telse\n"
					result += "\t\t\t" + adv_name + " <= frame_counter[27];\n"
					result += "\tend\n"

					result += "\talways@(posedge clk)\n"
					result += "\tbegin\n"
					result += "\t\tif(RESET_n==1'b0 | frame_ready_reset_" + strconv.Itoa(iw) + ") \n"
					result += "\t\t\tframe_ready_" + strconv.Itoa(iw) + " <= 1'b0;\n"
					result += "\t\telse if(" + adv_name + "!=frame_counter[27]) \n"
					result += "\t\t\tframe_ready_" + strconv.Itoa(iw) + " <= 1'b1;\n"
					result += "\tend\n"

					result += "\n"
					iw++

					result += "\treg [" + strconv.Itoa(int(bmach.Rsize)-1) + ":0] " + oname + "_d1;\n"
					result += "\talways@(posedge clk)\n"
					result += "\tbegin\n"
					result += "\t\tif(RESET_n==1'b0) \n"
					result += "\t\t\t" + oname + "_d1 <= 'b0;\n"
					result += "\t\telse\n"
					result += "\t\t\t" + oname + "_d1 <= " + oname + "_copy;\n"
					result += "\tend\n"

					result += "\talways@(posedge clk)\n"
					result += "\tbegin\n"
					result += "\t\tif(RESET_n==1'b0 | frame_ready_reset_" + strconv.Itoa(iw) + ") \n"
					result += "\t\t\tframe_ready_" + strconv.Itoa(iw) + " <= 1'b0;\n"
					result += "\t\telse if(" + oname + "_d1!=" + oname + "_copy) \n"
					result += "\t\tbegin\n"

					ods := strings.Split(odests[i], "-")
					for _, _ = range ods {
						result += "\t\t\tframe_ready_" + strconv.Itoa(iw) + " <= 1'b1;\n"
						iw++
					}
					result += "\t\tend\n"
					result += "\tend\n"

					result += "\n"
				}
			}

			result += "  //logic to enable the RX frame\n"

			result += "  reg enable_rx;\n"
			result += "  always@(posedge clk)\n"
			result += "  begin\n"
			result += "      if(RESET_n==1'b0)\n"
			result += "          enable_rx <= 'b0;\n"
			result += "      else\n"
			result += "          enable_rx <= frame_counter[27];\n"
			result += "  end\n"
			result += "  always@(posedge clk)\n"
			result += "  begin\n"
			result += "      if(RESET_n==1'b0 | frame_ready_reset_rx_0)\n"
			result += "          frame_ready_rx_0 <= 1'b0;\n"
			result += "      else if(enable_rx!=frame_counter[28])\n"
			result += "          frame_ready_rx_0 <= 1'b1;\n"
			result += "  end\n"
			result += "\n"
			result += "  (* keep = \"true\" *) reg [47:0] source_addess;\n"
			result += "  (* keep = \"true\" *) reg [47:0] destination_addess;\n"
			result += "  (* keep = \"true\" *) reg [15:0] ethernet_type;\n"
			result += "  (* keep = \"true\" *) reg [31:0] tagin;\n"
			result += "  (* keep = \"true\" *) reg [31:0] clusterid;\n"
			result += "  (* keep = \"true\" *) reg [31:0] nodeid;\n"

			result += "\n"
			result += "  always@(posedge clk)\n"
			result += "  begin\n"
			result += "    if(RESET_n==1'b0) begin\n"
			result += "        source_addess <= 'b0;\n"
			result += "        destination_addess <= 'b0;\n"
			result += "        ethernet_type <= 'b0;\n"
			result += "        i0_copy <= 8'b0;\n"
			result += "        tagin <= 32'b0;\n"
			result += "        nodeid <= 32'b0;\n"
			result += "        clusterid <= 32'b0;\n"
			result += "    end\n"
			result += "    else if(done_rx) begin\n"
			result += "        case (EthRx[ETHERNET_LENGTH-1-8-48*2:ETHERNET_LENGTH-8-48*2-16])\n"
			result += "            ETHERBOND_TYPE: begin\n"
			result += "                case (EthRx[ETHERNET_LENGTH-1-8-48*2-16:ETHERNET_LENGTH-8-48*2-16-8])\n"
			result += "                    IO_TR: begin\n"
			result += "                    tagin <= EthRx[ETHERNET_LENGTH-1-8-48*2-16-8:ETHERNET_LENGTH-8-48*2-16-8-32];\n"
			result += "                    clusterid <= EthRx[ETHERNET_LENGTH-1-8-48*2-16-8-32:ETHERNET_LENGTH-8-48*2-16-8-32-32];\n"
			result += "                    nodeid <= EthRx[ETHERNET_LENGTH-1-8-48*2-16-8-32-32:ETHERNET_LENGTH-8-48*2-16-8-32-32-32];\n"
			result += "                    case (EthRx[ETHERNET_LENGTH-1-8-48*2-16-8-32-32-32:ETHERNET_LENGTH-8-48*2-16-8-32-32-32-32])\n"
			result += "                        id_resource_i0: begin\n"
			result += "                            i0_copy <= EthRx[ETHERNET_LENGTH-1-8-48*2-16-8-32-32-32-32:ETHERNET_LENGTH-8-48*2-16-8-32-32-32-32-8];\n"
			result += "                        end\n"
			result += "                    endcase\n"
			result += "                    end\n"
			result += "                endcase\n"
			result += "            end\n"
			result += "        endcase\n"
			result += "        source_addess <= EthRx[ETHERNET_LENGTH-1-8-48:ETHERNET_LENGTH-8-48*2];\n"
			result += "        destination_addess <= EthRx[ETHERNET_LENGTH-1-8:ETHERNET_LENGTH-8-48];\n"
			result += "        ethernet_type <= EthRx[ETHERNET_LENGTH-1-8-48*2:ETHERNET_LENGTH-8-48*2-16];\n"
			result += "    end\n"
			result += "  end\n"

			result += "\n"
			result += "\n"
			result += "\n"
			result += "  TopModuleSPI TopModuleSPI_inst\n"
			result += "  (\n"
			result += "    .CLOCK(clk),\n"
			result += "    .SCK(sck),\n"
			result += "    .MOSI(mosi),\n"
			result += "    .CS_n(cs_n),\n"
			result += "    .MISO(miso),\n"
			result += "    .reset_n(RESET_n),\n"
			result += "    .start_conf(start_conf),\n"
			result += "    .end_conf(end_configuration),\n"
			result += "    .EthTx(EthTx),\n"
			result += "    .start_write(start_write),\n"
			result += "    .done_write(done_write),\n"
			result += "    .start_rx(start_rx),\n"
			result += "    .done_rx(done_rx),\n"
			result += "    .EthRx(EthRx)\n"
			result += "   );\n"

			//loopcheck := false
			for _, iname := range inames {
				if iname != "" {
					result += "\treg [" + strconv.Itoa(int(bmach.Rsize)-1) + ":0] " + iname + "_copy;\n"
					result += "\tassign " + iname + "=" + iname + "_copy;\n"
					//loopcheck = true
				}
			}

			//if loopcheck {
			//	result += "\talways @ (posedge clk) begin\n"
			//	for _, iname := range inames {
			//		if iname != "" {
			//			result += "\t\t" + iname + " <= " + iname + "_copy;\n"
			//		}
			//	}
			//	result += "\tend\n"
			//}

			result += "endmodule\n\n\n"
			return result, nil

		}
	}
	return "", errors.New("No etherbond module found")
}

func (bmach *Bondmachine) Write_verilog_udpbond(module_name string, flavor string, iomaps *IOmap, extramods []ExtraModule) (string, error) {
	for _, mod := range extramods {
		if mod.Get_Name() == "udpbond" {
			//			udpbond_params := mod.Get_Params().Params
			//			fmt.Println(udpbond_params)
			//
			//			//			cluster_id, _ := strconv.Atoi(udpbond_params["cluster_id"])
			//			//			peer_id, _ := strconv.Atoi(udpbond_params["peer_id"])
			//
			result := ""
			//			result += "module " + module_name + "_main(\n"
			//			result += "\n"
			//
			//			clk_name := "clk"
			//			rst_name := "reset"
			//
			//			result += "\tinput " + clk_name + ",\n"
			//			result += "\tinput " + rst_name + ",\n"
			//			result += "\toutput wifi_enable,\n"
			//			result += "\tinput wifi_rx,\n"
			//			result += "\toutput wifi_tx,\n"
			//
			//			inames := strings.Split(udpbond_params["inputs"], ",")
			//			//			iids := strings.Split(udpbond_params["input_ids"], ",")
			//
			//			onames := strings.Split(udpbond_params["outputs"], ",")
			//			//			oids := strings.Split(udpbond_params["output_ids"], ",")
			//			//			odests := strings.Split(udpbond_params["destinations"], ",")
			//
			//			for _, iname := range inames {
			//				if iname != "" {
			//					result += "\toutput [" + strconv.Itoa(int(bmach.Rsize)-1) + ":0] " + iname + ",\n"
			//				}
			//			}
			//
			//			for _, oname := range onames {
			//				if oname != "" {
			//					result += "\tinput [" + strconv.Itoa(int(bmach.Rsize)-1) + ":0] " + oname + ",\n"
			//				}
			//			}
			//
			//			result = result[0:len(result)-2] + "\n);\n\n"
			//
			//			result += "\n"
			//			result += "endmodule\n\n\n"

			return result, nil

		}
	}
	return "", errors.New("No udpbond module found")
}

func (bmach *Bondmachine) Write_verilog_board(module_name string, flavor string, iomaps *IOmap, extramods []ExtraModule) string {

	resolved_io := make(map[string]string)

	slow_module := false
	var slow_params map[string]string

	// etherbond and udpbond are exclusive

	etherbond_module := false
	var etherbond_params map[string]string

	udpbond_module := false
	var udpbond_params map[string]string

	var inames []string
	var onames []string

	basys3_7segment_module := false
	var basys3_7segment_params map[string]string
	basys3_7segment_mapped := ""

	result := ""
	result_headers := ""

	for _, mod := range extramods {

		result_headers += mod.Verilog_headers()

		result += mod.Static_verilog()

		if mod.Get_Name() == "slow" {
			slow_module = true
			slow_params = mod.Get_Params().Params
		}
		if mod.Get_Name() == "etherbond" {
			etherbond_module = true
			etherbond_params = mod.Get_Params().Params
			//fmt.Println(etherbond_params)
			if subresult, ok := bmach.Write_verilog_etherbond("etherbond", flavor, iomaps, extramods); ok == nil {
				result += subresult
			} else {
				// TODO proper error handling
			}
			inames = strings.Split(etherbond_params["inputs"], ",")
			onames = strings.Split(etherbond_params["outputs"], ",")
		}
		if mod.Get_Name() == "udpbond" {
			udpbond_module = true
			udpbond_params = mod.Get_Params().Params
			//fmt.Println(udpbond_params)
			if subresult, ok := bmach.Write_verilog_udpbond("udpbond", flavor, iomaps, extramods); ok == nil {
				result += subresult
			} else {
				// TODO proper error handling
			}
			inames = strings.Split(udpbond_params["inputs"], ",")
			onames = strings.Split(udpbond_params["outputs"], ",")
		}
		if mod.Get_Name() == "basys3_7segment" {
			basys3_7segment_module = true
			basys3_7segment_params = mod.Get_Params().Params
			if subresult, ok := bmach.Write_verilog_basys3_7segment("basys3_7segment", flavor, iomaps, extramods); ok == nil {
				result += subresult
			} else {
				// TODO proper error handling
			}
		}
	}

	result += "module " + module_name + "_main(\n"
	result += "\n"

	clk_name := "clk"

	if cname, ok := iomaps.Assoc["clk"]; ok {
		clk_name = cname
	}

	rst_name := "reset"

	if rname, ok := iomaps.Assoc["reset"]; ok {
		rst_name = rname
	}

	result += "\tinput " + clk_name + ",\n"
	result += "\tinput " + rst_name + ",\n"

	for i := 0; i < bmach.Inputs; i++ {
		iname := Get_input_name(i)
		resolved_io[iname] = "none"
		subresult := ""
		if rname, ok := iomaps.Assoc[iname]; ok {
			resolved_io[iname] = "board"
			subresult = "\tinput " + rname + ",\n"
		}
		result += subresult
	}

	for i := 0; i < bmach.Outputs; i++ {
		oname := Get_output_name(i)
		resolved_io[oname] = "none"
		subresult := ""
		if rname, ok := iomaps.Assoc[oname]; ok {
			resolved_io[oname] = "board"
			subresult = "\toutput reg " + rname + ",\n"
		}
		result += subresult
	}

	if basys3_7segment_module {
		result += "\toutput [6:0] segment,\n"
		result += "\toutput enable_D1,\n"
		result += "\toutput enable_D2,\n"
		result += "\toutput enable_D3,\n"
		result += "\toutput enable_D4,\n"
		result += "\toutput dp,\n"
	}

	if etherbond_module {
		result += "\toutput sck,\n"
		result += "\toutput mosi,\n"
		result += "\toutput cs_n,\n"
		result += "\tinput miso,\n"
		result += "\tinput int_n,\n"
	}

	if udpbond_module {
		result += "\toutput wifi_enable,\n"
		result += "\tinput wifi_rx,\n"
		result += "\toutput wifi_tx,\n"
	}

	result = result[0:len(result)-2] + "\n);\n\n"

	if rst_name != "reset" {
		result += "\tassign  reset = " + rst_name + ";\n"
	}

	// The External_inputs connected are defined as input port
	for i := 0; i < bmach.Inputs; i++ {
		iname := Get_input_name(i)
		if _, ok := iomaps.Assoc[iname]; ok {
			result += "\treg [" + strconv.Itoa(int(bmach.Rsize)-1) + ":0] Input" + strconv.Itoa(i) + ";\n"
			continue
		}

		if etherbond_module {
			for _, ethiname := range inames {
				if iname == ethiname {
					resolved_io[iname] = "etherbond"
					result += "\twire [" + strconv.Itoa(int(bmach.Rsize)-1) + ":0] Input" + strconv.Itoa(i) + ";\n"
					result += "\twire Input" + strconv.Itoa(i) + "_valid;\n"
					result += "\twire Input" + strconv.Itoa(i) + "_received;\n"
					break
				}
			}
		}

		if udpbond_module {
			for _, ethiname := range inames {
				if iname == ethiname {
					resolved_io[iname] = "udpbond"
					result += "\twire [" + strconv.Itoa(int(bmach.Rsize)-1) + ":0] Input" + strconv.Itoa(i) + ";\n"
					result += "\twire Input" + strconv.Itoa(i) + "_valid;\n"
					result += "\twire Input" + strconv.Itoa(i) + "_received;\n"
					break
				}
			}
		}
	}

	// The External_inputs connected are defined as input port
	for i := 0; i < bmach.Outputs; i++ {
		oname := Get_output_name(i)
		if _, ok := iomaps.Assoc[oname]; ok {
			result += "\twire [" + strconv.Itoa(int(bmach.Rsize)-1) + ":0] Output" + strconv.Itoa(i) + ";\n"
			continue
		}

		if etherbond_module {
			for _, ethoname := range onames {
				if oname == ethoname {
					resolved_io[oname] = "etherbond"
					result += "\twire [" + strconv.Itoa(int(bmach.Rsize)-1) + ":0] Output" + strconv.Itoa(i) + ";\n"
					result += "\twire Onput" + strconv.Itoa(i) + "_valid;\n"
					result += "\twire Onput" + strconv.Itoa(i) + "_received;\n"
					break
				}
			}
		}

		if udpbond_module {
			for _, ethoname := range onames {
				if oname == ethoname {
					resolved_io[oname] = "udpbond"
					result += "\twire [" + strconv.Itoa(int(bmach.Rsize)-1) + ":0] Output" + strconv.Itoa(i) + ";\n"
					result += "\twire Onput" + strconv.Itoa(i) + "_valid;\n"
					result += "\twire Onput" + strconv.Itoa(i) + "_received;\n"
					break
				}
			}
		}

		if basys3_7segment_module {
			if oname == basys3_7segment_params["mapped_output"] {
				resolved_io[oname] = "basys3_7segment"
				result += "\twire [" + strconv.Itoa(int(bmach.Rsize)-1) + ":0] Output" + strconv.Itoa(i) + ";\n"
				basys3_7segment_mapped = "Output" + strconv.Itoa(i)
			}
		}
	}

	result += "\n"

	clock_string := clk_name

	if slow_module {
		result += "\treg [31:0] divider;\n\n"
		clock_string = "divider[" + slow_params["slow_factor"] + "]"
	}

	result += "\t" + module_name + " " + module_name + "_inst " + "(" + clock_string + ", reset"

	// The External_inputs connected are defined as input port
	for i := 0; i < bmach.Inputs; i++ {
		result += ", Input" + strconv.Itoa(i)
	}

	// The External_inputs connected are defined as input port
	for i := 0; i < bmach.Outputs; i++ {
		result += ", Output" + strconv.Itoa(i)
	}

	result += ");\n\n"

	if etherbond_module {
		result += "\tetherbond_main etherbond_main_inst " + "(" + clk_name + ", reset"
		result += ", sck"
		result += ", mosi"
		result += ", cs_n"
		result += ", miso"
		result += ", int_n"
		for _, iname := range inames {
			for i := 0; i < bmach.Inputs; i++ {
				iiname := Get_input_name(i)
				if iiname == iname {
					result += ", Input" + strconv.Itoa(i)
				}
			}
		}

		for _, oname := range onames {
			for i := 0; i < bmach.Outputs; i++ {
				ooname := Get_output_name(i)
				if ooname == oname {
					result += ", Output" + strconv.Itoa(i)
				}
			}
		}

		result += ");\n"
	}

	if udpbond_module {
		result += "\tudpbond_main udpbond_main (.clk100(" + clk_name + "), .reset(reset), .wifi_enable(wifi_enable), .wifi_rx(wifi_rx), .wifi_tx(wifi_tx)"
		for _, iname := range inames {
			for i := 0; i < bmach.Inputs; i++ {
				iiname := Get_input_name(i)
				if iiname == iname {
					result += ", .input_" + iname + "(Input" + strconv.Itoa(i) + ")"
				}
			}
		}

		for _, oname := range onames {
			for i := 0; i < bmach.Outputs; i++ {
				ooname := Get_output_name(i)
				if ooname == oname {
					result += ", .output_" + oname + "(Output" + strconv.Itoa(i) + ")"
				}
			}
		}

		result += ");\n"
	}

	if basys3_7segment_module {
		result += "\tbond2seg bond2seg_inst(" + clk_name + ", reset, " + basys3_7segment_mapped + ", segment ,enable_D1, enable_D2, enable_D3, enable_D4, dp);\n"
	}

	for _, iores := range resolved_io {

		if iores == "board" {

			result += "\talways @ (posedge clk) begin\n"

			if slow_module {
				result += "\t\tdivider <= divider + 1;\n"
			}

			for i := 0; i < bmach.Inputs; i++ {
				iname := Get_input_name(i)
				tpname := "Input" + strconv.Itoa(i)
				aname := iname
				if rname, ok := iomaps.Assoc[iname]; ok {
					aname = rname
				}
				if aname != iname {
					for j := 0; j < int(bmach.Rsize); j++ {
						result += "\t\t" + tpname + "[" + strconv.Itoa(j) + "] <= " + nth_assoc(aname, j) + ";\n"
					}
				}
			}

			for i := 0; i < bmach.Outputs; i++ {
				oname := Get_output_name(i)
				tpname := "Output" + strconv.Itoa(i)
				aname := oname
				if rname, ok := iomaps.Assoc[oname]; ok {
					aname = rname
				}
				if aname != oname {
					for j := 0; j < int(bmach.Rsize); j++ {
						result += "\t\t" + nth_assoc(aname, j) + " <= " + tpname + "[" + strconv.Itoa(j) + "]" + ";\n"
					}
				}
			}

			result += "\tend\n"

			break
		}
	}

	result += "endmodule\n"
	return result_headers + result
}
