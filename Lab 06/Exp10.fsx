open System

printf "Enter a number: "
let n = int(Console.ReadLine())

let rec isPrime num divisor =
    if num <= 1 then
        false
    elif divisor * divisor > num then
        true
    elif num % divisor = 0 then
        false
    else
        isPrime num (divisor + 1)

if isPrime n 2 then
    printfn "%d is a Prime Number" n
else
    printfn "%d is Not a Prime Number" n