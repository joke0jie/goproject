package main

import (
	_ "bufio"
	"fmt"
	_ "io"
	"os"
	_ "strconv"
	_ "strings"
)

type Elf32_Section uint16

const EI_NIDENT = 16

type Elf32_Shdr struct {
	sh_name      uint32
	sh_type      uint32
	sh_flags     uint32
	sh_addr      uint32
	sh_offset    uint32
	sh_size      uint32
	sh_link      uint32
	sh_info      uint32
	sh_addralign uint32
	sh_entsize   uint32
}

type Elf32_Sym struct {
	st_name  uint32        /* Symbol name (string tbl index) */
	st_value uint32        /* Symbol value */
	st_size  uint32        /* Symbol size */
	st_info  uint8         /* Symbol type and binding */
	st_other uint8         /* Symbol visibility */
	st_shndx Elf32_Section /* Section index */
}

type Elf32_Chdr struct {
	ch_type      uint32
	ch_size      uint32
	ch_addralign uint32
}

type Elf32_Ehdr struct {
	E_ident     [EI_NIDENT]uint8
	E_type      uint16
	E_machine   uint16
	E_version   uint32
	E_entry     uint32
	E_phoff     uint32
	E_shoff     uint32
	E_flags     uint32
	E_ehsize    uint16
	E_phentsize uint16
	E_phnum     uint16
	E_shentsize uint16
	E_shnum     uint16
	E_shstrndx  uint16
}

type Sqw struct {
	Name string
	Age  int
}

type ExpStruct struct {
	Mi1 int
	Mf1 float32
}

type SoInfo struct {
	Elf_Ehdr Elf32_Ehdr
	Elf_Shdr [50]Elf32_Shdr
	Elf_Sym  Elf32_Sym
	Elf_Chdr Elf32_Chdr
}

func Uint32ToBytes(n uint32) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
		byte(n >> 16),
		byte(n >> 24),
	}
}

func Uint64ToBytes(n uint64) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
		byte(n >> 16),
		byte(n >> 24),
		byte(n >> 32),
		byte(n >> 40),
		byte(n >> 48),
		byte(n >> 56),
	}
}

func Uint16ToBytes(n uint16) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
	}
}

func BytesToUint16(array []byte) uint16 {
	var data uint16 = 0
	for i := 0; i < len(array); i++ {
		data = data + uint16(uint(array[i])<<uint(8*i))
	}

	return data
}

func BytesToUint32(array []byte) uint32 {
	var data uint32 = 0
	for i := 0; i < len(array); i++ {
		data = data + uint32(uint(array[i])<<uint(8*i))
	}

	return data
}

func parseHeader(fp *os.File, si *SoInfo) bool {

	buf := make([]byte, 1024)
	if _, err := fp.Read(buf); err == nil {

		for i := 0; i < 16; i++ {
			si.Elf_Ehdr.E_ident[i] |= buf[i]
		}

		si.Elf_Ehdr.E_type = BytesToUint16(buf[16:18])
		si.Elf_Ehdr.E_machine = BytesToUint16(buf[18:20])
		si.Elf_Ehdr.E_version = BytesToUint32(buf[20:24])

		si.Elf_Ehdr.E_entry = BytesToUint32(buf[24:28])
		si.Elf_Ehdr.E_phoff = BytesToUint32(buf[28:32])
		si.Elf_Ehdr.E_shoff = BytesToUint32(buf[32:36])
		si.Elf_Ehdr.E_flags = BytesToUint32(buf[36:40])

		si.Elf_Ehdr.E_ehsize = BytesToUint16(buf[40:42])
		si.Elf_Ehdr.E_phentsize = BytesToUint16(buf[42:44])
		si.Elf_Ehdr.E_phnum = BytesToUint16(buf[44:46])
		si.Elf_Ehdr.E_shentsize = BytesToUint16(buf[46:48])
		si.Elf_Ehdr.E_shnum = BytesToUint16(buf[48:50])
		si.Elf_Ehdr.E_shstrndx = BytesToUint16(buf[50:52])

		fmt.Printf("%x\n", si.Elf_Ehdr.E_ident)
		fmt.Printf("%x\n", si.Elf_Ehdr.E_type)
		fmt.Printf("%x\n", si.Elf_Ehdr.E_machine)
		fmt.Printf("%x\n", si.Elf_Ehdr.E_version)

		fmt.Printf("%x\n", si.Elf_Ehdr.E_entry)
		fmt.Printf("%x\n", si.Elf_Ehdr.E_phoff)
		fmt.Printf("%x\n", si.Elf_Ehdr.E_shoff)
		fmt.Printf("%x\n", si.Elf_Ehdr.E_flags)

		fmt.Printf("%x\n", si.Elf_Ehdr.E_ehsize)
		fmt.Printf("%x\n", si.Elf_Ehdr.E_phentsize)
		fmt.Printf("%x\n", si.Elf_Ehdr.E_phnum)
		fmt.Printf("%x\n", si.Elf_Ehdr.E_shentsize)
		fmt.Printf("%x\n", si.Elf_Ehdr.E_shnum)
		fmt.Printf("%x\n", si.Elf_Ehdr.E_shstrndx)

	}
	return true
}

func parseSection(fp *os.File, si *SoInfo) bool {
	shoff := si.Elf_Ehdr.E_shoff
	shnum := si.Elf_Ehdr.E_shnum
	//shentsize := si.Elf_Ehdr.E_shentsize

	var shdr Elf32_Shdr

	fmt.Println("shnum :", shnum)
	fp.Seek(int64(shoff), 0)
	fmt.Printf("shoff : 0x%Xh\n", shoff)

	buf := make([]byte, 40)
	for i := 0; i < int(shnum); i++ {
		if _, err := fp.Read(buf); err == nil {

			shdr.sh_name = BytesToUint32(buf[:4])
			shdr.sh_type = BytesToUint32(buf[4:8])
			shdr.sh_flags = BytesToUint32(buf[8:12])
			shdr.sh_addr = BytesToUint32(buf[12:16])
			shdr.sh_offset = BytesToUint32(buf[16:20])

			shdr.sh_size = BytesToUint32(buf[20:24])
			shdr.sh_link = BytesToUint32(buf[24:28])
			shdr.sh_info = BytesToUint32(buf[28:32])
			shdr.sh_addralign = BytesToUint32(buf[32:36])
			shdr.sh_entsize = BytesToUint32(buf[36:40])

			si.Elf_Shdr[i] = shdr

			fmt.Printf("sh_name :          0x%Xh\n", shdr.sh_name)
			fmt.Printf("sh_type :            0x%Xh\n", shdr.sh_type)
			fmt.Printf("sh_flags :           0x%Xh\n", shdr.sh_flags)
			fmt.Printf("sh_addr :            0x%08Xh\n", shdr.sh_addr)
			fmt.Printf("sh_offset :          0x%Xh\n", shdr.sh_offset)

			fmt.Printf("sh_size :             0x%Xh\n", shdr.sh_size)
			fmt.Printf("sh_link :              0x%Xh\n", shdr.sh_link)
			fmt.Printf("sh_info :              0x%Xh\n", shdr.sh_info)
			fmt.Printf("sh_addralign :     0x%Xh\n", shdr.sh_addralign)
			fmt.Printf("sh_entsize :         0x%Xh\n", shdr.sh_entsize)
			fmt.Printf("\n\n")

		}
	}

	return true
}
