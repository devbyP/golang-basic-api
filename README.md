# Sample REST api

for demonstration building a basic web rest api using go chi and pgx.

## Books Collection

will be build book collection api that collect book title and some simple info like author.

book attribute

- title: title of the book.
- author: author name.
- released_date: date of released.
- total_page: number of page.
- created_on: timestamp when the row is created.
- updated_on: timestamp when the row is last updated.

### Api operation

book api:

- GET All: /book - return all book in the collection.
    - can add get a number of items when using pagination.
- GET: /book/{bookId} - return 1 book info by id.
- PATCH: /book/{bookId} - update a book info.
- DELETE: /book/{bookId} - delete a book from collection.
- POST: /book - add new book to collection.
