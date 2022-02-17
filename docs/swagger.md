


# 42 Tutors API
An API for 42 tutors.
  

## Informations

### Version

0.0.1

### Contact

  

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json

### Produces
  * application/json

## Access control

### Security Schemes

#### OauthSecurity



> **Type**: oauth2
>
> **Flow**: accessCode
>
> **Authorization URL**: https://api.intra.42.fr/oauth/authorize
>
> **Token URL**: https://api.intra.42.fr/oauth/token
      

##### Scopes

Name | Description
-----|-------------
user | User scope

### Security Requirements
  * OauthSecurity: user

## All endpoints

###  authentication

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /auth/callback | [get auth callback](#get-auth-callback) | Return user token |
| GET | /login | [get login](#get-login) | Login user |
  


###  event

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /events | [create event](#create-event) | Create event |
| DELETE | /events/{id} | [delete event](#delete-event) | Delete event |
| GET | /events | [list event](#list-event) | List events |
| GET | /events/{id}/users | [list event users](#list-event-users) | List event users |
| GET | /events/{id} | [read event](#read-event) | Read event |
| PUT | /events/{id} | [update event](#update-event) | Update event |
  


###  user

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /users | [create user](#create-user) | Create user |
| DELETE | /users/{id} | [delete user](#delete-user) | Delete user |
| GET | /users | [list user](#list-user) | List users |
| GET | /users/{id}/events | [list user events](#list-user-events) | List user events |
| GET | /users/{id} | [read user](#read-user) | Read user |
| POST | /users/{userId}/events/{eventId} | [subscribe user](#subscribe-user) | Subscribe user |
| DELETE | /users/{userId}/events/{eventId} | [unsubscribe user](#unsubscribe-user) | Unsubscribe user |
| PUT | /users/{id} | [update user](#update-user) | Update user |
  


## Paths

### <span id="get-auth-callback"></span> Return user token (*GetAuthCallback*)

```
GET /auth/callback
```

Retrieve token from 42 API.

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-auth-callback-200) | OK | OK |  | [schema](#get-auth-callback-200-schema) |
| [500](#get-auth-callback-500) | Internal Server Error | Internal Server Error |  | [schema](#get-auth-callback-500-schema) |

#### Responses


##### <span id="get-auth-callback-200"></span> 200 - OK
Status: OK

###### <span id="get-auth-callback-200-schema"></span> Schema
   
  



##### <span id="get-auth-callback-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="get-auth-callback-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="get-login"></span> Login user (*GetLogin*)

```
GET /login
```

Login to 42 API with OAuth 2.0.

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [302](#get-login-302) | Found | Found |  | [schema](#get-login-302-schema) |
| [500](#get-login-500) | Internal Server Error | Internal Server Error |  | [schema](#get-login-500-schema) |

#### Responses


##### <span id="get-login-302"></span> 302 - Found
Status: Found

###### <span id="get-login-302-schema"></span> Schema

##### <span id="get-login-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="get-login-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="create-event"></span> Create event (*createEvent*)

```
POST /events
```

Create a new event.

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| event | `body` | [Event](#event) | `models.Event` | | ✓ | | Event to create. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-event-201) | Created | Created |  | [schema](#create-event-201-schema) |
| [400](#create-event-400) | Bad Request | Bad request |  | [schema](#create-event-400-schema) |
| [500](#create-event-500) | Internal Server Error | Internal Server Error |  | [schema](#create-event-500-schema) |

#### Responses


##### <span id="create-event-201"></span> 201 - Created
Status: Created

###### <span id="create-event-201-schema"></span> Schema
   
  

[Event](#event)

##### <span id="create-event-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="create-event-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="create-event-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="create-event-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="create-user"></span> Create user (*createUser*)

```
POST /users
```

Create a new user.

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| user | `body` | [User](#user) | `models.User` | | ✓ | | User to create. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-user-201) | Created | Created |  | [schema](#create-user-201-schema) |
| [400](#create-user-400) | Bad Request | Bad request |  | [schema](#create-user-400-schema) |
| [500](#create-user-500) | Internal Server Error | Internal Server Error |  | [schema](#create-user-500-schema) |

#### Responses


##### <span id="create-user-201"></span> 201 - Created
Status: Created

###### <span id="create-user-201-schema"></span> Schema
   
  

[User](#user)

##### <span id="create-user-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="create-user-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="create-user-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="create-user-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="delete-event"></span> Delete event (*deleteEvent*)

```
DELETE /events/{id}
```

Delete an event by ID.

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | Event ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#delete-event-204) | No Content | No Content |  | [schema](#delete-event-204-schema) |
| [400](#delete-event-400) | Bad Request | Bad request |  | [schema](#delete-event-400-schema) |
| [404](#delete-event-404) | Not Found | Not Found |  | [schema](#delete-event-404-schema) |
| [500](#delete-event-500) | Internal Server Error | Internal Server Error |  | [schema](#delete-event-500-schema) |

#### Responses


##### <span id="delete-event-204"></span> 204 - No Content
Status: No Content

###### <span id="delete-event-204-schema"></span> Schema

##### <span id="delete-event-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="delete-event-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="delete-event-404"></span> 404 - Not Found
Status: Not Found

###### <span id="delete-event-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="delete-event-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="delete-event-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="delete-user"></span> Delete user (*deleteUser*)

```
DELETE /users/{id}
```

Delete an user by ID.

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | User ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#delete-user-204) | No Content | No Content |  | [schema](#delete-user-204-schema) |
| [400](#delete-user-400) | Bad Request | Bad request |  | [schema](#delete-user-400-schema) |
| [404](#delete-user-404) | Not Found | Not Found |  | [schema](#delete-user-404-schema) |
| [500](#delete-user-500) | Internal Server Error | Internal Server Error |  | [schema](#delete-user-500-schema) |

#### Responses


##### <span id="delete-user-204"></span> 204 - No Content
Status: No Content

###### <span id="delete-user-204-schema"></span> Schema

##### <span id="delete-user-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="delete-user-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="delete-user-404"></span> 404 - Not Found
Status: Not Found

###### <span id="delete-user-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="delete-user-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="delete-user-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="list-event"></span> List events (*listEvent*)

```
GET /events
```

List all events.

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| day | `query` | string | `string` |  |  |  | Day filter. |
| month | `query` | string | `string` |  |  |  | Month filter. |
| week | `query` | string | `string` |  |  |  | Week filter. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-event-200) | OK | OK |  | [schema](#list-event-200-schema) |
| [400](#list-event-400) | Bad Request | Bad request |  | [schema](#list-event-400-schema) |
| [500](#list-event-500) | Internal Server Error | Internal Server Error |  | [schema](#list-event-500-schema) |

#### Responses


##### <span id="list-event-200"></span> 200 - OK
Status: OK

###### <span id="list-event-200-schema"></span> Schema
   
  

[][Event](#event)

##### <span id="list-event-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="list-event-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="list-event-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="list-event-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="list-event-users"></span> List event users (*listEventUsers*)

```
GET /events/{id}/users
```

List users subscribed to an event.

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | Event ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-event-users-200) | OK | OK |  | [schema](#list-event-users-200-schema) |
| [400](#list-event-users-400) | Bad Request | Bad request |  | [schema](#list-event-users-400-schema) |
| [404](#list-event-users-404) | Not Found | Not Found |  | [schema](#list-event-users-404-schema) |
| [500](#list-event-users-500) | Internal Server Error | Internal Server Error |  | [schema](#list-event-users-500-schema) |

#### Responses


##### <span id="list-event-users-200"></span> 200 - OK
Status: OK

###### <span id="list-event-users-200-schema"></span> Schema
   
  

[][User](#user)

##### <span id="list-event-users-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="list-event-users-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="list-event-users-404"></span> 404 - Not Found
Status: Not Found

###### <span id="list-event-users-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="list-event-users-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="list-event-users-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="list-user"></span> List users (*listUser*)

```
GET /users
```

List all users.

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tutor | `query` | boolean | `bool` |  |  |  | List all tutors. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-user-200) | OK | OK |  | [schema](#list-user-200-schema) |
| [400](#list-user-400) | Bad Request | Bad request |  | [schema](#list-user-400-schema) |
| [500](#list-user-500) | Internal Server Error | Internal Server Error |  | [schema](#list-user-500-schema) |

#### Responses


##### <span id="list-user-200"></span> 200 - OK
Status: OK

###### <span id="list-user-200-schema"></span> Schema
   
  

[][User](#user)

##### <span id="list-user-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="list-user-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="list-user-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="list-user-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="list-user-events"></span> List user events (*listUserEvents*)

```
GET /users/{id}/events
```

List user's subscribed events.

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | User ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-user-events-200) | OK | OK |  | [schema](#list-user-events-200-schema) |
| [400](#list-user-events-400) | Bad Request | Bad request |  | [schema](#list-user-events-400-schema) |
| [404](#list-user-events-404) | Not Found | Not Found |  | [schema](#list-user-events-404-schema) |
| [500](#list-user-events-500) | Internal Server Error | Internal Server Error |  | [schema](#list-user-events-500-schema) |

#### Responses


##### <span id="list-user-events-200"></span> 200 - OK
Status: OK

###### <span id="list-user-events-200-schema"></span> Schema
   
  

[][Event](#event)

##### <span id="list-user-events-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="list-user-events-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="list-user-events-404"></span> 404 - Not Found
Status: Not Found

###### <span id="list-user-events-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="list-user-events-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="list-user-events-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="read-event"></span> Read event (*readEvent*)

```
GET /events/{id}
```

Read an event by ID.

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | Event ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#read-event-200) | OK | OK |  | [schema](#read-event-200-schema) |
| [400](#read-event-400) | Bad Request | Bad request |  | [schema](#read-event-400-schema) |
| [404](#read-event-404) | Not Found | Not Found |  | [schema](#read-event-404-schema) |
| [500](#read-event-500) | Internal Server Error | Internal Server Error |  | [schema](#read-event-500-schema) |

#### Responses


##### <span id="read-event-200"></span> 200 - OK
Status: OK

###### <span id="read-event-200-schema"></span> Schema
   
  

[Event](#event)

##### <span id="read-event-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="read-event-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="read-event-404"></span> 404 - Not Found
Status: Not Found

###### <span id="read-event-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="read-event-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="read-event-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="read-user"></span> Read user (*readUser*)

```
GET /users/{id}
```

Read an user by ID.

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | User ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#read-user-200) | OK | OK |  | [schema](#read-user-200-schema) |
| [400](#read-user-400) | Bad Request | Bad request |  | [schema](#read-user-400-schema) |
| [404](#read-user-404) | Not Found | Not Found |  | [schema](#read-user-404-schema) |
| [500](#read-user-500) | Internal Server Error | Internal Server Error |  | [schema](#read-user-500-schema) |

#### Responses


##### <span id="read-user-200"></span> 200 - OK
Status: OK

###### <span id="read-user-200-schema"></span> Schema
   
  

[User](#user)

##### <span id="read-user-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="read-user-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="read-user-404"></span> 404 - Not Found
Status: Not Found

###### <span id="read-user-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="read-user-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="read-user-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="subscribe-user"></span> Subscribe user (*subscribeUser*)

```
POST /users/{userId}/events/{eventId}
```

Subscribe user to an event.

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| eventId | `path` | string | `string` |  | ✓ |  | Event ID. |
| userId | `path` | string | `string` |  | ✓ |  | User ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#subscribe-user-201) | Created | OK |  | [schema](#subscribe-user-201-schema) |
| [400](#subscribe-user-400) | Bad Request | Bad request |  | [schema](#subscribe-user-400-schema) |
| [404](#subscribe-user-404) | Not Found | Not Found |  | [schema](#subscribe-user-404-schema) |
| [500](#subscribe-user-500) | Internal Server Error | Internal Server Error |  | [schema](#subscribe-user-500-schema) |

#### Responses


##### <span id="subscribe-user-201"></span> 201 - OK
Status: Created

###### <span id="subscribe-user-201-schema"></span> Schema
   
  



##### <span id="subscribe-user-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="subscribe-user-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="subscribe-user-404"></span> 404 - Not Found
Status: Not Found

###### <span id="subscribe-user-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="subscribe-user-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="subscribe-user-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="unsubscribe-user"></span> Unsubscribe user (*unsubscribeUser*)

```
DELETE /users/{userId}/events/{eventId}
```

Unsubscribe user to an event.

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| eventId | `path` | string | `string` |  | ✓ |  | Event ID. |
| userId | `path` | string | `string` |  | ✓ |  | User ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#unsubscribe-user-204) | No Content | No Content |  | [schema](#unsubscribe-user-204-schema) |
| [400](#unsubscribe-user-400) | Bad Request | Bad request |  | [schema](#unsubscribe-user-400-schema) |
| [404](#unsubscribe-user-404) | Not Found | Not Found |  | [schema](#unsubscribe-user-404-schema) |
| [500](#unsubscribe-user-500) | Internal Server Error | Internal Server Error |  | [schema](#unsubscribe-user-500-schema) |

#### Responses


##### <span id="unsubscribe-user-204"></span> 204 - No Content
Status: No Content

###### <span id="unsubscribe-user-204-schema"></span> Schema

##### <span id="unsubscribe-user-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="unsubscribe-user-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="unsubscribe-user-404"></span> 404 - Not Found
Status: Not Found

###### <span id="unsubscribe-user-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="unsubscribe-user-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="unsubscribe-user-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="update-event"></span> Update event (*updateEvent*)

```
PUT /events/{id}
```

Update an event by ID.

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | Event ID. |
| event | `body` | [Event](#event) | `models.Event` | | ✓ | | Event to update. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-event-200) | OK | OK |  | [schema](#update-event-200-schema) |
| [400](#update-event-400) | Bad Request | Bad request |  | [schema](#update-event-400-schema) |
| [404](#update-event-404) | Not Found | Not Found |  | [schema](#update-event-404-schema) |
| [500](#update-event-500) | Internal Server Error | Internal Server Error |  | [schema](#update-event-500-schema) |

#### Responses


##### <span id="update-event-200"></span> 200 - OK
Status: OK

###### <span id="update-event-200-schema"></span> Schema
   
  

[Event](#event)

##### <span id="update-event-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="update-event-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="update-event-404"></span> 404 - Not Found
Status: Not Found

###### <span id="update-event-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="update-event-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="update-event-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="update-user"></span> Update user (*updateUser*)

```
PUT /users/{id}
```

Update an user by ID.

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | User ID. |
| user | `body` | [User](#user) | `models.User` | | ✓ | | User to update. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-user-200) | OK | OK |  | [schema](#update-user-200-schema) |
| [400](#update-user-400) | Bad Request | Bad request |  | [schema](#update-user-400-schema) |
| [404](#update-user-404) | Not Found | Not Found |  | [schema](#update-user-404-schema) |
| [500](#update-user-500) | Internal Server Error | Internal Server Error |  | [schema](#update-user-500-schema) |

#### Responses


##### <span id="update-user-200"></span> 200 - OK
Status: OK

###### <span id="update-user-200-schema"></span> Schema
   
  

[User](#user)

##### <span id="update-user-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="update-user-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="update-user-404"></span> 404 - Not Found
Status: Not Found

###### <span id="update-user-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="update-user-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="update-user-500-schema"></span> Schema
   
  

[Error](#error)

## Models

### <span id="error"></span> Error


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| code | integer| `int64` | ✓ | |  | `500` |
| message | string| `string` | ✓ | |  | `Internal Server Error` |



### <span id="event"></span> Event


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| createdAt | date-time (formatted string)| `strfmt.DateTime` |  | |  | `2022-02-15T09:00:00Z` |
| description | string (formatted string)| `string` |  | |  | `Exam Alone in the Dark - Exam Rank 2/3/4/5/6` |
| endAt | date-time (formatted string)| `strfmt.DateTime` |  | |  | `2022-02-15T13:00:00+01:00` |
| id | uuid (formatted string)| `strfmt.UUID` |  | |  | `123e4567-e89b-12d3-a456-426614174000` |
| name | string (formatted string)| `string` |  | |  | `Exam Stud` |
| startAt | date-time (formatted string)| `strfmt.DateTime` |  | |  | `2022-02-15T10:00:00+01:00` |
| tutorsRequired | int64 (formatted integer)| `int64` |  | |  | `3` |
| tutorsSubscribed | int64 (formatted integer)| `int64` |  | |  | `0` |
| users | [][User](#user)| `[]*User` |  | |  |  |
| walletsReward | int64 (formatted integer)| `int64` |  | |  | `200` |



### <span id="user"></span> User


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| adminScope | boolean| `bool` |  | |  | `false` |
| events | [][Event](#event)| `[]*Event` |  | |  |  |
| firstName | string (formatted string)| `string` |  | |  | `Gaëtan` |
| hoursDone | int64 (formatted integer)| `int64` |  | |  | `6` |
| id | uuid (formatted string)| `strfmt.UUID` |  | |  | `123e4567-e89b-12d3-a456-426614174000` |
| lastName | string (formatted string)| `string` |  | |  | `Marchal` |
| login | string (formatted string)| `string` |  | |  | `gamarcha` |
| tutorScope | boolean| `bool` |  | |  | `true` |



### <span id="principal"></span> principal


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| principal | string| string | |  |  |


