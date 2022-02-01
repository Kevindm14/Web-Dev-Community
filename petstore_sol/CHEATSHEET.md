# CheatSheet
This CheatSheet provides some common uses and tips for pop package.
### Install Pop
```go
go get github.com/gobuffalo/pop/...
```
### Install Soda
```go
go install github.com/gobuffalo/pop/soda
```

### Soda Commands
**Generate a database config file.**
```
soda g config
```
**Create the database for testing environment.**
```
soda create -e test
```
**Generate models.**
```
soda g model <model_name> <field1> <field2:type>
```
**Generate migrations.**
```go
soda  g migration pets // to generate fizz migrations
```
```go
soda  g sql pets // to generate sql migrations
```
**Run migrations.**
```
soda migrate -e test
```

### Pop Commands
**Select all records from pets table**
```sql
-- sql
select id, animal, breed, age, created_at, updated_at
from pets;
```

```go
// go
var pets []Pet
tx.All(&pets)
```

**Select all cats from pets table**
```sql
-- sql
select id, animal, breed, age, created_at, updated_at
from pets
where animal = 'cat';
```

```go
// go
var cats []Pet
tx.Where("animal = 'cat'").All(&pets)
```

**Select all cats with breed Abyssinian from pets table**
```sql
-- sql
select id, animal, breed, age, created_at, updated_at
from pets
where animal = 'cat' and breed = 'Abyssinian';
```

```go
// go
var cats []Pet
tx.Where("animal = 'cat'").Where("breed = 'Abyssinian'").All(&pets)
```

**Register a pet in database**
```sql
--sql
insert into pets(id, animal, breed, age, created_at, updated_at) values(1, 'cat', 'Abyssinian', 1, now(), now());
```

```go
// go
cat := Pet{
    ID: 1,
    Animal: "cat",
    Breed: "Abyssinian",
    Age: 1
}

tx.Create(&cat)
```

**Update a pet in database**
```sql
--sql
update pets set age = 2 where id = 1;
```

```go
// go
cat := Pet{
    ID: 1,
    Age: 2,
}

tx.Update(&cat)
```

**Delete a pet in database**
```sql
--sql
delete from pets where id = 1;
```

```go
// go
cat := Pet{
    ID: 1,
}

tx.Destroy(&cat)
```
### Run tests
```
./test.sh
```