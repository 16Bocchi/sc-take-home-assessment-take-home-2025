package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	res := folder.GetAllFolders()

	// example usage
	folderDriver := folder.NewDriver(res)
	orgFolder := folderDriver.GetFoldersByOrgID(orgID)
	childFolders, err := folderDriver.GetAllChildFolders(orgID, "stunning-horridus")

	folder.PrettyPrint(res)
	fmt.Printf("\n Folders for orgID: %s", orgID)
	folder.PrettyPrint(orgFolder)
	folder.PrettyPrint(childFolders)
	folder.PrettyPrint(err)

	// component 1
	// Implement the following methods:
	// GetAllChildFolders returns all child folders of a specific folder.

}
