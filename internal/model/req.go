package model

type PageReq struct {
	Page    int `json:"page" form:"page"`
	PerPage int `json:"per_page" form:"per_page"`
}

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func (p *PageReq) Validate() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PerPage < 1 {
		p.PerPage = MaxInt
	}
}

type FileHashMetadataReq struct {
	Name      string `json:"name"`
	Size      int64  `json:"size"`
	MD5       string `json:"md5"`
	SHA1      []byte `json:"sha1"`
	SHA256    []byte `json:"sha256"`
	GCID      []byte `json:"gcid"`
	ProofCode string `json:"proof_code"`
}
