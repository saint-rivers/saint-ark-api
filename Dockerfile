##
## STAGE 1 - BUILD
##
FROM golang:1.19.2 AS build
# RUN useradd -u 1001 -m iamuser

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . ./

# the scratch wasn't able to find the binary in the deployment stage
# unless I specified CGO, GOOS and GOARCH 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/rivers /app/main.go


##
## STAGE 2 - DEPLOY
##
FROM scratch
WORKDIR /

# copies the binary from the build stage
COPY --from=build /app/rivers /

# this cppies the static html pages for the site
# it will not be compiled. The code will reference it via relative directory 
# COPY --from=build /app/static /static 
COPY --from=build /app/uploads /uploads

## copy the non root user
# COPY --from=build /etc/passwd /etc/passwd

# USER 1001
EXPOSE 8080

# runs the imported binary
ENTRYPOINT ["/rivers"]