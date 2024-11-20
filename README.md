# simple-go-crm

## description

My cap stone project for the Udacity **Golang** course: a simple CRUD-backend to manage "customers".

## installation

- set up a local go environment (min. go1.23.0)
- clone this repo

## launch & usage

- within this folder, in a terminal, type: `go run simple-crm`
- the API will be exposed under `http://localhost:3000`

This project is backend-only, so you have to interact with the REST API in order to use it. Consider using an API client like Postman or Bruno.

### Available API operations

#### `GET /customers`

- response: [`Customer[]`](#customer)
- returns all saved customers

#### `GET /customers/{id}`

- response: [`Customer`](#customer)
- returns the customer matching the id
- if no customer matches the id, returns a 404 error
- if the id is not valid, returns a 400 error

#### `POST /customers`

- request payload: [`CustomerCreateDTO`](#customercreatedto)
- response: [`Customer`](#customer)
- creates a new customer with a random id and returns it
- if the payload is not valid, returns a 400 error

#### `PUT /customers/{id}`

**todo**

#### `DELETE /customers/{id}`

**todo**

## model

### `Customer`

```typescript
{
	"id":        UUID
	"name":      string
	"role":      string
	"email":     string
	"phone":     string
	"contacted": boolean
}
```

### `CustomerCreateDTO`

```typescript
{
	"name":      string
	"role":      string
	"email":     string
	"phone":     string
	"contacted": boolean
}
```
