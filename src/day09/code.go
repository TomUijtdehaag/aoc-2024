package day09

import (
	"aoc/utils"
	"fmt"
	"strings"
)

type File struct {
	id    int
	start int
	size  int
}

func (f *File) End() int {
	return f.start + f.size - 1
}

type Disk struct {
	diskMap []int
	blocks  []int
	files   []File
}

func NewDisk(input string) *Disk {
	diskMap := []int{}
	blocks := []int{}
	files := []File{}

	position := 0

	for i, blockSize := range strings.Split(input, "") {
		blockSizeInt := utils.ToInt(blockSize)
		diskMap = append(diskMap, blockSizeInt)
		if i%2 == 0 {
			id := i / 2
			for j := 0; j < blockSizeInt; j++ {
				blocks = append(blocks, id)
			}

			files = append(files, File{id: id, start: position, size: blockSizeInt})
		} else {
			for j := 0; j < blockSizeInt; j++ {
				blocks = append(blocks, -1)
			}
		}
		position += blockSizeInt
	}

	return &Disk{diskMap: diskMap, blocks: blocks, files: files}
}

func (d *Disk) PrintDiskMap() {
	fmt.Println("Disk Map:")
	for _, block := range d.diskMap {
		fmt.Printf("%v ", block)
	}
	fmt.Println()
}

func (d *Disk) PrintBlocks() {
	fmt.Println("Blocks:")
	for _, block := range d.blocks {
		if block == -1 {
			fmt.Print(". ")
		} else {
			fmt.Printf("%v ", block)
		}
	}
	fmt.Println()
}

func (d *Disk) SwapBlocks(from, to int) {
	d.blocks[from], d.blocks[to] = d.blocks[to], d.blocks[from]
}

func (d *Disk) MoveFile(id, to int) {
	file := &d.files[id]

	for i := 0; i < file.size; i++ {
		d.SwapBlocks(file.start+i, to+i)
	}

	file.start = to
}

func (d *Disk) Compact() {
	firstEmptyBlock := 0
	position := len(d.blocks) - 1

	for position > firstEmptyBlock {
		if d.blocks[firstEmptyBlock] != -1 {
			firstEmptyBlock++
			continue
		}

		if d.blocks[position] != -1 {
			d.SwapBlocks(firstEmptyBlock, position)
			firstEmptyBlock++
			position--
			continue
		}
		position--
	}
}

func (d *Disk) FileCompact() {
	for i := len(d.files) - 1; i >= 0; i-- {
		file := d.files[i]
		position := 0

		for position+file.size <= file.start {
			anyOccupied := false

			for block := 0; block < file.size; block++ {
				fileId := d.blocks[position+block]
				if fileId != -1 {
					anyOccupied = true
					position = d.files[fileId].End() + 1
					break
				}
			}

			if !anyOccupied {
				d.MoveFile(i, position)
				break
			}
		}
	}
}

func (d *Disk) GetChecksum() int {
	checksum := 0
	for i, block := range d.blocks {
		switch block {
		case -1:
			continue
		default:
			checksum += i * block
		}
	}
	return checksum
}

func part1(input string) int {
	disk := NewDisk(input)
	disk.Compact()
	return disk.GetChecksum()
}

func part2(input string) int {
	disk := NewDisk(input)
	disk.FileCompact()
	return disk.GetChecksum()
}

func Solve(day int) {
	input := utils.ReadInput(day)

	fmt.Printf("Day %d, Part 1: %v\n", day, part1(input))
	fmt.Printf("Day %d, Part 2: %v\n", day, part2(input))
}
