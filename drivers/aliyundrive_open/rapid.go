package aliyundrive_open

import (
	"context"
	"encoding/hex"
	"time"

	"github.com/OpenListTeam/OpenList/v4/drivers/base"
	"github.com/OpenListTeam/OpenList/v4/internal/model"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

func (d *AliyundriveOpen) RapidUpload(ctx context.Context, dstDir model.Obj, hash model.FileHashMetadata) (model.Obj, error) {
	if hash.ProofCode == "" {
		return nil, errors.New("proof_code is required for aliyundrive_open rapid upload")
	}

	const dateFormat = "2006-01-02T15:04:05.000Z"
	now := time.Now().UTC().Format(dateFormat)

	createData := base.Json{
		"drive_id":          d.DriveId,
		"parent_file_id":    dstDir.GetID(),
		"name":              hash.Name,
		"type":              "file",
		"check_name_mode":   "ignore",
		"local_modified_at": now,
		"local_created_at":  now,
		"size":              hash.Size,
		"proof_version":     "v1",
		"content_hash_name": "sha1",
		"content_hash":      hex.EncodeToString(hash.SHA1),
		"proof_code":        hash.ProofCode,
	}

	var createResp CreateResp
	_, err := d.request(ctx, limiterOther, "/adrive/v1.0/openFile/create", "POST", func(req *resty.Request) {
		req.SetBody(createData).SetResult(&createResp)
	})

	if err != nil {
		return nil, err
	}

	if createResp.RapidUpload {
		return &model.Object{
			ID:       createResp.FileId,
			Name:     hash.Name,
			Size:     hash.Size,
			IsFolder: false,
		}, nil
	}

	return nil, errors.New("rapid upload failed")
}
