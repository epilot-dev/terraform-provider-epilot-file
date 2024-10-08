// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type VerifyCustomDownloadURLPayload struct {
	// Custom external download url with signature and expiration time
	CustomDownloadURL string `json:"custom_download_url"`
}

func (o *VerifyCustomDownloadURLPayload) GetCustomDownloadURL() string {
	if o == nil {
		return ""
	}
	return o.CustomDownloadURL
}
