Project Brief:
    - This will be a Bookstore Management System 
    - We will interact with said system using APIs
    - The project will comprise of:
        - Database - MySQL
        - GORM
        - JSON marshall, unmarshall
        - Gorilla Mux
        - We will set-up a proper file structure for the project

Project Structure:
    - CMD
        - MAIN.GO
    - PKG
        - config folder
            - APP.GO - Will aid in connecting to database
        - controllers folder
            - BOOK-CONTROLLER - Will have the functions to process the response and the data
        - models folder
            - BOOK.GO - Will have the structs and models used by database
        - routes folder
            - BOOKSTORE-ROUTES
                - controller func | routes | method
                - GET BOOKS | /book/ | GET
                - CREATE BOOK | /book/ | POST
                - GET BOOK BY ID | /book/{bookId} | GET
                - UPDATE BOOKS | /book/{bookId} | PUT
                - DELETE BOOK | /book/{bookId} | DELETE
        - utils folder
            - UTILS.GO - Used for JSON marshalling and marshalling of requests

Project Packages:
    - MUX: go get "github.com/gorilla/mux"
    - GORM: go get "github.com/jinzhu/gorm"
    - GORM: go get "github.com/jinzhu/gorm/dialects/mysql"