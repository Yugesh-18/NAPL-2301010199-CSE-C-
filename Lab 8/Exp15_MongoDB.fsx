#r "nuget: MongoDB.Driver"

open System
open MongoDB.Driver
open MongoDB.Bson

let mongoClient =
    new MongoClient("mongodb://localhost:27017")

let collegeDatabase =
    mongoClient.GetDatabase("StudentDatabase")

let studentCollection =
    collegeDatabase.GetCollection<BsonDocument>("Students")

let searchRecord =
    Builders<BsonDocument>.Filter.Eq("studentId", 301)

studentCollection.DeleteMany(
    Builders<BsonDocument>.Filter.Empty
) |> ignore

let displayRecords () =
    let allRecords =
        studentCollection.Find(
            Builders<BsonDocument>.Filter.Empty
        ).ToList()

    for student in allRecords do
        printfn "Student ID: %d, Name: %s, Age: %d"
            (student.GetValue("studentId").AsInt32)
            (student.GetValue("name").AsString)
            (student.GetValue("age").AsInt32)

printfn "1. INSERT RECORD"

let newStudent = new BsonDocument()

newStudent.Add("studentId", BsonInt32(301))
newStudent.Add("name", BsonString("Rahul"))
newStudent.Add("age", BsonInt32(20))

studentCollection.InsertOne(newStudent)

displayRecords()

printfn "\n2. DISPLAY RECORD"

displayRecords()

printfn "\n3. UPDATE RECORD"

let changeAge =
    Builders<BsonDocument>.Update.Set("age", 22)

studentCollection.UpdateOne(searchRecord, changeAge) |> ignore

displayRecords()

printfn "\n4. DELETE RECORD"

studentCollection.DeleteOne(searchRecord) |> ignore

printfn "Student record deleted successfully."

printfn "\nRemaining Records:"
displayRecords()

printfn "\nAdding a new student record..."

let anotherStudent = new BsonDocument()

anotherStudent.Add("studentId", BsonInt32(302))
anotherStudent.Add("name", BsonString("Vikram"))
anotherStudent.Add("age", BsonInt32(23))

studentCollection.InsertOne(anotherStudent)

printfn "\nFinal Student Records:"
displayRecords()