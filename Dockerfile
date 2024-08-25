FROM golang:1.22.6

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY component/*.go ./component/
COPY gamestate/*.go ./gamestate/
COPY input/*.go ./input/
COPY loop/*.go ./loop/
COPY pieces/*.go ./pieces/
COPY screen/*.go ./screen/
COPY server/*.go ./server/
COPY game.go ./

RUN go build -o tetris_game

CMD ["./tetris_game"]