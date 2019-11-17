package hsutil

import (
	"path/filepath"
	"strings"
)

type FileTreeNode struct {
	Name     string

	IsFile   bool
	FullPath string

	Children []*FileTreeNode
}

func BuildTreeWalk(curnode *FileTreeNode, curpath []string, fullpath string) {
	if len(curpath) == 0 {
		return
	}

	// take the next bit off curpath
	var next string
	next, curpath = curpath[0], curpath[1:]

	// see if next already exists
	for _, node := range curnode.Children {
		if strings.ToLower(node.Name) == strings.ToLower(next) {
			BuildTreeWalk(node, curpath, fullpath) // node already exists, keep walking
			return
		}
	}

	// otherwise, add it
	isfile := len(curpath) == 0
	// find the index to add at
	// this logic ensures that dirs are on top of the list and files are on the bottom
	index := -1
	if !isfile {
		for i, node := range curnode.Children {
			if !isfile && node.IsFile {
				index = i
				break
			}
		}
	}

	newnode := &FileTreeNode{}
	if index == -1 {
		// if index is -1, it's a file or its a dir and we searched the whole list and found no files
		// so append it to the end
		curnode.Children = append(curnode.Children, newnode)
	} else {
		// insert the new node at a specific index
		curnode.Children = append(curnode.Children, nil)
		copy(curnode.Children[index + 1:], curnode.Children[index:])
		curnode.Children[index] = newnode
	}
	newnode.Name = next
	newnode.IsFile = isfile
	newnode.Children = make([]*FileTreeNode, 0) 
	if newnode.IsFile { // if it's a file, stop
		newnode.FullPath = fullpath
	} else { // otherwise, keep walking
		BuildTreeWalk(newnode, curpath, fullpath)
	}
}

func BuildFileTreeFromFileList(paths []string) *FileTreeNode {
	root := &FileTreeNode{}
	root.Name = "root"
	root.Children = make([]*FileTreeNode, 0)
	
	for _, p := range paths {
		pnames := strings.Split(p, string(filepath.Separator))
		BuildTreeWalk(root, pnames, p)
	}

	return root
}