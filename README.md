Message
===

A simple go service that replies with whatever it's environment variable `MESSAGE` has

Build `message`
```bash
go build -o message .
```

Get an echo message:

```bash
MESSAGE=lol ./message
```

then get your message

```bash
curl http://localhost:8090
```

that simple!

To start a distributed `message` try the `message-dist.yaml` file. It will start 4 message servers and try to ping all of them every second.

