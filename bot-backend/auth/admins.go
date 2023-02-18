package auth

type Admin struct {
	UserName string
	Scope    []Scope
}

type AdminManager struct {
	Admins []Admin
}

type Scope struct {
	Domain string
	Task   string
}

const (
	DomainDropbox = "dropbox"
)

const AllTasksPermission = "all"

// dropbox tasks
const (
	AddFileRequest = "add-file-request"
)

func NewAdminManager() *AdminManager {
	admins := []Admin{}
	admins = append(admins, Admin{
		UserName: "hjwusg",
		Scope: []Scope{
			{"dropbox", "all"},
		},
	})

	return &AdminManager{
		Admins: admins,
	}
}

func (am *AdminManager) CanPerformTask(username, domain, task string) bool {
	var admin Admin
	for _, v := range am.Admins {
		if v.UserName == username {
			admin = v
		}
	}

	if admin.UserName == "" {
		return false
	}

	return admin.CanPerformTask(domain, task)
}

func (a *Admin) FilterScope(domain string) []string {
	filtered := []string{}

	for _, scope := range a.Scope {
		if scope.Domain == domain {
			filtered = append(filtered, scope.Task)
		}
	}

	return filtered
}

func (a *Admin) CanPerformTask(domain, task string) bool {
	filtered := a.FilterScope(domain)
	for _, t := range filtered {
		if t == AllTasksPermission {
			return true
		}
		if t == task {
			return true
		}
	}
	return false
}
