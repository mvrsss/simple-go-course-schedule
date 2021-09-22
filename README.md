# Course Schedule Api with Go
Simple course scheduler using gorilla/mux and gorm

## Installation & Run
```
# Download the project
$ go get github.com/mvrsss/simple-go-course-schedule

# Download Gorilla Mux
$ go get github.com/gorilla/mux

# Download GORM
$ go get github.com/jinzhu/gorm
```
Setting DB in config/database.go
```
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "root",
			Password: "",
			Name:     "testdb",
			Charset:  "utf8",
		},
	}
}
```

## Structure
```
├── app
│   ├── routes.go                   // Router
│   ├── controllers                 // API core handlers
│   │   ├── PostCourses.go       
│   │   ├── ViewCourses.go       
│   └── models
│   |   └── Course.go     // APi Model
|   └── utils
|       ├── utils.go // generate id
|       ├── DataResponse.go // maps response code 
├── config
│   └── databases.go        // Configuration
├── migrate
|   └── course.go // Models for our application
└── main.go
```

## API
**/course**
Markup : * ```GET```: Get all courses
         * ```POST```: Create a new course

**/course/:id**
Markup : * ```GET```: Get a courses
         * ```PUT```: Update a course
         * ```DELETE```: Delete a course
         
Post Parameters
```
{
  "key": "Math 251",
  "title": "Discrete Mathematics",
  "description": "The main themes of the course are logic and proof, induction and recursion, discrete structures, set theory, combinatorics, algorithms, graph theory, and their applications"
}
```

