# Project Description

1.  **users_crud**: This directory contains the Users CRUD application using a MySQL database. The main.go file initializes the HTTP server and defines the routes for CRUD operations. The db.go file manages the MySQL database connection. The user.go file defines the User model and contains the handlers for various CRUD operations.


2. **articles_crud**: This directory contains the Articles CRUD application with in-memory persistence. The main.go file initializes the HTTP server and defines the routes for CRUD operations. The article.go file defines the Article model and contains the handlers for various CRUD operations.


3. **categories_crud**: This directory contains the Category CRUD application with file-based persistence. The main.go file initializes the HTTP server and defines the routes for CRUD operations. The category.go file defines the Category model and contains the handlers for various CRUD operations. The categories.txt file acts as the data storage for category records.


4. **docker**: This run.sh file defines the Docker Compose configuration for running the entire project. It includes MySQL, the Users CRUD service, the Articles CRUD service, and the Category CRUD service.

## Instructions
Clone the repository and navigate to the project root directory.
Ensure you have Docker and Docker Compose installed on your system.
Run docker-compose up to build and run the containers for MySQL, Users CRUD, Articles CRUD, and Category CRUD.
1. Access the Users CRUD API at http://localhost:8080/user.
2. Access the Articles CRUD API at http://localhost:8080/article.
3. Access the Category CRUD API at http://localhost:8080/category.


## API Endpoints
### Users CRUD:
1. **GET** /users<span style="white-space: pre">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span> Get all users
2. **GET** /users/{id}<span style="white-space: pre">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span> Get a user by ID
3. **POST** /users<span style="white-space: pre">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span> Create a new user
4. **PUT** /users/{id}<span style="white-space: pre">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span> Update an existing user
5. **DELETE** /users/{id}<span style="white-space: pre">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span> Delete a user by ID
### Articles CRUD:
1. **GET**     /articles<span style="white-space: pre">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span>Get all articles
2. **GET**     /articles/{id}<span style="white-space: pre">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span> Get an article by ID
3. **POST**    /articles<span style="white-space: pre">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span> Create a new article
4. **PUT**     /articles/{id}<span style="white-space: pre">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span> Update an existing article
5. **DELETE**  /articles/{id}<span style="white-space: pre">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span> Delete an article by ID
### Category CRUD:
1. **GET** /categories<span style="white-space: pre">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span> Get all categories
2. **GET** /categories/{id}<span style="white-space: pre">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span> Get a category by ID
3. **POST** /categories<span style="white-space: pre">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span> Create a new category
4. **PUT** /categories/{id}<span style="white-space: pre">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span> Update an existing category
5. **DELETE** /categories/{id}<span style="white-space: pre">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span> Delete a category by ID


## Project Setup



Follow these steps to run the project using Docker:



### Prerequisites



- Docker: [Installation Guide](https://docs.docker.com/get-docker/)






1. Clone the project repository:



   ```bash
   git clone https://github.com/nawid-salim/go-crud
   ```



2. Change to the project directory:



   ```bash
   cd go-crud
   ```


3. Build && run Project locally with local MySQL <span style="white-space: pre">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span> **Note** :: update config/dev_config.json file



   ```bash
   go build -o go-crud && ./go-crud
   ```

## OR

### Run the Docker Composer for complete project setup


3. Start a Docker Composer:



```bash
./run.sh
```



This command runs the Docker composer and forwards the host's port 8080 to the container's port 8080.



###  Open your web browser and access http://localhost:8080 to see the project in action.




## Configuration



All Configuration set in config folder && update values according to your environment



## Contributing



If you want to contribute to this project, please follow the guidelines below:



1. Fork the repository.
2. Create a new branch: `git checkout -b feature/your-feature`.
3. Make your changes and commit them: `git commit -m 'Add some feature'`.
4. Push to the branch: `git push origin feature/your-feature`.
5. Submit a pull request.



## License



Specify the license under which the project is distributed. For example:



This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).