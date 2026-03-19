package quark

import (
	"context"
	"encoding/hex"
	"io"

	"github.com/OpenListTeam/OpenList/v4/internal/model"
	"github.com/OpenListTeam/OpenList/v4/pkg/http_range"
	"github.com/pkg/errors"
)

func (d *QuarkOrUC) RapidUpload(ctx context.Context, dstDir model.Obj, hash model.FileHashMetadata) (model.Obj, error) {
	if len(hash.MD5) == 0 || len(hash.SHA1) == 0 {
		return nil, errors.New("md5 and sha1 are required for quark_uc rapid upload")
	}

	dummy := &model.Object{
		Name: hash.Name,
		Size: hash.Size,
	}
	
	pre, err := d.upPre(&dummyFileStreamer{Object: dummy}, dstDir.GetID())
	if err != nil {
		return nil, err
	}
	
	md5Str := hex.EncodeToString(hash.MD5)
	sha1Str := hex.EncodeToString(hash.SHA1)
	
	finish, err := d.upHash(md5Str, sha1Str, pre.Data.TaskId)
	if err != nil {
		return nil, err
	}
	
	if finish {
		return &model.Object{
			ID:       pre.Data.TaskId,
			Name:     hash.Name,
			Size:     hash.Size,
			IsFolder: false,
		}, nil
	}
	
	return nil, errors.New("rapid upload failed")
}

type dummyFileStreamer struct {
	*model.Object
}

func (d *dummyFileStreamer) Read(p []byte) (n int, err error) { return 0, io.EOF }
func (d *dummyFileStreamer) Close() error { return nil }
func (d *dummyFileStreamer) GetMimetype() string { return "application/octet-stream" }
func (d *dummyFileStreamer) NeedStore() bool { return false }
func (d *dummyFileStreamer) IsForceStreamUpload() bool { return false }
func (d *dummyFileStreamer) GetExist() model.Obj { return nil }
func (d *dummyFileStreamer) SetExist(model.Obj) {}
func (d *dummyFileStreamer) RangeRead(http_range.Range) (io.Reader, error) { return nil, errors.New("not supported") }
func (d *dummyFileStreamer) CacheFullAndWriter(up *model.UpdateProgress, writer io.Writer) (model.File, error) { return nil, errors.New("not supported") }
func (d *dummyFileStreamer) GetFile() model.File { return nil }
func (d *dummyFileStreamer) GetClosers() []io.Closer { return nil }
