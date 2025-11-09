package main
import (
	"fmt"
	"strings"

)

func Split(r rune) bool {
	return r == ':'|| r == '.' || r == ',' || r == ' ' || r == ';' || r == '!'
}


/*
Task1:  Word Frequency Count
Write a Go function that takes a string as input and returns a dictionary containing the frequency of each word in the string. Treat words in a case-insensitive manner and ignore punctuation marks.
[Optional]: Write test for your function

*/

func wordFreq(str string)map[string]int{
	// First I want to parse, and remove every punctuation mark and spaces between words
	// Convert sll words to lower case
	s := strings.FieldsFunc(str, Split)
	counter := make(map[string]int)

	for _, word := range s{
		key := strings.ToLower(word)
		counter[key] ++
	}

	return counter

}

/*
Task : Palindrome Check
Write a Go function that takes a string as input and checks whether it is a palindrome or not. A palindrome is a word, phrase, number, or other sequence of characters that reads the same forward and backward (ignoring spaces, punctuation, and capitalization).
[Optional]: Write test for your function
*/

func isPalindrome(str string)bool {
	
	// remove punctuations
	strippedStr := strings.FieldsFunc(str, Split)

	joinedStr := strings.Join(strippedStr, "")

	parsedStr := strings.ToLower(joinedStr)


	i, j := 0, len(parsedStr) - 1

	for i < j {
		if parsedStr[i] == parsedStr[j]{
			i++
			j--
		}else{
			return false
		}
	}

	return true
}






func main(){
	// Simple sentence
	str1 := "Hello world"
	fmt.Println(wordFreq(str1))

	// Case-insensitive count
	str2 := "Hello hello"
	fmt.Println(wordFreq(str2))

	// Test Case 3: Sentence with punctuation
	str3 := "Hello, world! Hello."
	fmt.Println(wordFreq(str3))
	// Expected Output: {"hello": 2, "world": 1}

	// Test Case 4: Empty string
	str4 := ""
	fmt.Println(wordFreq(str4))
	// Expected Output: {}

	// Test Case 5: Sentence with repeated words
	str5 := "The quick brown fox jumps over the lazy dog"
	fmt.Println(wordFreq(str5))
	// Expected Output: {"the": 2, "quick": 1, "brown": 1, "fox": 1, "jumps": 1, "over": 1, "lazy": 1, "dog": 1}

	// Test Case 6: Sentence with multiple spaces between words
	str6 :=  "hello hello world"
	fmt.Println(wordFreq(str6))
	// Expected Output: {"hello": 2, "world": 1}

	// Test Case 7: Sentence with only punctuation
	str7 :=  ",,,,,,,,,"
	fmt.Println(wordFreq(str7))
	// Expected Output: {}

	// Test Case 8: Sentence with mixed punctuation marks
	str8 := "This is a test; this, is a test!"
	fmt.Println(wordFreq(str8))
	// Expected Output: {"this": 2, "is": 2, "a": 2, "test": 2}

	// Test Case 9: Long sentence with varied punctuation and capitalization
	str9 := "It was a bright, sunny day. The birds were chirping; the sky was clear! Bright, sunny days are the best!"
	fmt.Println(wordFreq(str9)) 
	// Expected Output: {"it": 1, "was": 2, "a": 1, "bright": 2, "sunny": 2, "day": 1, "the": 3, "birds": 1, "were": 1, "chirping": 1, "sky": 1, "clear": 1, "days": 1, "are": 1, "best": 1}


	fmt.Println("==================== TESTING TASK 2 PALINDROME=====================")
	pStr := "camel.  "
	
	if isPalindrome(pStr){
		fmt.Printf("It is a Palindrome")
	}else{
		fmt.Printf("It is not a Palindrome")
	}
}




