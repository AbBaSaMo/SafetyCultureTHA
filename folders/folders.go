package folders

import (
	"errors"

	"github.com/gofrs/uuid"
)

// Returns a FetchFolderResponse containing all folders owned by the organisation
// specified in the request regardless of delete status.
//
// TODO:
//
// [x] vars 'err', 'f1', 'fs' are declared but unused, remove the latter 2
// [x] 'k' and 'k1' unused in for loops -> remove
// [x] does not handle errors that could be returned by func FetchAllFoldersByOrgID
// [x] add white sppace (\n) between sections of the function for readability
// [x] merge variable declarations with assignments
// [x] make var names more descriptive
// [x] 2nd for loop causing the same folder to be appended: address to v1 appended, but v1 updates so pointer does too
//
//	https://medium.com/swlh/use-pointer-of-for-range-loop-variable-in-go-3d3481f7ffc9
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	if req == nil {
		return nil, errors.New("empty request")
	}

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

// FetchAllFoldersByOrdId filters a list of folders by the organisation
// who owns them and returns an array of Folders owned by such organisation
// and information regarding any potential errors.
//
// List of folders is obtained from sample.json through GetSampleData call.
//
// TODO:
//
// [x] handle nil orgID
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	if orgID == uuid.Nil {
		return nil, errors.New("orgID is nil")
	}
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
