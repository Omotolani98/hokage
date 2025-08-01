package services

const HOME = "/Users/mac"

var FolderFileExtensions = map[string][]string{
	"Images": {
		".jpg",
		".png",
		".svg",
		".jpeg",
	},
	"Videos": {
		".mkv",
		".mp4",
		".mov",
	},
	"Documents": {
		".pdf",
		".docx",
		".txt",
		"md",
	},
	"Sheets": {
		".xlsx",
		".csv",
		".tsv",
	},
}
