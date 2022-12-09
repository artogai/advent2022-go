package day7

import (
	"advent2022/day7/cmd"
	"advent2022/day7/fs"
	"math"
)

func SizeDirs(filename string, atMostSize int) int {
	cmds := cmd.Read(filename)
	fs := buildFs(cmds)
	dirs := fs.SubDirsRec()

	size := 0
	for _, dir := range dirs {
		dirSize := dir.Size()
		if dirSize <= atMostSize {
			size += dirSize
		}
	}
	return size
}

func FindMinDeleteDirSize(filename string, maxFsSize int, updateSize int) int {
	cmds := cmd.Read(filename)
	fs := buildFs(cmds)
	dirs := fs.SubDirsRec()
	freeSize := maxFsSize - fs.Size()
	requiredSize := updateSize - freeSize

	minSize := math.MaxInt
	for _, dir := range dirs {
		dirSize := dir.Size()
		if dirSize >= requiredSize && dirSize < minSize {
			minSize = dirSize
		}
	}

	return minSize
}

func buildFs(cmds []cmd.Cmd) *fs.Directory {
	baseDir := fs.NewDirectory("/", nil)
	currDir := baseDir

	for _, c := range cmds {
		switch ct := c.(type) {
		case cmd.Cd:
			switch ct.Path {
			case "/":
				currDir = baseDir
			case "..":
				currDir = currDir.Parent()
			default:
				currDir = currDir.FindSubDir(ct.Path)
			}
		case cmd.Ls:
			for _, entity := range ct.Content {
				entity.SetParent(currDir)
			}
			currDir.Content = ct.Content

		}
	}
	return baseDir
}
