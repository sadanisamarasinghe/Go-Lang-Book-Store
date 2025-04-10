# Create
curl -X POST http://localhost:3000/books -H "Content-Type: application/json" -d '{"title":"GoLang 101", "author":"Sandani"}'

# Read All
curl http://localhost:3000/books

# Read One
curl http://localhost:3000/books/1

# Update
curl -X PUT http://localhost:3000/books/1 -H "Content-Type: application/json" -d '{"title":"Updated Book", "author":"Sandani"}'

# Delete
curl -X DELETE http://localhost:3000/books/1
