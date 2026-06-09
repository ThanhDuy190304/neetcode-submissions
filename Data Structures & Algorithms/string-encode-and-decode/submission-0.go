type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    var out strings.Builder

	for _, str := range strs {
		out.WriteString(strconv.Itoa(len(str)))
		out.WriteByte('#')
		out.WriteString(str)
	}

	return out.String()

}

func (s *Solution) Decode(encoded string) []string {
    result := []string{}

	for i := 0; i < len(encoded); {
		j := i

		for encoded[j] != '#' {
			j++
		}

		length, _ := strconv.Atoi(encoded[i:j])

		start := j + 1
		end := start + length

		result = append(result, encoded[start:end])
		i = end
	}

	return result

}
