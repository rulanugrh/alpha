<h1 align="center"> ALPHA PROJECT </h1>
<div align="center">
    <img src="https://wallpapercave.com/wp/wp2763910.gif" />
</div>

<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif">

Heyyo omaetachi, this opportunity I will make a project to learn about microservices. In short this project is about a library system named `Alpha`. It consists of 3 services namely `book`, `order`, and `user` service. Each service sends notifications to each other via a message broker, and is centralized to one point, namely the api gateway using `nginx`.

| Feature             | Status              | 
|---------------------|---------------------|
| Authentication      | :heavy_check_mark:  |
| Message Broker      | :heavy_check_mark:  |
| Book                | :heavy_check_mark:  |
| User                | :heavy_check_mark:  |
| Dockerize           | :heavy_check_mark:  |
| API Gateway         | :heavy_check_mark:  |
| Frontend            | Ongoing             |
| Pipeline            | Ongoing             |
| Order               | :heavy_check_mark:  |
| Kurbenetes          | Ongoing             |
| Monitoring          | Ongoing             |

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

> First, installastion kurbenetes if u want use it
```console
alpha@chentauri:~$ sudo adduser kurbenetes
alpha@chentauri:~$ sudo usermod -aG sudo kurbenetes
alpha@chentauri:~$ su kurbenetes
alpha@chentauri:~$ sudo apt install docker docker.io docker-compose -y
```

> Step 2: Install k8s with wget
```console
alpha@chentauri:~$ sudo wget https://github.com/kubernetes/minikube/releases/download/v1.24.0/minikube-linux-amd64
alpha@chentauri:~$ sudo mv minikube-linux-amd64 minikube
```

> Step 3: Change Permission minikube n save to /usr/local/bin/minukube
```console
alpha@chentauri:~$ sudo mv minikube /usr/local/bin/minikube
```

> Step 4: Configure kubectl
```console
alpha@chentauri:~$ sudo curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
alpha@chentauri:~$ sudo chmod +x kubectl
alpha@chentauri:~$ sudo mv kubectl /user/local/bin/kubectl
```

> Step 5: minikube can start n you can see detail
```console
alpha@chentauri:~$ minikube start
* minikube v1.24.0 on Debian 10.2
* Using the docker driver based on existing profile

X The requested memory allocation of 1970MiB does not leave room for system overhead (total system memory: 1970MiB). You may face stability issues.
* Suggestion: Start minikube with less memory allocated: 'minikube start --memory=1970mb'
---

alpha@chentauri:~$ kubectl version
Client Version: version.Info{Major:"1", Minor:"22", GitVersion:"v1.22.3", GitCommit:"c92036820499fedefec0f847e2054d824aea6cd1", GitTreeState:"clean", BuildDate:"2023-25-27T18:41:28Z", GoVersion:"go1.16.9", Compiler:"gc", Platform:"linux/amd64"}
Server Version: version.Info{Major:"1", Minor:"22", GitVersion:"v1.22.3", GitCommit:"c92036820499fedefec0f847e2054d824aea6cd1", GitTreeState:"clean", BuildDate:"2023-25-27T18:35:25Z", GoVersion:"go1.16.9", Compiler:"gc", Platform:"linux/amd64"}
```

> If you install from docker you can must do this
```console
alpha@chentauri:~$ sudo git clone https://github.com/rulanugrh/alpha
alpha@chentauri:~$ chmod +x install
alpha@chentauri:~$ ./install app
```

<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif">


> I learning by opinion public, so if you have something opinion you can create issue, thanks :u 