'use strict'

// device
let audioinput = document.querySelector("select#audioinput");
let audiooutput = document.querySelector("select#audiooutput");
let videoinput = document.querySelector("select#videoinput");
let filter = document.querySelector('select#filter');
// filter
// tools
let snapshot = document.querySelector('button#snapshot');
let record = document.querySelector('button#record');
let play = document.querySelector('button#play');
let download = document.querySelector('button#download');
// snapshot
let picture = document.querySelector('canvas#picture');
picture.width = 320;
picture.height = 240;
// video
let video = document.querySelector('video#player');
let recordVideo = document.querySelector('video#record');
let displayVideo = document.querySelector('video#displayVideo');
// audio
let audio = document.querySelector('audio#audio');
// constraint
let constraint = document.querySelector('div#constraint');

function start() {
    if (!navigator.mediaDevices ||
        !navigator.mediaDevices.getUserMedia) {
        console.log("getUserMedia is not supported!");
    } else {
        let deviceId = videoinput.value
        let constraint = {
            video: {
                width: 320,
                height: 240,
                frameRate: 60,
            },
            audio: {
                noiseSuppression: true,
                echoCancellation: true
            },
            deviceId: deviceId ? deviceId : undefined
        }
        navigator.mediaDevices.getUserMedia(constraint)
            .then(gotMediaSteam)
            .then(gotDevice)
            .catch(handleError);
    }
}

function gotDevice(deviceInfos) {
    deviceInfos.forEach(element => {
        let option = document.createElement("option");
        option.text = element.label;
        option.value = element.deviceId;

        switch (element.kind) {
            case "audioinput":
                audioinput.appendChild(option)
                break;
            case "audiooutput":
                audiooutput.appendChild(option)
                break;
            case "videoinput":
                videoinput.appendChild(option)
                break;
            default:
        }
    });
}

function gotMediaSteam(stream) {
    // 获取配置信息
    let videoTrack = stream.getVideoTracks()[0];
    let videoConstraint = videoTrack.getSettings();
    constraint.textContent = JSON.stringify(videoConstraint, null, 2);

    window.stream = stream
    // 视频流
    video.srcObject = stream
    // 音频流
    audio.srcObject = stream

    return navigator.mediaDevices.enumerateDevices()
}

function handleError(err) {
    console.log("getUserMedia error:", err);
}

start();

if (!navigator.mediaDevices ||
    !navigator.mediaDevices.getDisplayMedia) {
    console.log("getDisplayMedia is not supported!");
} else {
    let constraint = {
        video: {
            width: 2400,
            height: 1800,
            frameRate: 60,
        },
        audio: {
            noiseSuppression: true,
            echoCancellation: true
        },
    }
    navigator.mediaDevices.getDisplayMedia(constraint)
        .then(gotDisplayMediaSteam)
        .catch(handleError);
}

function gotDisplayMediaSteam(stream) {
    // 视频流
    displayVideo.srcObject = stream
}

videoinput.onchange = start;

filter.onchange = function () {
    video.className = filter.value
}

// snapshot

snapshot.onclick = function () {
    picture.className = filter.value
    picture.getContext('2d').drawImage(video, 0, 0, picture.width, picture.height)
}

// record

record.onclick = () => {
    if (record.textContent === 'record') {
        startRecord();
        record.textContent = 'recording...';
        play.disabled = true;
        download.disabled = true;
    } else {
        stopRecord();
        record.textContent = 'record';
        play.disabled = false;
        download.disabled = false;
    }
}

let buffer;

function handleDataAvailable(e) {
    if (e && e.data && e.data.size > 0) {
        buffer.push(e.data);
    }
}

let mediaRecorder

function startRecord() {
    let options = {
        miniType: 'video/webm;codecs=vp8'
    }
    buffer = [];
    if (!MediaRecorder.isTypeSupported(options.miniType)) {
        return console.error(`${op.miniType} is not supported!`);
    }

    try {
        mediaRecorder = new MediaRecorder(window.stream, options);
    } catch (e) {
        console.error('Failed to create MediaRecorder:', e);
    }

    mediaRecorder.ondataavailable = handleDataAvailable;
    mediaRecorder.start(10)
}

function stopRecord() {
    mediaRecorder.stop();
}

play.onclick = () => {
    let blob = new Blob(buffer, {type: 'video/webm'});
    recordVideo.src = window.URL.createObjectURL(blob);
    recordVideo.srcObject = null;
    recordVideo.controls = true;
    recordVideo.play();
}

download.onclick = function() {
    let blob = new Blob(buffer, {type: 'video/webm'});
    let url = window.URL.createObjectURL(blob);
    let a = document.createElement('a')
    a.href = url
    a.style.display = 'none';
    a.download = 'aaa.webm';
    a.click();
}