package webhook

import "time"

type Issues struct {
	ObjectKind string `json:"object_kind"`
	EventType  string `json:"event_type"`
	User       struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
		Email     string `json:"email"`
	} `json:"user"`
	Project struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		WebURL            string      `json:"web_url"`
		AvatarURL         interface{} `json:"avatar_url"`
		GitSSHURL         string      `json:"git_ssh_url"`
		GitHTTPURL        string      `json:"git_http_url"`
		Namespace         string      `json:"namespace"`
		VisibilityLevel   int         `json:"visibility_level"`
		PathWithNamespace string      `json:"path_with_namespace"`
		DefaultBranch     string      `json:"default_branch"`
		CiConfigPath      interface{} `json:"ci_config_path"`
		Homepage          string      `json:"homepage"`
		URL               string      `json:"url"`
		SSHURL            string      `json:"ssh_url"`
		HTTPURL           string      `json:"http_url"`
	} `json:"project"`
	ObjectAttributes struct {
		ID                  int         `json:"id"`
		Title               string      `json:"title"`
		AssigneeIds         []int       `json:"assignee_ids"`
		AssigneeID          int         `json:"assignee_id"`
		AuthorID            int         `json:"author_id"`
		ProjectID           int         `json:"project_id"`
		CreatedAt           time.Time   `json:"created_at"`
		UpdatedAt           time.Time   `json:"updated_at"`
		UpdatedByID         int         `json:"updated_by_id"`
		LastEditedAt        interface{} `json:"last_edited_at"`
		LastEditedByID      interface{} `json:"last_edited_by_id"`
		RelativePosition    int         `json:"relative_position"`
		Description         string      `json:"description"`
		MilestoneID         interface{} `json:"milestone_id"`
		StateID             int         `json:"state_id"`
		Confidential        bool        `json:"confidential"`
		DiscussionLocked    bool        `json:"discussion_locked"`
		DueDate             interface{} `json:"due_date"`
		MovedToID           interface{} `json:"moved_to_id"`
		DuplicatedToID      interface{} `json:"duplicated_to_id"`
		TimeEstimate        int         `json:"time_estimate"`
		TotalTimeSpent      int         `json:"total_time_spent"`
		TimeChange          int         `json:"time_change"`
		HumanTotalTimeSpent interface{} `json:"human_total_time_spent"`
		HumanTimeEstimate   interface{} `json:"human_time_estimate"`
		HumanTimeChange     interface{} `json:"human_time_change"`
		Weight              interface{} `json:"weight"`
		Iid                 int         `json:"iid"`
		URL                 string      `json:"url"`
		State               string      `json:"state"`
		Action              string      `json:"action"`
		Severity            string      `json:"severity"`
		EscalationStatus    string      `json:"escalation_status"`
		EscalationPolicy    struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"escalation_policy"`
		Labels []struct {
			ID          int       `json:"id"`
			Title       string    `json:"title"`
			Color       string    `json:"color"`
			ProjectID   int       `json:"project_id"`
			CreatedAt   time.Time `json:"created_at"`
			UpdatedAt   time.Time `json:"updated_at"`
			Template    bool      `json:"template"`
			Description string    `json:"description"`
			Type        string    `json:"type"`
			GroupID     int       `json:"group_id"`
		} `json:"labels"`
	} `json:"object_attributes"`
	Repository struct {
		Name        string `json:"name"`
		URL         string `json:"url"`
		Description string `json:"description"`
		Homepage    string `json:"homepage"`
	} `json:"repository"`
	Assignees []struct {
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
	} `json:"assignees"`
	Assignee struct {
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
	} `json:"assignee"`
	Labels []struct {
		ID          int       `json:"id"`
		Title       string    `json:"title"`
		Color       string    `json:"color"`
		ProjectID   int       `json:"project_id"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		Template    bool      `json:"template"`
		Description string    `json:"description"`
		Type        string    `json:"type"`
		GroupID     int       `json:"group_id"`
	} `json:"labels"`
	Changes struct {
		UpdatedByID struct {
			Previous interface{} `json:"previous"`
			Current  int         `json:"current"`
		} `json:"updated_by_id"`
		UpdatedAt struct {
			Previous string `json:"previous"`
			Current  string `json:"current"`
		} `json:"updated_at"`
		Labels struct {
			Previous []struct {
				ID          int       `json:"id"`
				Title       string    `json:"title"`
				Color       string    `json:"color"`
				ProjectID   int       `json:"project_id"`
				CreatedAt   time.Time `json:"created_at"`
				UpdatedAt   time.Time `json:"updated_at"`
				Template    bool      `json:"template"`
				Description string    `json:"description"`
				Type        string    `json:"type"`
				GroupID     int       `json:"group_id"`
			} `json:"previous"`
			Current []struct {
				ID          int       `json:"id"`
				Title       string    `json:"title"`
				Color       string    `json:"color"`
				ProjectID   int       `json:"project_id"`
				CreatedAt   time.Time `json:"created_at"`
				UpdatedAt   time.Time `json:"updated_at"`
				Template    bool      `json:"template"`
				Description string    `json:"description"`
				Type        string    `json:"type"`
				GroupID     int       `json:"group_id"`
			} `json:"current"`
		} `json:"labels"`
	} `json:"changes"`
}
