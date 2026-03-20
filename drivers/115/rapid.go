package _115

import (
	"context"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/OpenListTeam/OpenList/v4/internal/model"
)

func (d *Pan115) RapidUpload(ctx context.Context, dstDir model.Obj, hash model.FileHashMetadataReq) (model.Obj, error) {
	if err := d.WaitLimit(ctx); err != nil {
		return nil, err
	}
	if ok, err := d.client.UploadAvailable(); err != nil || !ok {
		return nil, err
	}
	if hash.Size > d.client.UploadMetaInfo.SizeLimit {
		return nil, nil // Or return a specific error
	}

	fullHash := strings.ToUpper(hex.EncodeToString(hash.SHA1))
	// 115 rapid upload usually requires a pre-hash (first 128KB SHA1).
	// If it's not provided, we try with the full hash as pre-hash,
	// though it might fail if the server strictly requires the 128KB one.
	// In a pure hash-based rapid upload, we might not have the 128KB hash unless it's stored.
	preHash := fullHash

	dirID := dstDir.GetID()
	fastInfo, err := d.rapidUpload(hash.Size, hash.Name, dirID, preHash, fullHash, nil)
	if err != nil {
		return nil, err
	}
	if matched, err := fastInfo.Ok(); err != nil {
		return nil, err
	} else if matched {
		f, err := d.getNewFileByPickCode(fastInfo.PickCode)
		if err != nil {
			return nil, nil
		}
		return f, nil
	}
	return nil, errors.New("rapid upload failed: not matched")
}
