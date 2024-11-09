// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"errors"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-file/internal/sdk/internal/utils"
)

type SaveFilePayloadV2Type string

const (
	SaveFilePayloadV2TypeSaveS3FilePayload            SaveFilePayloadV2Type = "SaveS3FilePayload"
	SaveFilePayloadV2TypeSaveFileFromSourceURLPayload SaveFilePayloadV2Type = "SaveFileFromSourceURLPayload"
	SaveFilePayloadV2TypeSaveCustomFilePayload        SaveFilePayloadV2Type = "SaveCustomFilePayload"
)

type SaveFilePayloadV2 struct {
	SaveS3FilePayload            *SaveS3FilePayload
	SaveFileFromSourceURLPayload *SaveFileFromSourceURLPayload
	SaveCustomFilePayload        *SaveCustomFilePayload

	Type SaveFilePayloadV2Type
}

func CreateSaveFilePayloadV2SaveS3FilePayload(saveS3FilePayload SaveS3FilePayload) SaveFilePayloadV2 {
	typ := SaveFilePayloadV2TypeSaveS3FilePayload

	return SaveFilePayloadV2{
		SaveS3FilePayload: &saveS3FilePayload,
		Type:              typ,
	}
}

func CreateSaveFilePayloadV2SaveFileFromSourceURLPayload(saveFileFromSourceURLPayload SaveFileFromSourceURLPayload) SaveFilePayloadV2 {
	typ := SaveFilePayloadV2TypeSaveFileFromSourceURLPayload

	return SaveFilePayloadV2{
		SaveFileFromSourceURLPayload: &saveFileFromSourceURLPayload,
		Type:                         typ,
	}
}

func CreateSaveFilePayloadV2SaveCustomFilePayload(saveCustomFilePayload SaveCustomFilePayload) SaveFilePayloadV2 {
	typ := SaveFilePayloadV2TypeSaveCustomFilePayload

	return SaveFilePayloadV2{
		SaveCustomFilePayload: &saveCustomFilePayload,
		Type:                  typ,
	}
}

func (u *SaveFilePayloadV2) UnmarshalJSON(data []byte) error {

	var saveCustomFilePayload SaveCustomFilePayload = SaveCustomFilePayload{}
	if err := utils.UnmarshalJSON(data, &saveCustomFilePayload, "", true, true); err == nil {
		u.SaveCustomFilePayload = &saveCustomFilePayload
		u.Type = SaveFilePayloadV2TypeSaveCustomFilePayload
		return nil
	}

	var saveS3FilePayload SaveS3FilePayload = SaveS3FilePayload{}
	if err := utils.UnmarshalJSON(data, &saveS3FilePayload, "", true, true); err == nil {
		u.SaveS3FilePayload = &saveS3FilePayload
		u.Type = SaveFilePayloadV2TypeSaveS3FilePayload
		return nil
	}

	var saveFileFromSourceURLPayload SaveFileFromSourceURLPayload = SaveFileFromSourceURLPayload{}
	if err := utils.UnmarshalJSON(data, &saveFileFromSourceURLPayload, "", true, true); err == nil {
		u.SaveFileFromSourceURLPayload = &saveFileFromSourceURLPayload
		u.Type = SaveFilePayloadV2TypeSaveFileFromSourceURLPayload
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SaveFilePayloadV2", string(data))
}

func (u SaveFilePayloadV2) MarshalJSON() ([]byte, error) {
	if u.SaveS3FilePayload != nil {
		return utils.MarshalJSON(u.SaveS3FilePayload, "", true)
	}

	if u.SaveFileFromSourceURLPayload != nil {
		return utils.MarshalJSON(u.SaveFileFromSourceURLPayload, "", true)
	}

	if u.SaveCustomFilePayload != nil {
		return utils.MarshalJSON(u.SaveCustomFilePayload, "", true)
	}

	return nil, errors.New("could not marshal union type SaveFilePayloadV2: all fields are null")
}
