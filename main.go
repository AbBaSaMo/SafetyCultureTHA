package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
)

// obtains and print the default organisation's folders
// if any errors are encointered, logs them
func main() {
	req := &folders.FetchFolderRequest{
		OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
	}

	res, err := folders.GetAllFolders(req)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	folders.PrettyPrint(res)
}
