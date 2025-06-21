package swagger

import (
	"bytes"
	"text/template"
)

type Meta struct {
	Title   string
	Version string
	Host    string
}

func GetSwagger(meta Meta) (string, error) {
	tmpl, err := template.New("swagger").Parse(templateContent)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, meta); err != nil {
		return "", err
	}

	return buf.String(), nil
}

var templateContent = `swagger: "2.0"
info:
  title: '{{.Title}}'
  version: '{{.Version}}'
tags:
  - name: AdminRolesService
  - name: AdminUsersService
  - name: AuthenticationService
  - name: PermissionsService
  - name: SomeService
host: '{{.Host}}'
schemes:
  - https
consumes:
  - application/json
produces:
  - application/json
paths:
  /v1/admin/authz/roles:
    get:
      summary: GetAllRoles возвращает список всех ролей по всем доменам.
      operationId: AdminRolesService_GetAllRoles
      responses:
        "200":
          description: A successful response.
          schema:
            $ref: '#/definitions/v1GetAllRolesResponse'
        default:
          description: An unexpected error response.
          schema:
            $ref: '#/definitions/rpcStatus'
      tags:
        - AdminRolesService
    post:
      summary: CreateRole создает новую роль в указанном домене.
      operationId: AdminRolesService_CreateRole
      responses:
        "200":
          description: A successful response.
          schema:
            $ref: '#/definitions/v1CreateRoleResponse'
        default:
          description: An unexpected error response.
          schema:
            $ref: '#/definitions/rpcStatus'
      parameters:
        - name: body
          description: CreateRoleRequest запрос на создание новой роли.
          in: body
          required: true
          schema:
            $ref: '#/definitions/v1CreateRoleRequest'
      tags:
        - AdminRolesService
    put:
      summary: UpdateRole обновляет существующую роль.
      operationId: AdminRolesService_UpdateRole
      responses:
        "200":
          description: A successful response.
          schema:
            $ref: '#/definitions/v1UpdateRoleResponse'
        default:
          description: An unexpected error response.
          schema:
            $ref: '#/definitions/rpcStatus'
      parameters:
        - name: body
          description: UpdateRoleRequest запрос на обновление роли.
          in: body
          required: true
          schema:
            $ref: '#/definitions/v1UpdateRoleRequest'
      tags:
        - AdminRolesService
  /v1/admin/authz/roles/{roleId}:
    delete:
      summary: DeleteRole удаляет роль по её идентификатору.
      operationId: AdminRolesService_DeleteRole
      responses:
        "200":
          description: A successful response.
          schema:
            $ref: '#/definitions/v1DeleteRoleResponse'
        default:
          description: An unexpected error response.
          schema:
            $ref: '#/definitions/rpcStatus'
      parameters:
        - name: roleId
          description: RoleId идентификатор роли для удаления.
          in: path
          required: true
          type: string
      tags:
        - AdminRolesService
  /v1/admin/users:
    get:
      summary: GetUsers возвращает список всех пользователей.
      operationId: AdminUsersService_GetUsers
      responses:
        "200":
          description: A successful response.
          schema:
            $ref: '#/definitions/v1GetUsersResponse'
        default:
          description: An unexpected error response.
          schema:
            $ref: '#/definitions/rpcStatus'
      tags:
        - AdminUsersService
    post:
      summary: CreateUser создает нового пользователя.
      operationId: AdminUsersService_CreateUser
      responses:
        "200":
          description: A successful response.
          schema:
            $ref: '#/definitions/v1CreateUserResponse'
        default:
          description: An unexpected error response.
          schema:
            $ref: '#/definitions/rpcStatus'
      parameters:
        - name: body
          description: CreateUserRequest запрос на создание пользователя.
          in: body
          required: true
          schema:
            $ref: '#/definitions/v1CreateUserRequest'
      tags:
        - AdminUsersService
    put:
      summary: UpdateUser обновляет информацию о пользователе.
      operationId: AdminUsersService_UpdateUser
      responses:
        "200":
          description: A successful response.
          schema:
            $ref: '#/definitions/v1UpdateUserResponse'
        default:
          description: An unexpected error response.
          schema:
            $ref: '#/definitions/rpcStatus'
      parameters:
        - name: body
          description: UpdateUserRequest запрос на обновление пользователя.
          in: body
          required: true
          schema:
            $ref: '#/definitions/v1UpdateUserRequest'
      tags:
        - AdminUsersService
  /v1/admin/users/{userId}:
    delete:
      summary: DeleteUser удаляет пользователя.
      operationId: AdminUsersService_DeleteUser
      responses:
        "200":
          description: A successful response.
          schema:
            $ref: '#/definitions/v1DeleteUserResponse'
        default:
          description: An unexpected error response.
          schema:
            $ref: '#/definitions/rpcStatus'
      parameters:
        - name: userId
          description: UserId идентификатор пользователя.
          in: path
          required: true
          type: string
      tags:
        - AdminUsersService
  /v1/admin/users/{userId}/domains:
    post:
      summary: AssignUserToDomain назначает пользователю роль в домене.
      operationId: AdminUsersService_AssignUserToDomain
      responses:
        "200":
          description: A successful response.
          schema:
            $ref: '#/definitions/v1AssignUserToDomainResponse'
        default:
          description: An unexpected error response.
          schema:
            $ref: '#/definitions/rpcStatus'
      parameters:
        - name: userId
          description: UserId идентификатор пользователя.
          in: path
          required: true
          type: string
        - name: body
          in: body
          required: true
          schema:
            $ref: '#/definitions/AdminUsersServiceAssignUserToDomainBody'
      tags:
        - AdminUsersService
  /v1/admin/users/{userId}/domains/{domainId}:
    delete:
      summary: RemoveUserFromDomain удаляет пользователя из домена.
      operationId: AdminUsersService_RemoveUserFromDomain
      responses:
        "200":
          description: A successful response.
          schema:
            $ref: '#/definitions/v1RemoveUserFromDomainResponse'
        default:
          description: An unexpected error response.
          schema:
            $ref: '#/definitions/rpcStatus'
      parameters:
        - name: userId
          description: UserId идентификатор пользователя.
          in: path
          required: true
          type: string
        - name: domainId
          description: DomainId идентификатор домена.
          in: path
          required: true
          type: string
      tags:
        - AdminUsersService
  /v1/admin/users/{userId}/domains/{domainId}/role:
    put:
      summary: UpdateUserDomainRole обновляет роль пользователя в домене.
      operationId: AdminUsersService_UpdateUserDomainRole
      responses:
        "200":
          description: A successful response.
          schema:
            $ref: '#/definitions/v1UpdateUserDomainRoleResponse'
        default:
          description: An unexpected error response.
          schema:
            $ref: '#/definitions/rpcStatus'
      parameters:
        - name: userId
          description: UserId идентификатор пользователя.
          in: path
          required: true
          type: string
        - name: domainId
          description: DomainId идентификатор домена.
          in: path
          required: true
          type: string
        - name: body
          in: body
          required: true
          schema:
            $ref: '#/definitions/AdminUsersServiceUpdateUserDomainRoleBody'
      tags:
        - AdminUsersService
  /v1/auth/authorize:
    post:
      summary: Authorize проверяет валидность токена и возвращает информацию о пользователе.
      operationId: AuthenticationService_Authorize
      responses:
        "200":
          description: A successful response.
          schema:
            $ref: '#/definitions/v1AuthorizeResponse'
        default:
          description: An unexpected error response.
          schema:
            $ref: '#/definitions/rpcStatus'
      parameters:
        - name: body
          description: AuthorizeRequest запрос на проверку токена.
          in: body
          required: true
          schema:
            $ref: '#/definitions/v1AuthorizeRequest'
      tags:
        - AuthenticationService
  /v1/auth/refresh:
    post:
      summary: RefreshToken обновляет истекающий access токен.
      operationId: AuthenticationService_RefreshToken
      responses:
        "200":
          description: A successful response.
          schema:
            $ref: '#/definitions/v1RefreshTokenResponse'
        default:
          description: An unexpected error response.
          schema:
            $ref: '#/definitions/rpcStatus'
      parameters:
        - name: body
          description: RefreshTokenRequest запрос на обновление токена.
          in: body
          required: true
          schema:
            $ref: '#/definitions/v1RefreshTokenRequest'
      tags:
        - AuthenticationService
  /v1/auth/signin:
    post:
      summary: SignIn выполняет аутентификацию существующего пользователя.
      operationId: AuthenticationService_SignIn
      responses:
        "200":
          description: A successful response.
          schema:
            $ref: '#/definitions/v1SignInResponse'
        default:
          description: An unexpected error response.
          schema:
            $ref: '#/definitions/rpcStatus'
      parameters:
        - name: body
          description: SignInRequest запрос на вход в систему.
          in: body
          required: true
          schema:
            $ref: '#/definitions/v1SignInRequest'
      tags:
        - AuthenticationService
  /v1/auth/signup:
    post:
      summary: SignUp выполняет регистрацию нового пользователя в системе.
      operationId: AuthenticationService_SignUp
      responses:
        "200":
          description: A successful response.
          schema:
            $ref: '#/definitions/v1SignUpResponse'
        default:
          description: An unexpected error response.
          schema:
            $ref: '#/definitions/rpcStatus'
      parameters:
        - name: body
          description: SignUpRequest запрос на регистрацию нового пользователя.
          in: body
          required: true
          schema:
            $ref: '#/definitions/v1SignUpRequest'
      tags:
        - AuthenticationService
  /v1/authz/permissions:
    get:
      operationId: PermissionsService_GetAllPermissions
      responses:
        "200":
          description: A successful response.
          schema:
            $ref: '#/definitions/v1GetAllPermissionsResponse'
        default:
          description: An unexpected error response.
          schema:
            $ref: '#/definitions/rpcStatus'
      tags:
        - PermissionsService
  /v1/some-service/example:
    get:
      summary: SomeExampleMethod method for example
      operationId: SomeService_SomeExampleMethod
      responses:
        "200":
          description: A successful response.
          schema:
            $ref: '#/definitions/v1SomeExampleMethodResponse'
        default:
          description: An unexpected error response.
          schema:
            $ref: '#/definitions/rpcStatus'
      parameters:
        - name: someText
          description: SomeText field to provide some text
          in: query
          required: false
          type: string
        - name: someUnsignedIntegerValue
          description: SomeUnsignedIntegerValue field to provide some integer with unsigned value
          in: query
          required: false
          type: integer
          format: int64
      tags:
        - SomeService
  /v1/some-service/example/another:
    get:
      summary: SomeExampleMethod method for example
      operationId: SomeService_AnotherExampleMethod
      responses:
        "200":
          description: A successful response.
          schema:
            $ref: '#/definitions/v1SomeExampleMethodResponse'
        default:
          description: An unexpected error response.
          schema:
            $ref: '#/definitions/rpcStatus'
      parameters:
        - name: someText
          description: SomeText field to provide some text
          in: query
          required: false
          type: string
        - name: someUnsignedIntegerValue
          description: SomeUnsignedIntegerValue field to provide some integer with unsigned value
          in: query
          required: false
          type: integer
          format: int64
      tags:
        - SomeService
definitions:
  AdminUsersServiceAssignUserToDomainBody:
    type: object
    properties:
      domainId:
        type: string
        description: DomainId идентификатор домена.
      roleId:
        type: string
        description: RoleId идентификатор роли.
    description: AssignUserToDomainRequest запрос на назначение пользователю роли в домене.
  AdminUsersServiceUpdateUserDomainRoleBody:
    type: object
    properties:
      roleId:
        type: string
        description: RoleId идентификатор новой роли.
    description: UpdateUserDomainRoleRequest запрос на обновление роли пользователя в домене.
  commonDomain:
    type: object
    properties:
      id:
        type: string
      name:
        type: string
  permissionsPermissionMeta:
    type: object
    properties:
      alias:
        type: string
      description:
        type: string
  protobufAny:
    type: object
    properties:
      '@type':
        type: string
    additionalProperties: {}
  rpcStatus:
    type: object
    properties:
      code:
        type: integer
        format: int32
      message:
        type: string
      details:
        type: array
        items:
          type: object
          $ref: '#/definitions/protobufAny'
  v1AssignUserToDomainResponse:
    type: object
    properties:
      userDomains:
        $ref: '#/definitions/v1UserDomains'
        description: UserDomains обновленная информация о пользователе.
    description: AssignUserToDomainResponse ответ на назначение пользователю роли в домене.
  v1AuthorizeRequest:
    type: object
    properties:
      accessToken:
        type: string
        description: AccessToken JWT токен для проверки.
    description: AuthorizeRequest запрос на проверку токена.
  v1AuthorizeResponse:
    type: object
    properties:
      userId:
        type: string
        description: UserId уникальный идентификатор пользователя.
      email:
        type: string
        description: Email адрес электронной почты пользователя.
      currentDomainId:
        type: string
        description: CurrentDomainId идентификатор текущего домена пользователя.
      currentRoleId:
        type: string
        title: CurrentRoleId
      permissions:
        type: array
        items:
          type: string
        title: Permissions
    description: AuthorizeResponse ответ с данными пользователя из токена.
  v1CreateRoleRequest:
    type: object
    properties:
      role:
        $ref: '#/definitions/v1Role'
        description: Role роль для создания.
    description: CreateRoleRequest запрос на создание новой роли.
  v1CreateRoleResponse:
    type: object
    properties:
      domainRoles:
        $ref: '#/definitions/v1DomainRoles'
        description: DomainRoles обновленный список ролей домена.
    description: CreateRoleResponse ответ на создание роли.
  v1CreateUserRequest:
    type: object
    properties:
      user:
        $ref: '#/definitions/v1User'
        description: User данные нового пользователя.
    description: CreateUserRequest запрос на создание пользователя.
  v1CreateUserResponse:
    type: object
    properties:
      userDomains:
        $ref: '#/definitions/v1UserDomains'
        description: UserDomains информация о созданном пользователе.
    description: CreateUserResponse ответ на создание пользователя.
  v1DeleteRoleResponse:
    type: object
    properties:
      domainRoles:
        $ref: '#/definitions/v1DomainRoles'
        description: DomainRoles обновленный список ролей домена.
    description: DeleteRoleResponse ответ на удаление роли.
  v1DeleteUserResponse:
    type: object
    description: DeleteUserResponse ответ на удаление пользователя.
  v1DomainRole:
    type: object
    properties:
      domain:
        $ref: '#/definitions/commonDomain'
        description: Domain информация о домене.
      role:
        $ref: '#/definitions/v1Role'
        description: Role роль пользователя в домене.
    description: DomainRole связывает домен и роль пользователя.
  v1DomainRoles:
    type: object
    properties:
      domain:
        $ref: '#/definitions/commonDomain'
        description: Domain информация о домене.
      roles:
        type: array
        items:
          type: object
          $ref: '#/definitions/v1Role'
        description: Roles список ролей в домене.
    description: DomainRoles группирует роли по домену.
  v1GetAllPermissionsResponse:
    type: object
    properties:
      permissions:
        type: array
        items:
          type: object
          $ref: '#/definitions/permissionsPermissionMeta'
  v1GetAllRolesResponse:
    type: object
    properties:
      domainRoles:
        type: array
        items:
          type: object
          $ref: '#/definitions/v1DomainRoles'
        description: DomainRoles список ролей сгруппированных по доменам.
    description: GetAllRolesResponse ответ со списком всех ролей по доменам.
  v1GetUsersResponse:
    type: object
    properties:
      users:
        type: array
        items:
          type: object
          $ref: '#/definitions/v1UserDomains'
        description: Users список пользователей с их ролями в доменах.
    description: GetUsersResponse ответ со списком пользователей.
  v1RefreshTokenRequest:
    type: object
    properties:
      refreshToken:
        type: string
        description: RefreshToken токен для обновления access токена.
    description: RefreshTokenRequest запрос на обновление токена.
  v1RefreshTokenResponse:
    type: object
    properties:
      accessToken:
        type: string
        description: AccessToken новый JWT токен доступа.
      refreshToken:
        type: string
        description: RefreshToken новый токен для обновления.
    description: RefreshTokenResponse ответ с новыми токенами.
  v1RemoveUserFromDomainResponse:
    type: object
    properties:
      userDomains:
        $ref: '#/definitions/v1UserDomains'
        description: UserDomains обновленная информация о пользователе.
    description: RemoveUserFromDomainResponse ответ на удаление пользователя из домена.
  v1Role:
    type: object
    properties:
      id:
        type: string
        title: Id идентификатор роли
      name:
        type: string
        description: Name название роли.
      description:
        type: string
        description: Description описание роли.
      permissions:
        type: array
        items:
          type: string
        description: Permissions список разрешений роли.
      domainId:
        type: string
        description: DomainId идентификатор домена, которому принадлежит роль.
    description: Role описывает роль в системе.
  v1SignInRequest:
    type: object
    properties:
      email:
        type: string
        description: Email адрес электронной почты пользователя.
      password:
        type: string
        description: Password пароль пользователя.
    description: SignInRequest запрос на вход в систему.
  v1SignInResponse:
    type: object
    properties:
      accessToken:
        type: string
        description: AccessToken JWT токен доступа.
      refreshToken:
        type: string
        description: RefreshToken токен для обновления access токена.
    description: SignInResponse ответ на успешную аутентификацию.
  v1SignUpRequest:
    type: object
    properties:
      email:
        type: string
        description: Email адрес электронной почты пользователя.
      password:
        type: string
        description: Password пароль пользователя.
      name:
        type: string
        description: Name имя пользователя.
    description: SignUpRequest запрос на регистрацию нового пользователя.
  v1SignUpResponse:
    type: object
    properties:
      accessToken:
        type: string
        description: AccessToken JWT токен доступа.
      refreshToken:
        type: string
        description: RefreshToken токен для обновления access токена.
    description: SignUpResponse ответ на успешную регистрацию.
  v1SomeExampleMethodResponse:
    type: object
    properties:
      someResponse:
        type: string
        title: SomeResponse field provides some text will returned from service
  v1UpdateRoleRequest:
    type: object
    properties:
      role:
        $ref: '#/definitions/v1Role'
        description: Role обновленная роль.
    description: UpdateRoleRequest запрос на обновление роли.
  v1UpdateRoleResponse:
    type: object
    properties:
      domainRoles:
        $ref: '#/definitions/v1DomainRoles'
        description: DomainRoles обновленный список ролей домена.
    description: UpdateRoleResponse ответ на обновление роли.
  v1UpdateUserDomainRoleResponse:
    type: object
    properties:
      userDomains:
        $ref: '#/definitions/v1UserDomains'
        description: UserDomains обновленная информация о пользователе.
    description: UpdateUserDomainRoleResponse ответ на обновление роли пользователя в домене.
  v1UpdateUserRequest:
    type: object
    properties:
      user:
        $ref: '#/definitions/v1User'
        description: User обновленные данные пользователя.
    description: UpdateUserRequest запрос на обновление пользователя.
  v1UpdateUserResponse:
    type: object
    properties:
      userDomains:
        $ref: '#/definitions/v1UserDomains'
        description: UserDomains обновленная информация о пользователе.
    description: UpdateUserResponse ответ на обновление пользователя.
  v1User:
    type: object
    properties:
      id:
        type: string
        description: Id уникальный идентификатор пользователя.
      name:
        type: string
        description: Name имя пользователя.
      email:
        type: string
        description: Email адрес электронной почты пользователя.
      passwordHash:
        type: string
        description: PasswordHash хэш пароля пользователя.
      currentDomainId:
        type: string
        description: CurrentDomainId идентификатор текущего домена пользователя.
    description: User описывает пользователя в системе.
  v1UserDomains:
    type: object
    properties:
      user:
        $ref: '#/definitions/v1User'
        description: User информация о пользователе.
      domainRoles:
        type: array
        items:
          type: object
          $ref: '#/definitions/v1DomainRole'
        description: DomainRoles список ролей пользователя по доменам.
    description: UserDomains группирует информацию о пользователе и его ролях в доменах.
securityDefinitions:
  Bearer:
    type: apiKey
    description: 'Используйте формат: Bearer <JWT token>'
    name: Authorization
    in: header
`
