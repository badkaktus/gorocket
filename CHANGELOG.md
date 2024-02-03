# CHANGELOG

## [v0.1.4] - 2024-02-03
- Updated the description of all methods according to the golang convention
- Comment few fields in GroupMessage struct (will fix it in next release)
- Fix InviteGroupResponse struct
- Fix RespInfo struct
- Fix RespDirectory struct
- Fix UsersInfoResponse struct
- Write tests for all methods
- Added logo 😄 

## [v0.1.3] - 2022-10-27
- Bugfix: count/offset/sort overwrite other params
- Add receive group messages (GroupMessage)

## [v0.1.2] - 2022-10-02
- Add pagination

## [v0.1.1] - 2022-09-20

- Add semver version `v0.1.1` (see semver.org)
- Add `NewWithOptions(url string, opts ...Option)` for custom client settings
- Rollback `go.mod` package name to `github.com/badkaktus/gorocket`
