package auth

type Admin struct {
	UserName string  `json:"username"`
	Scope    []Scope `json:"scope"`
}

type AdminManager struct {
	Admins []Admin `json:"data"`
}

type Scope struct {
	Domain string `json:"domain"`
	Task   string `json:"task"`
}

const (
	DomainDropbox = "dropbox"
)

const AllTasksPermission = "all"

// dropbox tasks
const (
	AddFileRequest = "add-file-request"
)
