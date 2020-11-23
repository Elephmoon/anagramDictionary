# anagramDictionary
This application is started using [docker](https://www.docker.com/) and [docker-compose](https://docs.docker.com/compose/).  
## First step
Before start you need copy from .env.example to .env and edit it.
### Table with parameters from .env file
| Parameter     | Value         | Description  |
| :-------------: |:-------------:| -----|
| DEBUG         | true/false    | When set to true, additional information is displayed in the logs. |
| DB_HOST       | postgres-dict | If you want to run the application from a third-party DB you need to edit |
| DB_PORT       | 5433          | If you want to run the application from a third-party DB you need to edit |
| DB_NAME       | dictionary    | If you want to run the application from a third-party DB you need to edit |
| DB_USER       |               | You need to specify a user to connect to DB |
| DB_PASSWORD   |               | You need to specify a password to connect to DB |
| API_PORT      | 3000          | The port that will listen to the application. Make sure that it is free in your system |
| API_HOST      | localhost     | Necessary for OpenAPI specification to work |
## Second step
1. You need start the application using docker-compose.  
In the directory with the project  
`sudo docker-compos up --build`  
2. You need to make sure that the application is running  
   In the console you should see a similar message about the start  
`dictionary       | time="2020-11-23T17:16:58Z" level=info msg="server start at :3000" program=anagramDictionary`  
3. Now you can go to http://localhost:3000/swagger-ui/ and use the application 
## Using API without swagger-ui
If you want to send requests directly to the application, you need to use **http://localhost:3000/api/v1/** also endpoints and data types from specification  
You can send requests using [Postman](https://www.postman.com/) or other applications
