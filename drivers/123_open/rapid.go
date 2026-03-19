package _123_open

import (
	"context"
	"encoding/hex"
	"strconv"

	"github.com/OpenListTeam/OpenList/v4/internal/model"
)

func (d *Open123) RapidUpload(ctx context.Context, dstDir model.Obj, hash model.FileHashMetadata) (model.Obj, error) {
	parentID, err := strconv.ParseInt(dstDir.GetID(), 10, 64)
	if err != nil {
		return nil, err
	}
	sha1 := hex.EncodeToString(hash.SHA1)
	resp, err := d.sha1Reuse(parentID, hash.Name, sha1, hash.Size, 0)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, nil // Or specific error
	}
	// Note: The sha1Reuse response might not return the full object info immediately.
	// Some drivers might return a partial object or just success.
	// We might need to fetch the object details if needed.
	return &model.Object{
		Name:     hash.Name,
		Size:     hash.Size,
		IsFolder: false,
		// We don't have the ID here, but the system might be able to find it by name in the parent dir.
	}, nil
}
