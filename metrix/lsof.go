package metrix

import (
	"bufio"
	"io"
	"sort"

	"github.com/dynport/dgtk/stats"
)

type OpenFiles struct {
	Files []*File
}

type File struct {
	FileAccessMode          string
	FileName                string
	FileType                string
	FileLockStatus          string
	FileDescriptor          string
	FileInodeNumber         string
	FileDeviceNumber        string
	LinkCount               string
	FileSize                string
	FileFlags               string
	ParentProcessId         string
	ProcessLoginName        string
	ProcessId               string
	ProcessGroupId          string
	ProcessUserId           string
	ProcessCommandName      string
	FileOffset              string
	FileDeciceCharacterCode string
	TcpInformation          string
	ProtocolName            string
	TaskId                  string
}

const (
	fileAccessMode = iota + 1
	fileName
	fileType
	fileLockStatus
	fileDescriptor
	fileInodeNumber
	fileDeviceNumber
	linkCount
	fileSize
	fileFlags
	parentProcessId
	processLoginName
	processId
	processGroupId
	processUserId
	processCommandName
	fileOffset
	fileDeciceCharacterCode
	tcpInformation
	protocolName
	taskId
)

var fileMapping = map[string]int{
	"a": fileAccessMode,
	"n": fileName,
	"t": fileType,
	"l": fileLockStatus,
	"f": fileDescriptor,
	"i": fileInodeNumber,
	"D": fileDeviceNumber,
	"k": linkCount,
	"s": fileSize,
	"G": fileFlags,
	"R": parentProcessId,
	"L": processLoginName,
	"p": processId,
	"g": processGroupId,
	"u": processUserId,
	"c": processCommandName,
	"o": fileOffset,
	"d": fileDeciceCharacterCode,
	"T": tcpInformation,
	"P": protocolName,
	"K": taskId,
}

func (r *OpenFiles) Load(in io.Reader) error {
	scanner := bufio.NewScanner(in)
	stats := stats.Map{}
	var f *File
	for scanner.Scan() {
		line := scanner.Text()
		first := line[0]
		rest := line[1:]
		switch fileMapping[string(first)] {
		case processId:
			f = &File{ProcessId: rest}
			r.Files = append(r.Files, f)
		case parentProcessId:
			f.ParentProcessId = rest
		case fileAccessMode:
			f.FileAccessMode = rest
		case fileName:
			f.FileName = rest
		case fileType:
			f.FileType = rest
		case fileLockStatus:
			f.FileLockStatus = rest
		case fileDescriptor:
			f.FileDescriptor = rest
		case fileInodeNumber:
			f.FileInodeNumber = rest
		case fileDeviceNumber:
			f.FileDeviceNumber = rest
		case linkCount:
			f.LinkCount = rest
		case fileSize:
			f.FileSize = rest
		case fileFlags:
			f.FileFlags = rest
		case processLoginName:
			f.ProcessLoginName = rest
		case processGroupId:
			f.ProcessGroupId = rest
		case processUserId:
			f.ProcessUserId = rest
		case processCommandName:
			f.ProcessCommandName = rest
		case fileOffset:
			f.FileOffset = rest
		case fileDeciceCharacterCode:
			f.FileDeciceCharacterCode = rest
		case tcpInformation:
			f.TcpInformation = rest
		case protocolName:
			f.ProtocolName = rest
		case taskId:
			f.TaskId = rest
		default:
			logger.Printf("do not know how to handle %q", string(first))
		}
	}
	e := scanner.Err()
	if e != nil {
		return e
	}
	values := stats.Values()
	sort.Sort(sort.Reverse(values))
	for _, v := range values {
		logger.Printf("%s: %d", v.Key, v.Value)
	}
	return nil
}
