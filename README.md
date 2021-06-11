# Description of the project
This is `grpc-gateway` project for working with library via database. You can get information about a book, you can add a book to the database, also you can delete or uprdate info about a book.

## Information about ports
`grpc-server` runs on 8080 port. 
`Gateway` runs on 8081 port. 
`Database` runs on 5432 port.
To work with this application you need to use URL = `localhost:8081/api/v1/{path}`

## How you can run this project
To run this application locally you need to clone this project and you need to have active docker server.

Application for get books with such operations as in table below:


|             Path            | Method | Description                           | Body example                                                                                                                                                                                                                     |
|:---------------------------:|--------|---------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| /books/api/v1                  | GET    | get all books                      | ```[{"BookId":1,"AuthorId":2,"PublisherId":1,"NameOfBook":"Belka","YearOfPublication":"2020-10-10", "BookVolume":20, "Number":1},{"BookId":2,"AuthorId":1,"PublisherId":4,"NameOfBook":"Strelka","YearOfPublication":"2021-12-21", "BookVolume":220, "Number":11},{"BookId":2,"AuthorId":3,"PublisherId":4,"NameOfBook":"Space","YearOfPublication":"2010-10-10", "BookVolume":202, "Number":11}]``` |
| /books/api/v1                   | POST   | create new book                    |                                                                                                                                                                                                                                  |
| /books/api/v1/{id}              | GET    | get book by the id                 | ```{"BookId":1,"AuthorId":2,"PublisherId":1,"NameOfBook":"Belka","YearOfPublication":"2020-10-10", "BookVolume":20, "Number":1}```                                                                                                                                  |
| /books/api/v1/{id}/{unit_price} | PUT    | update book's price by the id |                                                                                                                                                                                                                                  |
| /books/api/v1/{id}              | DELETE | delete book by the id              |                                                                                                                                                                                             

## How to generate go code from protobuf
To generate go code from protobuf, you need a command to enter into the terminal while you are in the project root:
`protoc -I . user.proto --grpc-gateway_out . --go_out=plugins=grpc:.`

## How to run application via cluster
To run application via cluster, you need create a cluster:
`k3d cluster create hello --port 1514:1514/UDP@loadbalancer`

Then, use command:
`kubectl apply -f ./kube-config`
