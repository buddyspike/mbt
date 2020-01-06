package git

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

const DataPath string = "./.data"

func setup() {
	os.RemoveAll(DataPath)
	os.Mkdir(DataPath, 0744)
}

func TestVersion(t *testing.T) {
	setup()
	c := NewCLI(DataPath)
	v, err := c.Version()
	assert.NoError(t, err)
	assert.Contains(t, v, "git version")
}

func TestInit(t *testing.T) {
	setup()
	c := NewCLI(DataPath)
	err := c.Init()
	assert.NoError(t, err)
	s, err := c.Status()
	assert.Empty(t, s)
}

func TestStatus(t *testing.T) {
	setup()
	c := NewCLI(DataPath)
	err := c.Init()
	assert.NoError(t, err)
	assert.NoError(t, ioutil.WriteFile(path.Join(DataPath, "readme.md"), []byte("#hey"), 0744))
	s, err := c.Status()
	assert.NoError(t, err)
	assert.Equal(t, "readme.md", s[0].Path)
	assert.Equal(t, StatusUntracked, s[0].Status)
}

func TestIndexAll(t *testing.T) {
	setup()
	c := NewCLI(DataPath)
	err := c.Init()
	assert.NoError(t, err)
	assert.NoError(t, ioutil.WriteFile(path.Join(DataPath, "readme.md"), []byte("#hey"), 0744))
	assert.NoError(t, c.IndexAll())
	s, err := c.Status()
	assert.NoError(t, err)
	assert.Equal(t, "readme.md", s[0].Path)
	assert.Equal(t, StatusAdded, s[0].Status)
}

func TestCommit(t *testing.T) {
	setup()
	c := NewCLI(DataPath)
	err := c.Init()
	assert.NoError(t, err)
	assert.NoError(t, c.ConfigureUser("barry", "barry@starlabs.org"))
	assert.NoError(t, ioutil.WriteFile(path.Join(DataPath, "readme.md"), []byte("#hey"), 0744))
	assert.NoError(t, c.IndexAll())
	assert.NoError(t, c.Commit("first"))
	s, err := c.Status()
	assert.NoError(t, err)
	assert.Empty(t, s)
}

func TestLog(t *testing.T) {
	setup()
	c := NewCLI(DataPath)
	err := c.Init()
	assert.NoError(t, err)
	assert.NoError(t, c.ConfigureUser("barry", "barry@starlabs.org"))
	assert.NoError(t, ioutil.WriteFile(path.Join(DataPath, "readme.md"), []byte("#hey"), 0744))
	assert.NoError(t, c.IndexAll())
	assert.NoError(t, c.Commit("first"))

	l, err := c.Log()
	assert.NoError(t, err)
	assert.Len(t, l, 1)

	// Create the second commit
	assert.NoError(t, ioutil.WriteFile(path.Join(DataPath, "readme.md"), []byte("#hey hey"), 0744))
	assert.NoError(t, c.IndexAll())
	assert.NoError(t, c.Commit("second"))
	l, err = c.Log()
	assert.NoError(t, err)
	assert.Len(t, l, 2)
}

func TestRevParse(t *testing.T) {
	setup()
	c := NewCLI(DataPath)
	assert.NoError(t, c.Init())
	assert.NoError(t, c.ConfigureUser("barry", "barry@starlabs.org"))
	assert.NoError(t, ioutil.WriteFile(path.Join(DataPath, "readme.md"), []byte("#hey"), 0744))
	assert.NoError(t, c.IndexAll())
	assert.NoError(t, c.Commit("first"))
	sha, err := c.RevParse("HEAD")
	assert.NoError(t, err)
	assert.NotEmpty(t, sha)
}

func TestCreateBranch(t *testing.T) {
	setup()
	c := NewCLI(DataPath)
	assert.NoError(t, c.Init())
	assert.NoError(t, c.ConfigureUser("barry", "barry@starlabs.org"))
	assert.NoError(t, ioutil.WriteFile(path.Join(DataPath, "readme.md"), []byte("#hey"), 0744))
	assert.NoError(t, c.IndexAll())
	assert.NoError(t, c.Commit("first"))

	assert.NoError(t, c.CreateBranch("topic"))
	assert.NoError(t, ioutil.WriteFile(path.Join(DataPath, "readme.md"), []byte("#hey hey"), 0744))
	assert.NoError(t, c.IndexAll())
	assert.NoError(t, c.Commit("second"))

	log, err := c.Log()
	assert.NoError(t, err)
	assert.Len(t, log, 2)

	assert.NoError(t, c.Checkout("master"))
	log, err = c.Log()
	assert.NoError(t, err)
	assert.Len(t, log, 1)
}

func TestDiff(t *testing.T) {
	setup()
	c := NewCLI(DataPath)
	assert.NoError(t, c.Init())
	assert.NoError(t, c.ConfigureUser("barry", "barry@starlabs.org"))
	assert.NoError(t, ioutil.WriteFile(path.Join(DataPath, "readme.md"), []byte("#hey"), 0744))
	assert.NoError(t, c.IndexAll())
	assert.NoError(t, c.Commit("first"))

	assert.NoError(t, os.Mkdir(path.Join(DataPath, "docs"), 0744))
	assert.NoError(t, ioutil.WriteFile(path.Join(DataPath, "docs/hello.md"), []byte("#hey"), 0744))
	assert.NoError(t, c.IndexAll())
	assert.NoError(t, c.Commit("second"))

	log, err := c.Log()
	assert.NoError(t, err)

	diff, err := c.Diff(log[0], log[1])
	assert.NoError(t, err)

	assert.Len(t, diff, 1)
	assert.Equal(t, "docs/hello.md", diff[0])
}

func TestBlob(t *testing.T) {
	setup()
	c := NewCLI(DataPath)
	assert.NoError(t, c.Init())
	assert.NoError(t, c.ConfigureUser("barry", "barry@starlabs.org"))
	assert.NoError(t, ioutil.WriteFile(path.Join(DataPath, "readme.md"), []byte("#hey"), 0744))
	assert.NoError(t, c.IndexAll())
	assert.NoError(t, c.Commit("first"))

	log, err := c.Log()
	assert.NoError(t, err)

	content, err := c.Blob(log[0], "readme.md")
	assert.NoError(t, err)
	assert.Equal(t, "#hey", content)
}

func TestLsTree(t *testing.T) {
	setup()
	c := NewCLI(DataPath)
	assert.NoError(t, c.Init())
	assert.NoError(t, c.ConfigureUser("barry", "barry@starlabs.org"))
	assert.NoError(t, ioutil.WriteFile(path.Join(DataPath, "readme.md"), []byte("#hey"), 0744))
	assert.NoError(t, c.IndexAll())
	assert.NoError(t, c.Commit("first"))

	assert.NoError(t, os.Mkdir(path.Join(DataPath, "docs"), 0744))
	assert.NoError(t, ioutil.WriteFile(path.Join(DataPath, "docs/hello.md"), []byte("#hey"), 0744))
	assert.NoError(t, c.IndexAll())
	assert.NoError(t, c.Commit("second"))

	log, err := c.Log()
	assert.NoError(t, err)

	ls, err := c.LsTree(log[0])
	assert.NoError(t, err)
	assert.Len(t, ls, 2)
	assert.Equal(t, "docs/hello.md", ls[0])
	assert.Equal(t, "readme.md", ls[1])

	ls, err = c.LsTree(log[1])
	assert.NoError(t, err)
	assert.Len(t, ls, 1)
	assert.Equal(t, "readme.md", ls[0])
}
