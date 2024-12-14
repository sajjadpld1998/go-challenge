package helpers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"regexp"
	"skeleton/config"
	"skeleton/error_handler"
	"skeleton/services"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type requestJsonMapper struct {
	mapper interface{}
}

func (obj *requestJsonMapper) getValueByKey(key string) (value interface{}, ok bool) {
	keys, ok := obj.mapper.(map[string]interface{})

	if ok {
		value, ok = keys[key]
	}

	return
}

func (obj *requestJsonMapper) StringValue(key string) (stringValue string) {
	value, ok := obj.getValueByKey(key)

	if ok {
		stringValue = fmt.Sprint(value)
	} else {
		stringValue = ""
	}

	return
}

func (obj *requestJsonMapper) ArrayStringValue(key string) (arr []string) {
	value, ok := obj.getValueByKey(key)

	if ok {

		interfaceArray, notInterfaceArray := value.([]interface{})

		if notInterfaceArray {
			for _, element := range interfaceArray {
				stringElement, notStringElement := element.(string)

				if notStringElement {
					arr = append(arr, stringElement)
				}
			}
		}
	}

	return
}

func (obj *requestJsonMapper) JsonToStringValue(key string) (stringValue string) {
	value, ok := obj.getValueByKey(key)

	if ok {
		jsonByte, err := json.Marshal(value)

		if err != nil {
			stringValue = ""
		} else {
			stringValue = string(jsonByte)
		}
	} else {
		stringValue = ""
	}

	return
}

func (obj *requestJsonMapper) IntValue(key string) (intValue int) {
	value, ok := obj.getValueByKey(key)

	if ok {
		intValue, _ = strconv.Atoi(fmt.Sprint(value))
	} else {
		intValue = 0
	}

	return
}

func RequestJsonBody(context *gin.Context) (request requestJsonMapper) {
	// Generate copy of c.Request.Body
	bodyCopy := new(bytes.Buffer)
	_, err := io.Copy(bodyCopy, context.Request.Body)
	if err != nil {
		return
	}

	// Write copy to bodyData
	bodyData := bodyCopy.Bytes()

	// Replace the body with a reader that reads from the buffer
	context.Request.Body = io.NopCloser(bytes.NewReader(bodyData))

	// Set body again for c.ShouldBindJSON
	context.Request.Body = io.NopCloser(bytes.NewReader(bodyData))

	var jsonx interface{}
	err = context.ShouldBindJSON(&jsonx)
	if err != nil {
		return
	}

	request.mapper = jsonx

	return
}

func StringArrayToArrayString(stringArray string) (arrayString []string) {
	var images []string
	unmarshalErr := json.Unmarshal([]byte(stringArray), &images)

	if unmarshalErr == nil {
		arrayString = images
	}

	return
}

func ArrayStringIntersection(arr1, arr2 []string) (intersection []string) {
	exists := false

	for _, arr1Element := range arr1 {
		for _, arr2Element := range arr2 {
			if arr1Element == arr2Element {
				exists = true
				break
			}
		}

		if exists {
			intersection = append(intersection, arr1Element)
		}

		exists = false
	}
	return
}

func PointerStringToValueString(value *string) string {
	if value != nil {
		return *value
	}

	return ""
}

func PointerInt64ToValueInt64(value *int64) int64 {
	if value != nil {
		return *value
	}

	return 0
}

func PointerIntToValueInt(value *int) int {
	if value != nil {
		return *value
	}

	return 0
}

func PointerBoolToValueBool(value *bool) bool {
	if value != nil {
		return *value
	}

	return false
}

func EmptyStringToNil(value string) *string {
	if len(value) == 0 {
		return nil
	}

	return &value
}

func EmptyStringToInt64Pointer(value string) *int64 {
	if len(value) == 0 {
		return nil
	}

	convertetInt64, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		return nil
	}

	return &convertetInt64
}

func EmptyStringToIntPointer(value string) *int {
	if len(value) == 0 {
		return nil
	}

	convertetInt, err := strconv.Atoi(value)

	if err != nil {
		return nil
	}

	return &convertetInt
}

