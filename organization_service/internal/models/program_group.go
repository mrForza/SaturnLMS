package models

type ProgramGroup struct {
	Number       uint16 `db:"number"`
	Name         string `db:"name"`
	CourseNumber uint8  `db:"course_number"`
	ProgramName  string `db:"program_name"`
}
