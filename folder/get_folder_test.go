package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	useFolders := folder.GetSampleData()
	useOrgID := uuid.FromStringOrNil("b20c2cfb-04c4-474d-ba67-03a8a7453578")
	useDriver := folder.NewDriver(useFolders)

	tests := [...]struct {
		name       string
		orgID      uuid.UUID
		folderName string
		want       []folder.Folder
		err        error
	}{
		{
			name:       "Case 1: Parent with multiple children, valid orgID",
			orgID:      useOrgID,
			folderName: "alpha",
			want: []folder.Folder{
				{Name: "beta", OrgId: useOrgID, Paths: "alpha.beta"},
				{Name: "gamma", OrgId: useOrgID, Paths: "alpha.beta.gamma"},
				{Name: "delta", OrgId: useOrgID, Paths: "alpha.beta.delta"},
				{Name: "epsilon", OrgId: useOrgID, Paths: "alpha.beta.epsilon"},
				{Name: "zeta", OrgId: useOrgID, Paths: "alpha.zeta"},
				{Name: "eta", OrgId: useOrgID, Paths: "alpha.eta"},
			},
			err: nil,
		},
		{
			name:       "Case 2: Parent with no children, valid orgID",
			orgID:      useOrgID,
			folderName: "epsilon",
			want:       []folder.Folder{},
			err:        nil,
		},
		{
			name:       "Case 3: Parent with multiple children, invalid orgID",
			orgID:      uuid.Must(uuid.NewV4()),
			folderName: "alpha",
			want:       nil,
			err:        folder.ErrOrganizationNotFound,
		},
		{
			name:       "Case 4: Empty folder name",
			orgID:      useOrgID,
			folderName: "",
			want:       nil,
			err:        folder.ErrFolderNameEmpty,
		},
		{
			name:       "Case 5: Folder name contains dot",
			orgID:      useOrgID,
			folderName: "alpha.beta",
			want:       nil,
			err:        folder.ErrFolderNameContainsDot,
		},
		{
			name:       "Case 6: Parent doesn't exist, valid orgID",
			orgID:      useOrgID,
			folderName: "omega",
			want:       nil,
			err:        folder.ErrFolderNotFound,
		},
		{
			name:       "Case 7: Parent with single child, valid orgID",
			orgID:      useOrgID,
			folderName: "kappa",
			want: []folder.Folder{
				{Name: "lambda", OrgId: useOrgID, Paths: "theta.kappa.lambda"},
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret, err := useDriver.GetAllChildFolders(tt.orgID, tt.folderName)
			if err != nil {
				if err != tt.err {
					t.Errorf("got error %v, want %v", err, tt.err)
				} else {
					t.Logf("got error %v, want %v", err, tt.err)
				}
			}
			if (ret != nil) && (tt.want != nil) {
				assert.Equal(t, tt.want, ret)
			} else {
				t.Logf("got %v, want %v", ret, tt.want)
			}
			// f := folder.NewDriver(tt.folders)
			// get := f.GetFoldersByOrgID(tt.orgID)

		})
	}
}