func EmptyBooleanToBoolPointer(value string) *bool {
	tmptBool := true

	if value == "true" || value == "1" || value == "True" || value == "TRUE" {
		tmptBool = true

		return &tmptBool
	}

	if value == "false" || value == "0" || value == "False" || value == "FALSE" {
		tmptBool = false

		return &tmptBool
	}

	return nil
}

func StringToInt64(value string) int64 {
	if len(value) == 0 {
		return 0
	}

	convertetInt64, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		return 0
	}

	return convertetInt64
}

func StringToNullString(value string) (nullable sql.NullString) {
	if len(value) > 0 {
		nullable.Valid = true
		nullable.String = value
	}

	return
}

func ArrayStringToJsonArray(arr []string) (value string) {
	if len(arr) > 0 {
		byteValue, err := json.Marshal(arr)

		if err == nil {
			value = string(byteValue)
		}
	}

	return
}

func StoreProductImages(images []*multipart.FileHeader, name string) (imagesArrayJson string) {
	Path := config.GetConfig().Path.ProductStorePathDirectory + "/" + config.GetConfig().Path.ProductImagesStorePathDirectory

	imagesArrayJson, err := UploadImages(images, name, Path)

	if err != nil {
		error_handler.ThrowServerError(err)
	}

	return
}

func StoreKitBlockImages(images []*multipart.FileHeader, name string) (imagesArrayJson string) {
	Path := config.GetConfig().Path.KitBlockStorePathDirectory + "/" + config.GetConfig().Path.KitBlockImagesStorePathDirectory

	imagesArrayJson, err := UploadImages(images, name, Path)

	if err != nil {
		error_handler.ThrowServerError(err)
	}

	return
}

func StoreKitBlockFile(file *multipart.FileHeader, name string) (fullPath string) {
	Path := config.GetConfig().Path.KitBlockStorePathDirectory + "/" + config.GetConfig().Path.KitBlockFileStorePathDirectory

	fullPath, err := UploadFile(file, name, Path)

	if err != nil {
		error_handler.ThrowServerError(err)
	}

	return
}

func UploadFile(file *multipart.FileHeader, name string, path string) (fullPath string, err error) {
	err = StringToSlug(&name)

	if err != nil {
		return
	}

	Path := path + "/" + name

	fullPath, err = StoreUploadedFile(file, Path)

	if err == nil {
		fullPath = generateFileUrl(fullPath)
	}

	return
}

func UploadImages(images []*multipart.FileHeader, name string, path string) (imagesArrayJson string, err error) {
	err = StringToSlug(&name)

	if err != nil {
		return
	}

	Path := path + "/" + name

	var imagesArray []string

	for _, image := range images {
		fullPath, storEerr := StoreUploadedFile(image, Path)
		if storEerr != nil {
			err = storEerr
			return
		}

		fullPath = generateFileUrl(fullPath)

		imagesArray = append(imagesArray, fullPath)
	}

	if len(imagesArray) == 0 {
		return
	}

	jsonArray, err := json.Marshal(imagesArray)

	if err != nil {
		return
	}

	imagesArrayJson = string(jsonArray)

	return
}

