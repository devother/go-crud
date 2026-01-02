## Introduction  
A lightweight, easily integrable Swagger setup for Go (Golang) with intuitive baseline functionality, including CRUD operations and authentication from various sources.

## Description  
This project provides a starter template for building RESTful APIs in Go with built-in Swagger/OpenAPI documentation, database integration, and extensible authentication mechanisms.

### Progress  
- [x] Project structure  
- [x] Database connection  
- [x] Entity models  
- [ ] CRUD operations  
- [ ] JWT implementation  
- [ ] Authentication system  

### Quick Start 
```bash
# Clone the repository
git clone https://github.com/devother/go-crud.git
cd go-crud

# Install dependencies
go mod download

# Run in development mode (with hot reload)
make dev

# Or run normally
make run
```

### Technologies Used  
- **Go**: 1.24  
- **Framework**: Gin  
- **Database**: SQLite with GORM  
- **Hot Reload**: CompileDaemon  
- **Validation**: go-playground/validator  
- **Environment**: godotenv  

#### Key Dependencies  
- `github.com/gin-gonic/gin` – Web framework  
- `gorm.io/gorm` – ORM  
- `gorm.io/driver/sqlite` – SQLite driver  
- `github.com/go-playground/validator/v10` – Request validation  
- `github.com/joho/godotenv` – Environment configuration  
- `github.com/githubnemo/CompileDaemon` – Live reload utility  

*(A complete list of dependencies is available in `go.mod`.)*
