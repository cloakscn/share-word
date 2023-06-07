const { Server } = require("socket.io");
const { createServer } = require("http");

const httpServer = createServer();
const io = new Server(httpServer, {
  cors: {
    origin: "http://127.0.0.1:8080"
  }
});

io.on("connection", (socket) => {
    socket.on('join', (room) => {
        socket.join(room);
        let myRoom = io.sockets.adapter.rooms;
        // let users = Object.keys(myRoom.sockets).length;
        // console.log(users);
        // 仅回复本人
        socket.emit('joined', room, socket.id);
        // 除自己之外所有人
        // socket.to(room).emit('joined', room, socket.id);
        // 房间内所有人
        // io.in(room).emit('joined', room, socket.id);
        // 全站广播
        // socket.broadcast.emit('joined', room, socket.id);
    })

    socket.on('leaved', (room) => {
        let myRoom = io.sockets.adapter.rooms(room);
        let users = Object.keys(myRoom.sockets).length;
        socket.leave(room);
        users -= 1;
        console.log(users);
        // 仅回复本人
        // socket.emit('joined', room, socket.id);
        // 除自己之外所有人
        socket.to(room).emit('leaved', room, socket.id);
        // 房间内所有人
        // io.in(room).emit('joined', room, socket.id);
        // 全站广播
        // socket.broadcast.emit('joined', room, socket.id);
    })
});

io.listen(3000);