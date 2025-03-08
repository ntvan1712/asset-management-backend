package app_utils

import (
	"path/filepath"
	"regexp"
	"strings"
	"time"

	// "com.pegatech.faceswap/common/enums"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

var contentTypes = map[string]string{
	// Image
	".apng": "image/apng",
	".avif": "image/avif",
	".gif":  "image/gif",
	".jpeg": "image/jpeg",
	".jpg":  "image/jpeg",
	".png":  "image/png",
	".svg":  "image/svg+xml",
	".webp": "image/webp",
	".heic": "image/heic",
	".heif": "image/heif",
	".bmp":  "image/bmp",
	// Video
	".mp4":  "video/mp4",
	".mov":  "video/quicktime",
	".avi":  "video/x-msvideo",
	".mkv":  "video/x-matroska",
	".mpeg": "video/mpeg",
	".ogv":  "video/ogg",

	// Audio
	".mp3":  "audio/mpeg",
	".wav":  "audio/wav",
	".weba": "audio/webm",
	".m4a":  "audio/mp4",
	".aac":  "audio/aac",
	".mid":  "audio/midi",
	".midi": "audio/midi",
	".oga":  "audio/ogg",
	".opus": "audio/ogg",

	".pdf":  "application/pdf",
	".doc":  "application/msword",
	".docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	".xls":  "application/vnd.ms-excel",
	".xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	".txt":  "text/plain",
	".html": "text/html",
}

func GetContentType(fileName string) string {
	ext := strings.ToLower(filepath.Ext(fileName))
	if contentType, ok := contentTypes[ext]; ok {
		return contentType
	}
	return "application/octet-stream"
}

// Return file type by file name, "image", "video", "audio", "unknown"
// func GetFileTypeFromFileName(fileName string) string {
// 	fileExt := strings.ToLower(fileName[strings.LastIndex(fileName, ".")+1:])
// 	switch fileExt {
// 	case "jpg", "jpeg", "png", "webp", "svg", "bmp", "heif", "hevc", "heic":
// 		return enums.MediaType.Image
// 	case "mp4", "mkv", "avi", "mov", "wmv", "flv", "webm", "mpeg":
// 		return enums.MediaType.Video
// 	case "mp3", "wav", "aac", "flac", "ogg", "m4a", "wma", "alac":
// 		return enums.MediaType.Audio
// 	default:
// 		return "unknown"
// 	}
// }

// Chuẩn hóa filename với input là filename và return 1 filename được chuẩn hóa với
func StandardizedFilename(filename string) string {
	filename = strings.ReplaceAll(filename, " ", "_")
	re := regexp.MustCompile(`[^\w\d_\-\.]+`)
	return re.ReplaceAllString(filename, "_")
}


// Hàm kiểm tra list strings có các phần tử duy nhất hay không
func IsUniqueListStrings(list []string) bool {
	seen := make(map[string]bool)

	for _, item := range list {
		if seen[item] {
			// Nếu phần tử đã tồn tại trong map, danh sách không duy nhất
			return false
		}
		seen[item] = true
	}

	// Nếu không có phần tử nào bị lặp lại, danh sách là duy nhất
	return true
}

// At secs
func UnixTimeNowUTCAtSecs() int64 {
	return time.Now().UTC().Unix()
}

func UnixTimeNowUTCAtMillis() int64 {
	return time.Now().UTC().UnixMilli()
}


func ChangeFileExtension(filename string, newExt string) string {
	// Lấy extension hiện tại
	oldExt := filepath.Ext(filename)

	// Xóa extension cũ ra khỏi filename
	nameWithoutExt := strings.TrimSuffix(filename, oldExt)

	// Ghép filename mới với extension mới
	newFilename := nameWithoutExt + newExt

	return newFilename
}

func FileNameWithoutExtension(filename string) string {
	// Lấy extension hiện tại
	oldExt := filepath.Ext(filename)

	// Xóa extension cũ ra khỏi filename
	return strings.TrimSuffix(filename, oldExt)
}

func IsElementOfSlice(slice []string, item string) bool {
    for _, value := range slice {
        if value == item {
            return true
        }
    }
    return false
}