package main

import (
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/sirupsen/logrus"
)

func main() {
	// orgID := uuid.FromStringOrNil("b20c2cfb-04c4-474d-ba67-03a8a7453578")

	res := folder.GetAllFolders()

	// example usage
	folderDriver := folder.NewDriver(res)
	// orgFolder := folderDriver.GetFoldersByOrgID(orgID)
	src := "beta"
	dst := "zeta"
	res, err := folderDriver.MoveFolder(src, dst)
	if err != nil {
		logrus.Error(err)
	}

	// folder.PrettyPrint(res)
	// fmt.Printf("\n Folders for orgID: %s", orgID)
	// folder.PrettyPrint(orgFolder)
	if res != nil {
		logrus.Info("\n Folders after moving folder: ")
		folder.PrettyPrint(res)
	}

	// folder.PrettyPrint(err)

	// component 1
	// Implement the following methods:
	// GetAllChildFolders returns all child folders of a specific folder.

}
