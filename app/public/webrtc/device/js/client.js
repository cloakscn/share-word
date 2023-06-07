'use strict'

let audioSource = document.getElementById("audioSource");
let audioOutPut = document.getElementById("audioOutPut");
let videoSource = document.getElementById("videoSource");

if (!navigator.mediaDevices ||
    !navigator.mediaDevices.enumerateDevices) {
    console.log("enumerateDevices is not supported!");
} else {
    navigator.mediaDevices.enumerateDevices()
        .then(gotDevices)
        .catch(handleError);
}

function gotDevices(deviceInfos) {
    deviceInfos.forEach(element => {
        let option = document.createElement("option");
        option.text = element.label;
        option.value = element.deviceId;

        switch (element.kind) {
            case "audioinput":
                audioSource.appendChild(option)
                break;
            case "audiooutput":
                audioOutPut.appendChild(option)
                break;
            case "videoinput":
                videoSource.appendChild(option)
                break;
            default:
        }
    });
}

function handleError(err) {
    console.log("enumerateDevices error:", err);
}