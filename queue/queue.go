package queue

var x []string

func Push(str string) {
	x = append(x, str)
}

func Pop() string {
	numOfElements := len(x)
	if numOfElements == 0 {
		return ""
	}
	popElem := x[0]
	x = x[1:numOfElements]
	return popElem

}
