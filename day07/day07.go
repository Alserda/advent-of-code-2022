package day07

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/Alserda/advent-of-code-2022/utils"
)

//go:embed input.txt
var input string

type File struct {
	Name string
	Size int
}

type Directory struct {
	Entry    bool
	Name     string
	Parent   *Directory
	Children []*Directory
	Files    []File
}

func Run() {
	fmt.Printf(part1(input))
	fmt.Printf(part2(input))
}

func part1(input string) string {
	dir := parse(input)

	total := dir.sumPart1(0)
	return fmt.Sprintf("Sum of directory sizes: %d", total)
}

func part2(input string) string {
	dir := parse(input)

	diskSpace, requiredSpace, totalSize := 70000000, 30000000, dir.size()
	cleanupSize := requiredSpace - (diskSpace - totalSize)

	candidate := dir.delete(cleanupSize, nil)
	size := candidate.size()

	return fmt.Sprintf("Size of directory to delete: %d", size)
}

func parse(input string) Directory {
	lines := utils.SplitRows(input)
	dir := &Directory{
		Name:  "/",
		Entry: true,
	}

	dir.do(lines)

	return *dir
}

func (dir Directory) delete(requiredSpace int, candidate *Directory) *Directory {
	size := dir.size()

	if size > requiredSpace {
		if candidate == nil || size < candidate.size() {
			candidate = &dir
		}
	}

	for _, child := range dir.Children {
		candidate = child.delete(requiredSpace, candidate)
	}

	return candidate
}

func (dir Directory) sumPart1(sum int) int {
	size := dir.size()

	for _, child := range dir.Children {
		sum = child.sumPart1(sum)
	}

	if size < 100000 {
		sum += size
	}

	return sum
}

func (dir Directory) size() int {
	size := 0

	for _, file := range dir.Files {
		size += file.Size
	}

	for _, child := range dir.Children {
		size += child.size()
	}

	return size
}

func (dir *Directory) do(lines []string) {
	i := 0

	for i < len(lines) {
		parts := strings.Split(lines[i], " ")

		if parts[0] == "$" {
			switch parts[1] {
			case "cd":
				if parts[2] != "/" {
					dir = dir.cd(parts[2])
				}

			case "ls":
				i += dir.ls(lines[i+1:])
			}
		}

		i++
	}
}

func (dir *Directory) cd(arg string) *Directory {
	if arg == ".." {
		return dir.Parent
	}

	for _, child := range dir.Children {
		if child.Name == arg {
			return child
		}
	}

	panic("could not find child to cd in")
}

func (dir *Directory) ls(lines []string) int {
	i := 0
	for _, line := range lines {
		if string(line[0]) == "$" {
			return i
		}

		parts := strings.Split(line, " ")
		if parts[0] == "dir" {
			dir.Children = append(dir.Children, &Directory{
				Parent: dir,
				Name:   parts[1],
			})
		} else {
			dir.Files = append(dir.Files, File{
				Size: utils.ParseInt(parts[0]),
				Name: parts[1],
			})
		}

		i++
	}

	return i
}
