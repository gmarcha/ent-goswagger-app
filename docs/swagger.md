


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
public | Public access
event | Read events
event:write | Read-write events
user | Read users
user:subscription | Subscribe to events
user:write | Read-write users

## All endpoints

###  authentication

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /v2/auth/callback | [callback](#callback) | Receive token |
| GET | /v2/auth/login | [login](#login) | Login user |
| GET | /v2/auth/token/info | [token info](#token-info) | Send token information |
| GET | /v2/auth/token/refresh | [token refresh](#token-refresh) | Refresh token |
  


###  event

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /v2/events | [create event](#create-event) | Create event |
| DELETE | /v2/events/{id} | [delete event](#delete-event) | Delete event |
| GET | /v2/events/{id}/types | [event type](#event-type) | Event type |
| GET | /v2/events | [list event](#list-event) | List events |
| GET | /v2/events/{id}/users | [list event users](#list-event-users) | List event users |
| GET | /v2/events/{id} | [read event](#read-event) | Read event |
| PUT | /v2/events/{id} | [update event](#update-event) | Update event |
  


###  event_type

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /v2/events/types | [create type](#create-type) | Create event type |
| DELETE | /v2/events/types/{id} | [delete type](#delete-type) | Delete event type |
| GET | /v2/events/types | [list type](#list-type) | List event types |
| GET | /v2/events/types/{id} | [read type](#read-type) | Read event type |
| POST | /v2/events/types/{id} | [update type](#update-type) | Update event type |
  


###  role

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /v2/users/{id}/role/admin | [add admin](#add-admin) | Add admin |
| POST | /v2/users/{id}/role/calendar | [add calendar](#add-calendar) | Add calendar |
| POST | /v2/users/{id}/role/tutor | [add tutor](#add-tutor) | Add tutor |
| DELETE | /v2/users/{id}/role/admin | [remove admin](#remove-admin) | Remove admin |
| DELETE | /v2/users/{id}/role/calendar | [remove calendar](#remove-calendar) | Remove calendar |
| DELETE | /v2/users/{id}/role/tutor | [remove tutor](#remove-tutor) | Remove tutor |
  


###  user

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /v2/users | [create user](#create-user) | Create user |
| DELETE | /v2/users/me | [delete me](#delete-me) | Delete authenticated user |
| DELETE | /v2/users/{id} | [delete user](#delete-user) | Delete user |
| GET | /v2/users/me/events | [list me events](#list-me-events) | List authenticated user events |
| GET | /v2/users/me/roles | [list me roles](#list-me-roles) | List authenticated user roles |
| GET | /v2/users | [list user](#list-user) | List users |
| GET | /v2/users/{id}/events | [list user events](#list-user-events) | List user events |
| GET | /v2/users/{id}/roles | [list user roles](#list-user-roles) | List user roles |
| GET | /v2/users/me | [read me](#read-me) | Read authenticated user |
| GET | /v2/users/{id} | [read user](#read-user) | Read user |
| POST | /v2/users/me/events/{id} | [subscribe me](#subscribe-me) | Subscribe authenticated user |
| POST | /v2/users/{userId}/events/{eventId} | [subscribe user](#subscribe-user) | Subscribe user |
| DELETE | /v2/users/me/events/{id} | [unsubscribe me](#unsubscribe-me) | Unsubscribe authenticated user |
| DELETE | /v2/users/{userId}/events/{eventId} | [unsubscribe user](#unsubscribe-user) | Unsubscribe user |
| PUT | /v2/users/me | [update me](#update-me) | Update authenticated user |
| PUT | /v2/users/{id} | [update user](#update-user) | Update user |
  


## Paths

### <span id="add-admin"></span> Add admin (*addAdmin*)

```
POST /v2/users/{id}/role/admin
```

Add admin role to a user.

#### Security Requirements
  * OAuth2: public, user, user:write

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | User ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#add-admin-201) | Created | Created |  | [schema](#add-admin-201-schema) |
| [400](#add-admin-400) | Bad Request | Bad request |  | [schema](#add-admin-400-schema) |
| [401](#add-admin-401) | Unauthorized | Unauthorized |  | [schema](#add-admin-401-schema) |
| [403](#add-admin-403) | Forbidden | Forbidden |  | [schema](#add-admin-403-schema) |
| [404](#add-admin-404) | Not Found | Not Found |  | [schema](#add-admin-404-schema) |
| [500](#add-admin-500) | Internal Server Error | Internal Server Error |  | [schema](#add-admin-500-schema) |

#### Responses


##### <span id="add-admin-201"></span> 201 - Created
Status: Created

###### <span id="add-admin-201-schema"></span> Schema
   
  

[][Role](#role)

##### <span id="add-admin-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="add-admin-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="add-admin-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="add-admin-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="add-admin-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="add-admin-403-schema"></span> Schema
   
  

[Error](#error)

##### <span id="add-admin-404"></span> 404 - Not Found
Status: Not Found

###### <span id="add-admin-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="add-admin-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="add-admin-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="add-calendar"></span> Add calendar (*addCalendar*)

```
POST /v2/users/{id}/role/calendar
```

Add calendar role to a user.

#### Security Requirements
  * OAuth2: public, user, user:write

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | User ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#add-calendar-201) | Created | Created |  | [schema](#add-calendar-201-schema) |
| [400](#add-calendar-400) | Bad Request | Bad request |  | [schema](#add-calendar-400-schema) |
| [401](#add-calendar-401) | Unauthorized | Unauthorized |  | [schema](#add-calendar-401-schema) |
| [403](#add-calendar-403) | Forbidden | Forbidden |  | [schema](#add-calendar-403-schema) |
| [404](#add-calendar-404) | Not Found | Not Found |  | [schema](#add-calendar-404-schema) |
| [500](#add-calendar-500) | Internal Server Error | Internal Server Error |  | [schema](#add-calendar-500-schema) |

#### Responses


##### <span id="add-calendar-201"></span> 201 - Created
Status: Created

###### <span id="add-calendar-201-schema"></span> Schema
   
  

[][Role](#role)

##### <span id="add-calendar-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="add-calendar-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="add-calendar-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="add-calendar-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="add-calendar-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="add-calendar-403-schema"></span> Schema
   
  

[Error](#error)

##### <span id="add-calendar-404"></span> 404 - Not Found
Status: Not Found

###### <span id="add-calendar-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="add-calendar-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="add-calendar-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="add-tutor"></span> Add tutor (*addTutor*)

```
POST /v2/users/{id}/role/tutor
```

Add tutor role to a user.

#### Security Requirements
  * OAuth2: public, user, user:write

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | User ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#add-tutor-201) | Created | Created |  | [schema](#add-tutor-201-schema) |
| [400](#add-tutor-400) | Bad Request | Bad request |  | [schema](#add-tutor-400-schema) |
| [401](#add-tutor-401) | Unauthorized | Unauthorized |  | [schema](#add-tutor-401-schema) |
| [403](#add-tutor-403) | Forbidden | Forbidden |  | [schema](#add-tutor-403-schema) |
| [404](#add-tutor-404) | Not Found | Not Found |  | [schema](#add-tutor-404-schema) |
| [500](#add-tutor-500) | Internal Server Error | Internal Server Error |  | [schema](#add-tutor-500-schema) |

#### Responses


##### <span id="add-tutor-201"></span> 201 - Created
Status: Created

###### <span id="add-tutor-201-schema"></span> Schema
   
  

[][Role](#role)

##### <span id="add-tutor-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="add-tutor-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="add-tutor-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="add-tutor-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="add-tutor-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="add-tutor-403-schema"></span> Schema
   
  

[Error](#error)

##### <span id="add-tutor-404"></span> 404 - Not Found
Status: Not Found

###### <span id="add-tutor-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="add-tutor-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="add-tutor-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="callback"></span> Receive token (*callback*)

```
GET /v2/auth/callback
```

Receive token as a response from 42 API.

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#callback-200) | OK | OK |  | [schema](#callback-200-schema) |
| [422](#callback-422) | Unprocessable Entity | Unprocessable Entity |  | [schema](#callback-422-schema) |
| [500](#callback-500) | Internal Server Error | Internal Server Error |  | [schema](#callback-500-schema) |

#### Responses


##### <span id="callback-200"></span> 200 - OK
Status: OK

###### <span id="callback-200-schema"></span> Schema
   
  

[Token](#token)

##### <span id="callback-422"></span> 422 - Unprocessable Entity
Status: Unprocessable Entity

###### <span id="callback-422-schema"></span> Schema
   
  

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
  * OAuth2: event, event:write, public

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| event | `body` | [Event](#event) | `ent.Event` | | ✓ | | Event to create. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-event-201) | Created | Created |  | [schema](#create-event-201-schema) |
| [400](#create-event-400) | Bad Request | Bad request |  | [schema](#create-event-400-schema) |
| [401](#create-event-401) | Unauthorized | Unauthorized |  | [schema](#create-event-401-schema) |
| [403](#create-event-403) | Forbidden | Forbidden |  | [schema](#create-event-403-schema) |
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

##### <span id="create-event-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="create-event-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="create-event-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="create-event-403-schema"></span> Schema
   
  

[Error](#error)

##### <span id="create-event-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="create-event-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="create-type"></span> Create event type (*createType*)

```
POST /v2/events/types
```

Create a new event category.

#### Security Requirements
  * OAuth2: event, event:write, public

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| type | `body` | [EventType](#event-type) | `ent.EventType` | | ✓ | | Event category to create. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-type-201) | Created | Created |  | [schema](#create-type-201-schema) |
| [400](#create-type-400) | Bad Request | Bad request |  | [schema](#create-type-400-schema) |
| [401](#create-type-401) | Unauthorized | Unauthorized |  | [schema](#create-type-401-schema) |
| [403](#create-type-403) | Forbidden | Forbidden |  | [schema](#create-type-403-schema) |
| [500](#create-type-500) | Internal Server Error | Internal Server Error |  | [schema](#create-type-500-schema) |

#### Responses


##### <span id="create-type-201"></span> 201 - Created
Status: Created

###### <span id="create-type-201-schema"></span> Schema
   
  

[EventType](#event-type)

##### <span id="create-type-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="create-type-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="create-type-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="create-type-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="create-type-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="create-type-403-schema"></span> Schema
   
  

[Error](#error)

##### <span id="create-type-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="create-type-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="create-user"></span> Create user (*createUser*)

```
POST /v2/users
```

Create a new user.

#### Security Requirements
  * OAuth2: public, user, user:write

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| user | `body` | [User](#user) | `ent.User` | | ✓ | | User to create. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-user-201) | Created | Created |  | [schema](#create-user-201-schema) |
| [400](#create-user-400) | Bad Request | Bad request |  | [schema](#create-user-400-schema) |
| [401](#create-user-401) | Unauthorized | Unauthorized |  | [schema](#create-user-401-schema) |
| [403](#create-user-403) | Forbidden | Forbidden |  | [schema](#create-user-403-schema) |
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

##### <span id="create-user-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="create-user-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="create-user-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="create-user-403-schema"></span> Schema
   
  

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
  * OAuth2: event, event:write, public

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | Event ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#delete-event-204) | No Content | No Content |  | [schema](#delete-event-204-schema) |
| [400](#delete-event-400) | Bad Request | Bad request |  | [schema](#delete-event-400-schema) |
| [401](#delete-event-401) | Unauthorized | Unauthorized |  | [schema](#delete-event-401-schema) |
| [403](#delete-event-403) | Forbidden | Forbidden |  | [schema](#delete-event-403-schema) |
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

##### <span id="delete-event-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="delete-event-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="delete-event-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="delete-event-403-schema"></span> Schema
   
  

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
  * OAuth2: public

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#delete-me-204) | No Content | No Content |  | [schema](#delete-me-204-schema) |
| [401](#delete-me-401) | Unauthorized | Unauthorized |  | [schema](#delete-me-401-schema) |
| [500](#delete-me-500) | Internal Server Error | Internal Server Error |  | [schema](#delete-me-500-schema) |

#### Responses


##### <span id="delete-me-204"></span> 204 - No Content
Status: No Content

###### <span id="delete-me-204-schema"></span> Schema

##### <span id="delete-me-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="delete-me-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="delete-me-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="delete-me-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="delete-type"></span> Delete event type (*deleteType*)

```
DELETE /v2/events/types/{id}
```

Update an event category.

#### Security Requirements
  * OAuth2: event, event:write, public

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | ID of event category. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#delete-type-204) | No Content | No content |  | [schema](#delete-type-204-schema) |
| [400](#delete-type-400) | Bad Request | Bad request |  | [schema](#delete-type-400-schema) |
| [401](#delete-type-401) | Unauthorized | Unauthorized |  | [schema](#delete-type-401-schema) |
| [403](#delete-type-403) | Forbidden | Forbidden |  | [schema](#delete-type-403-schema) |
| [404](#delete-type-404) | Not Found | Not Found |  | [schema](#delete-type-404-schema) |
| [500](#delete-type-500) | Internal Server Error | Internal Server Error |  | [schema](#delete-type-500-schema) |

#### Responses


##### <span id="delete-type-204"></span> 204 - No content
Status: No Content

###### <span id="delete-type-204-schema"></span> Schema

##### <span id="delete-type-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="delete-type-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="delete-type-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="delete-type-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="delete-type-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="delete-type-403-schema"></span> Schema
   
  

[Error](#error)

##### <span id="delete-type-404"></span> 404 - Not Found
Status: Not Found

###### <span id="delete-type-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="delete-type-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="delete-type-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="delete-user"></span> Delete user (*deleteUser*)

```
DELETE /v2/users/{id}
```

Delete an user by ID.

#### Security Requirements
  * OAuth2: public, user, user:write

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | User ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#delete-user-204) | No Content | No Content |  | [schema](#delete-user-204-schema) |
| [400](#delete-user-400) | Bad Request | Bad request |  | [schema](#delete-user-400-schema) |
| [401](#delete-user-401) | Unauthorized | Unauthorized |  | [schema](#delete-user-401-schema) |
| [403](#delete-user-403) | Forbidden | Forbidden |  | [schema](#delete-user-403-schema) |
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

##### <span id="delete-user-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="delete-user-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="delete-user-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="delete-user-403-schema"></span> Schema
   
  

[Error](#error)

##### <span id="delete-user-404"></span> 404 - Not Found
Status: Not Found

###### <span id="delete-user-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="delete-user-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="delete-user-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="event-type"></span> Event type (*eventType*)

```
GET /v2/events/{id}/types
```

Retrieve event type.

#### Security Requirements
  * OAuth2: event, public

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | Event ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#event-type-200) | OK | OK |  | [schema](#event-type-200-schema) |
| [400](#event-type-400) | Bad Request | Bad request |  | [schema](#event-type-400-schema) |
| [401](#event-type-401) | Unauthorized | Unauthorized |  | [schema](#event-type-401-schema) |
| [403](#event-type-403) | Forbidden | Forbidden |  | [schema](#event-type-403-schema) |
| [404](#event-type-404) | Not Found | Not Found |  | [schema](#event-type-404-schema) |
| [500](#event-type-500) | Internal Server Error | Internal Server Error |  | [schema](#event-type-500-schema) |

#### Responses


##### <span id="event-type-200"></span> 200 - OK
Status: OK

###### <span id="event-type-200-schema"></span> Schema
   
  

[EventType](#event-type)

##### <span id="event-type-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="event-type-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="event-type-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="event-type-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="event-type-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="event-type-403-schema"></span> Schema
   
  

[Error](#error)

##### <span id="event-type-404"></span> 404 - Not Found
Status: Not Found

###### <span id="event-type-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="event-type-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="event-type-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="list-event"></span> List events (*listEvent*)

```
GET /v2/events
```

List all events.

#### Security Requirements
  * OAuth2: public

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
| [401](#list-event-401) | Unauthorized | Unauthorized |  | [schema](#list-event-401-schema) |
| [403](#list-event-403) | Forbidden | Forbidden |  | [schema](#list-event-403-schema) |
| [500](#list-event-500) | Internal Server Error | Internal Server Error |  | [schema](#list-event-500-schema) |

#### Responses


##### <span id="list-event-200"></span> 200 - OK
Status: OK

###### <span id="list-event-200-schema"></span> Schema
   
  

[][Event](#event)

##### <span id="list-event-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="list-event-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="list-event-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="list-event-403-schema"></span> Schema
   
  

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
  * OAuth2: event, public, user

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | Event ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-event-users-200) | OK | OK |  | [schema](#list-event-users-200-schema) |
| [400](#list-event-users-400) | Bad Request | Bad request |  | [schema](#list-event-users-400-schema) |
| [401](#list-event-users-401) | Unauthorized | Unauthorized |  | [schema](#list-event-users-401-schema) |
| [403](#list-event-users-403) | Forbidden | Forbidden |  | [schema](#list-event-users-403-schema) |
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

##### <span id="list-event-users-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="list-event-users-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="list-event-users-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="list-event-users-403-schema"></span> Schema
   
  

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
  * OAuth2: event, public, user

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-me-events-200) | OK | OK |  | [schema](#list-me-events-200-schema) |
| [401](#list-me-events-401) | Unauthorized | Unauthorized |  | [schema](#list-me-events-401-schema) |
| [403](#list-me-events-403) | Forbidden | Forbidden |  | [schema](#list-me-events-403-schema) |
| [500](#list-me-events-500) | Internal Server Error | Internal Server Error |  | [schema](#list-me-events-500-schema) |

#### Responses


##### <span id="list-me-events-200"></span> 200 - OK
Status: OK

###### <span id="list-me-events-200-schema"></span> Schema
   
  

[][Event](#event)

##### <span id="list-me-events-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="list-me-events-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="list-me-events-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="list-me-events-403-schema"></span> Schema
   
  

[Error](#error)

##### <span id="list-me-events-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="list-me-events-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="list-me-roles"></span> List authenticated user roles (*listMeRoles*)

```
GET /v2/users/me/roles
```

List the authenticated user's roles.

#### Security Requirements
  * OAuth2: public, user

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-me-roles-200) | OK | OK |  | [schema](#list-me-roles-200-schema) |
| [401](#list-me-roles-401) | Unauthorized | Unauthorized |  | [schema](#list-me-roles-401-schema) |
| [403](#list-me-roles-403) | Forbidden | Forbidden |  | [schema](#list-me-roles-403-schema) |
| [500](#list-me-roles-500) | Internal Server Error | Internal Server Error |  | [schema](#list-me-roles-500-schema) |

#### Responses


##### <span id="list-me-roles-200"></span> 200 - OK
Status: OK

###### <span id="list-me-roles-200-schema"></span> Schema
   
  

[][Role](#role)

##### <span id="list-me-roles-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="list-me-roles-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="list-me-roles-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="list-me-roles-403-schema"></span> Schema
   
  

[Error](#error)

##### <span id="list-me-roles-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="list-me-roles-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="list-type"></span> List event types (*listType*)

```
GET /v2/events/types
```

List all event categories.

#### Security Requirements
  * OAuth2: event, public

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-type-200) | OK | OK |  | [schema](#list-type-200-schema) |
| [401](#list-type-401) | Unauthorized | Unauthorized |  | [schema](#list-type-401-schema) |
| [403](#list-type-403) | Forbidden | Forbidden |  | [schema](#list-type-403-schema) |
| [500](#list-type-500) | Internal Server Error | Internal Server Error |  | [schema](#list-type-500-schema) |

#### Responses


##### <span id="list-type-200"></span> 200 - OK
Status: OK

###### <span id="list-type-200-schema"></span> Schema
   
  

[][EventType](#event-type)

##### <span id="list-type-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="list-type-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="list-type-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="list-type-403-schema"></span> Schema
   
  

[Error](#error)

##### <span id="list-type-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="list-type-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="list-user"></span> List users (*listUser*)

```
GET /v2/users
```

List all users.

#### Security Requirements
  * OAuth2: public

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tutor | `query` | boolean | `bool` |  |  |  | Tutor filter. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-user-200) | OK | OK |  | [schema](#list-user-200-schema) |
| [401](#list-user-401) | Unauthorized | Unauthorized |  | [schema](#list-user-401-schema) |
| [500](#list-user-500) | Internal Server Error | Internal Server Error |  | [schema](#list-user-500-schema) |

#### Responses


##### <span id="list-user-200"></span> 200 - OK
Status: OK

###### <span id="list-user-200-schema"></span> Schema
   
  

[][User](#user)

##### <span id="list-user-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="list-user-401-schema"></span> Schema
   
  

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
  * OAuth2: event, public, user

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | User ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-user-events-200) | OK | OK |  | [schema](#list-user-events-200-schema) |
| [400](#list-user-events-400) | Bad Request | Bad request |  | [schema](#list-user-events-400-schema) |
| [401](#list-user-events-401) | Unauthorized | Unauthorized |  | [schema](#list-user-events-401-schema) |
| [403](#list-user-events-403) | Forbidden | Forbidden |  | [schema](#list-user-events-403-schema) |
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

##### <span id="list-user-events-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="list-user-events-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="list-user-events-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="list-user-events-403-schema"></span> Schema
   
  

[Error](#error)

##### <span id="list-user-events-404"></span> 404 - Not Found
Status: Not Found

###### <span id="list-user-events-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="list-user-events-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="list-user-events-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="list-user-roles"></span> List user roles (*listUserRoles*)

```
GET /v2/users/{id}/roles
```

List the user's roles.

#### Security Requirements
  * OAuth2: public, user

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | User ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-user-roles-200) | OK | OK |  | [schema](#list-user-roles-200-schema) |
| [400](#list-user-roles-400) | Bad Request | Bad request |  | [schema](#list-user-roles-400-schema) |
| [401](#list-user-roles-401) | Unauthorized | Unauthorized |  | [schema](#list-user-roles-401-schema) |
| [403](#list-user-roles-403) | Forbidden | Forbidden |  | [schema](#list-user-roles-403-schema) |
| [404](#list-user-roles-404) | Not Found | Not Found |  | [schema](#list-user-roles-404-schema) |
| [500](#list-user-roles-500) | Internal Server Error | Internal Server Error |  | [schema](#list-user-roles-500-schema) |

#### Responses


##### <span id="list-user-roles-200"></span> 200 - OK
Status: OK

###### <span id="list-user-roles-200-schema"></span> Schema
   
  

[][Role](#role)

##### <span id="list-user-roles-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="list-user-roles-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="list-user-roles-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="list-user-roles-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="list-user-roles-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="list-user-roles-403-schema"></span> Schema
   
  

[Error](#error)

##### <span id="list-user-roles-404"></span> 404 - Not Found
Status: Not Found

###### <span id="list-user-roles-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="list-user-roles-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="list-user-roles-500-schema"></span> Schema
   
  

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
  * OAuth2: event, public

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | Event ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#read-event-200) | OK | OK |  | [schema](#read-event-200-schema) |
| [400](#read-event-400) | Bad Request | Bad request |  | [schema](#read-event-400-schema) |
| [401](#read-event-401) | Unauthorized | Unauthorized |  | [schema](#read-event-401-schema) |
| [403](#read-event-403) | Forbidden | Forbidden |  | [schema](#read-event-403-schema) |
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

##### <span id="read-event-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="read-event-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="read-event-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="read-event-403-schema"></span> Schema
   
  

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
  * OAuth2: public

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#read-me-200) | OK | OK |  | [schema](#read-me-200-schema) |
| [401](#read-me-401) | Unauthorized | Unauthorized |  | [schema](#read-me-401-schema) |
| [500](#read-me-500) | Internal Server Error | Internal Server Error |  | [schema](#read-me-500-schema) |

#### Responses


##### <span id="read-me-200"></span> 200 - OK
Status: OK

###### <span id="read-me-200-schema"></span> Schema
   
  

[User](#user)

##### <span id="read-me-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="read-me-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="read-me-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="read-me-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="read-type"></span> Read event type (*readType*)

```
GET /v2/events/types/{id}
```

Read an event category.

#### Security Requirements
  * OAuth2: event, public

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | ID of event category. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#read-type-200) | OK | OK |  | [schema](#read-type-200-schema) |
| [400](#read-type-400) | Bad Request | Bad request |  | [schema](#read-type-400-schema) |
| [401](#read-type-401) | Unauthorized | Unauthorized |  | [schema](#read-type-401-schema) |
| [403](#read-type-403) | Forbidden | Forbidden |  | [schema](#read-type-403-schema) |
| [404](#read-type-404) | Not Found | Not Found |  | [schema](#read-type-404-schema) |
| [500](#read-type-500) | Internal Server Error | Internal Server Error |  | [schema](#read-type-500-schema) |

#### Responses


##### <span id="read-type-200"></span> 200 - OK
Status: OK

###### <span id="read-type-200-schema"></span> Schema
   
  

[EventType](#event-type)

##### <span id="read-type-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="read-type-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="read-type-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="read-type-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="read-type-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="read-type-403-schema"></span> Schema
   
  

[Error](#error)

##### <span id="read-type-404"></span> 404 - Not Found
Status: Not Found

###### <span id="read-type-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="read-type-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="read-type-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="read-user"></span> Read user (*readUser*)

```
GET /v2/users/{id}
```

Read an user by ID.

#### Security Requirements
  * OAuth2: public, user

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | User ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#read-user-200) | OK | OK |  | [schema](#read-user-200-schema) |
| [400](#read-user-400) | Bad Request | Bad request |  | [schema](#read-user-400-schema) |
| [401](#read-user-401) | Unauthorized | Unauthorized |  | [schema](#read-user-401-schema) |
| [403](#read-user-403) | Forbidden | Forbidden |  | [schema](#read-user-403-schema) |
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

##### <span id="read-user-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="read-user-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="read-user-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="read-user-403-schema"></span> Schema
   
  

[Error](#error)

##### <span id="read-user-404"></span> 404 - Not Found
Status: Not Found

###### <span id="read-user-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="read-user-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="read-user-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="remove-admin"></span> Remove admin (*removeAdmin*)

```
DELETE /v2/users/{id}/role/admin
```

Remove admin role to a user.

#### Security Requirements
  * OAuth2: public, user, user:write

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | User ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#remove-admin-204) | No Content | No Content |  | [schema](#remove-admin-204-schema) |
| [400](#remove-admin-400) | Bad Request | Bad request |  | [schema](#remove-admin-400-schema) |
| [401](#remove-admin-401) | Unauthorized | Unauthorized |  | [schema](#remove-admin-401-schema) |
| [403](#remove-admin-403) | Forbidden | Forbidden |  | [schema](#remove-admin-403-schema) |
| [404](#remove-admin-404) | Not Found | Not Found |  | [schema](#remove-admin-404-schema) |
| [500](#remove-admin-500) | Internal Server Error | Internal Server Error |  | [schema](#remove-admin-500-schema) |

#### Responses


##### <span id="remove-admin-204"></span> 204 - No Content
Status: No Content

###### <span id="remove-admin-204-schema"></span> Schema

##### <span id="remove-admin-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="remove-admin-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="remove-admin-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="remove-admin-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="remove-admin-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="remove-admin-403-schema"></span> Schema
   
  

[Error](#error)

##### <span id="remove-admin-404"></span> 404 - Not Found
Status: Not Found

###### <span id="remove-admin-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="remove-admin-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="remove-admin-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="remove-calendar"></span> Remove calendar (*removeCalendar*)

```
DELETE /v2/users/{id}/role/calendar
```

Remove calendar role to a user.

#### Security Requirements
  * OAuth2: public, user, user:write

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | User ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#remove-calendar-204) | No Content | No Content |  | [schema](#remove-calendar-204-schema) |
| [400](#remove-calendar-400) | Bad Request | Bad request |  | [schema](#remove-calendar-400-schema) |
| [401](#remove-calendar-401) | Unauthorized | Unauthorized |  | [schema](#remove-calendar-401-schema) |
| [403](#remove-calendar-403) | Forbidden | Forbidden |  | [schema](#remove-calendar-403-schema) |
| [404](#remove-calendar-404) | Not Found | Not Found |  | [schema](#remove-calendar-404-schema) |
| [500](#remove-calendar-500) | Internal Server Error | Internal Server Error |  | [schema](#remove-calendar-500-schema) |

#### Responses


##### <span id="remove-calendar-204"></span> 204 - No Content
Status: No Content

###### <span id="remove-calendar-204-schema"></span> Schema

##### <span id="remove-calendar-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="remove-calendar-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="remove-calendar-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="remove-calendar-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="remove-calendar-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="remove-calendar-403-schema"></span> Schema
   
  

[Error](#error)

##### <span id="remove-calendar-404"></span> 404 - Not Found
Status: Not Found

###### <span id="remove-calendar-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="remove-calendar-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="remove-calendar-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="remove-tutor"></span> Remove tutor (*removeTutor*)

```
DELETE /v2/users/{id}/role/tutor
```

Remove tutor role to a user.

#### Security Requirements
  * OAuth2: public, user, user:write

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | User ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#remove-tutor-204) | No Content | No Content |  | [schema](#remove-tutor-204-schema) |
| [400](#remove-tutor-400) | Bad Request | Bad request |  | [schema](#remove-tutor-400-schema) |
| [401](#remove-tutor-401) | Unauthorized | Unauthorized |  | [schema](#remove-tutor-401-schema) |
| [403](#remove-tutor-403) | Forbidden | Forbidden |  | [schema](#remove-tutor-403-schema) |
| [404](#remove-tutor-404) | Not Found | Not Found |  | [schema](#remove-tutor-404-schema) |
| [500](#remove-tutor-500) | Internal Server Error | Internal Server Error |  | [schema](#remove-tutor-500-schema) |

#### Responses


##### <span id="remove-tutor-204"></span> 204 - No Content
Status: No Content

###### <span id="remove-tutor-204-schema"></span> Schema

##### <span id="remove-tutor-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="remove-tutor-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="remove-tutor-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="remove-tutor-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="remove-tutor-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="remove-tutor-403-schema"></span> Schema
   
  

[Error](#error)

##### <span id="remove-tutor-404"></span> 404 - Not Found
Status: Not Found

###### <span id="remove-tutor-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="remove-tutor-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="remove-tutor-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="subscribe-me"></span> Subscribe authenticated user (*subscribeMe*)

```
POST /v2/users/me/events/{id}
```

Subscribe an authenticated user to an event.

#### Security Requirements
  * OAuth2: event, public, user, user:subscription

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | Event ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#subscribe-me-201) | Created | Created |  | [schema](#subscribe-me-201-schema) |
| [400](#subscribe-me-400) | Bad Request | Bad request |  | [schema](#subscribe-me-400-schema) |
| [401](#subscribe-me-401) | Unauthorized | Unauthorized |  | [schema](#subscribe-me-401-schema) |
| [403](#subscribe-me-403) | Forbidden | Forbidden |  | [schema](#subscribe-me-403-schema) |
| [404](#subscribe-me-404) | Not Found | Not Found |  | [schema](#subscribe-me-404-schema) |
| [500](#subscribe-me-500) | Internal Server Error | Internal Server Error |  | [schema](#subscribe-me-500-schema) |

#### Responses


##### <span id="subscribe-me-201"></span> 201 - Created
Status: Created

###### <span id="subscribe-me-201-schema"></span> Schema
   
  

[][Event](#event)

##### <span id="subscribe-me-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="subscribe-me-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="subscribe-me-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="subscribe-me-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="subscribe-me-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="subscribe-me-403-schema"></span> Schema
   
  

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
  * OAuth2: event, public, user, user:subscription, user:write

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| eventId | `path` | string | `string` |  | ✓ |  | Event ID. |
| userId | `path` | string | `string` |  | ✓ |  | User ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#subscribe-user-201) | Created | Created |  | [schema](#subscribe-user-201-schema) |
| [400](#subscribe-user-400) | Bad Request | Bad request |  | [schema](#subscribe-user-400-schema) |
| [401](#subscribe-user-401) | Unauthorized | Unauthorized |  | [schema](#subscribe-user-401-schema) |
| [403](#subscribe-user-403) | Forbidden | Forbidden |  | [schema](#subscribe-user-403-schema) |
| [404](#subscribe-user-404) | Not Found | Not Found |  | [schema](#subscribe-user-404-schema) |
| [500](#subscribe-user-500) | Internal Server Error | Internal Server Error |  | [schema](#subscribe-user-500-schema) |

#### Responses


##### <span id="subscribe-user-201"></span> 201 - Created
Status: Created

###### <span id="subscribe-user-201-schema"></span> Schema
   
  

[][Event](#event)

##### <span id="subscribe-user-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="subscribe-user-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="subscribe-user-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="subscribe-user-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="subscribe-user-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="subscribe-user-403-schema"></span> Schema
   
  

[Error](#error)

##### <span id="subscribe-user-404"></span> 404 - Not Found
Status: Not Found

###### <span id="subscribe-user-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="subscribe-user-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="subscribe-user-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="token-info"></span> Send token information (*tokenInfo*)

```
GET /v2/auth/token/info
```

Send token information or unauthorized error response.

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Authorization | `header` | string | `string` |  | ✓ |  | Bearer access token |

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
   
  

[TokenInfoOKBody](#token-info-o-k-body)

##### <span id="token-info-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="token-info-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="token-info-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="token-info-500-schema"></span> Schema
   
  

[Error](#error)

###### Inlined models

**<span id="token-info-o-k-body"></span> TokenInfoOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| expiresAt | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| login | string| `string` |  | | username |  |



### <span id="token-refresh"></span> Refresh token (*tokenRefresh*)

```
GET /v2/auth/token/refresh
```

Refresh expired token.

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Authorization | `header` | string | `string` |  | ✓ |  | Bearer refresh token |

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
   
  

[Token](#token)

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
  * OAuth2: event, public, user, user:subscription

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | Event ID. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#unsubscribe-me-204) | No Content | No Content |  | [schema](#unsubscribe-me-204-schema) |
| [400](#unsubscribe-me-400) | Bad Request | Bad request |  | [schema](#unsubscribe-me-400-schema) |
| [401](#unsubscribe-me-401) | Unauthorized | Unauthorized |  | [schema](#unsubscribe-me-401-schema) |
| [403](#unsubscribe-me-403) | Forbidden | Forbidden |  | [schema](#unsubscribe-me-403-schema) |
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

##### <span id="unsubscribe-me-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="unsubscribe-me-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="unsubscribe-me-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="unsubscribe-me-403-schema"></span> Schema
   
  

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
  * OAuth2: event, public, user, user:subscription, user:write

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
| [401](#unsubscribe-user-401) | Unauthorized | Unauthorized |  | [schema](#unsubscribe-user-401-schema) |
| [403](#unsubscribe-user-403) | Forbidden | Forbidden |  | [schema](#unsubscribe-user-403-schema) |
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

##### <span id="unsubscribe-user-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="unsubscribe-user-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="unsubscribe-user-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="unsubscribe-user-403-schema"></span> Schema
   
  

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
  * OAuth2: event, event:write, public

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
| [401](#update-event-401) | Unauthorized | Unauthorized |  | [schema](#update-event-401-schema) |
| [403](#update-event-403) | Forbidden | Forbidden |  | [schema](#update-event-403-schema) |
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

##### <span id="update-event-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="update-event-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="update-event-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="update-event-403-schema"></span> Schema
   
  

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
  * OAuth2: public

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| user | `body` | [User](#user) | `ent.User` | | ✓ | | User to update. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-me-200) | OK | OK |  | [schema](#update-me-200-schema) |
| [400](#update-me-400) | Bad Request | Bad request |  | [schema](#update-me-400-schema) |
| [401](#update-me-401) | Unauthorized | Unauthorized |  | [schema](#update-me-401-schema) |
| [500](#update-me-500) | Internal Server Error | Internal Server Error |  | [schema](#update-me-500-schema) |

#### Responses


##### <span id="update-me-200"></span> 200 - OK
Status: OK

###### <span id="update-me-200-schema"></span> Schema
   
  

[User](#user)

##### <span id="update-me-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="update-me-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="update-me-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="update-me-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="update-me-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="update-me-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="update-type"></span> Update event type (*updateType*)

```
POST /v2/events/types/{id}
```

Update an event category.

#### Security Requirements
  * OAuth2: event, event:write, public

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | ID of event category. |
| type | `body` | [EventType](#event-type) | `ent.EventType` | | ✓ | | Event category to update. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-type-200) | OK | OK |  | [schema](#update-type-200-schema) |
| [400](#update-type-400) | Bad Request | Bad request |  | [schema](#update-type-400-schema) |
| [401](#update-type-401) | Unauthorized | Unauthorized |  | [schema](#update-type-401-schema) |
| [403](#update-type-403) | Forbidden | Forbidden |  | [schema](#update-type-403-schema) |
| [404](#update-type-404) | Not Found | Not Found |  | [schema](#update-type-404-schema) |
| [500](#update-type-500) | Internal Server Error | Internal Server Error |  | [schema](#update-type-500-schema) |

#### Responses


##### <span id="update-type-200"></span> 200 - OK
Status: OK

###### <span id="update-type-200-schema"></span> Schema
   
  

[EventType](#event-type)

##### <span id="update-type-400"></span> 400 - Bad request
Status: Bad Request

###### <span id="update-type-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="update-type-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="update-type-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="update-type-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="update-type-403-schema"></span> Schema
   
  

[Error](#error)

##### <span id="update-type-404"></span> 404 - Not Found
Status: Not Found

###### <span id="update-type-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="update-type-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="update-type-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="update-user"></span> Update user (*updateUser*)

```
PUT /v2/users/{id}
```

Update an user by ID.

#### Security Requirements
  * OAuth2: public, user, user:write

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
| [401](#update-user-401) | Unauthorized | Unauthorized |  | [schema](#update-user-401-schema) |
| [403](#update-user-403) | Forbidden | Forbidden |  | [schema](#update-user-403-schema) |
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

##### <span id="update-user-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="update-user-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="update-user-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="update-user-403-schema"></span> Schema
   
  

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



### <span id="token"></span> Token


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| accessToken | string| `string` |  | |  |  |
| refreshToken | string| `string` |  | |  |  |



### <span id="principal"></span> principal


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| principal | string| string | |  |  |


