package main

import (
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	res := folder.GetAllFolders()

	// example usage
	folderDriver := folder.NewDriver(res)
	// orgFolder := folderDriver.GetFoldersByOrgID(orgID)
	name := "settling-hobgoblin"
	childFolders, err := folderDriver.GetAllChildFolders(orgID, name)
	if err != nil {
		logrus.Error(err)
	}

	// folder.PrettyPrint(res)
	// fmt.Printf("\n Folders for orgID: %s", orgID)
	// folder.PrettyPrint(orgFolder)
	if childFolders != nil {
		logrus.Info("\n Child folders for orgID and name: ", orgID, name)
		folder.PrettyPrint(childFolders)
	}

	// folder.PrettyPrint(err)

	// component 1
	// Implement the following methods:
	// GetAllChildFolders returns all child folders of a specific folder.

}
