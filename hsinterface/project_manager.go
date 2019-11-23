package hsinterface

import "github.com/OpenDiablo2/HellSpawner/hsutil"

type ProjectManager interface {
	OnFileSelected(filename string, mpqpath hsutil.MpqPath)
	OnViewMpqFileDetails(mpqpath string)
}
