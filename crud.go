package main

var id int

func Create(stus *[]Stu, new Stu) bool {
	id += 1
	//(*stus)[len(*stus)] = new
	*stus = append((*stus), new)
	(*stus)[len(*stus)-1].id = id
	return true
}

func Query(stus *[]Stu, name string) Stu {
	for _, e := range *stus {
		if e.name == name {
			return e
		}
	}
	return Stu{}
}

func Delete(stus *[]Stu, name string) bool {
	index := -1
	for i, e := range *stus {
		if e.name == name {
			index = i
			break
		}
	}
	if index == -1 {
		return false
	} else {
		*stus = append((*stus)[:index], (*stus)[index+1:]...)
		return true
	}
}

func Update(stus *[]Stu, name string, new Stu) bool {
	index := -1
	for i, e := range *stus {
		if e.name == name {
			e = new
			index = i
		}
	}
	if index != -1 {
		(*stus)[index] = new
	} else {
		return false
	}
	return false
}
