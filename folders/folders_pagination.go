package folders

import (
	"github.com/gofrs/uuid"
)

// TODO
func GetAllFoldersPaginated(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	folders := []Folder{}
	foldersByOrgIdRes, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err
	}

	for _, f := range foldersByOrgIdRes {
		folders = append(folders, *f)
	}

	var foldersPointers []*Folder
	for _, f := range folders {
		folder := f
		foldersPointers = append(foldersPointers, &folder)
	}

	var fetchFolderResponse *FetchFolderResponse = &FetchFolderResponse{Folders: foldersPointers}
	return fetchFolderResponse, nil
}

// TODO
func FetchAllFoldersByOrgIDPaginated(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
