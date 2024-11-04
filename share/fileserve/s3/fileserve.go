package s3

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"strings"
	"wellbe-common/settings"
	constants "wellbe-common/share/commonsettings/constants"
	errordef "wellbe-common/share/errordef"
	log "wellbe-common/share/log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

func PutFile(fileBytes io.ReadSeeker, size int64, fileType string, key string) (string, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    uuidId, _ := uuid.NewRandom()
    bucketName := settings.GetS3StaticBucketName()
    regionName := settings.GetAWSRegionName()
    configure := settings.GetAWSConfigure()
    imageDomain := settings.GetImageDomainName()
    extentions, _ := mime.ExtensionsByType(fileType)
    extention := ""
    if len(extentions) > 0 {
        extention = extentions[0]
    }
    path := key + uuidId.String() + extention

    sess := session.Must(session.NewSession(configure))

    svc := s3.New(sess, &aws.Config{
        Region: aws.String(regionName),
    })

    params := &s3.PutObjectInput{
        Bucket:        aws.String(bucketName),
        Key:           aws.String(path),
        Body:          fileBytes,
        ContentLength: aws.Int64(size),
        ContentType:   aws.String(fileType),
    }
    _, err := svc.PutObject(params)
    if err != nil {
        logger.Error(err.Error())
        return "", &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
    }
    
	return constants.HTTPS_PROTOCOL + imageDomain + "/" + path, nil
}

func PutFileWithFileName(fileBytes []byte, fileType string, fullpath string) (string, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    bucketName := settings.GetS3StaticBucketName()
    regionName := settings.GetAWSRegionName()
    configure := settings.GetAWSConfigure()
    key := GetS3Key(fullpath)

    sess := session.Must(session.NewSession(configure))

    svc := s3.New(sess, &aws.Config{
        Region: aws.String(regionName),
    })

    body := bytes.NewReader(fileBytes)

    params := &s3.PutObjectInput{
        Bucket:        aws.String(bucketName),
        Key:           aws.String(key),
        Body:          bytes.NewReader(fileBytes),
        ContentLength: aws.Int64(body.Size()),
        ContentType:   aws.String(fileType),
    }
    _, err := svc.PutObject(params)
    if err != nil {
        logger.Error(err.Error())
        return "", &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
    }
    
	return GetS3ImagePath(fullpath), nil
}

func DeleteFile(fullpath string) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    bucketName := settings.GetS3StaticBucketName()
    regionName := settings.GetAWSRegionName()
    configure := settings.GetAWSConfigure()
    key := GetS3Key(fullpath)

    sess := session.Must(session.NewSession(configure))

    svc := s3.New(sess, &aws.Config{
        Region: aws.String(regionName),
    })

    params := &s3.DeleteObjectInput{
        Bucket:        aws.String(bucketName),
        Key:           aws.String(key),
    }
    _, err := svc.DeleteObject(params)
    if err != nil {
        logger.Error(err.Error())
        return &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
    }
    
	return nil
}

func ExistsObject(fullpath string) bool {
    logger := log.GetLogger()
    defer logger.Sync()
    bucketName := settings.GetS3StaticBucketName()
    regionName := settings.GetAWSRegionName()
    configure := settings.GetAWSConfigure()
    key := GetS3Key(fullpath)

    sess := session.Must(session.NewSession(configure))

    svc := s3.New(sess, &aws.Config{
        Region: aws.String(regionName),
    })
    _, err := svc.HeadObject(&s3.HeadObjectInput{
      Bucket: aws.String(bucketName),
      Key:    aws.String(key),
    })
    return err == nil
  }

  func PublishObject(fullpath string) (string, *errordef.LogicError) {
      logger := log.GetLogger()
      defer logger.Sync()
      bucketName := settings.GetS3StaticBucketName()
      regionName := settings.GetAWSRegionName()
      configure := settings.GetAWSConfigure()
      sourceKey := GetS3Key(fullpath)
      publishedKey := GetS3PublishedKey(fullpath)
  
      sess := session.Must(session.NewSession(configure))
  
      svc := s3.New(sess, &aws.Config{
          Region: aws.String(regionName),
      })
      _, err := svc.CopyObject(&s3.CopyObjectInput{
        Bucket: aws.String(bucketName),
        CopySource: aws.String(fmt.Sprintf("%v/%v", bucketName, sourceKey)),
        Key:    aws.String(publishedKey),
      })
      if err != nil {
          logger.Error(err.Error())
          return "", &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
      }

      return GetS3ImagePublishedPath(fullpath), nil
    }

  func GetS3ImagePath(fullpath string) string {
    imageDomain := settings.GetImageDomainName()
    path := GetS3Key(fullpath)
    
    return constants.HTTPS_PROTOCOL + imageDomain + "/" + path
  }

  func GetS3ImagePublishedPath(fullpath string) string {
    imageDomain := settings.GetImageDomainName()
    path := GetS3PublishedKey(fullpath)
    
    return constants.HTTPS_PROTOCOL + imageDomain + "/" + path
  }

  func GetS3Key(fullpath string) string {
    splFullPath := strings.Split(fullpath, "/")
    return constants.S3_IMAGE_OBJECT_KEY + splFullPath[len(splFullPath)-1]
  }

  func GetS3PublishedKey(fullpath string) string {
    splFullPath := strings.Split(fullpath, "/")
    return constants.S3_IMAGE_OBJECT_KEY + constants.S3_IMAGE_OBJECT_KEY_PUBLISHED + splFullPath[len(splFullPath)-1]
  }