'use strict'

let username = document.querySelector('input#username');
let room = document.querySelector('input#room');
let roomBtn = document.querySelector('button#roomBtn');
let content = document.querySelector('textarea#content');
let input = document.querySelector('textarea#input');
let inputBtn = document.querySelector('button#inputBtn');

let socket;
let roomObj;
roomBtn.onclick = () => {
    // connect
    socket = io.connect('http://127.0.0.1:3000');
    // message
    socket.on('joined', (room, id)=> {
        roomBtn.disabled = true;
        input.disabled = false;
        inputBtn.disabled = false;
    });

    socket.on('leaved', (room, id)=> {
        roomBtn.disabled = false;
        input.disabled = true;
        inputBtn.disabled = true;
    });

    socket.on('message', (room, id, data)=> {
        content.value += data + '\n';
    });
    // send
    roomObj = room.value
    socket.emit('join', roomObj)
}

inputBtn.onclick = () => {
    const data = input.value;
    data = username.value + ':' + data;
    socket.emit('message', roomObj, data);
    input.value = '';
}
