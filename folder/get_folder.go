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

	// Define errors here! For use later in code
	var ErrFolderNotFound = errors.New("folder not found")
	var ErrOrganizationNotFound = errors.New("organization not found")
	var ErrFolderNameEmpty = errors.New("folder name cannot be empty")
	var ErrFolderNameContainsDot = errors.New("folder name cannot contain '.'")

	// Define variables to store parent and child folders
	childFolders := []Folder{}
	orgFolders := f.GetFoldersByOrgID(orgID)
	// println("orgFolders: ", orgFolders)
	if len(orgFolders) == 0 {
		return nil, ErrOrganizationNotFound
	}

	if name == "" {
		return nil, ErrFolderNameEmpty
	}

	if strings.Count(name, ".") > 0 {
		return nil, ErrFolderNameContainsDot
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
		return nil, ErrFolderNotFound
	}
	// parent folder path is prefix for child folders
	parentPath := parentFolder.Paths + "."

	// loop through orgFolders, if folder path starts with parentPath, append to childFolders
	for _, folder := range orgFolders {
		if strings.HasPrefix(folder.Paths, parentPath) && folder.Paths != parentPath {
			childFolders = append(childFolders, folder)
		}
	}

	return childFolders, nil
}
