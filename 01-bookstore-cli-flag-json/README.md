## Book Store Cli

## clone a repo
git clone https://github.com/akilans/golang-mini-projects.git

## go to the 01-bookstore-cli-flag-json dir
cd 01-bookstore-cli-flag-json

## build
go build

## run

## get books
./bookstore get --all
./bookstore get --id 5

## add a book with id ,title, author, price, image_url
./bookstore add --id 6 --title test-book --author akilan --price 200

## update a book with id ,title, author, price, image_url
./bookstore update --id 6 --title test-book-1 --author akilan1 --price 2001 

## delete a book by --id
./bookstore delete --id 6