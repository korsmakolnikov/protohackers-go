#Env
FROM archlinux:base
RUN pacman -Sy --noconfirm go

#App setup
COPY ./ /code
WORKDIR /code
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app

#Runtime settigns
EXPOSE 8080

#Run
CMD [ "/app" ]
