# Appointy-Instagram-API-task
## Requirements
* Go lang
* MongoDB
* Postman
* Text editor

## Operations performed by the API
* Create an user
* Get an user by ID
* Create a post
* Get a post by post ID
* List all the posts posted by a particular user

## How to run:
#### Start the MongoDB server
* Open the command prompt and type `mongo`

![mongo](https://user-images.githubusercontent.com/69240053/136656817-c4b20cb9-3eda-4fdb-9e92-376438051213.png)
* Now run the program by typing the command:
```
go run main.go
```
![image](https://user-images.githubusercontent.com/69240053/136656839-85210188-578b-4a0f-97d7-e39ba68087ac.png)

#### Open Postman
* Create a user
  * Select `POST` method and type `127.0.0.1:8080/users`. Select `body` & switch to `JSON` and type the `user` details and hit `send`
  
  ![POST user](https://user-images.githubusercontent.com/69240053/136656330-85ee7e0e-f906-4446-bf70-e1a0f3fbe12d.png)
* Get an user by ID
  * Select `GET` method and type `127.0.0.1:8080/users/1` and hit `send`
  * Password will be `hashed` after creating a user
  
  ![GET userById](https://user-images.githubusercontent.com/69240053/136656589-b5754cbf-29cd-437a-8ce9-77def8c82a01.png)
* Create a post
  * Select `POST` method and type `127.0.0.1:8080/posts`. Select `body` & switch to `JSON` and type the `post` details and hit `send`
  
  ![POST post](https://user-images.githubusercontent.com/69240053/136656667-3ddaa08b-a0b2-4ea3-9f82-f448793918d6.png)
* Get a post by PID
  * Select `GET` method and type `127.0.0.1:8080/posts/3` and hit `send`
  
  ![GET postById](https://user-images.githubusercontent.com/69240053/136656698-7745e771-7ea4-4046-a01f-6cf2d751af9d.png)
* List all the posts posted by a user
  * Select `GET` method and type `127.0.0.1:8080/posts/users/1` and hit `send`
  
  ![GET postsByUserById](https://user-images.githubusercontent.com/69240053/136656743-185e2a14-70f3-4f88-b315-93c02747940b.png)

#### Threads are made safe by using `mutex.Lock()` and `mutex.Unlock`
## Testing:
`Unit Tests` are created in `main_test.go` file.
To run the tests type the following command:
```
go test -v
```
