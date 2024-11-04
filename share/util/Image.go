package util

import (
	"io"
	"io/ioutil"
	"os"
	constants "wellbe-common/share/commonsettings/constants"
	errordef "wellbe-common/share/errordef"
	messages "wellbe-common/share/messages"
)

type IoType struct {
    Image io.ReadSeeker
	Size int64
}
func ConvertByte2Io(image []byte) (*IoType, *errordef.LogicError) {
	ioTmp := &IoType{}
	tmpFile, err := ioutil.TempFile("", "image")
	if err != nil {
		return nil, &errordef.LogicError{Code: constants.LOGIC_ERROR_CODE_IMAGE_CONVERT, Msg: messages.MESSAGE_EN_IMAGE_CONVERT_ERROR}
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write(image)
	if err != nil {
		return nil, &errordef.LogicError{Code: constants.LOGIC_ERROR_CODE_IMAGE_CONVERT, Msg: messages.MESSAGE_EN_IMAGE_CONVERT_ERROR}
	}

	tmpFile.Seek(0, io.SeekStart)
	imageReader := io.ReadSeeker(tmpFile)

	fileInfo, err := tmpFile.Stat()
	if err != nil {
		return nil, &errordef.LogicError{Code: constants.LOGIC_ERROR_CODE_IMAGE_CONVERT, Msg: messages.MESSAGE_EN_IMAGE_CONVERT_ERROR}
	}
	ioTmp.Image = imageReader
	ioTmp.Size = fileInfo.Size()


	return ioTmp, nil
}
