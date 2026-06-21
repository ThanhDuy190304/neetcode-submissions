type Stack struct{
	s []int
	i []int
}

func NewStack() Stack{
	return Stack{}
}

func (this *Stack) TopVal() int {
	return this.s[len(this.s) - 1]
}
func (this *Stack) TopIndex() int{
	return this.i[len(this.i) - 1]
}

func (this *Stack) Pop (){
	this.s = this.s[:len(this.s) - 1]
	this.i = this.i[:len(this.i) - 1]
}

func (this *Stack) Push (val int, index int){
	this.s = append(this.s, val)
	this.i = append(this.i, index)

}

func (this *Stack) Len() int{
	return len(this.s)
}

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))

	st := NewStack()

	for i := len(temperatures) - 1; i >=0; i--{
		for {
			if st.Len() == 0{
				result[i] = 0
				st.Push(temperatures[i], i)
				break
			}
			if st.TopVal() <=  temperatures[i]{
				st.Pop()
			}else{
				result[i] = st.TopIndex() - i
				st.Push(temperatures[i], i)
				break
			}

		}
	} 
	return result

}
