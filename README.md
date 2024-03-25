# Final_Project MyGram

## Installation

1. Clone the repository.
2. Install dependencies using `go mod tidy`

## Set Up Database
1. Open terminal and type
   mysql-u (your username) -p (your pass)
2. Create database
   create database mygramm

## Running the Application

To start the application, use the following command:

go run main.go

The application will run on http://localhost:8080 by default.


## Endpoints API

### Users

- **Register User:** `POST /users/register`
- **Login User:** `POST /users/login`
- **Update User:** `PUT /users/:userId`
- **Delete User:** `DELETE /users/:userId`

### Photos

- **Create Photo:** `POST /photos`
- **Get Photos:** `GET /photos`
- **Update Photo:** `PUT /photos/:photoId`
- **Delete Photo:** `DELETE /photos/:photoId`

### Comments

- **Create Comment:** `POST /comments`
- **Get Comments:** `GET /comments`
- **Update Comment:** `PUT /comments/:commentId`
- **Delete Comment:** `DELETE /comments/:commentId`

### Social Medias

- **Create Social Media:** `POST /socialmedias`
- **Get Social Medias:** `GET /socialmedias`
- **Update Social Media:** `PUT /socialmedias/:socialMediaId`
- **Delete Social Media:** `DELETE /socialmedias/:socialMediaId`

## Koleksi Postman

 Dapat menemukan koleksi Postman untuk API di [sini](https://api.postman.com/collections/23870699-e42c594b-4593-4cbf-b742-599ae0152801?access_key=PMAT-01HSV6GXN08KM9GWGH73AK535T).

