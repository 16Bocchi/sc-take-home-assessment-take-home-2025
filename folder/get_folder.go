package folder

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	// Your code here...
	// get all folders using orgID
	orgFolders := f.GetFoldersByOrgID(orgID)

	// new var for parent folder
	var parentFolder *Folder

	// like python for loop -> loop through orgFolders
	// if find folder name match, assign to parent folder
	for _, folder := range orgFolders {
		if folder.Name == name {
			parentFolder = &folder
			break
		}
	}

	// if no parent folder, return error and nil -> not empty slice
	if parentFolder == nil {
		return nil, errors.New("Folder does not exist within the organization: " + orgID.String())
	}
	// parent folder path is prefix for child folders
	parentPath := parentFolder.Paths + "."

	// new var for child folders
	var childFolders []Folder

	for _, folder := range orgFolders {
		if strings.HasPrefix(folder.Paths, parentPath) {
			childFolders = append(childFolders, folder)
		}
	}

	return childFolders, nil
}
