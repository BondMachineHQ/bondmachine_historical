package procbuilder

// TODO This is the ROM, change it to halndle also the RAM case

import (
	"math/rand"
	"strconv"
)

type J struct{}

func (op J) Op_get_name() string {
	return "j"
}

func (op J) Op_get_desc() string {
	return "Jump to a program location"
}

func (op J) Op_show_assembler(arch *Arch) string {
	opbits := arch.Opcodes_bits()
	result := "j [" + strconv.Itoa(int(arch.O)) + "(Location)]	// Jump to a program location [" + strconv.Itoa(opbits+int(arch.O)) + "]\n"
	return result
}

func (op J) Op_get_instruction_len(arch *Arch) int {
	opbits := arch.Opcodes_bits()
	return opbits + int(arch.O) // The bits for the opcode + bits for a location
}

func (op J) Op_instruction_verilog_header(conf *Config, arch *Arch, flavor string) string {
	return ""
}

func (Op J) Op_instruction_verilog_reset(arch *Arch, flavor string) string {
	return ""
}

func (Op J) Op_instruction_verilog_internal_state(arch *Arch, flavor string) string {
	return ""
}

func (Op J) Op_instruction_verilog_default_state(arch *Arch, flavor string) string {
	return ""
}

func (op J) Op_instruction_verilog_state_machine(arch *Arch, flavor string) string {
	rom_word := arch.Max_word()
	opbits := arch.Opcodes_bits()

	result := ""
	result += "					J: begin\n"
	if arch.O == 1 {
		result += "						_pc <= #1 rom_value[" + strconv.Itoa(rom_word-opbits-1) + "];\n"
		result += "						$display(\"J \", rom_value[" + strconv.Itoa(rom_word-opbits-1) + "]);\n"
	} else {
		result += "						_pc <= #1 rom_value[" + strconv.Itoa(rom_word-opbits-1) + ":" + strconv.Itoa(rom_word-opbits-int(arch.O)) + "];\n"
		result += "						$display(\"J \", rom_value[" + strconv.Itoa(rom_word-opbits-1) + ":" + strconv.Itoa(rom_word-opbits-int(arch.O)) + "]);\n"
	}
	result += "					end\n"
	return result
}

func (op J) Op_instruction_verilog_footer(arch *Arch, flavor string) string {
	return ""
}

func (op J) Assembler(arch *Arch, words []string) (string, error) {
	opbits := arch.Opcodes_bits()
	rom_word := arch.Max_word()

	if len(words) != 1 {
		return "", Prerror{"Wrong arguments number"}
	}

	result := ""
	if partial, err := Process_number(words[0]); err == nil {
		result += zeros_prefix(int(arch.O), partial)
	} else {
		return "", Prerror{err.Error()}
	}

	for i := opbits + int(arch.O); i < rom_word; i++ {
		result += "0"
	}

	return result, nil
}

func (op J) Disassembler(arch *Arch, instr string) (string, error) {
	value := get_id(instr[:arch.O])
	result := strconv.Itoa(value)
	return result, nil
}

func (op J) Simulate(vm *VM, instr string) error {
	value := get_id(instr[:vm.Mach.O])
	if value < len(vm.Mach.Slocs) {
		vm.Pc = uint64(value)
	} else {
		vm.Pc = vm.Pc + 1
	}
	return nil
}

func (op J) Generate(arch *Arch) string {
	max_value := 1 << arch.O
	value := rand.Intn(max_value)
	return zeros_prefix(int(arch.O), get_binary(value))
}

func (op J) Required_shared() (bool, []string) {
	return false, []string{}
}

func (op J) Required_modes() (bool, []string) {
	return false, []string{}
}

func (op J) Forbidden_modes() (bool, []string) {
	return false, []string{}
}

func (Op J) Op_instruction_verilog_extra_modules(arch *Arch, flavor string) ([]string, []string) {
	return []string{}, []string{}
}

func (Op J) Abstract_Assembler(arch *Arch, words []string) ([]UsageNotify, error) {
	result := make([]UsageNotify, 1)
	newnot := UsageNotify{C_OPCODE, "j", I_NIL}
	result[0] = newnot
	return result, nil
}

func (Op J) Op_instruction_verilog_extra_block(arch *Arch, flavor string, level uint8, blockname string, objects []string) string {
	result := ""
	switch blockname {
	default:
		result = ""
	}
	return result
}
