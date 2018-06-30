package procbuilder

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
)

type Clr struct{}

func (op Clr) Op_get_name() string {
	return "clr"
}

func (op Clr) Op_get_desc() string {
	return "Clear register"
}

func (op Clr) Op_show_assembler(arch *Arch) string {
	opbits := arch.Opcodes_bits()
	result := "clr [" + strconv.Itoa(int(arch.R)) + "(Reg)]	// Set a register to 0 [" + strconv.Itoa(opbits+int(arch.R)) + "]\n"
	return result
}

func (op Clr) Op_get_instruction_len(arch *Arch) int {
	opbits := arch.Opcodes_bits()
	return opbits + int(arch.R) // The bits for the opcode + bits for a register
}

func (op Clr) Op_instruction_verilog_header(conf *Config, arch *Arch, flavor string) string {
	return ""
}

func (Op Clr) Op_instruction_verilog_reset(arch *Arch, flavor string) string {
	return ""
}
func (Op Clr) Op_instruction_verilog_internal_state(arch *Arch, flavor string) string {
	return ""
}

func (Op Clr) Op_instruction_verilog_default_state(arch *Arch, flavor string) string {
	return ""
}

func (op Clr) Op_instruction_verilog_state_machine(arch *Arch, flavor string) string {
	rom_word := arch.Max_word()
	opbits := arch.Opcodes_bits()

	reg_num := 1 << arch.R

	result := ""
	result += "					CLR: begin\n"
	if arch.R == 1 {
		result += "						case (rom_value[" + strconv.Itoa(rom_word-opbits-1) + "])\n"
	} else {
		result += "						case (rom_value[" + strconv.Itoa(rom_word-opbits-1) + ":" + strconv.Itoa(rom_word-opbits-int(arch.R)) + "])\n"
	}
	for i := 0; i < reg_num; i++ {
		result += "						" + strings.ToUpper(Get_register_name(i)) + " : begin\n"
		result += "							_" + strings.ToLower(Get_register_name(i)) + " <= #1 'b0;\n"
		result += "							$display(\"CLR " + strings.ToUpper(Get_register_name(i)) + "\");\n"
		result += "						end\n"
	}
	result += "						endcase\n"
	result += "						_pc <= #1 _pc + 1'b1 ;\n"
	result += "					end\n"
	return result
}

func (op Clr) Op_instruction_verilog_footer(arch *Arch, flavor string) string {
	return ""
}

func (op Clr) Assembler(arch *Arch, words []string) (string, error) {
	opbits := arch.Opcodes_bits()
	rom_word := arch.Max_word()

	reg_num := 1 << arch.R

	if len(words) != 1 {
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

	for i := opbits + int(arch.R); i < rom_word; i++ {
		result += "0"
	}

	return result, nil
}

func (op Clr) Disassembler(arch *Arch, instr string) (string, error) {
	reg_id := get_id(instr[:arch.R])
	result := strings.ToLower(Get_register_name(reg_id)) + " "
	return result, nil
}

func (op Clr) Simulate(vm *VM, instr string) error {
	reg_bits := vm.Mach.R
	reg := get_id(instr[:reg_bits])
	vm.Registers[reg] = 0
	vm.Pc = vm.Pc + 1
	return nil
}

func (op Clr) Generate(arch *Arch) string {
	reg_num := 1 << arch.R
	reg := rand.Intn(reg_num)
	return zeros_prefix(int(arch.R), get_binary(reg))
}

func (op Clr) Required_shared() (bool, []string) {
	return false, []string{}
}

func (op Clr) Required_modes() (bool, []string) {
	return false, []string{}
}

func (op Clr) Forbidden_modes() (bool, []string) {
	return false, []string{}
}

func (Op Clr) Op_instruction_verilog_extra_modules(arch *Arch, flavor string) ([]string, []string) {
	return []string{}, []string{}
}

func (Op Clr) Abstract_Assembler(arch *Arch, words []string) ([]UsageNotify, error) {
	seq, types := Sequence_to_0(words[0])
	if len(seq) > 0 && types == O_REGISTER {

		result := make([]UsageNotify, 2)
		newnot0 := UsageNotify{C_OPCODE, "clr", I_NIL}
		result[0] = newnot0
		newnot1 := UsageNotify{C_REGSIZE, S_NIL, len(seq)}
		result[1] = newnot1

		return result, nil
	}

	return []UsageNotify{}, errors.New("Wrong register")

}

func (Op Clr) Op_instruction_verilog_extra_block(arch *Arch, flavor string, level uint8, blockname string, objects []string) string {
	result := ""
	switch blockname {
	default:
		result = ""
	}
	return result
}
