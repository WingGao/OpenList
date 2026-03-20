package _189pc

import (
	"context"
	"fmt"

	"github.com/OpenListTeam/OpenList/v4/internal/model"
	"github.com/pkg/errors"
)

func (y *Cloud189PC) RapidUpload(ctx context.Context, dstDir model.Obj, hash model.FileHashMetadataReq) (model.Obj, error) {
	fileMd5 := hash.MD5
	if len(fileMd5) == 0 {
		return nil, errors.New("invalid hash")
	}

	isFamily := y.isFamily()
	uploadInfo, err := y.OldUploadCreate(ctx, dstDir.GetID(), fileMd5, hash.Name, fmt.Sprint(hash.Size), isFamily)
	if err != nil {
		return nil, err
	}

	if uploadInfo.FileDataExists != 1 {
		return nil, errors.New("rapid upload fail")
	}

	return y.OldUploadCommit(ctx, uploadInfo.FileCommitUrl, uploadInfo.UploadFileId, isFamily, true)
}
