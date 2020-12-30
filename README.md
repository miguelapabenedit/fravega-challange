El propósito de esta API es proveer un servicio para la Alta y Consulta de sucursales (Branches), como también localizar la ubicación mas cercana de un deliver dada una posición.

## Contents
 - [Stack Tecnológico](#stack)
 - [Como utilizarlo](#como-utilizarlo)
 - [Comentarios Generales Desarrollo](#comentarios)
 - [Agradecimientos](#agradecimientos)
 
 ## Stack Tecnológico (#stack)
 
 1. [Docker - Docker Compose](https://www.docker.com/ 
 2. Testing y Mock con [Stretchr/Testify](https://github.com/stretchr/testify)
 4. Back Go 1.15
 3. Swagger con [Swag](https://github.com/swaggo/swag)
 5. Routing [Gorilla](https://github.com/gorilla)
 6. [Sql 2017](https://hub.docker.com/_/microsoft-mssql-server)
 
 ## Cómo utilizarlo
 1. Download fravega-challange utilizando el comando (#como-utilizarlo).
 ```sh
 $ go get -u https://github.com/miguelapabenedit/fravega-challange
 ```
 2.En el root del repositorio correr el comando 
 ```sh
 $ docker-compose up -d
 ```
 3.Esperar a que los contenedores se inicialicen y correr swagger 
 ```sh
http://localhost:8080/swagger/index.html
 ```
 4.Las configuraciones son manejadas dentro del archivo Dockerfile situados en
 ```sh
 ./Dockefile -- Go Api
 ./infrastructure/Dockerfile -- Sql server
 ```
 ## Comentarios Generales
 
 Fue un desafío interesante que me di un tiempo de 2 días para resolver:
 1. Diseño de la aplicación me llevo 25% del tiempo
 2. Swagger: nunca use swagger con golang por lo que tuve un par de inconvenientes al momento de entender como mostrarlo en el mismo dominio que la api, tuve un par de 
 ncuentros con swagger-go/swagger-go pero termine inclinandome por Swag. Gaste un 40% para poder resolver ese tema( un montón lo se), dado que en .Net es un poco mas sensillo.
 4. Docker y docker-compose: conseguir la información y armar los scripts fue relativamente fácil, pero lo que me genero un dolor de cabeza fue intentar seedear la base de datos 
 en conjunto con el build. La mayoría del tiempo lo gaste intentando resolver el seed 35%
 
 Para el diseño del algoritmo de encontrar el punto mas cercano me base en este [articulo escrito por Jan Philip Matuschek](http://janmatuschek.de/LatitudeLongitudeBoundingCoordinates) 
 y algunas busquedas sobre terminos en wikipedia.
 
 ## Agradecimientos
 1. Gracias a [Federico Alonso Alomá](https://www.linkedin.com/in/federico-alonso-aloma/) por el challange y el contacto
 2. Gracias a la persona de IT que dedique tiempo en leer mi codigo este fragmento de código y en especial si tiene feedback sobre el mismo.
 3. Gracias a [Rodrigo Auche](https://www.linkedin.com/in/rodrigoauche/) que me ayudo con un blockeo que tuve con Docker.
