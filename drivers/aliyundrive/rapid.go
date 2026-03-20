package aliyundrive

import (
	"context"
	"encoding/hex"

	"github.com/OpenListTeam/OpenList/v4/drivers/base"
	"github.com/OpenListTeam/OpenList/v4/internal/model"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

func (d *AliDrive) RapidUpload(ctx context.Context, dstDir model.Obj, hash model.FileHashMetadataReq) (model.Obj, error) {
	if hash.ProofCode == "" {
		return nil, errors.New("proof_code is required for aliyundrive rapid upload")
	}

	reqBody := base.Json{
		"check_name_mode":   "overwrite",
		"drive_id":          d.DriveId,
		"name":              hash.Name,
		"parent_file_id":    dstDir.GetID(),
		"size":              hash.Size,
		"type":              "file",
		"content_hash":      hex.EncodeToString(hash.SHA1),
		"content_hash_name": "sha1",
		"proof_version":     "v1",
		"proof_code":        hash.ProofCode,
	}

	var resp UploadResp
	_, err, _ := d.request("https://api.alipan.com/adrive/v2/file/createWithFolders", "POST", func(req *resty.Request) {
		req.SetBody(reqBody)
	}, &resp)

	if err != nil {
		return nil, err
	}

	if resp.RapidUpload {
		return &model.Object{
			ID:       resp.FileId,
			Name:     hash.Name,
			Size:     hash.Size,
			IsFolder: false,
		}, nil
	}

	return nil, errors.New("rapid upload failed")
}
