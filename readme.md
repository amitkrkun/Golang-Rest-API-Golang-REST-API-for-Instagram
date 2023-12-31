![Logo](https://miro.medium.com/max/920/1*CdjOgfolLt_GNJYBzI-1QQ.jpeg)

# Go REST API

A web based API made using Native Golang .

Routes include

- User Routes -> These routes include functionality such as fetchUserById and AddUser.

- Authentication Route -> Authentication Implimented using JWT

- File Upload -> Achived using Buckets with gridfs and using Methods like OpenUploadStream

- File Download-> Implimented using DownloadStreamByName Methods

For Database connectivity MongoDB has been used .

MongoDB can be accessed via a container for easy access the syntax for which is given below

```bash
docker run --name RESTApiDB -p 27017:27017 -d mongo:latest
```

## Run Locally

Clone the project

```bash
  git clone https://github.com/NimeshJohari02/Golang-RestAPI.git
```

Go to the project directory

```bash
  cd Golang-RestAPI
```

Install dependencies

```bash
go mod init
go mod tidy
```

Start the RestAPI

```bash
go run main.go
```

## API Reference

#### Get all items

```http
  GET /user/getUserById
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of user to fetch |

Dummy Response

```
Response :
{
    "Email": "\"john.doe@gmail.com\"",
    "Password": "$2a$14$C0pdzLah2gIHdttH2kbiOOf55mwHgEdlewV1Jlt2nyK8E8Jo.PSga",
    "Posts": null,
    "_id": "6161ad1923c42f4188930f8a",
    "id": "7a62cbbc-7f70-48db-b740-7be5fef57328",
    "name": "\"John Doe\""
}
```

#### Add User To Database

```http
  POST /user/addUser/
```

| Parameter   | Type     | Description                      |
| :---------- | :------- | :------------------------------- |
| `Name `     | `string` | Your Name as query Parameter     |
| `Email `    | `string` | Your Email as query Parameter    |
| `Password ` | `string` | Your Password as query Parameter |

Note: The Password must be hashed from the source of input . Functions for that have been provided in the API . The crypto library was used for this purpose and HashPassword Function has been implimented for
the same

#### Adding a POST by a Particular User

To add A post Under a user with a given Id following are the query parma

| Parameter     | Type     | Description                                               |
| :------------ | :------- | :-------------------------------------------------------- |
| `Id `         | `string` | UUID Generated By uuid package                            |
| `Title `      | `string` | Title to Your Post passed as a String                     |
| `Description` | `string` | A paragraph long Description of the post                  |
| `urlToImage`  | `string` | URI to the local image location                           |
| `publishedAt` | `string` | TimeStamp Generated using time.now()                      |
| `fileName`    | `string` | Unique fileName Generated by using gridFS                 |
| `userId`      | `string` | id that links the post and the user that created the post |

## Running Tests

To run tests, run the following command
The unit tests can be found nested inside the test directory

```bash
  go test -v
```

## TODO

- [ ] Containerization Of Application using Docker

- [ ] Deployment of the API on PostMan

