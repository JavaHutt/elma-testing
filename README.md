# Elma testing
Final test for elma's golang crash course
## Usage
Run the program in bash
```bash
go run main.go
```
Then input one or many URLs, use **space** for separate them!
For example:```bash
http://godoc.org http://google.com
```
If all the URLs are correct, you'll see counters of **"Go"** for each web page in the output.
Then you may enter more URLs, or end the program by tapping Ctrl + C or by typing `quit` in the console.
Note that total counter will accumulate all the occurrences from one session. 