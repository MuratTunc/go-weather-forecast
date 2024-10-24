# WEATHER SERVICE APPLICATION VIA GOLANG

In this study, I will try to introduce the simplest postgreSQL applicaiton using the Golang programming language. Intermediate level knowledge of Golang language will be sufficient.

> ## Working environment
>
> - Ubuntu Operating System (Linux)
> - Golang sw language
> - PostgreSQL
> - DBeaver
> - VS Code IDE

## General hierarchy of the project
![hierarchy.PNG](pictures/Weather-service-hiearchy.PNG)




## Start the program
Go under projects directory.
Type _sudo make up_build_ comment from terminal.
After this command, binary files of _Broker_ and _weather_ services will be created.

Type _make up_ command from terminal. After this command, docker images will be started in the background.

The last lines of the console output will be as follows

 ```
 project_broker-service_1 is up-to-date
 project_weather-service_1 is up-to-date
 Docker images started!
  ```

Type _make start_ command from terminal. After this command, the front end binary will be created and run.
```
Building front end binary...
cd ../front-end && env CGO_ENABLED=0 /usr/local/go/bin/go build -o frontApp ./cmd/web
Done!
Starting front end
cd ../front-end && ./frontApp
WEB PAGE FRONT END SERVICES IS STARTED!!! AT LOCALHOST PORT :8090
```

# Github source code


In this project there a 2 services which are Broker and Weather.
Broker service is the main entrance os the HTTP request from user web application.
After getting user action, it will send this to Weather service.

Broker service bahaves like a Gateaway services.
__**I wish you good coding, thinking that it is educational and fun.**__


author:Murat Tunç
17.10.2024
