package config

type LocalUpload struct {
	UploadFilePath string `yaml:"upload_file_path"`
	Size           int    `yaml:"size"`
}
