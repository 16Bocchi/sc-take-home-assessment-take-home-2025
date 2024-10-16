package folder

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

// Function to retrieve all folders within an organization with a given ID.
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

// Function to retrieve all child folders of a folder with a given name.
// The function should return an error if the folder with the given name does not exist within the organization.
// Function signature has been modified to return an error.
func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {

	childFolders := []Folder{}
	orgFolders := f.GetFoldersByOrgID(orgID)
	// println("orgFolders: ", orgFolders)
	if len(orgFolders) == 0 {
		return nil, errors.New("Organization does not exist: " + orgID.String())
	}

	if name == "" {
		return nil, errors.New("Folder name cannot be empty")
	}

	if strings.Count(name, ".") > 0 {
		return nil, errors.New("Folder name cannot contain '.'")
	}

	// new var for parent folder
	var parentFolder Folder

	// like python for loop -> loop through orgFolders
	// if find folder name match, assign to parent folder
	for _, folder := range orgFolders {
		if folder.Name == name {
			parentFolder = folder
			break
		}
	}

	// if no parent folder, return error and nil -> not empty slice
	if parentFolder.Name == "" {
		return nil, errors.New("Folder does not exist within the organization: " + orgID.String())
	}
	// parent folder path is prefix for child folders
	parentPath := parentFolder.Paths + "."

	for _, folder := range orgFolders {
		if strings.HasPrefix(folder.Paths, parentPath) && folder.Paths != parentPath {
			childFolders = append(childFolders, folder)
		}
	}

	return childFolders, nil
}