func generateRandomStringBytes(length int) string {
	var letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func generateFileUrl(path string) (fullPath string) {
	storageDriver := config.GetConfig().FileSystem.Default

	if storageDriver == "local" {
		fullPath = config.GetConfig().FileSystem.Disks.Local.Url + path + "?hash=" + generateRandomStringBytes(32)
	}

	return
}

func StoreUploadedFile(file *multipart.FileHeader, path string) (fullPath string, err error) {
	StringWithSlashInPrefixAndSuffix(&path)

	directoryPath := path

	fullPath = path + file.Filename

	storageDriver := config.GetConfig().FileSystem.Default

	switch storageDriver {
	case "s3":
		aws := services.Aws{}
		s3Error := aws.S3().Store(file, fullPath)
		if s3Error != nil {
			err = s3Error
		}
	case "local":
		localStorage := services.LocalStorage{}
		localError := localStorage.Store(file, directoryPath)
		if localError != nil {
			err = localError
		}
	}

	return
}

func DelteOldFile(oldFilePath, newFilePath string) {
	if oldFilePath != newFilePath {
		err := DeleteFile(oldFilePath)

		if err != nil {
			error_handler.ThrowServerError(err)
		}
	}
}

func DeleteOldImages(oldImages, newImages string) {
	var err error
	var oldImagesArray []string

	err = json.Unmarshal([]byte(oldImages), &oldImagesArray)

	if err != nil {
		error_handler.ThrowServerError(err)
	}

	var newImagesArray []string

	err = json.Unmarshal([]byte(newImages), &newImagesArray)

	if err != nil {
		error_handler.ThrowServerError(err)
	}

	if len(oldImagesArray) > 0 {
		for _, oldImage := range oldImagesArray {
			exists := false

			for _, newImage := range newImagesArray {
				if oldImage == newImage {
					exists = true
				}
			}

			if exists {
				err = DeleteFile(oldImage)

				if err != nil {
					error_handler.ThrowServerError(err)
				}
			}
		}
	}
}

func DeleteFile(fullPath string) (err error) {
	storageDriver := config.GetConfig().FileSystem.Default

	switch storageDriver {
	case "s3":
		aws := services.Aws{}
		s3Error := aws.S3().Delete(fullPath)
		if s3Error != nil {
			err = s3Error
		}
	}

	return
}

func StringWithSlashInPrefixAndSuffix(path *string) {
	if !strings.HasPrefix(*path, "/") {
		*path = "/" + *path
	}

	if !strings.HasSuffix(*path, "/") {
		*path = *path + "/"
	}
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func StringToSlug(s *string) error {
	str := []byte(strings.ToLower(*s))

	// convert all spaces to dash
	regE := regexp.MustCompile("[[:space:]]")
	str = regE.ReplaceAll(str, []byte("-"))

	// remove all blanks such as tab
	regE = regexp.MustCompile("[[:blank:]]")
	str = regE.ReplaceAll(str, []byte(""))

	// remove all punctuations with the exception of dash

	regE = regexp.MustCompile("[!/:-@[-`{-~]")
	str = regE.ReplaceAll(str, []byte(""))

	regE = regexp.MustCompile("/[^\x20-\x7F]/")
	str = regE.ReplaceAll(str, []byte(""))

	regE = regexp.MustCompile("`&(amp;)?#?[a-z0-9]+;`i")
	str = regE.ReplaceAll(str, []byte("-"))

	regE = regexp.MustCompile("`&([a-z])(acute|uml|circ|grave|ring|cedil|slash|tilde|caron|lig|quot|rsquo);`i")
	str = regE.ReplaceAll(str, []byte("\\1"))

	regE = regexp.MustCompile("`[^a-z0-9]`i")
	str = regE.ReplaceAll(str, []byte("-"))

	regE = regexp.MustCompile("`[-]+`")
	str = regE.ReplaceAll(str, []byte("-"))

	strReplaced := strings.Replace(string(str), "&", "", -1)
	strReplaced = strings.Replace(strReplaced, `"`, "", -1)
	strReplaced = strings.Replace(strReplaced, "&", "-", -1)
	strReplaced = strings.Replace(strReplaced, "--", "-", -1)

	if strings.HasPrefix(strReplaced, "-") || strings.HasPrefix(strReplaced, "--") {
		strReplaced = strings.TrimPrefix(strReplaced, "-")
		strReplaced = strings.TrimPrefix(strReplaced, "--")
	}

	if strings.HasSuffix(strReplaced, "-") || strings.HasSuffix(strReplaced, "--") {
		strReplaced = strings.TrimSuffix(strReplaced, "-")
		strReplaced = strings.TrimSuffix(strReplaced, "--")
	}

	// normalize unicode strings and remove all diacritical/accents marks
	// see https://www.socketloop.com/tutorials/golang-normalize-unicode-strings-for-comparison-purpose

	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	slug, _, err := transform.String(t, strReplaced)

	if err != nil {
		return err
	}

	*s = strings.TrimSpace(slug)

	return nil
}
