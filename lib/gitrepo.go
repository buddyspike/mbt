package lib

import (
	"github.com/mbtproject/mbt/e"
	"github.com/mbtproject/mbt/git"
)

type gitRepo struct {
	path   string
	client git.CLI
	Log    Log
}

// Commit is a reference to a git commit object
type cliCommit struct {
	sha string
}

func (c *cliCommit) ID() string {
	return c.sha
}

func (c *cliCommit) String() string {
	return c.sha
}

type cliReference struct {
	name         string
	symbolicName string
}

func (c *cliReference) Name() string {
	return c.name
}

func (c *cliReference) SymbolicName() string {
	return c.symbolicName
}

// NewGitRepo creates a gitRepo instance that interacts with
// git repository using git cli
func NewGitRepo(path string, log Log) (Repo, error) {
	return &gitRepo{
		path: path,
		Log:  log,
	}, nil
}

func (r *gitRepo) GetCommit(commitSha string) (Commit, error) {
	return nil, e.NewError(ErrClassInternal, "Not implemented")
}

func (r *gitRepo) Path() string {
	return r.path
}

func (r *gitRepo) Diff(a, b Commit) ([]*DiffDelta, error) {
	return nil, nil
}

func (r *gitRepo) DiffMergeBase(from, to Commit) ([]*DiffDelta, error) {
	return nil, nil
}

func (r *gitRepo) DiffWorkspace() ([]*DiffDelta, error) {
	return nil, nil
}

func (r *gitRepo) Changes(c Commit) ([]*DiffDelta, error) {
	return nil, nil
}

func (r *gitRepo) WalkBlobs(commit Commit, callback BlobWalkCallback) error {
	return nil
}

func (r *gitRepo) BlobContents(blob Blob) ([]byte, error) {
	return nil, nil
}

func (r *gitRepo) EntryID(commit Commit, path string) (string, error) {
	return "", nil
}

func (r *gitRepo) BranchCommit(name string) (Commit, error) {
	return nil, nil
}

func (r *gitRepo) CurrentBranch() (string, error) {
	return "", nil
}

func (r *gitRepo) CurrentBranchCommit() (Commit, error) {
	return nil, nil
}

func (r *gitRepo) IsEmpty() (bool, error) {
	return false, nil
}

func (r *gitRepo) FindAllFilesInWorkspace(pathSpec []string) ([]string, error) {
	return nil, nil
}

func (r *gitRepo) EnsureSafeWorkspace() error {
	return nil
}

func (r *gitRepo) BlobContentsFromTree(commit Commit, path string) ([]byte, error) {
	return nil, nil
}

func (r *gitRepo) Checkout(commit Commit) (Reference, error) {
	return nil, nil
}

func (r *gitRepo) CheckoutReference(reference Reference) error {
	return nil
}

func (r *gitRepo) MergeBase(a, b Commit) (Commit, error) {
	return nil, nil
}
