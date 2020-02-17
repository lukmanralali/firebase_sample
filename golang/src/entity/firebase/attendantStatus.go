package response

type AttendantStatus struct {
	isAttended boolean `json:"is_attended"`
}

type UserAttendant struct {
	user1 AttendantStatus `json:"user_1"`
	user2 AttendantStatus `json:"user_2"`
}