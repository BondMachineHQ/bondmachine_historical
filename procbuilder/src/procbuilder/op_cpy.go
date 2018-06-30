package procbuilder

import (
	"strconv"
	"strings"
)

type Cpy struct{}

func (op Cpy) Op_get_name() string {
	return "cpy"
}

func (op Cpy) Op_get_desc() string {
	return "Copy from a register to another"
}

func (op Cpy) Op_show_assembler(arch *Arch) string {
	opbits := arch.Opcodes_bits()
	result := "chc [" + strconv.Itoa(int(arch.R)) + "(Reg)] [" + strconv.Itoa(int(arch.R)) + "(Reg)]	// Copy the value of a register to another [" + strconv.Itoa(opbits+int(arch.R)+int(arch.R)) + "]\n"
	return result
}

func (op Cpy) Op_get_instruction_len(arch *Arch) int {
	opbits := arch.Opcodes_bits()
	return opbits + int(arch.R) + int(arch.R) // The bits for the opcode + bits for a register + bits for another register
}

func (op Cpy) Op_instruction_verilog_header(conf *Config, arch *Arch, flavor string) string {
	return ""
}

func (op Cpy) Op_instruction_verilog_state_machine(arch *Arch, flavor string) string {
	rom_word := arch.Max_word()
	opbits := arch.Opcodes_bits()

	reg_num := 1 << arch.R

	result := ""
	result += "					CPY: begin\n"
	if arch.R == 1 {
		result += "						case (rom_value[" + strconv.Itoa(rom_word-opbits-1) + "])\n"
	} else {
		result += "						case (rom_value[" + strconv.Itoa(rom_word-opbits-1) + ":" + strconv.Itoa(rom_word-opbits-int(arch.R)) + "])\n"
	}
	for i := 0; i < reg_num; i++ {
		result += "						" + strings.ToUpper(Get_register_name(i)) + " : begin\n"

		if arch.R == 1 {
			result += "							case (rom_value[" + strconv.Itoa(rom_word-opbits-int(arch.R)-1) + "])\n"
		} else {
			result += "							case (rom_value[" + strconv.Itoa(rom_word-opbits-int(arch.R)-1) + ":" + strconv.Itoa(rom_word-opbits-int(arch.R)-int(arch.R)) + "])\n"
		}

		for j := 0; j < reg_num; j++ {
			result += "							" + strings.ToUpper(Get_register_name(j)) + " : begin\n"
			result += "								_" + strings.ToLower(Get_register_name(i)) + " <= #1 _" + strings.ToLower(Get_register_name(j)) + ";\n"
			result += "								$display(\"CPY " + strings.ToUpper(Get_register_name(i)) + " " + strings.ToUpper(Get_register_name(j)) + "\");\n"
			result += "							end\n"

		}
		result += "							endcase\n"
		result += "						end\n"
	}
	result += "						endcase\n"
	result += "						_pc <= #1 _pc + 1'b1 ;\n"
	result += "					end\n"
	return result
}

func (op Cpy) Op_instruction_verilog_footer(arch *Arch, flavor string) string {
	return ""
}

func (op Cpy) Assembler(arch *Arch, words []string) (string, error) {
	opbits := arch.Opcodes_bits()
	rom_word := arch.Max_word()

	reg_num := 1 << arch.R

	if len(words) != 2 {
		return "", Prerror{"Wrong arguments number"}
	}

	result := ""
	for i := 0; i < reg_num; i++ {
		if words[0] == strings.ToLower(Get_register_name(i)) {
			result += zeros_prefix(int(arch.R), get_binary(i))
			break
		}
	}

	if result == "" {
		return "", Prerror{"Unknown register name " + words[0]}
	}

	partial := ""
	for i := 0; i < reg_num; i++ {
		if words[1] == strings.ToLower(Get_register_name(i)) {
			partial += zeros_prefix(int(arch.R), get_binary(i))
			break
		}
	}

	if partial == "" {
		return "", Prerror{"Unknown register name " + words[1]}
	}

	result += partial

	for i := opbits + 2*int(arch.R); i < rom_word; i++ {
		result += "0"
	}

	return result, nil
}

func (op Cpy) Disassembler(arch *Arch, instr string) (string, error) {
	reg_id := get_id(instr[:arch.R])
	result := strings.ToLower(Get_register_name(reg_id)) + " "
	reg_id = get_id(instr[arch.R : 2*int(arch.R)])
	result += strings.ToLower(Get_register_name(reg_id))
	return result, nil
}

// The simulation does nothing
func (op Cpy) Simulate(vm *VM, instr string) error {
	// TODO
	vm.Pc = vm.Pc + 1
	return nil
}

// The random genaration does nothing
func (op Cpy) Generate(arch *Arch) string {
	// TODO
	return ""
}

func (op Cpy) Required_shared() (bool, []string) {
	return false, []string{}
}

func (op Cpy) Required_modes() (bool, []string) {
	return false, []string{}
}

func (op Cpy) Forbidden_modes() (bool, []string) {
	return false, []string{}
}

func (op Cpy) Op_instruction_internal_state(arch *Arch, flavor string) string {
	return ""
}

func (Op Cpy) Op_instruction_verilog_reset(arch *Arch, flavor string) string {
	return ""
}

func (Op Cpy) Op_instruction_verilog_default_state(arch *Arch, flavor string) string {
	return ""
}

func (Op Cpy) Op_instruction_verilog_internal_state(arch *Arch, flavor string) string {
	return ""
}

func (Op Cpy) Op_instruction_verilog_extra_modules(arch *Arch, flavor string) ([]string, []string) {
	return []string{}, []string{}
}

func (Op Cpy) Abstract_Assembler(arch *Arch, words []string) ([]UsageNotify, error) {
	result := make([]UsageNotify, 0)
	return result, nil
}

func (Op Cpy) Op_instruction_verilog_extra_block(arch *Arch, flavor string, level uint8, blockname string, objects []string) string {
	result := ""
	switch blockname {
	default:
		result = ""
	}
	return result
}
