package converter

type ConverterTypePair struct {
	SourceType string
	TargetType string
}

type Converter interface {
	ToFastImage(inputFile string, outputFile string) error
	ToPrettyPdf(inputFile string, outputFile string) error
	Destory()
}
