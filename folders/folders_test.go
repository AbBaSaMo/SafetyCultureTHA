package folders

import (
	"errors"
	"testing"

	"github.com/gofrs/uuid"
)

// res: &FetchFolderResponse{Folders: []*Folder{
// 	{
// 		Id:      uuid.FromStringOrNil("1167c1ac-911b-4a1f-b460-a98f724c7289"),
// 		Name:    "heroic-bella",
// 		OrgId:   uuid.FromStringOrNil("4212d618-66ff-468a-862d-ea49fef5e183"),
// 		Deleted: true,
// 	},
// }},

func Test_GetAllFolders(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		// nil input
		res, err := GetAllFolders(nil)
		if res != nil || err == nil {
			t.Errorf("got res %v err %v expected res %v err %v", res, err, nil, errors.New("empty request"))
		}

		// input with nil uuid
		res, err = GetAllFolders(&FetchFolderRequest{OrgID: uuid.Nil})
		if res != nil || err == nil {
			t.Errorf("got res %v err %v expected res %v err %v", res, err, nil, errors.New("orgID is nil"))
		}

		// non-existant org uuid
		res, err = GetAllFolders(&FetchFolderRequest{OrgID: uuid.FromStringOrNil("5212d622-88ff-468a-862d-ea49fef5e183")})
		if len(res.Folders) != 0 || err != nil {
			t.Errorf("got res %v err %v expected res %v err %v", *res, err, FetchFolderResponse{Folders: []*Folder{}}, nil)
		}

		// existant org uuid
		res, err = GetAllFolders(&FetchFolderRequest{OrgID: uuid.FromStringOrNil("4212d618-66ff-468a-862d-ea49fef5e183")})
		if err != nil {
			t.Errorf("got res %v err %v expected res %v err %v", *res, err, FetchFolderResponse{Folders: []*Folder{}}, nil)
		}
	})
}
