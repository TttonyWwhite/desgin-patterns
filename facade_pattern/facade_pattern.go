package facade_pattern

// 视频转换库中的相关类，用户不需要关心其具体实现，开发者也无法进行简化
type File interface {
	Name() string
}

type VideoFile struct {
	name string
}

func NewVideoFile(filename string) VideoFile {
	return VideoFile{name: filename}
}

func (v VideoFile) Name() string {
	return v.name
}

type OggCompressionCodec struct {
}

type MPEG4CompressionCodec struct {
}

type CodecFactory struct {
}

func (c *CodecFactory) extract(file File) string {
	// 对文件进行解压
	return ""
}

type BitrateReader struct {
}

func (b *BitrateReader) read(filename, source string) []byte {
	return nil
}

func (b *BitrateReader) convert(buffer []byte, destination interface{}) interface{} {
	return nil
}

type AudioMixer struct {
}

func (a *AudioMixer) fix(toMix interface{}) interface{} {
	return nil
}

type VideoConverter struct {
}

func (v *VideoConverter) Convert(filename string, format string) File {
	file := NewVideoFile(filename)
	factory := CodecFactory{}
	sourceCodec := factory.extract(file)
	var destinationCodec interface{}
	if format == "mp4" {
		destinationCodec = MPEG4CompressionCodec{}
	} else {
		destinationCodec = OggCompressionCodec{}
	}
	bitrateReader := BitrateReader{}
	audioMixer := AudioMixer{}
	buffer := bitrateReader.read(filename, sourceCodec)
	result := bitrateReader.convert(buffer, destinationCodec)
	result = audioMixer.fix(result)

	// return result
	return nil
}
