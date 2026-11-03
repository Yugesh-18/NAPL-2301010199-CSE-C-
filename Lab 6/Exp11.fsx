open System
open System.Text.RegularExpressions

// Function to reverse a string
let reverseString (str: string) =
    new string(Array.rev (str.ToCharArray()))

// Function to check palindrome
let isPalindrome (str: string) =
    str = reverseString str

// Function to transform string
let transformString (str: string) =
    if str |> Seq.forall Char.IsLetter then
        str.ToUpper()
    elif str |> Seq.forall Char.IsDigit then
        let num = Int32.Parse(str)
        string (num * 2)
    else
        str

// Function to validate email
let isValidEmail (email: string) =
    Regex.IsMatch(email, @"^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$")

// Main Program
printfn "STRING PROCESSING"

// Reverse and Palindrome
printfn "\n1. Reverse and Palindrome Check"
["madam"; "hello"; "level"]
|> List.iter (fun s ->
    printfn "String: %s, Reverse: %s, Palindrome: %b"
        s (reverseString s) (isPalindrome s))

// String Transformation
printfn "\n2. String Transformation"
["hello"; "123"; "abc123"]
|> List.iter (fun s ->
    printfn "%s -> %s" s (transformString s))

// Email Validation
printfn "\n3. Email Validation"
["test@gmail.com"; "hello123"; "user@yahoo.in"]
|> List.iter (fun email ->
    printfn "%s -> %b" email (isValidEmail email))