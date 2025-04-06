package enums

type fileTypeEnum struct {
	Image string
	Video string
}

var FileTypeEnum = fileTypeEnum{
	Image: "image",
	Video: "video",
}
