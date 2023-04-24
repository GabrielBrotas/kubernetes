FROM golang:1.15

COPY . .

# gerar executavel 
RUN go build -o server .

CMD [ "./server" ]