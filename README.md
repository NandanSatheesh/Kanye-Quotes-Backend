# Kanye-Quotes-Backend
Backend REST API for Kanye Quotes App (Go and MySQL)


## Installation:

Please set up tables on your local MySQL cluster before running quotes.go   
Add MySQL connection parameters wherever necessary in quotes.go and writedata.go 

Schema followed - 
```
CREATE TABLE `quotes_data`.`quotesq` (
  `idquotes` INT NOT NULL AUTO_INCREMENT,
  `data` VARCHAR(45) NULL,
  PRIMARY KEY (`idquotes`));
```
Now run these
```
git clone https://github.com/NandanSatheesh/Kanye-Quotes-Backend.git
cd Kanye-Quotes-Backend
cd database 
go run writedata.go
cd ..
go run quotes.go
```
## Run:

Go to -  http://localhost:8080/quote in your browser or hit the URL using Postman (GET Request)

Sample Outputs 

```
{
quote: "People only get jealous when they care."
}
```
