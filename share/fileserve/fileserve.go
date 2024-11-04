package fileserve

import (
	"io"
	constants "wellbe-common/share/commonsettings/constants"
	errordef "wellbe-common/share/errordef"
	imageflux "wellbe-common/share/fileserve/imageflux"
	imgix "wellbe-common/share/fileserve/imgix"
	s3 "wellbe-common/share/fileserve/s3"
)


func PutImageFile(fileBytes io.ReadSeeker, size int64, fileType string) (string, *errordef.LogicError) {
	return s3.PutFile(fileBytes, size, fileType, constants.S3_IMAGE_OBJECT_KEY)
}

func PutFileWithFileName(fileBytes  []byte, fileType string, fullpath string) (string, *errordef.LogicError) {
	return s3.PutFileWithFileName(fileBytes, fileType, fullpath)
}

func DeleteImageFile(fullpath string) *errordef.LogicError {
	return s3.DeleteFile(fullpath)
}

func ExistsObject(fullpath string) bool {
	return s3.ExistsObject(fullpath)
}

func PublishObject(fullpath string) (string, *errordef.LogicError) {
	return s3.PublishObject(fullpath)
}

func GetS3ImagePath(fullpath string) (string) {
	return s3.GetS3ImagePath(fullpath)
}

func GetImigxImagePublishedPath(fullpath string) (string) {
	return imgix.GetImigxImagePublishedPath(fullpath)
}

func GetImageFluxImagePublishedPath(fullpath string, width string) (string) {
	return imageflux.GetImageFluxImagePublishedPath(fullpath, width)
}