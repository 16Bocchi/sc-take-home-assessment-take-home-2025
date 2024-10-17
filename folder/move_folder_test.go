package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// Custom function to copy folders.

// This function was created because the tests were modifying the folders in the driver.
func copyFolders(folders []folder.Folder) []folder.Folder {
	copyFolders := make([]folder.Folder, len(folders))
	copy(copyFolders, folders)
	return copyFolders
}

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()
	useFolders := folder.GetSampleData()
	useOrgID := uuid.FromStringOrNil("b20c2cfb-04c4-474d-ba67-03a8a7453578")
	// useDriver := folder.NewDriver(useFolders) // called again later in test, each test has its own driver
	altOrgID := uuid.FromStringOrNil("452b5b49-5762-4d10-877a-e84b6a8beb76")
	useInitialFolders := copyFolders(useFolders)

	tests := [...]struct {
		name           string
		src            string
		dst            string
		orgID          uuid.UUID
		want           []folder.Folder
		err            string
		initialFolders []folder.Folder
	}{
		{
			name:  "Case 1: Parent with multiple children",
			orgID: useOrgID,
			src:   "beta",
			dst:   "zeta",
			want: []folder.Folder{
				{Name: "alpha", OrgId: useOrgID, Paths: "alpha"},
				{Name: "beta", OrgId: useOrgID, Paths: "alpha.zeta.beta"},
				{Name: "gamma", OrgId: useOrgID, Paths: "alpha.zeta.beta.gamma"},
				{Name: "delta", OrgId: useOrgID, Paths: "alpha.zeta.beta.delta"},
				{Name: "epsilon", OrgId: useOrgID, Paths: "alpha.zeta.beta.epsilon"},
				{Name: "zeta", OrgId: useOrgID, Paths: "alpha.zeta"},
				{Name: "eta", OrgId: useOrgID, Paths: "alpha.eta"},
				{Name: "theta", OrgId: useOrgID, Paths: "theta"},
				{Name: "iota", OrgId: useOrgID, Paths: "theta.iota"},
				{Name: "kappa", OrgId: useOrgID, Paths: "theta.kappa"},
				{Name: "lambda", OrgId: useOrgID, Paths: "theta.kappa.lambda"},
				{Name: "mu", OrgId: useOrgID, Paths: "theta.mu"},
				{Name: "nu", OrgId: altOrgID, Paths: "nu"},
				{Name: "xi", OrgId: altOrgID, Paths: "xi"},
				{Name: "omicron", OrgId: altOrgID, Paths: "nu.omicron"},
				{Name: "pi", OrgId: altOrgID, Paths: "nu.omicron.pi"},
				{Name: "rho", OrgId: altOrgID, Paths: "nu.omicron.pi.rho"},
				{Name: "sigma", OrgId: altOrgID, Paths: "alpha.beta.gamma.sigma"},
				{Name: "tau", OrgId: altOrgID, Paths: "nu.tau"},
				{Name: "upsilon", OrgId: altOrgID, Paths: "upsilon"},
				{Name: "phi", OrgId: altOrgID, Paths: "upsilon.phi"},
				{Name: "chi", OrgId: altOrgID, Paths: "upsilon.phi.chi"},
				{Name: "psi", OrgId: altOrgID, Paths: "upsilon.phi.chi.psi"},
				{Name: "omega", OrgId: altOrgID, Paths: "omega"},
			},
			err:            "",
			initialFolders: useInitialFolders,
		},
		{
			name:  "Case 2: Parent with no children",
			orgID: useOrgID,
			src:   "epsilon",
			dst:   "zeta",
			want: []folder.Folder{
				{Name: "alpha", OrgId: useOrgID, Paths: "alpha"},
				{Name: "beta", OrgId: useOrgID, Paths: "alpha.beta"},
				{Name: "gamma", OrgId: useOrgID, Paths: "alpha.beta.gamma"},
				{Name: "delta", OrgId: useOrgID, Paths: "alpha.beta.delta"},
				{Name: "epsilon", OrgId: useOrgID, Paths: "alpha.zeta.epsilon"},
				{Name: "zeta", OrgId: useOrgID, Paths: "alpha.zeta"},
				{Name: "eta", OrgId: useOrgID, Paths: "alpha.eta"},
				{Name: "theta", OrgId: useOrgID, Paths: "theta"},
				{Name: "iota", OrgId: useOrgID, Paths: "theta.iota"},
				{Name: "kappa", OrgId: useOrgID, Paths: "theta.kappa"},
				{Name: "lambda", OrgId: useOrgID, Paths: "theta.kappa.lambda"},
				{Name: "mu", OrgId: useOrgID, Paths: "theta.mu"},
				{Name: "nu", OrgId: altOrgID, Paths: "nu"},
				{Name: "xi", OrgId: altOrgID, Paths: "xi"},
				{Name: "omicron", OrgId: altOrgID, Paths: "nu.omicron"},
				{Name: "pi", OrgId: altOrgID, Paths: "nu.omicron.pi"},
				{Name: "rho", OrgId: altOrgID, Paths: "nu.omicron.pi.rho"},
				{Name: "sigma", OrgId: altOrgID, Paths: "alpha.beta.gamma.sigma"},
				{Name: "tau", OrgId: altOrgID, Paths: "nu.tau"},
				{Name: "upsilon", OrgId: altOrgID, Paths: "upsilon"},
				{Name: "phi", OrgId: altOrgID, Paths: "upsilon.phi"},
				{Name: "chi", OrgId: altOrgID, Paths: "upsilon.phi.chi"},
				{Name: "psi", OrgId: altOrgID, Paths: "upsilon.phi.chi.psi"},
				{Name: "omega", OrgId: altOrgID, Paths: "omega"},
			},
			err:            "",
			initialFolders: useInitialFolders,
		},
		{
			name:  "Case 3: no source folder",
			orgID: useOrgID,
			src:   "test",
			dst:   "zeta",
			want: []folder.Folder{
				{Name: "alpha", OrgId: useOrgID, Paths: "alpha"},
				{Name: "beta", OrgId: useOrgID, Paths: "alpha.beta"},
				{Name: "gamma", OrgId: useOrgID, Paths: "alpha.beta.gamma"},
				{Name: "delta", OrgId: useOrgID, Paths: "alpha.beta.delta"},
				{Name: "epsilon", OrgId: useOrgID, Paths: "alpha.beta.epsilon"},
				{Name: "zeta", OrgId: useOrgID, Paths: "alpha.zeta"},
				{Name: "eta", OrgId: useOrgID, Paths: "alpha.eta"},
				{Name: "theta", OrgId: useOrgID, Paths: "theta"},
				{Name: "iota", OrgId: useOrgID, Paths: "theta.iota"},
				{Name: "kappa", OrgId: useOrgID, Paths: "theta.kappa"},
				{Name: "lambda", OrgId: useOrgID, Paths: "theta.kappa.lambda"},
				{Name: "mu", OrgId: useOrgID, Paths: "theta.mu"},
				{Name: "nu", OrgId: altOrgID, Paths: "nu"},
				{Name: "xi", OrgId: altOrgID, Paths: "xi"},
				{Name: "omicron", OrgId: altOrgID, Paths: "nu.omicron"},
				{Name: "pi", OrgId: altOrgID, Paths: "nu.omicron.pi"},
				{Name: "rho", OrgId: altOrgID, Paths: "nu.omicron.pi.rho"},
				{Name: "sigma", OrgId: altOrgID, Paths: "alpha.beta.gamma.sigma"},
				{Name: "tau", OrgId: altOrgID, Paths: "nu.tau"},
				{Name: "upsilon", OrgId: altOrgID, Paths: "upsilon"},
				{Name: "phi", OrgId: altOrgID, Paths: "upsilon.phi"},
				{Name: "chi", OrgId: altOrgID, Paths: "upsilon.phi.chi"},
				{Name: "psi", OrgId: altOrgID, Paths: "upsilon.phi.chi.psi"},
				{Name: "omega", OrgId: altOrgID, Paths: "omega"},
			},
			err:            "folder not found",
			initialFolders: useInitialFolders,
		},
		{
			name:  "Case 4: no destination folder",
			orgID: useOrgID,
			src:   "epsilon",
			dst:   "test",
			want: []folder.Folder{
				{Name: "alpha", OrgId: useOrgID, Paths: "alpha"},
				{Name: "beta", OrgId: useOrgID, Paths: "alpha.beta"},
				{Name: "gamma", OrgId: useOrgID, Paths: "alpha.beta.gamma"},
				{Name: "delta", OrgId: useOrgID, Paths: "alpha.beta.delta"},
				{Name: "epsilon", OrgId: useOrgID, Paths: "alpha.beta.epsilon"},
				{Name: "zeta", OrgId: useOrgID, Paths: "alpha.zeta"},
				{Name: "eta", OrgId: useOrgID, Paths: "alpha.eta"},
				{Name: "theta", OrgId: useOrgID, Paths: "theta"},
				{Name: "iota", OrgId: useOrgID, Paths: "theta.iota"},
				{Name: "kappa", OrgId: useOrgID, Paths: "theta.kappa"},
				{Name: "lambda", OrgId: useOrgID, Paths: "theta.kappa.lambda"},
				{Name: "mu", OrgId: useOrgID, Paths: "theta.mu"},
				{Name: "nu", OrgId: altOrgID, Paths: "nu"},
				{Name: "xi", OrgId: altOrgID, Paths: "xi"},
				{Name: "omicron", OrgId: altOrgID, Paths: "nu.omicron"},
				{Name: "pi", OrgId: altOrgID, Paths: "nu.omicron.pi"},
				{Name: "rho", OrgId: altOrgID, Paths: "nu.omicron.pi.rho"},
				{Name: "sigma", OrgId: altOrgID, Paths: "alpha.beta.gamma.sigma"},
				{Name: "tau", OrgId: altOrgID, Paths: "nu.tau"},
				{Name: "upsilon", OrgId: altOrgID, Paths: "upsilon"},
				{Name: "phi", OrgId: altOrgID, Paths: "upsilon.phi"},
				{Name: "chi", OrgId: altOrgID, Paths: "upsilon.phi.chi"},
				{Name: "psi", OrgId: altOrgID, Paths: "upsilon.phi.chi.psi"},
				{Name: "omega", OrgId: altOrgID, Paths: "omega"},
			},
			err:            "destination folder not found",
			initialFolders: useInitialFolders,
		},
		{
			name:  "Case 5: orgID mismatch",
			orgID: useOrgID,
			src:   "mu",
			dst:   "tau",
			want: []folder.Folder{
				{Name: "alpha", OrgId: useOrgID, Paths: "alpha"},
				{Name: "beta", OrgId: useOrgID, Paths: "alpha.beta"},
				{Name: "gamma", OrgId: useOrgID, Paths: "alpha.beta.gamma"},
				{Name: "delta", OrgId: useOrgID, Paths: "alpha.beta.delta"},
				{Name: "epsilon", OrgId: useOrgID, Paths: "alpha.beta.epsilon"},
				{Name: "zeta", OrgId: useOrgID, Paths: "alpha.zeta"},
				{Name: "eta", OrgId: useOrgID, Paths: "alpha.eta"},
				{Name: "theta", OrgId: useOrgID, Paths: "theta"},
				{Name: "iota", OrgId: useOrgID, Paths: "theta.iota"},
				{Name: "kappa", OrgId: useOrgID, Paths: "theta.kappa"},
				{Name: "lambda", OrgId: useOrgID, Paths: "theta.kappa.lambda"},
				{Name: "mu", OrgId: useOrgID, Paths: "theta.mu"},
				{Name: "nu", OrgId: altOrgID, Paths: "nu"},
				{Name: "xi", OrgId: altOrgID, Paths: "xi"},
				{Name: "omicron", OrgId: altOrgID, Paths: "nu.omicron"},
				{Name: "pi", OrgId: altOrgID, Paths: "nu.omicron.pi"},
				{Name: "rho", OrgId: altOrgID, Paths: "nu.omicron.pi.rho"},
				{Name: "sigma", OrgId: altOrgID, Paths: "alpha.beta.gamma.sigma"},
				{Name: "tau", OrgId: altOrgID, Paths: "nu.tau"},
				{Name: "upsilon", OrgId: altOrgID, Paths: "upsilon"},
				{Name: "phi", OrgId: altOrgID, Paths: "upsilon.phi"},
				{Name: "chi", OrgId: altOrgID, Paths: "upsilon.phi.chi"},
				{Name: "psi", OrgId: altOrgID, Paths: "upsilon.phi.chi.psi"},
				{Name: "omega", OrgId: altOrgID, Paths: "omega"},
			},
			err:            "organisation ID mismatch",
			initialFolders: useInitialFolders,
		},
		{
			name:  "Case 6: source matches destination",
			orgID: useOrgID,
			src:   "epsilon",
			dst:   "epsilon",
			want: []folder.Folder{
				{Name: "alpha", OrgId: useOrgID, Paths: "alpha"},
				{Name: "beta", OrgId: useOrgID, Paths: "alpha.beta"},
				{Name: "gamma", OrgId: useOrgID, Paths: "alpha.beta.gamma"},
				{Name: "delta", OrgId: useOrgID, Paths: "alpha.beta.delta"},
				{Name: "epsilon", OrgId: useOrgID, Paths: "alpha.beta.epsilon"},
				{Name: "zeta", OrgId: useOrgID, Paths: "alpha.zeta"},
				{Name: "eta", OrgId: useOrgID, Paths: "alpha.eta"},
				{Name: "theta", OrgId: useOrgID, Paths: "theta"},
				{Name: "iota", OrgId: useOrgID, Paths: "theta.iota"},
				{Name: "kappa", OrgId: useOrgID, Paths: "theta.kappa"},
				{Name: "lambda", OrgId: useOrgID, Paths: "theta.kappa.lambda"},
				{Name: "mu", OrgId: useOrgID, Paths: "theta.mu"},
				{Name: "nu", OrgId: altOrgID, Paths: "nu"},
				{Name: "xi", OrgId: altOrgID, Paths: "xi"},
				{Name: "omicron", OrgId: altOrgID, Paths: "nu.omicron"},
				{Name: "pi", OrgId: altOrgID, Paths: "nu.omicron.pi"},
				{Name: "rho", OrgId: altOrgID, Paths: "nu.omicron.pi.rho"},
				{Name: "sigma", OrgId: altOrgID, Paths: "alpha.beta.gamma.sigma"},
				{Name: "tau", OrgId: altOrgID, Paths: "nu.tau"},
				{Name: "upsilon", OrgId: altOrgID, Paths: "upsilon"},
				{Name: "phi", OrgId: altOrgID, Paths: "upsilon.phi"},
				{Name: "chi", OrgId: altOrgID, Paths: "upsilon.phi.chi"},
				{Name: "psi", OrgId: altOrgID, Paths: "upsilon.phi.chi.psi"},
				{Name: "omega", OrgId: altOrgID, Paths: "omega"},
			},
			err:            "source folder and destination folder can not be the same",
			initialFolders: useInitialFolders,
		},
		{
			name:  "Case 7: source is child of destination",
			orgID: useOrgID,
			src:   "beta",
			dst:   "gamma",
			want: []folder.Folder{
				{Name: "alpha", OrgId: useOrgID, Paths: "alpha"},
				{Name: "beta", OrgId: useOrgID, Paths: "alpha.beta"},
				{Name: "gamma", OrgId: useOrgID, Paths: "alpha.beta.gamma"},
				{Name: "delta", OrgId: useOrgID, Paths: "alpha.beta.delta"},
				{Name: "epsilon", OrgId: useOrgID, Paths: "alpha.beta.epsilon"},
				{Name: "zeta", OrgId: useOrgID, Paths: "alpha.zeta"},
				{Name: "eta", OrgId: useOrgID, Paths: "alpha.eta"},
				{Name: "theta", OrgId: useOrgID, Paths: "theta"},
				{Name: "iota", OrgId: useOrgID, Paths: "theta.iota"},
				{Name: "kappa", OrgId: useOrgID, Paths: "theta.kappa"},
				{Name: "lambda", OrgId: useOrgID, Paths: "theta.kappa.lambda"},
				{Name: "mu", OrgId: useOrgID, Paths: "theta.mu"},
				{Name: "nu", OrgId: altOrgID, Paths: "nu"},
				{Name: "xi", OrgId: altOrgID, Paths: "xi"},
				{Name: "omicron", OrgId: altOrgID, Paths: "nu.omicron"},
				{Name: "pi", OrgId: altOrgID, Paths: "nu.omicron.pi"},
				{Name: "rho", OrgId: altOrgID, Paths: "nu.omicron.pi.rho"},
				{Name: "sigma", OrgId: altOrgID, Paths: "alpha.beta.gamma.sigma"},
				{Name: "tau", OrgId: altOrgID, Paths: "nu.tau"},
				{Name: "upsilon", OrgId: altOrgID, Paths: "upsilon"},
				{Name: "phi", OrgId: altOrgID, Paths: "upsilon.phi"},
				{Name: "chi", OrgId: altOrgID, Paths: "upsilon.phi.chi"},
				{Name: "psi", OrgId: altOrgID, Paths: "upsilon.phi.chi.psi"},
				{Name: "omega", OrgId: altOrgID, Paths: "omega"},
			},
			err:            "destination folder is a child of the source folder",
			initialFolders: useInitialFolders,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useDriver := folder.NewDriver(copyFolders(tt.initialFolders))
			ret, err := useDriver.MoveFolder(tt.src, tt.dst)
			if tt.err != "" {
				if err == nil || err.Error() != tt.err {
					t.Errorf("expected: %v, got: %v", tt.err, err)
				} else {
					// This log message will appear if the error matches as expected
					t.Logf("got expected error %v", err)
				}
			}

			if (ret != nil) && (tt.want != nil) {
				assert.Equal(t, tt.want, ret)
			} else {
				t.Logf("got %v, want %v", ret, tt.want)
			}
		})
	}

}
