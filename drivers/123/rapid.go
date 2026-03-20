package _123

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/OpenListTeam/OpenList/v4/drivers/base"
	"github.com/OpenListTeam/OpenList/v4/internal/model"
	"github.com/go-resty/resty/v2"
)

func (d *Pan123) RapidUpload(ctx context.Context, dstDir model.Obj, hash model.FileHashMetadataReq) (model.Obj, error) {
	etag := hash.MD5
	if len(etag) == 0 {
		return nil, errors.New("invalid hash")
	}

	data := base.Json{
		"driveId":      0,
		"duplicate":    2, // 2->覆盖 1->重命名 0->默认
		"etag":         strings.ToLower(etag),
		"fileName":     hash.Name,
		"parentFileId": dstDir.GetID(),
		"size":         hash.Size,
		"type":         0,
	}
	var resp UploadResp
	_, err := d.Request(UploadRequest, http.MethodPost, func(req *resty.Request) {
		req.SetBody(data).SetContext(ctx)
	}, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Data.Reuse || resp.Data.Key == "" {
		return File{
			FileName: hash.Name,
			Size:     hash.Size,
			FileId:   resp.Data.FileId,
			Etag:     etag,
			Type:     0,
			UpdateAt: time.Now(),
		}, nil
	}
	return nil, errors.New("rapid upload failed")
}
