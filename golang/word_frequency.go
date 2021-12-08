// create a map with key as the word and value as the frequency of the word
// read the text file and store the words in the map
// print the map

type WordFrequency struct {
	word string
	freq int
}

// caculate the frequency of the words in the text file with go routines


func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var words map[string]int
	words = make(map[string]int)

	for scanner.Scan() {
		words[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for k, v := range words {
		fmt.Printf("%s: %d\n", k, v)
	}
}
