type OrderedStream struct {
	Pointer int
	Body    []string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		Pointer: 0,
		Body:    make([]string, n),
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	chunk := []string{}
	this.Body[idKey-1] = value

	for this.Pointer < len(this.Body) && this.Body[this.Pointer] != "" {
		chunk = append(chunk, this.Body[this.Pointer])
		fmt.Println(this.Pointer)

		this.Pointer += 1
	}

	return chunk
}
