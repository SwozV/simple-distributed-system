package business

func init() {
	students = []Student{
		{
			ID:   1,
			Name: "Kanou Yoshiki",
			Grades: []Grade{
				{
					Title: "Japenese",
					Score: 11,
				},
				{
					Title: "Chinese",
					Score: 22,
				},
			},
		},
	}
}
