open System

let firstTask =
    async {
        printfn "First Task Started"
        do! Async.Sleep(2000)
        printfn "First Task Completed"
        return "First Task Result"
    }

let secondTask =
    async {
        printfn "Second Task Started"
        do! Async.Sleep(1000)
        printfn "Second Task Completed"
        return "Second Task Result"
    }

let thirdTask =
    async {
        printfn "Third Task Started"
        do! Async.Sleep(1500)
        printfn "Third Task Completed"
        return "Third Task Result"
    }

printfn "Starting All Tasks..."

let output =
    [firstTask; secondTask; thirdTask]
    |> Async.Parallel
    |> Async.RunSynchronously

printfn "\nAll Tasks Finished!"

printfn "\nDisplaying Results:"

output
|> Array.iter (fun result ->
    printfn "%s" result
)

printfn "\nProgram Completed."