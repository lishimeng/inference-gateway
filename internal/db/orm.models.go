package model

func Tables() (t []interface{}) {
	t = append(t,
		new(Driver),
	)
	return
}
