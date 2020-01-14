package lib

import (
	"fmt"
	"github.com/mbtproject/mbt/e"
)

type gitRepo struct {
	path string
	Log  Log
}

// Commit is a reference to a git commit object
type Commit struct {
	ID string
	String string
}

// Reference to a git tree object
type Reference struct {
	Name string
	SymbolicName string
}

// NewGitRepo creates a gitRepo instance that interacts with 
// git repository using git cli
func NewGitRepo(path string, log Log) (Repo, error) {
	return &gitRepo{
		path: path,
		Log: log,
	}, nil
}

func (r *gitRepo) GetCommit(commitSha string) (Commit, error) {
	return nil, nil
}

func (r *gitRepo) Path() string {
	return r.path
}

func (r *gitRepo) Diff(a, b Commit) ([]*DiffDelta, error) {
	diff, err := diff(r.Repo, a, b)
	if err != nil {
		return nil, e.Wrap(ErrClassInternal, err)
	}

	return deltas(diff)
}

func (r *gitRepo) DiffMergeBase(from, to Commit) ([]*DiffDelta, error) {
	bc, err := r.MergeBase(from, to)
	if err != nil {
		return nil, err
	}

	diff, err := diff(r.Repo, bc, to)
	if err != nil {
		return nil, e.Wrap(ErrClassInternal, err)
	}

	return deltas(diff)
}

func (r *gitRepo) DiffWorkspace() ([]*DiffDelta, error) {
	index, err := r.Repo.Index()
	if err != nil {
		return nil, e.Wrap(ErrClassInternal, err)
	}

	// Diff flags below are essential to get a list of
	// untracked files (including the ones inside new directories)
	// in the diff.
	// Without git.DiffRecurseUntracked option, if a new file is added inside
	// a new directory, we only get the path to the directory.
	// This option is same as running git status -uall
	diff, err := r.Repo.DiffIndexToWorkdir(index, &git.DiffOptions{
		Flags: git.DiffIncludeUntracked | git.DiffRecurseUntracked,
	})

	if err != nil {
		return nil, e.Wrap(ErrClassInternal, err)
	}

	return deltas(diff)
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
