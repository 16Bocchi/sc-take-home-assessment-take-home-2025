package folder

import (
	"errors"
	"strings"
)

/*
Function to move a folder with a given name to a new parent folder with a given name.

The function should return an error if:

- The folder with the given name does not exist within the organization.

- The destination folder does not exist within the organization.

- The organization ID of the source folder does not match the organization ID of the destination folder.

- The source folder and destination folder have the same name.

- The destination folder is a child of the source folder.
*/
func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {

	// Define variables to store source and destination folders
	var srcF *Folder
	var dstF *Folder
	var srcIndex int

	var ErrSrcNotFound = errors.New("folder not found")
	var ErrDstNotFound = errors.New("destination folder not found")
	var ErrOrgIDMismatch = errors.New("organisation ID mismatch")
	var ErrSrcMatchDst = errors.New("source folder and destination folder can not be the same")
	var ErrDstIsChild = errors.New("destination folder is a child of the source folder")

	// Look for source folder
	for index, folder := range f.folders {
		if folder.Name == name {
			srcF = &folder
			srcIndex = index
			break
		}
	}

	// If source folder not found, return error
	if srcF == nil {
		return nil, ErrSrcNotFound
	}

	// Look for destination folder
	for _, folder := range f.folders {
		if folder.Name == dst {
			dstF = &folder
			break
		}
	}

	// Error handling time
	if dstF == nil {
		return nil, ErrDstNotFound
	}

	if srcF.OrgId != dstF.OrgId {
		return nil, ErrOrgIDMismatch
	}

	if srcF.Name == dstF.Name {
		return nil, ErrSrcMatchDst
	}

	if strings.HasPrefix(dstF.Paths, srcF.Paths) {
		return nil, ErrDstIsChild
	}

	// Define new paths
	newPathStr := dstF.Paths + "." + srcF.Name
	oldPathStr := srcF.Paths

	// update source folder path
	f.folders[srcIndex].Paths = newPathStr

	// update child folders path
	for i, folder := range f.folders {
		if strings.HasPrefix(folder.Paths, oldPathStr) && folder.OrgId == srcF.OrgId {
			folder.Paths = strings.Replace(folder.Paths, oldPathStr, newPathStr, 1) // One is here to replace first occurence
			f.folders[i] = folder
		}
	}

	return f.folders, nil
}
