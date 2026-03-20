package baidu_netdisk

import (
	"context"
	stdpath "path"

	"github.com/OpenListTeam/OpenList/v4/internal/model"
	"github.com/OpenListTeam/OpenList/v4/pkg/utils"
	"github.com/pkg/errors"
)

func (d *BaiduNetdisk) RapidUpload(ctx context.Context, dstDir model.Obj, hash model.FileHashMetadataReq) (model.Obj, error) {
	contentMd5 := hash.MD5
	if len(contentMd5) == 0 {
		return nil, errors.New("invalid hash")
	}

	path := stdpath.Join(dstDir.GetPath(), hash.Name)
	blockList, _ := utils.Json.MarshalToString([]string{contentMd5})

	var newFile File
	_, err := d.create(path, hash.Size, 0, "", blockList, &newFile, 0, 0)
	if err != nil {
		return nil, err
	}
	return fileToObj(newFile), nil
}
