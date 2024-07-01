Certainly! Here's the entire project code formatted as Markdown for your `README.md` file:

```markdown
# INI Parser using Go

## Project Description

This Go library provides an INI file parser with methods to parse, access, edit, save, and write INI files. It adheres to specific INI file specifications:

- Assumes no global keys; every key must be part of a section.
- Uses = as the key-value separator.
- Trims spaces from keys and values.
- Allows comments only at the beginning of lines.

### Features

- Load and parse INI configuration from strings and files.
- Retrieve section names and their key-value pairs.
- Serialize parsed data to a map-like structure.
- Get and set values for keys within sections.
- Convert INI data to a string representation.
- Write parsed data back to an INI file.

## How to Use

1. **Import Package**

```go
import "github.com/codescalersinternships/INI_Parser_Mohamed_Riyad"
```

2. **Create a New Parser**

```go
parser := INIParser.NewParser()
```

3. **Load from a File**

```go
_ = parser.LoadFromFile("./test-files/test_file_1.ini")
val, _ := parser.Get("owner", "name")
fmt.Println(val) // Output: John Doe
```

4. **Load from a String**

```go
validStringInput := "[owner]\nname = John Doe\n"
_ = parser.LoadFromString(validStringInput)
```

5. **Get a Key Value from a Section**

```go
val, _ := parser.Get("owner", "name")
fmt.Println(val) // Output: John Doe
```

6. **Set a Value for a Key in a Section**

```go
_ = parser.Set("owner", "name", "Ryad")
val, _ := parser.Get("owner", "name")
fmt.Println(val) // Output: Eyad
```

7. **Get Section Names**

```go
sectionsNames, _ := parser.GetSectionNames()
```

8. **Get Parsed Data**

```go
sections, _ := parser.GetSections()
```

9. **Convert Data to String**

```go
str, _ := parser.ToString()
```

10. **Save Data to File**

```go
_ = parser.LoadFromFile("./test-files/test_file_1.ini")
_ = parser.Set("owner", "name", "Ryad")
_ = parser.SaveToFile("./test-files/output.ini")
```

## How to Test

To run tests, navigate to the root directory of your project and use:

```sh
go test ./...
```

Add the `-v` flag for more detailed test output:

```sh
go test -v ./...
```

Feel free to contribute by submitting issues or pull requests for improvements.

