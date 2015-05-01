package model

type Emvideo struct {
	VideoId, Provider string
}

type EmvideoError struct {
	message string
}

func (e EmvideoError) Error() string {
	return e.message
}

func EmvideoForNodeField(cckFieldType *CCKFieldType, cckFieldData map[CCKField]interface{}) (*Emvideo, error) {
	videoId, vidok := cckFieldData[CCKField{cckFieldType.Name, "value", "varchar"}]
	provider, providerok := cckFieldData[CCKField{cckFieldType.Name, "provider", "varchar"}]
	if !vidok || !providerok {
		return nil, EmvideoError{"emvideo data not found for field"}
	}
	return &Emvideo{videoId.(string), provider.(string)}, nil
}
