package channels

import "strings"

// SendRange sends all integers between start and end (inclusive) on the output channel and then closes the channel.
// For example, if the input is start=3 end=7 it will send 3, 4, 5, 6, 7
// If start is > end, closes output without sending anything.
func SendRange(start, end int, output chan int) {
	// TODO implementation
	if start <= end {
		for i := start; i <= end; i++ {
			output <- i
		}
	}
	close(output)
}

// BuildString receives string chunks on the input channel and concatenates them together into one long string.
// Once the input channel is closed, it sends the completed concatenation on result and closes result.
// For example, for input "hi, " "how are" " you?" {close} it should return "hi, how are you?"
func BuildString(input chan string, result chan string) {
	// TODO implementation
	concatenatedString := ""
	for inputString := range input {
		concatenatedString += inputString
	}
	result <- concatenatedString
	close(result)
}

// ConvertToLowercase converts all strings on the input channel to lowercase and sends them on the output channel.
// It continues to do this until one of two things happens:
//  1. The input channel is closed
//  2. A message comes on the interrupt channel
//
// After one of these two things happens, it closes the output channel and returns.
// Notes:
//
//	You can check if a channel is closed by reading like this: val, ok := <-input
//	Check the "strings" package for  string utility functions.
func ConvertToLowercase(input <-chan string, output chan<- string, interrupt <-chan struct{}) {
	// TODO implementation
	for {
		select {
		case inputString, ok := <-input:
			if !ok {
				// input channel closed, close the output channel and return
				close(output)
				return
			}
			lowerCaseInputString := strings.ToLower(inputString)
			output <- lowerCaseInputString
		case <-interrupt:
			// interrupt received, close the output channel and return
			close(output)
			return
		}
	}
}
