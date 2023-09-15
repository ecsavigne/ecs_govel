package helpers

// import "os"

// var DocHandlerExts = make(map[string]string, 40)

// var APIExts = make(map[string]string, 40)

// var CurrDir string

// func init() {

// 	CurrDir, _ = os.Getwd()

// 	DocHandlerExts["application/vnd.ms-excel"] = "xls"
// 	DocHandlerExts["application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"] = "xlsx"
// 	DocHandlerExts["application/vnd.openxmlformats-officedocument.wordprocessingml.document"] = "docx"
// 	DocHandlerExts["application/msword"] = "doc"
// 	DocHandlerExts["application/vnd.ms-powerpoint"] = "ppt"
// 	DocHandlerExts["application/vnd.openxmlformats-officedocument.presentationml.presentation"] = "pptx"

// 	DocHandlerExts["application/vnd.oasis.opendocument.text"] = "odt"
// 	DocHandlerExts["application/vnd.oasis.opendocument.spreadsheet"] = "ods"
// 	DocHandlerExts["application/vnd.oasis.opendocument.presentation"] = "odp"

// 	DocHandlerExts["application/epub+zip"] = "epub"
// 	DocHandlerExts["text/plain"] = "text"

// 	DocHandlerExts["text/csv"] = "csv"

// 	DocHandlerExts["video/mp4"] = "mp4"

// 	DocHandlerExts["application/pdf"] = "pdf"
// 	DocHandlerExts["application/zip"] = "zip"

// 	DocHandlerExts["image/png"] = "png"
// 	DocHandlerExts["image/jpg"] = "jpg"
// 	DocHandlerExts["image/jpeg"] = "jpeg"
// 	DocHandlerExts["image/gif"] = "gif"

// 	DocHandlerExts["video/mpeg"] = "mpeg"
// 	DocHandlerExts["image/webp"] = "webp"

// 	DocHandlerExts["audio/mpeg"] = "mp3"
// 	DocHandlerExts["audio/ogg; codecs=opus"] = "ogg"
// 	DocHandlerExts["audio/mpeg; codecs=opus"] = "mpeg"
// 	DocHandlerExts["audio/mp4; codecs=opus"] = "mp4"
// 	DocHandlerExts["audio/aac; codecs=opus"] = "aac"

// 	APIExts["png"] = "image/png"
// 	APIExts["jpg"] = "image/jpeg"
// 	APIExts["jpeg"] = "image/jpeg"
// 	// added by JR
// 	APIExts["webp"] = "image/webp"

// 	APIExts["docx"] = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
// 	APIExts["doc"] = "application/msword"
// 	APIExts["pdf"] = "application/pdf"
// 	APIExts["epub"] = "application/epub+zip"

// 	APIExts["ppt"] = "application/vnd.ms-powerpoint"
// 	APIExts["pptx"] = "application/vnd.openxmlformats-officedocument.presentationml.presentation"

// 	APIExts["gif"] = "image/gif"
// 	APIExts["csv"] = "application/csv" // APIExts["csv"] = "text/csv"
// 	APIExts["xlsx"] = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
// 	APIExts["xls"] = "application/excel"

// 	// video format
// 	APIExts["mp4"] = "video/mp4"

// 	// audio formats
// 	APIExts["ogg"] = "audio/ogg; codecs=opus"
// 	APIExts["mp3"] = "audio/mpeg"

// 	APIExts["txt"] = "text/plain"
// 	APIExts["odt"] = "application/vnd.oasis.opendocument.text"
// 	APIExts["zip"] = "application/zip"
// }
