# Cloud Identity Manager

Cloud Identity Manager  (cidm) will take care of Authentication and Authorization mechanism as a service for a suite of services.

[![Go Report Card](https://goreportcard.com/badge/github.com/craguilar/cidm)](https://goreportcard.com/report/github.com/craguilar/cidm)

## Overview 
Actors

* User client . Actions executed via Web Browser for login only , this actions get translated to REST calls.
* Service clients (micro services), executed via SSL or any other secure method inside the Customer instance.
High level flow

* User client enters their login credentials
* Identity Manager Server verifies the credentials are correct and returns a signed JWT token .
* This token is stored client-side, most commonly in Web Browser local storage .
* Subsequent User clients requests to the API Gateway server include this token as an additional Authorization header.
* API Gateway server handover the request to Identity Manager Server decodes the JWT and if the token is valid processes the request
* Once a user logs out, the token is destroyed client-side, no interaction with the server is necessary.

### Authentication

Token-based authentication is stateless. The server does not keep a record of which users are logged in or which JWTs have 
been issued (at least not explicitly). Instead, every request to the server is accompanied by a token which the server uses to verify the authenticity of the request. The token is generally sent as an addition Authorization header in form of Bearer {JWT}
JWT has three distinct parts that are URL encoded for transport:

* Header: The header contains the metadata for the token and at a minimal contains the type of the signature and/or encryption algorithm
* Claims: The claims contains any information that you want signed
* JSON Web Signature (JWS): The headers and claims digitally signed using the algorithm in the specified in the header

Example of JWT Token

````
//header
{
    "alg": "HS256", //algorithm
    "typ": "JWT" //denotes the type (shorthand typ) of token this is
}
 
//claims
{
    "sub": "tom@stormpath.com",
    "name": "Tom Abbott",
    "role": "user"
}
````

REST API, full programatic access to the internals makes it easy to manage your API users, keys and API Configuration from within your systems


### Authorization 

Implement simple schema for authorization, based on this principle : who can use Role X to Action on Resoruce Y .

Roles:

* (GET)roles.Read (resource uri level)
* (PUT) roles.Update
* (POST)roles.Create
* (DELETE)roles.Delete

Resource: defined by Service level

## Key Features

* Authentication with OAuth 2 
* Basic schema for authorization .

## Installation

### Storage 

Install postgresql, reference 

Connect to the data base and run the data base creation script available in 

/resource/sql/initdb.sql

### Setup of OAuth 2  

In Google cloud console go to API & Services > OAuth 2.0 client IDs

### Code dependencies
Update dependencies with dep

````
dep ensure
````

### Deployment with Docker

Build docker

````
docker build -t cimd .
````

Run docker 

````
docker run --publish 30030:8080  --name cimd --rm cimd
````
### Deployment with 

Run install 
````
go install github.com/craguilar/cidm/cmd/cidm-server
````

## TODO 

See code hints TODO , for pending actions on this project.