package user

// Profile is their "business card" in the system.
// Through their name, avatar, status, and profile visibility,
// the user presents themselves to others in the messenger.
type Profile struct {
	Nickname     string `json:"nickname"`
	AvatarURL    string `json:"avatar_url,omitempty"`
	Status       string `json:"status,omitempty"`
	Bio          string `json:"bio,omitempty"`
	LastSeenAt   int64  `json:"last_seen_at"`
	IsPrivate    bool   `json:"is_private"`
	ShowLastSeen bool   `json:"show_last_seen"`
	UserID       string `json:"user_id"` // Внешний ключ к учетке
}

func NewProfile(nickname string, avatarURL string, status string, bio string, lastSeenAt int64, isPrivate bool, showLastSeen bool, userID string) *Profile {
	return &Profile{
		Nickname:     nickname,
		AvatarURL:    avatarURL,
		Status:       status,
		Bio:          bio,
		LastSeenAt:   lastSeenAt,
		IsPrivate:    isPrivate,
		ShowLastSeen: showLastSeen,
		UserID:       userID,
	}
}
