package quark_open

import (
	"context"
	"encoding/hex"
	"time"

	"github.com/OpenListTeam/OpenList/v4/drivers/base"
	"github.com/OpenListTeam/OpenList/v4/internal/model"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

// TODO: 有bug
func (d *QuarkOpen) RapidUpload(ctx context.Context, dstDir model.Obj, hash model.FileHashMetadata) (model.Obj, error) {
	if hash.ProofCode == "" {
		return nil, errors.New("proof_code is required for quark_open rapid upload")
	}
	if len(hash.MD5) == 0 || len(hash.SHA1) == 0 {
		return nil, errors.New("md5 and sha1 are required for quark_open rapid upload")
	}

	now := time.Now()
	httpMethod := "POST"
	apiPath := "/open/v1/file/upload_pre"
	tm, xPanToken, reqID := d.generateReqSign(httpMethod, apiPath, d.Addition.SignKey)

	data := base.Json{
		"file_name":       hash.Name,
		"size":            hash.Size,
		"format_type":     "application/octet-stream",
		"md5":             hex.EncodeToString(hash.MD5),
		"sha1":            hex.EncodeToString(hash.SHA1),
		"l_created_at":    now.UnixMilli(),
		"l_updated_at":    now.UnixMilli(),
		"pdir_fid":        dstDir.GetID(),
		"same_path_reuse": true,
		"proof_version":   "v1",
		"proof_code":      hash.ProofCode,
	}

	// We also need proof_seed1 and proof_seed2
	data["proof_seed1"] = d.generateProofSeed1(xPanToken)
	data["proof_seed2"] = d.generateProofSeed2(hash.Size)

	var resp UpPreResp
	manualSign := &ManualSign{
		Tm:    tm,
		Token: xPanToken,
		ReqID: reqID,
	}

	_, err := d.request(ctx, "/open/v1/file/upload_pre", "", func(req *resty.Request) {
		req.SetBody(data)
	}, &resp, manualSign)

	if err != nil {
		return nil, err
	}

	if resp.Data.Finish {
		return &model.Object{
			ID:       resp.Data.Fid,
			Name:     hash.Name,
			Size:     hash.Size,
			IsFolder: false,
		}, nil
	}

	return nil, errors.New("rapid upload failed")
}
