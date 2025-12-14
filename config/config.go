package config

var cfg *Config

type Config struct {
	ChatModel  *ModelConfig
	ImageModel *ModelConfig
	VideoModel *ModelConfig
	AudioModel *ModelConfig
	VoiceModel *ModelConfig // audio for role conversion
}

type ModelConfig struct {
	Type    string
	BaseUrl string
	ApiKey  string
	Model   string
}
