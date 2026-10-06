package service

import (
	"context"
	"errors"
	"io"
	"mime/multipart"
	"strings"
	"testing"
	"time"

	"github.com/Tencent/WeKnora/internal/types"
	"github.com/Tencent/WeKnora/internal/types/interfaces"
)

type fakeFileService struct {
	saved     map[string][]byte
	saveErr   error
	seq       int
	tenantIDs []uint64
}

func (f *fakeFileService) CheckConnectivity(_ context.Context) error { return nil }
func (f *fakeFileService) SaveFile(_ context.Context, _ *multipart.FileHeader, _ uint64, _ string) (string, error) {
	panic("unused")
}
func (f *fakeFileService) SaveBytes(_ context.Context, data []byte, tenantID uint64, fileName string, _ bool) (string, error) {
	return "", nil
}
func (f *fakeFileService) GetFile(_ context.Context, _ string) (io.ReadCloser, error) {
	panic("unused")
}
func (f *fakeFileService) GetFileURL(_ context.Context, _ string) (string, error) {
	panic("unused")
}
func (f *fakeFileService) DeleteFile(_ context.Context, _ string) error { return nil }
func (f *fakeFileService) CopyFile(_ context.Context, _ string, _ uint64, _ string) (string, error) {
	panic("unused")
}

type fakeCatalog struct {
	binds            []bindCall
	bindErr          error
	releaseRemaining map[string]int64
	releaseErr       error
	releases         []string
}

func (c *fakeCatalog) Register(context.Context, uint64, string, interfaces.ResourceRegistration) (string, error) {
	return "", nil
}
func (c *fakeCatalog) Resolve(context.Context, string) (*types.StoredResource, error) {
	return nil, nil
}
func (c *fakeCatalog) ResolvePath(_ context.Context, v string) (string, *types.StoredResource, error) {
	return v, nil, nil
}
func (c *fakeCatalog) Bind(_ context.Context, ref, ownerType, ownerID, relation string) error {
	c.binds = append(c.binds, bindCall{ref, ownerType, ownerID, relation})
	return c.bindErr
}
func (c *fakeCatalog) MarkDeleted(context.Context, string) error { return nil }

func (c *fakeCatalog) Release(_ context.Context, ref, ownerType, ownerID string) (int64, error) {
	c.releases = append(c.releases, ref+"|"+ownerType+"|"+ownerID)
	if c.releaseErr != nil {
		return -1, c.releaseErr
	}
	if remaining, ok := c.releaseRemaining[ref]; ok {
		return remaining, nil
	}
	return -1, nil
}
func (c *fakeCatalog) CreateAccessGrant(context.Context, string, time.Duration) (string, error) {
	return "", nil
}
func (c *fakeCatalog) ResolveAccessGrant(context.Context, string) (*types.StoredResource, error) {
	return nil, nil
}

// deleteRecorder captures which files a cleanup actually removed.
type deleteRecorder struct {
	fakeFileService
	deleted []string
}

func (f *deleteRecorder) DeleteFile(_ context.Context, filePath string) error {
	f.deleted = append(f.deleted, filePath)
	return nil
}

func handleRef(char string) string {
	return types.BuildResourcePath(strings.Repeat(char, types.ResourceHandleLength))
}

func TestDeleteExtractedImagesKeepsFilesAnotherOwnerStillClaims(t *testing.T) {
	shared := handleRef("a")    // also shown by the chat message it was saved from
	exclusive := handleRef("b") // only this knowledge references it
	legacy := "local://7/exports/old.png"

	catalog := &fakeCatalog{releaseRemaining: map[string]int64{shared: 1, exclusive: 0}}
	files := &deleteRecorder{}

	deleteExtractedImages(
		context.Background(), files,
		knowledgeResourceOwners(catalog, "kn-1"),
		[]string{shared, exclusive, legacy},
	)

	want := []string{exclusive, legacy}
	if len(files.deleted) != len(want) {
		t.Fatalf("deleted %v, want %v", files.deleted, want)
	}
	for i, url := range want {
		if files.deleted[i] != url {
			t.Fatalf("deleted[%d] = %q, want %q", i, files.deleted[i], url)
		}
	}
	if len(catalog.releases) != 3 {
		t.Fatalf("released %v, want one call per reference", catalog.releases)
	}
	if got := catalog.releases[0]; got != shared+"|"+types.ResourceOwnerKnowledge+"|kn-1" {
		t.Fatalf("unexpected release call %q", got)
	}
}

// A knowledge base delete releases every entry's claim before deciding, so a
// file shared between two entries of the same base is still removed.
func TestDeleteExtractedImagesReleasesEveryOwnerBeforeDeciding(t *testing.T) {
	shared := handleRef("c")
	catalog := &fakeCatalog{releaseRemaining: map[string]int64{shared: 0}}
	files := &deleteRecorder{}

	deleteExtractedImages(
		context.Background(), files,
		knowledgeResourceOwners(catalog, "kn-1", "kn-2"),
		[]string{shared},
	)

	if len(files.deleted) != 1 {
		t.Fatalf("deleted %v, want the shared file removed once", files.deleted)
	}
	if len(catalog.releases) != 2 {
		t.Fatalf("released %v, want both owners released", catalog.releases)
	}
}

// An unreadable binding count must not destroy bytes: an orphaned blob can be
// reclaimed later, an image missing from a document nobody deleted cannot.
func TestDeleteExtractedImagesKeepsFileWhenReleaseFails(t *testing.T) {
	catalog := &fakeCatalog{releaseErr: errors.New("db down")}
	files := &deleteRecorder{}

	deleteExtractedImages(
		context.Background(), files,
		knowledgeResourceOwners(catalog, "kn-1"),
		[]string{handleRef("d")},
	)

	if len(files.deleted) != 0 {
		t.Fatalf("deleted %v, want nothing deleted while the count is unknown", files.deleted)
	}
}

// Without a catalog the guard is inert and cleanup behaves as it always did.
func TestDeleteExtractedImagesWithoutCatalogDeletesEverything(t *testing.T) {
	files := &deleteRecorder{}
	urls := []string{handleRef("e"), "local://7/exports/x.png"}

	deleteExtractedImages(context.Background(), files, knowledgeResourceOwners(nil, "kn-1"), urls)

	if len(files.deleted) != len(urls) {
		t.Fatalf("deleted %v, want %v", files.deleted, urls)
	}
}
