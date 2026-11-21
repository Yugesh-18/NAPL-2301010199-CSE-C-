open System
open System.Text.RegularExpressions

let reverseText (text: string) =
    new string(Array.rev (text.ToCharArray()))

let checkPalindrome (text: string) =
    text = reverseText text

let changeString (text: string) =
    if text |> Seq.forall Char.IsLetter then
        text.ToUpper()
    elif text |> Seq.forall Char.IsDigit then
        let value = Int32.Parse(text)
        string (value * 2)
    else
        text

let checkEmail (address: string) =
    Regex.IsMatch(address, @"^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$")

printfn "STRING PROCESSING PROGRAM"

printfn "\n1. Reverse and Palindrome Check"

["madam"; "computer"; "level"]
|> List.iter (fun word ->
    printfn "Original: %s" word
    printfn "Reversed: %s" (reverseText word)
    printfn "Palindrome: %b" (checkPalindrome word)
)

printfn "\n2. String Transformation"

["welcome"; "456"; "abc123"]
|> List.iter (fun value ->
    printfn "%s -> %s" value (changeString value)
)

printfn "\n3. Email Validation"

["student@gmail.com"; "invalid123"; "user@yahoo.in"]
|> List.iter (fun mail ->
    if checkEmail mail then
        printfn "%s : Valid Email" mail
    else
        printfn "%s : Invalid Email" mail
)