package main

import (
	"fmt"
)

func IsPalindrome(s string) bool {
    if len(s) <= 1 {
        return true
    }

    if s[0] != s[len(s)-1] {
        return false
    }

    return IsPalindrome(s[1 : len(s)-1])
}

func main() {

    fmt.Println("Is 'kasur rusak' palindrome?", IsPalindrome("kasur rusak")) // Harusnya: true
    fmt.Println("Is 'level' palindrome?", IsPalindrome("level"))         // Harusnya: true
    fmt.Println("Is 'hello' palindrome?", IsPalindrome("hello"))         // Harusnya: false
    fmt.Println("Is 'a' palindrome?", IsPalindrome("a"))                 // Harusnya: true
    fmt.Println("Is '' palindrome?", IsPalindrome(""))                   // Harusnya: true
    fmt.Println("Is 'racecar' palindrome?", IsPalindrome("racecar"))     // Harusnya: true
    fmt.Println("Is 'ab' palindrome?", IsPalindrome("ab"))               // Harusnya: false
    fmt.Println("Is 'aba' palindrome?", IsPalindrome("aba"))             // Harusnya: true


}
