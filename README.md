<div id="top"></div>



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="https://play-lh.googleusercontent.com/Qr1iuQY4wWmcpfuKt0t_7Yn82UDuZSFloOHSmRsQA1IgYggwBYvnfThNJ3JCP_rlB6A" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">BiTaksi Backend</h3>

  <p align="center">
    There are two APIs in this project, Rider and Matching.
    <br />
  </p>
</div>

<!-- ABOUT THE PROJECT -->
## About The Project

I have used more than one package in this project. The project was completed in 13 days. At each stage, I learned new things and improved my knowledge of the golang. My biggest motivation when developing this project was to develop a project with Golang and get a job at Bitaksi. Now just below I will refer to some packages that I use in this project.




### Packages Used

List of main packages used

* [GoFiber](https://github.com/gofiber/fiber)
* [MongoDB](https://www.mongodb.com/languages/golang)
* [Zap](https://github.com/uber-go/zap)
* [Jwt-Go](github.com/dgrijalva/jwt-go)
* [Swagger-Fiber](https://github.com/arsmn/fiber-swagger)
* [Testify](https://github.com/stretchr/testify)
* [GoDotEnv](https://github.com/joho/godotenv)




<!-- GETTING STARTED -->
## Getting Started


Download or fork projects to your computer. Then open the project in editors like Goland vs Vscode. I recommend goland editor.then open terminal screen type.

* go mod
  ```sh
  go mod tidy
  ```
  
For documentation run the following command

* swagger
  ```sh
  swag init --parseDependency --parseInternal
  ```
For Test

* Test
  ```sh
  go test testname_test.go  -v 
  ```
  
For Run

* Test
  ```sh
  go run main.go
  ```
  
  
<!-- USAGE EXAMPLES -->
## Usage

we have two api available and two swaggers. We have 6 endpoints in Driver Api, only one of them needs token usage, otherwise it will give an error. We need to use match endpoint with token. the token we need to use is just below.

* Token
  ```sh
  Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQWhtZXQifQ._B6OV9qIyBevZi7oyfeQwTvF4SNDFD6LKgkLFGDZfwI
  ```
  I am gonna show an example from the matching api
  [![image](https://r.resimlink.com/hkgfsbFq.png)](https://resimlink.com/hkgfsbFq)

* New let's see our two api documentation
[![image](https://r.resimlink.com/oflkBK8O.png)](https://resimlink.com/oflkBK8O)
<a href="https://resimlink.com/Ss0Ena3RXuh" title="ResimLink - Resim Yükle"><img src="https://r.resimlink.com/Ss0Ena3RXuh.png" title="ResimLink - Resim Yükle" alt="ResimLink - Resim Yükle"></a>



