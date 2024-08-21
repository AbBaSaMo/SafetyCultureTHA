package folders

import (
	"errors"
	"strconv"
	"github.com/gofrs/uuid"
)

type PaginatedFetchFolderResponse struct {
	Folders []*Folder
	// number of folders in the Folders field
	ChunkSize string
	// given the ChunkSize, the offset from first chunk to obtain the
	// set of folders subsequent to this current one
	NextOffset string
}

// Returns chunkSize number of folders from the list of an organisations folders given an
// offset from the start of the list where offset starts from 0.
func GetFoldersByPage(req *FetchFolderRequest, chunkSize int, offset int) (*PaginatedFetchFolderResponse, error) {
	foldersByOrgIdRes, err := FetchAllOrgIDFoldersPaginated(req.OrgID, &chunkSize)
	if err != nil {
		return nil, err
	}

	if offset > len(foldersByOrgIdRes) - 1 {
		return nil, errors.New("offset exceeds number of pages")
	}

	var foldersPointers []*Folder
	for _, f := range foldersByOrgIdRes[offset] {
		folder := f
		foldersPointers = append(foldersPointers, folder)
	}

	var nextOffset string
	if offset < len(foldersByOrgIdRes)-1 {
		nextOffset = strconv.Itoa(offset + 1)
	}

	var fetchFolderResponse *PaginatedFetchFolderResponse = &PaginatedFetchFolderResponse {
		Folders: foldersPointers, 
		ChunkSize: strconv.Itoa(chunkSize), 
		NextOffset: nextOffset,
	}
	return fetchFolderResponse, nil
}

// Fetch all folders for a given OrgId and split them into chunks as defined by the chunk size.
// If folders cannot evenly be split by chunk, the remainder at the end is grouped as it's own chunk.
func FetchAllOrgIDFoldersPaginated(orgID uuid.UUID, chunkSize *int) ([][]*Folder, error) {
	allFolders := GetSampleData()

	orgFolder := []*Folder{}
	for _, folder := range allFolders {
		if folder.OrgId == orgID {
			orgFolder = append(orgFolder, folder)
		}
	}

	if *chunkSize > len(orgFolder) {
		*chunkSize = len(orgFolder)
	}

	// split into even chunks
	numChunks := len(orgFolder)/(*chunkSize)
	paginatedFolders := [][]*Folder{}

	for i:=0; i < numChunks; i++ {
		start := i*(*chunkSize)
		end   := start + (*chunkSize)
		chunk := orgFolder[start: end]

		paginatedFolders = append(paginatedFolders, chunk)
	}

	// append any remainder in separate chunk
	rem := len(orgFolder)%(*chunkSize)
	if rem > 0 {
		start := len(orgFolder) - rem
		lastChunk := orgFolder[start:len(orgFolder)]
		paginatedFolders = append(paginatedFolders, lastChunk)
	}

	return paginatedFolders, nil
}
