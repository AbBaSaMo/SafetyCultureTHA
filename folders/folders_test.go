package folders_test

import (
	"testing"
)

func Test_GetAllFolders(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		// your test/s here
		var tests = []struct {
			req *FetchFolderRequest
			res *FetchFolderResponse
		}{
			// nil request
			nil,
		}

		for _, test := range tests {
			if GetAllFolders(test.req) != test.res {

			}

		}

	})
}

// nil *FetchFolderRequest
// empty/nil uuid
// non-sensical uuid

// uuid not found in the list
// uuid found in the list
/*
{
	"id": "1167c1ac-911b-4a1f-b460-a98f724c7289",
	"name": "heroic-bella",
	"org_id": "4212d618-66ff-468a-862d-ea49fef5e183",
	"deleted": true
}
*/
