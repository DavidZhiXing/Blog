// reading 16gb file in seconds

func main() {
	s := time.Now()
	args := os.Args
	if len(args) != 6 {
		fmt.Println("Usage: LogExtractor.exe -s <start_time> -e <end_time> -f <file_name>")
		return
	}
	startTimeArg := args[1]
	finishTimeArg := args[3]
	fileName := args[5]

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	qureryStartTime, err := strconv.ParseInt(startTimeArg, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	filestat, err := file.Stat()
	offset := fileSize - 1
	lastLineSize := 0

	for {
		b := make([]byte, 1)
		n, err := file.ReadAt(b, offset)
		if err != nil {
			fmt.Println("error reading file ",err)
			return
		}
		char := string(b[0])
		if char == "\n" {
			break
		}
		offset--
		lastLineSize += n
	}

	lastLine := make([]byte, lastLineSize)
	_, err = file.ReadAt(lastLine, offset+1)

	if err != nil {
		fmt.Println("error reading file ",err)
		return
	}

	logSlice := strings.Split(string(lastLine), ",", 2)
	logCreationTimeString := logSlice[0]
	logCreationTime, err := time.Parse("2006-01-02 15:04:05", logCreationTimeString)
	if err != nil {
		fmt.Println(err)
		return
	}

	if lastLogCreationTime.Before(queryFinishTime) && lastLogCreationTime.After(qureryStartTime) {
		Process(file, qureryStartTime, queryFinishTime)
	}
	fmt.Println("time taken: ", time.Since(s))
}

func Process(file *os.File, startTime time.Time, finishTime time.Time) {
	linePool := sync.Pool{
		New: func() interface{} {
			return make([]byte, 250*1024)
		}
	}

	stringPool := sync.Pool{
		New: func() interface{} {
			lines := ""
			return lines
		}
	}

	r := bufio.NewReader(file)

	var wg sync.WaitGroup

	for {
		buf := linePool.Get().([]byte)

		n, err := r.Read(buf)
		buf = buf[:n]

		if n == 0 {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}

		nextUntillNewLine := r.ReadBytes('\n')

		if err != io.EOF {
			buf = append(buf, nextUntillNewLine...)
		}

		wg.Add(1)
		go func() {
			ProcessChunk(buf, &linePool, &stringPool, startTime, finishTime)
			wg.Done()
		}()
	}
	wg.Wait()
	return nil
}

func ProcessChunk(buf []byte, linePool *sync.Pool, stringPool *sync.Pool, startTime time.Time, finishTime time.Time) {
	var wg2 sync.WaitGroup

	logs := stringPool.Get().(string)
	logs = string(buf)

	linePool.Put(buf)

	logsSlice := strings.Split(logs, "\n")

	stringPool.Put(logs)

	chunkSize := 300
	n := len(logsSlice)
	noOfThreads := n / chunkSize

	if n%chunkSize != 0 {
		noOfThreads++
	}

	for i := 0; i < noOfThreads; i++ {
		wg2.Add(1)
		go func(s int, e int) {
			defer wg2.Done()
			for j := s; j < e; j++ {
				text := logsSlice[j]
				if len(text) == 0 {
					continue
				}
				logSlice := strings.Split(text, ",", 2)
				logCreationTimeString := logSlice[0]

				logCreationTime, err := time.Parse("2006-01-02 15:04:05", logCreationTimeString)
				if err != nil {
					fmt.Println(err)
					return
				}

				if logCreationTime.Before(finishTime) && logCreationTime.After(startTime) {
					fmt.Println(text)
				}
			}
		}(i * chunkSize, int(math.Min(float64(i*chunkSize+chunkSize), float64(len(logsSlice)))))
	}
	wg2.Wait()
	logsSlice = nil
}