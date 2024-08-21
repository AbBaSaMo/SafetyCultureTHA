package folders

import (
	// "fmt"
	"github.com/gofrs/uuid"
)

type PaginatedFetchFolderResponse struct {
	Folders []*Folder
	// number of folders in Folders
	ChunkSize int
	// given the Limit, the offset from first chunk to obtain the
	// set of folders subsequent to this set
	NextOffset int
}


// TODO
// Returns a chunk of data as defined by 
// func GetAllFoldersPaginated(req *FetchFolderRequest, chunkSize int, offset int) (*PaginatedFetchFolderResponse, error) {
// 	foldersByOrgIdRes, err := FetchAllFoldersByOrgIDPaginated(req.OrgID, chunkSize)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// fetch the chunk we want
// 	PrettyPrint(foldersByOrgIdRes)


// 	// prep return

// 	var fetchFolderResponse *PaginatedFetchFolderResponse = &PaginatedFetchFolderResponse {
// 		Folders: []*Folder{}, 
// 		ChunkSize: chunkSize, 
// 		NextOffset: offset + 1, // need to modulo with size of it all
// 	}
// 	return fetchFolderResponse, nil
// }

// Fetch all folders for a given OrgId and split them into chunks as defined by the chunk size.
// If folders cannot evenly be split by chunk, the remainder at the end is grouped as it's own chunk.
func FetchAllFoldersByOrgIDPaginated(orgID uuid.UUID, chunkSize int) ([][]*Folder, error) {
	folders := GetSampleData()

	// obtain all folders by id
	orgFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			orgFolder = append(orgFolder, folder)
		}
	}

	// split into even chunks
	numChunks := len(orgFolder)/chunkSize
	paginatedFolders := [][]*Folder{}

	for i:=0; i < numChunks; i++ {
		start := i*chunkSize
		end   := start + chunkSize
		chunk := orgFolder[start: end]

		paginatedFolders = append(paginatedFolders, chunk)
	}

	// append any remainder in separate chunk
	rem := len(orgFolder)%chunkSize
	if rem > 0 {
		start := len(orgFolder) - rem
		lastChunk := orgFolder[start:len(orgFolder)]
		paginatedFolders = append(paginatedFolders, lastChunk)
	}

	return paginatedFolders, nil
}
