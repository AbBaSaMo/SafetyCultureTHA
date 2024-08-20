package folders

import (
	"github.com/gofrs/uuid"
)

// Returns a FetchFolderResponse containing all folders owned by the organisation
// specified in the request regardless of delete status.
//
// TODO:
//
//		[ ] vars 'err', 'f1', 'fs' are declared but unused, these are 'redeclared' as f, fp and err is not handled
//		[ ] 'k' and 'k1' unused in for loops -> remove definition
//		[ ] does not handle errors that could be returned by func FetchAllFoldersByOrgID
//	    [x] add white sppace (\n) between sections of the function for readability
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	var (
		err error
		f1  Folder
		fs  []*Folder
	)

	f := []Folder{}
	r, _ := FetchAllFoldersByOrgID(req.OrgID)
	for k, v := range r {
		f = append(f, *v)
	}

	var fp []*Folder
	for k1, v1 := range f {
		fp = append(fp, &v1)
	}

	var ffr *FetchFolderResponse
	ffr = &FetchFolderResponse{Folders: fp}
	return ffr, nil
}

// FetchAllFoldersByOrdId filters a list of folders by the organisation
// who owns them and returns an array of Folders owned by such organisation
// and information regarding any potential errors.
//
// List of folders is obtained from sample.json through GetSampleData call.
//
// TODO:
//
//	[ ] Does not appear to handle errors and returns 'nil' in all cases.
//	[ ] Look into common practices when returning errors
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
