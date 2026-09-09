package progress

import "github.com/google/uuid"

// LessonProgressView is trả về bởi GET /courses/:courseId/progress — 1 item
// mỗi lesson trong course, kèm field tính toán Unlocked (không lưu DB).
type LessonProgressView struct {
	LessonID  uuid.UUID `json:"lesson_id"`
	Completed bool      `json:"completed"`
	Unlocked  bool      `json:"unlocked"`
}