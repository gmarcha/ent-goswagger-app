


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

#### OAuth2



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
tutor | Tutor scope
event | Event scope
admin | Admin scope

## All endpoints

###  authentication

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /v2/auth/callback | [callback](#callback) | Receive token |
| GET | /v2/auth/login | [login](#login) | Login user |
| GET | /v2/auth/token/info | [token info](#token-info) | Send user information |
| GET | /v2/auth/token/refresh | [token refresh](#token-refresh) | Refresh token |
  


###  event

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /v2/events | [create event](#create-event) | Create event |
| DELETE | /v2/events/{id} | [delete event](#delete-event) | Delete event |
| GET | /v2/events | [list event](#list-event) | List events |
| GET | /v2/events/{id}/users | [list event users](#list-event-users) | List event users |
| GET | /v2/events/{id} | [read event](#read-event) | Read event |
| PUT | /v2/events/{id} | [update event](#update-event) | Update event |
  


###  user

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /v2/users | [create user](#create-user) | Create user |
| DELETE | /v2/users/me | [delete me](#delete-me) | Delete authenticated user |
| DELETE | /v2/users/{id} | [delete user](#delete-user) | Delete user |
| GET | /v2/users/me/events | [list me events](#list-me-events) | List authenticated user events |
| GET | /v2/users | [list user](#list-user) | List users |
| GET | /v2/users/{id}/events | [list user events](#list-user-events) | List user events |
| GET | /v2/users/me | [read me](#read-me) | Read authenticated user |
| GET | /v2/users/{id} | [read user](#read-user) | Read user |
| POST | /v2/users/me/events/{id} | [subscribe me](#subscribe-me) | Subscribe authenticated user |
| POST | /v2/users/{userId}/events/{eventId} | [subscribe user](#subscribe-user) | Subscribe user |
| DELETE | /v2/users/me/events/{id} | [unsubscribe me](#unsubscribe-me) | Unsubscribe authenticated user |
| DELETE | /v2/users/{userId}/events/{eventId} | [unsubscribe user](#unsubscribe-user) | Unsubscribe user |
| PUT | /v2/users/me | [update me](#update-me) | Update authenticated user |
| PUT | /v2/users/{id} | [update user](#update-user) | Update user |
  


## Paths

### <span id="callback"></span> Receive token (*callback*)

```
GET /v2/auth/callback
```

Receive token as a response from 42 API.

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#callback-200) | OK | OK |  | [schema](#callback-200-schema) |
| [401](#callback-401) | Unauthorized | Unauthorized |  | [schema](#callback-401-schema) |
| [500](#callback-500) | Internal Server Error | Internal Server Error |  | [schema](#callback-500-schema) |

#### Responses


##### <span id="callback-200"></span> 200 - OK
Status: OK

###### <span id="callback-200-schema"></span> Schema
   
  



##### <span id="callback-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="callback-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="callback-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="callback-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="create-event"></span> Create event (*createEvent*)

```
POST /v2/events
```

Create a new event.

#### Security Requirements
  * OAuth2: event

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| event | `body` | [Event](#event) | `ent.Event` | | ✓ | | Event to create. |

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
POST /v2/users
```

Create a new user.

#### Security Requirements
  * OAuth2: admin

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| user | `body` | [User](#user) | `ent.User` | | ✓ | | User to create. |

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
DELETE /v2/events/{id}
```

Delete an event by ID.

#### Security Requirements
  * OAuth2: event

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

### <span id="delete-me"></span> Delete authenticated user (*deleteMe*)

```
DELETE /v2/users/me
```

Delete the authenticated user.

#### Security Requirements
  * OAuth2: tutor

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#delete-me-204) | No Content | No Content |  | [schema](#delete-me-204-schema) |
| [500](#delete-me-500) | Internal Server Error | Internal Server Error |  | [schema](#delete-me-500-schema) |

#### Responses


##### <span id="delete-me-204"></span> 204 - No Content
Status: No Content

###### <span id="delete-me-204-schema"></span> Schema

##### <span id="delete-me-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="delete-me-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="delete-user"></span> Delete user (*deleteUser*)

```
DELETE /v2/users/{id}
```

Delete an user by ID.

#### Security Requirements
  * OAuth2: admin

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
GET /v2/events
```

List all events.

#### Security Requirements
  * OAuth2: tutor

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
GET /v2/events/{id}/users
```

List users subscribed to an event.

#### Security Requirements
  * OAuth2: tutor

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

### <span id="list-me-events"></span> List authenticated user events (*listMeEvents*)

```
GET /v2/users/me/events
```

List the authenticated user's subscribed events.

#### Security Requirements
  * OAuth2: tutor

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-me-events-200) | OK | OK |  | [schema](#list-me-events-200-schema) |
| [500](#list-me-events-500) | Internal Server Error | Internal Server Error |  | [schema](#list-me-events-500-schema) |

#### Responses


##### <span id="list-me-events-200"></span> 200 - OK
Status: OK

###### <span id="list-me-events-200-schema"></span> Schema
   
  

[][Event](#event)

##### <span id="list-me-events-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="list-me-events-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="list-user"></span> List users (*listUser*)

```
GET /v2/users
```

List all users.

#### Security Requirements
  * OAuth2: tutor

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
GET /v2/users/{id}/events
```

List user's subscribed events.

#### Security Requirements
  * OAuth2: tutor

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

### <span id="login"></span> Login user (*login*)

```
GET /v2/auth/login
```

Login a user with 42 API.

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [302](#login-302) | Found | Found |  | [schema](#login-302-schema) |
| [500](#login-500) | Internal Server Error | Internal Server Error |  | [schema](#login-500-schema) |

#### Responses


##### <span id="login-302"></span> 302 - Found
Status: Found

###### <span id="login-302-schema"></span> Schema

##### <span id="login-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="login-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="read-event"></span> Read event (*readEvent*)

```
GET /v2/events/{id}
```

Read an event by ID.

#### Security Requirements
  * OAuth2: tutor

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

### <span id="read-me"></span> Read authenticated user (*readMe*)

```
GET /v2/users/me
```

Read the authenticated user.

#### Security Requirements
  * OAuth2: tutor

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#read-me-200) | OK | OK |  | [schema](#read-me-200-schema) |
| [500](#read-me-500) | Internal Server Error | Internal Server Error |  | [schema](#read-me-500-schema) |

#### Responses


##### <span id="read-me-200"></span> 200 - OK
Status: OK

###### <span id="read-me-200-schema"></span> Schema
   
  

[User](#user)

##### <span id="read-me-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="read-me-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="read-user"></span> Read user (*readUser*)

```
GET /v2/users/{id}
```

Read an user by ID.

#### Security Requirements
  * OAuth2: tutor

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

### <span id="subscribe-me"></span> Subscribe authenticated user (*subscribeMe*)

```
POST /v2/users/me/events/{id}
```

Subscribe an authenticated user to an event.

#### Security Requirements
  * OAuth2: tutor

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | Event ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#subscribe-me-201) | Created | OK |  | [schema](#subscribe-me-201-schema) |
| [400](#subscribe-me-400) | Bad Request | Bad request |  | [schema](#subscribe-me-400-schema) |
| [404](#subscribe-me-404) | Not Found | Not Found |  | [schema](#subscribe-me-404-schema) |
| [500](#subscribe-me-500) | Internal Server Error | Internal Server Error |  | [schema](#subscribe-me-500-schema) |

#### Responses


##### <span id="subscribe-me-201"></span> 201 - OK
Status: Created

###### <span id="subscribe-me-201-schema"></span> Schema
   
  



##### <span id="subscribe-me-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="subscribe-me-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="subscribe-me-404"></span> 404 - Not Found
Status: Not Found

###### <span id="subscribe-me-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="subscribe-me-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="subscribe-me-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="subscribe-user"></span> Subscribe user (*subscribeUser*)

```
POST /v2/users/{userId}/events/{eventId}
```

Subscribe a user to an event.

#### Security Requirements
  * OAuth2: admin

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

### <span id="token-info"></span> Send user information (*tokenInfo*)

```
GET /v2/auth/token/info
```

Send authenticated user information or unauthorized error response.

#### Security Requirements
  * OAuth2: tutor

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#token-info-200) | OK | OK |  | [schema](#token-info-200-schema) |
| [401](#token-info-401) | Unauthorized | Unauthorized |  | [schema](#token-info-401-schema) |
| [500](#token-info-500) | Internal Server Error | Internal Server Error |  | [schema](#token-info-500-schema) |

#### Responses


##### <span id="token-info-200"></span> 200 - OK
Status: OK

###### <span id="token-info-200-schema"></span> Schema
   
  

[User](#user)

##### <span id="token-info-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="token-info-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="token-info-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="token-info-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="token-refresh"></span> Refresh token (*tokenRefresh*)

```
GET /v2/auth/token/refresh
```

Refresh access token if refresh token is still valid.

#### Security Requirements
  * OAuth2: tutor

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#token-refresh-200) | OK | OK |  | [schema](#token-refresh-200-schema) |
| [401](#token-refresh-401) | Unauthorized | Unauthorized |  | [schema](#token-refresh-401-schema) |
| [500](#token-refresh-500) | Internal Server Error | Internal Server Error |  | [schema](#token-refresh-500-schema) |

#### Responses


##### <span id="token-refresh-200"></span> 200 - OK
Status: OK

###### <span id="token-refresh-200-schema"></span> Schema
   
  



##### <span id="token-refresh-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="token-refresh-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="token-refresh-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="token-refresh-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="unsubscribe-me"></span> Unsubscribe authenticated user (*unsubscribeMe*)

```
DELETE /v2/users/me/events/{id}
```

Unsubscribe an authenticated user to an event.

#### Security Requirements
  * OAuth2: tutor

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | Event ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#unsubscribe-me-204) | No Content | No Content |  | [schema](#unsubscribe-me-204-schema) |
| [400](#unsubscribe-me-400) | Bad Request | Bad request |  | [schema](#unsubscribe-me-400-schema) |
| [404](#unsubscribe-me-404) | Not Found | Not Found |  | [schema](#unsubscribe-me-404-schema) |
| [500](#unsubscribe-me-500) | Internal Server Error | Internal Server Error |  | [schema](#unsubscribe-me-500-schema) |

#### Responses


##### <span id="unsubscribe-me-204"></span> 204 - No Content
Status: No Content

###### <span id="unsubscribe-me-204-schema"></span> Schema

##### <span id="unsubscribe-me-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="unsubscribe-me-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="unsubscribe-me-404"></span> 404 - Not Found
Status: Not Found

###### <span id="unsubscribe-me-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="unsubscribe-me-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="unsubscribe-me-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="unsubscribe-user"></span> Unsubscribe user (*unsubscribeUser*)

```
DELETE /v2/users/{userId}/events/{eventId}
```

Unsubscribe a user to an event.

#### Security Requirements
  * OAuth2: admin

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
PUT /v2/events/{id}
```

Update an event by ID.

#### Security Requirements
  * OAuth2: event

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | Event ID. |
| event | `body` | [Event](#event) | `ent.Event` | | ✓ | | Event to update. |

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

### <span id="update-me"></span> Update authenticated user (*updateMe*)

```
PUT /v2/users/me
```

Update the authenticated user.

#### Security Requirements
  * OAuth2: tutor

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| user | `body` | [User](#user) | `ent.User` | | ✓ | | User to update. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-me-200) | OK | OK |  | [schema](#update-me-200-schema) |
| [500](#update-me-500) | Internal Server Error | Internal Server Error |  | [schema](#update-me-500-schema) |

#### Responses


##### <span id="update-me-200"></span> 200 - OK
Status: OK

###### <span id="update-me-200-schema"></span> Schema
   
  

[User](#user)

##### <span id="update-me-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="update-me-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="update-user"></span> Update user (*updateUser*)

```
PUT /v2/users/{id}
```

Update an user by ID.

#### Security Requirements
  * OAuth2: admin

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | User ID. |
| user | `body` | [User](#user) | `ent.User` | | ✓ | | User to update. |

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
| message | string| `string` | ✓ | |  | `Explicit error message` |
| status | string| `string` | ✓ | |  | `Internal Server Error` |



### <span id="principal"></span> principal


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| principal | string| string | |  |  |


