package main

import (
	"fmt"
	"strings"
)

func main(){
	// reader := bufio.NewReader(os.Stdin);
	// input,_ := reader.ReadString('\n');

	   // Compare
    fmt.Println(strings.Compare("apple", "banana")) // -1 (apple < banana)
    fmt.Println(strings.Compare("banana", "apple")) // 1  (banana > apple)
    fmt.Println(strings.Compare("apple", "apple"))  // 0  (equal)

    // Contains
    fmt.Println(strings.Contains("Hello World", "World")) // true
    fmt.Println(strings.Contains("Hello World", "world")) // false (case-sensitive)

    // ContainsAny
    fmt.Println(strings.ContainsAny("Hello", "xyz")) // false
    fmt.Println(strings.ContainsAny("Hello", "aeiou")) // true (contains 'e' and 'o')

    // ContainsRune
    fmt.Println(strings.ContainsRune("Hello", 'H')) // true
    fmt.Println(strings.ContainsRune("Hello", 'z')) // false

    // HasPrefix
    fmt.Println(strings.HasPrefix("Hello World", "He")) // true
    fmt.Println(strings.HasPrefix("Hello World", "Wo")) // false

    // HasSuffix
    fmt.Println(strings.HasSuffix("Hello World", "ld")) // true
    fmt.Println(strings.HasSuffix("Hello World", "lo")) // false

    // EqualFold (case-insensitive compare)
    fmt.Println(strings.EqualFold("GoLang", "golang")) // true
    fmt.Println(strings.EqualFold("GoLang", "Golang!")) // false

    fmt.Println(strings.Count("banana", "a")) // 3
    fmt.Println(strings.Index("chicken", "ken")) // 4

    fmt.Println(strings.IndexAny("team", "aeiou")) // 1 ('e')
    fmt.Println(strings.IndexRune("Hello", 'o')) // 4
    fmt.Println(strings.LastIndex("go gopher go", "go")) // 10
    fmt.Println(strings.Split("a,b,c", ",")) // [a b c]
    fmt.Println(strings.SplitN("a,b,c", ",", 2)) // [a b,c]
    fmt.Println(strings.Fields(" foo bar baz ")) // [foo bar baz]
    f := func(r rune) bool { return r == ',' || r == ';' }
    fmt.Println(strings.FieldsFunc("a,b;c", f)) // [a b c]
    before, after, found := strings.Cut("foo:bar", ":")
    fmt.Println(before, after, found) // foo bar true
    fmt.Println(strings.Replace("oink oink oink", "oink", "moo", 2))// moo moo oink
    fmt.Println(strings.ReplaceAll("oink oink", "oink", "moo"))// moo moo
    mapping := func(r rune) rune {
    if r == 'a' { return 'x' }
    return r
    }
    fmt.Println(strings.Map(mapping, "banana")) // bxnxnx

    fmt.Println(strings.Trim("¡¡¡Hello!!!", "!¡")) // Hello
    fmt.Println(strings.TrimSpace("   Hello World   ")) // Hello World
    fmt.Println(strings.TrimPrefix("HelloWorld", "Hello")) // World
    fmt.Println(strings.TrimSuffix("HelloWorld", "World")) // Hello

    s := []string{"foo", "bar", "baz"}
    fmt.Println(strings.Join(s, ", ")) // foo, bar, baz
    fmt.Println(strings.Repeat("na", 3)) // nanana

    fmt.Println(strings.ToLower("GoLang")) // golang
    fmt.Println(strings.ToUpper("GoLang")) // GOLANG
    fmt.Println(strings.Title("go language")) // Go Language
    c := strings.Clone("Hello")
    fmt.Println(c == "Hello") // true
}