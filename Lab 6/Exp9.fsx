open System

printf "Enter marks of Subject 1: "
let a = int(Console.ReadLine())

printf "Enter marks of Subject 2: "
let b = int(Console.ReadLine())

printf "Enter marks of Subject 3: "
let c = int(Console.ReadLine())

printf "Enter marks of Subject 4: "
let d = int(Console.ReadLine())

let total = a + b + c + d
let percentage = float total / 4.0

printfn "Total Marks = %d" total
printfn "Percentage = %.2f%%" percentage

match percentage with
| p when p >= 60.0 -> printfn "Division = First Division"
| p when p >= 50.0 -> printfn "Division = Second Division"
| p when p >= 40.0 -> printfn "Division = Third Division"
| _ -> printfn "Division = Fail"