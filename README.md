# CI Full Stack

## SonarCloud

## SonarQube
### Run SonarQube
```bash
docker run -d --name sonarqube -e SONAR_ES_BOOTSTRAP_CHECKS_DISABLE=true -p 9000:9000 sonarqube:latest
```

### Create Coverage Report
```bash
go test -coverprofile=coverage.out ./...
```

### Executing the Scanner
```bash
 sonar-scanner \                                                                                                 ─╯
  -Dsonar.projectKey=go-ci-full-cycle \
  -Dsonar.sources=. \
  -Dsonar.host.url=http://localhost:9000 \
  -Dsonar.token=sqp_b86bebd34aebfe4b7f2046258ef036c55d4dc20c
```

### Sonar Properties
```properties
sonar.projectKey=go-ci-full-cycle

sonar.sources=.
sonar.exclusion=**/*_test.go

sonar.test=.
sonar.test.inclusions=**/*_test.go
sonar.test.exclusions=**/main.go
sonar.host.url=http://localhost:9000
sonar.login=sqp_b86bebd34aebfe4b7f2046258ef036c55d4dc20c
sonar.go.coverage.reportPaths=coverage.out
```

