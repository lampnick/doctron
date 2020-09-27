package common

// error code defined
const (
	Success                             = 0
	AuthFailed                          = 10000000
	InvalidParams                       = 10000001
	InvalidUrl                          = 10000002
	ApiRateLimitExceeded                = 10000003
	ConvertPdfFailed                    = 20000000
	ConvertPdfWriteBytesFailed          = 20000001
	ConvertPdfUploadFailed              = 20000002
	ConvertHtml2ImageFailed             = 30000000
	ConvertHtml2ImageWriteBytesFailed   = 30000001
	ConvertHtml2ImageUploadFailed       = 30000002
	ConvertPdf2ImageFailed              = 40000000
	ConvertPdf2ImageWriteBytesFailed    = 40000001
	ConvertPdf2ImageUploadFailed        = 40000002
	ConvertPdfWatermarkFailed           = 50000000
	ConvertPdfWatermarkWriteBytesFailed = 50000001
	ConvertPdfWatermarkUploadFailed     = 50000002
)

// error msg map defined
var ErrMsg = map[int]string{
	Success:                             "",
	AuthFailed:                          "invalid authorization",
	InvalidParams:                       "invalid params",
	InvalidUrl:                          "invalid url",
	ApiRateLimitExceeded:                "api rate limit exceeded",
	ConvertPdfFailed:                    "failed convert html to pdf",
	ConvertPdfWriteBytesFailed:          "failed convert html to pdf. write bytes failed",
	ConvertPdfUploadFailed:              "failed convert html to pdf. upload failed",
	ConvertHtml2ImageFailed:             "failed convert html to image",
	ConvertHtml2ImageWriteBytesFailed:   "failed convert html to image. write bytes failed",
	ConvertHtml2ImageUploadFailed:       "failed convert html to image. upload failed",
	ConvertPdf2ImageFailed:              "failed convert pdf to image",
	ConvertPdf2ImageWriteBytesFailed:    "failed convert pdf to image. write bytes failed",
	ConvertPdf2ImageUploadFailed:        "failed convert pdf to image. upload failed",
	ConvertPdfWatermarkFailed:           "failed add watermark on pdf",
	ConvertPdfWatermarkWriteBytesFailed: "failed add watermark on pdf. write bytes failed",
	ConvertPdfWatermarkUploadFailed:     "failed add watermark on pdf. upload failed",
}
