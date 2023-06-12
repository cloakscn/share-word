'use strict'

let peerA = document.querySelector('video#peera')
let peerB = document.querySelector('video#peerb')

let startBtn = document.querySelector('button#start')
let callBtn = document.querySelector('button#call')
let stopBtn = document.querySelector('button#stop')

let sdpA = document.querySelector('textarea#sdpa')
let sdpB = document.querySelector('textarea#sdpb')

let peerACoon
let peerBCoon
let localStream

startBtn.onclick = start;
callBtn.onclick = call;
stopBtn.onclick = stop;

function stop() {
    peerACoon.close();
    peerBCoon.close();
    peerA.srcObject = null;
    peerB.srcObject = null;
}

function getRemoteStream(e) {
    peerB.srcObject = e.streams[0];
}

function call() {
    callBtn.disabled = true

    peerACoon = new RTCPeerConnection();
    peerBCoon = new RTCPeerConnection();
    // 收集 candidate
    peerACoon.onicecandidate = (e) => {
        peerBCoon.addIceCandidate(e.candidate)
    };
    peerBCoon.onicecandidate = (e) => {
        peerACoon.addIceCandidate(e.candidate)
    };

    peerBCoon.ontrack = getRemoteStream;

    localStream.getTracks().forEach(element => {
        peerACoon.addTrack(element, localStream);
    });

    // 媒体协商
    var offerOptions = {
        offerToReceiveAudio: 1,
        offerToReceiveVideo: 1,
    };
    peerACoon.createOffer(offerOptions)
        .then(getOffer)
        .catch(handleDescError);
}

function getOffer(desc) {
    peerACoon.setLocalDescription(desc);

    sdpA.value = desc.sdp;

    // send desc to signal
    // receive desc from signal

    peerBCoon.setRemoteDescription(desc);

    peerBCoon.createAnswer()
        .then(getAnswer)
        .catch(handleAnswerError);
}

function getAnswer(desc) {
    peerBCoon.setLocalDescription(desc);

    sdpB.value = desc.sdp;

    // send desc to signal
    // receive desc from signal

    peerACoon.setRemoteDescription(desc);
}

function handleAnswerError(err) {
    console.log("getAnswer error:", err);
    callBtn.disabled = false
}

function handleDescError(err) {
    console.log("getLocalDescription error:", err);
    callBtn.disabled = false
}

function start() {
    if (!navigator.mediaDevices ||
        !navigator.mediaDevices.getUserMedia) {
        console.log("getUserMedia is not supported!");
    } else {
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
        }
        navigator.mediaDevices.getUserMedia(constraint)
            .then(gotUserMediaSteam)
            .catch(handleUserError);
    }
}

function gotUserMediaSteam(stream) {
    peerA.srcObject = stream
    localStream = stream
    startBtn.disabled = true
    callBtn.disabled = false
    stopBtn.disabled = false
}

function handleUserError(err) {
    console.log("getUserMedia error:", err);
    startBtn.disabled = false
}