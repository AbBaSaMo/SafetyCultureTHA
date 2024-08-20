package folders_test

import (
	"errors"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"

	"github.com/gofrs/uuid"
)

var tests = []struct {
	req *folders.FetchFolderRequest
	res *folders.FetchFolderResponse
	err error
}{
	{
		req: nil,
		res: nil,
		err: errors.New("empty request"),
	},
	{
		req: &folders.FetchFolderRequest{OrgID: uuid.Nil},
		res: nil,
		err: errors.New("orgID is nil"),
	},
	{
		req: &folders.FetchFolderRequest{OrgID: uuid.FromStringOrNil("5212d622-88ff-468a-862d-ea49fef5e183")},
		res: &folders.FetchFolderResponse{Folders: []*folders.Folder{}},
		err: nil,
	},
	{
		req: &folders.FetchFolderRequest{OrgID: uuid.FromStringOrNil("4212d618-66ff-468a-862d-ea49fef5e183")},
		res: &folders.FetchFolderResponse{Folders: []*folders.Folder{
			{
				Id:      uuid.FromStringOrNil("1167c1ac-911b-4a1f-b460-a98f724c7289"),
				Name:    "heroic-bella",
				OrgId:   uuid.FromStringOrNil("4212d618-66ff-468a-862d-ea49fef5e183"),
				Deleted: true,
			},
		}},
		err: nil,
	},
}

func Test_GetAllFolders(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		// your test/s here

	})
}
