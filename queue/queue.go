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
	popElem := x[numOfElements-numOfElements]
	x = x[1:numOfElements]
	return popElem

}
