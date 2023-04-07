<h1 align="center"> ALPHA PROJECT </h1>
<div align="center">
    <img src="https://wallpapercave.com/wp/wp2763910.gif" />
</div>

<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif">

Heyyo omaetachi, this opportunity I will make a project to learn about microservices. In short this project is about a library system named `Alpha`. It consists of 3 services namely `book`, `order`, and `user` service. Each service sends notifications to each other via a message broker, and is centralized to one point, namely the api gateway using `nginx`.

| Feature             | Status              | 
|---------------------|---------------------|
| Authentication      | :heavy_check_mark:  |
| Message Broker      | Ongoing             |
| Book                | :heavy_check_mark:  |
| User                | :heavy_check_mark:  |
| Dockerize           | :heavy_check_mark:  |
| API Gateway         | :heavy_check_mark:  |
| Frontend            | Ongoing             |
| Pipeline            | Ongoing             |
| Order               | Ongoing             |
| Kurbenetes          | Ongoing             |

# Technology Used
- [`Docker`](https://www.docker.com) - Main tools to build this project
- [`Golang`](https://go.dev) - The progamming language i use to build order service and book services
- [`Javascript`](https://developer.mozilla.org/en-US/docs/Web/javascript) - The programming language i use to build user-services
- [`Nginx`](https://www.nginx.com/) - API Gateway I used to build this project
- [`Gorm`](https://gorm.io/gorm) - Framework i use to client for database mysql
- [`Sequelize`](https://sequelize.org) - Framework for javascript to client database mysql
- [`Mysql`](https://www.mysql.com) - The main database for this project
- [`Gin-Gonic`](https://gin-gonic.com) - Framework for golang web application
- [`Express`](https://google.com) - Framework for javascript web application

# Getting Started
This project is ongoing, and will continue to be updated at any time. For now there are 2 services namely `user` and `book`. For installation is quite easy 

> First, you must git clone this repository
```
git clone https://github.com/rulanugrh/alpha
```
> Second, if you not install dependecy you can run this command
```
chmod +x install
./install require
```
> If you have install all dependecy you can run this command
```
./install app
```
<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif">
