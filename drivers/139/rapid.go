package _139

import (
	"context"
	"encoding/hex"
	"strings"

	"github.com/OpenListTeam/OpenList/v4/drivers/base"
	"github.com/OpenListTeam/OpenList/v4/internal/model"
	"github.com/pkg/errors"
)

func (d *Yun139) RapidUpload(ctx context.Context, dstDir model.Obj, hash model.FileHashMetadataReq) (model.Obj, error) {
	if len(hash.SHA256) == 0 {
		return nil, errors.New("sha256 is required for 139 rapid upload")
	}
	fullHash := strings.ToUpper(hex.EncodeToString(hash.SHA256))

	data := base.Json{
		"contentHash":          fullHash,
		"contentHashAlgorithm": "SHA256",
		"contentType":          "application/octet-stream",
		"parallelUpload":       false,
		"size":                 hash.Size,
		"parentFileId":         dstDir.GetID(),
		"name":                 hash.Name,
		"type":                 "file",
		"fileRenameMode":       "auto_rename",
	}
	pathname := "/file/create"
	var resp PersonalUploadResp
	_, err := d.personalPost(pathname, data, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Data.Exist || resp.Data.RapidUpload {
		return &model.Object{
			ID:       resp.Data.FileId,
			Name:     resp.Data.FileName,
			Size:     hash.Size,
			IsFolder: false,
		}, nil
	}

	return nil, errors.New("rapid upload failed")
}
