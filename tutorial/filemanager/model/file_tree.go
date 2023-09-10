package model

type FileType int

const (
	FileType_UnKnown FileType = 0
	FileType_Dir     FileType = 1
	FileType_Curl    FileType = 2
)

type FileNode struct {
	Name       string
	Type       FileType
	Path       string
	ParentNode *FileNode
	Level      int
}

func (f *FileNode) GetName() string {
	if f == nil {
		return ""
	}
	return f.Name
}

func (f *FileNode) GetType() FileType {
	if f == nil {
		return FileType_UnKnown
	}
	return f.Type
}

func (f *FileNode) GetPath() string {
	if f == nil {
		return ""
	}
	return f.Path
}

func (f *FileNode) GetParent() *FileNode {
	if f == nil {
		return nil
	}
	return f.ParentNode
}

func (f *FileNode) GetLevel() int {
	if f == nil {
		return 0
	}
	return f.Level
}

func (f *FileNode) GetFullName() string {
	if f == nil {
		return ""
	}
	return f.Path + "/" + f.Name
}

func (f *FileNode) IsDir() bool {
	if f == nil {
		return false
	}
	return f.Type == FileType_Dir
}
