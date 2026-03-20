package pikpak

import (
	"context"
	"encoding/hex"
	"strings"

	"github.com/OpenListTeam/OpenList/v4/drivers/base"
	"github.com/OpenListTeam/OpenList/v4/internal/model"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

// TODO 有bug
func (d *PikPak) RapidUpload(ctx context.Context, dstDir model.Obj, hash model.FileHashMetadataReq) (model.Obj, error) {
	if len(hash.GCID) == 0 {
		return nil, errors.New("gcid is required for pikpak rapid upload")
	}
	gcidStr := strings.ToUpper(hex.EncodeToString(hash.GCID))

	var resp UploadTaskData
	_, err := d.request("https://api-drive.mypikpak.net/drive/v1/files", "POST", func(req *resty.Request) {
		req.SetBody(base.Json{
			"kind":        "drive#file",
			"name":        hash.Name,
			"size":        hash.Size,
			"hash":        gcidStr,
			"upload_type": "UPLOAD_TYPE_RESUMABLE",
			"objProvider": base.Json{"provider": "UPLOAD_TYPE_UNKNOWN"},
			"parent_id":   dstDir.GetID(),
			"folder_type": "NORMAL",
		})
	}, &resp)

	if err != nil {
		return nil, err
	}

	// Rapid upload success if Resumable is nil
	if resp.Resumable == nil {
		return &model.Object{
			//ID:       resp.File.ID,
			Name:     hash.Name,
			Size:     hash.Size,
			IsFolder: false,
		}, nil
	}

	return nil, errors.New("rapid upload failed")
}
