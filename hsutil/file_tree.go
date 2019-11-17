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
	newnode := &FileTreeNode{}
	curnode.Children = append(curnode.Children, newnode)
	newnode.Name = next
	newnode.IsFile = len(curpath) == 0
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