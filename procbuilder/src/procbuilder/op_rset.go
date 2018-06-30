package procbuilder

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Rset struct{}

func (op Rset) Op_get_name() string {
	return "rset"
}

func (op Rset) Op_get_desc() string {
	return "Register set value"
}

func (op Rset) Op_show_assembler(arch *Arch) string {
	opbits := arch.Opcodes_bits()
	result := "rset [" + strconv.Itoa(int(arch.R)) + "(Reg)] [" + strconv.Itoa(int(arch.Rsize)) + "(Value)]	// Set a register to the given value [" + strconv.Itoa(opbits+int(arch.R)+int(arch.Rsize)) + "]\n"
	return result
}

func (op Rset) Op_get_instruction_len(arch *Arch) int {
	opbits := arch.Opcodes_bits()
	return opbits + int(arch.R) + int(arch.Rsize) // The bits for the opcode + bits for a register + register size bit
}

func (op Rset) Op_instruction_verilog_header(conf *Config, arch *Arch, flavor string) string {
	return ""
}

func (op Rset) Op_instruction_internal_state(arch *Arch, flavor string) string {
	return ""
}

func (op Rset) Op_instruction_verilog_state_machine(arch *Arch, flavor string) string {
	rom_word := arch.Max_word()
	opbits := arch.Opcodes_bits()

	reg_num := 1 << arch.R

	result := ""
	result += "					RSET: begin\n"
	if arch.R == 1 {
		result += "						case (rom_value[" + strconv.Itoa(rom_word-opbits-1) + "])\n"
	} else {
		result += "						case (rom_value[" + strconv.Itoa(rom_word-opbits-1) + ":" + strconv.Itoa(rom_word-opbits-int(arch.R)) + "])\n"
	}
	for i := 0; i < reg_num; i++ {
		result += "						" + strings.ToUpper(Get_register_name(i)) + " : begin\n"
		result += "							_" + strings.ToLower(Get_register_name(i)) + " <= #1 rom_value[" + strconv.Itoa(rom_word-opbits-1-int(arch.R)) + ":" + strconv.Itoa(rom_word-opbits-int(arch.Rsize)-int(arch.R)) + "];\n"
		result += "							$display(\"RSET " + strings.ToUpper(Get_register_name(i)) + " \",_" + strings.ToLower(Get_register_name(i)) + ");\n"
		result += "						end\n"
	}
	result += "						endcase\n"
	result += "						_pc <= #1 _pc + 1'b1 ;\n"
	result += "					end\n"
	return result
}

func (op Rset) Op_instruction_verilog_footer(arch *Arch, flavor string) string {
	return ""
}

func (op Rset) Assembler(arch *Arch, words []string) (string, error) {
	opbits := arch.Opcodes_bits()
	rom_word := arch.Max_word()
	rsize := int(arch.Rsize)

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

	if partial, err := Process_number(words[1]); err == nil {
		result += zeros_prefix(rsize, partial)
	} else {
		return "", Prerror{err.Error()}
	}

	for i := opbits + int(arch.R) + rsize; i < rom_word; i++ {
		result += "0"
	}

	return result, nil
}

func (op Rset) Disassembler(arch *Arch, instr string) (string, error) {
	rsize := int(arch.Rsize)
	reg_id := get_id(instr[:arch.R])
	result := strings.ToLower(Get_register_name(reg_id)) + " "
	value := get_id(instr[arch.R : int(arch.R)+rsize])
	result += strconv.Itoa(value)
	return result, nil
}

func (op Rset) Simulate(vm *VM, instr string) error {
	reg_bits := vm.Mach.R
	reg := get_id(instr[:reg_bits])
	memval := get_id(instr[reg_bits : reg_bits+vm.Mach.Rsize])
	switch vm.Mach.Rsize {
	case 8:
		vm.Registers[reg] = uint8(memval)
	case 16:
		vm.Registers[reg] = uint16(memval)
	default:
		// TODO Fix
	}
	vm.Pc = vm.Pc + 1
	return nil
}

func (op Rset) Generate(arch *Arch) string {
	reg_num := 1 << arch.R
	reg := rand.Intn(reg_num)
	value := rand.Intn(1 << arch.Rsize)
	fmt.Println(reg, value)
	return zeros_prefix(int(arch.R), get_binary(reg)) + zeros_prefix(int(arch.Rsize), get_binary(value))
}

func (op Rset) Required_shared() (bool, []string) {
	return false, []string{}
}

func (op Rset) Required_modes() (bool, []string) {
	return false, []string{}
}

func (op Rset) Forbidden_modes() (bool, []string) {
	return false, []string{}
}

func (Op Rset) Op_instruction_verilog_reset(arch *Arch, flavor string) string {
	return ""
}

func (Op Rset) Op_instruction_verilog_default_state(arch *Arch, flavor string) string {
	return ""
}

func (Op Rset) Op_instruction_verilog_internal_state(arch *Arch, flavor string) string {
	return ""
}

func (Op Rset) Op_instruction_verilog_extra_modules(arch *Arch, flavor string) ([]string, []string) {
	return []string{}, []string{}
}

func (Op Rset) Abstract_Assembler(arch *Arch, words []string) ([]UsageNotify, error) {
	seq, types := Sequence_to_0(words[0])
	if len(seq) > 0 && types == O_REGISTER {

		result := make([]UsageNotify, 2)
		newnot0 := UsageNotify{C_OPCODE, "rset", I_NIL}
		result[0] = newnot0
		newnot1 := UsageNotify{C_REGSIZE, S_NIL, len(seq)}
		result[1] = newnot1

		return result, nil
	}

	return []UsageNotify{}, errors.New("Wrong register")
}

func (Op Rset) Op_instruction_verilog_extra_block(arch *Arch, flavor string, level uint8, blockname string, objects []string) string {
	result := ""
	switch blockname {
	default:
		result = ""
	}
	return result
}
