
terminal = document.querySelector(".card-content");
terminal.scrollTop = terminal.scrollHeight;
window.SpeechRecognition = window.webkitSpeechRecognition || window.SpeechRecognition;

function sendVoice(str) {
  var xhttp = new XMLHttpRequest();
  xhttp.open("POST", "/voiceCtrl", true);
  xhttp.onreadystatechange = function() {
    if (!(this.readyState == 4 && this.status == 200)) {
        // alert("You said: " + str + "\n\n" + "Server did not respond");
    }
  };
  xhttp.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  xhttp.send("voiceStr=" + str);
}


function startVoice() {
  var textDisp = document.querySelector("#cmd");
  console.log("Starting Voice Recording...");
  if (window.hasOwnProperty('webkitSpeechRecognition')) {

    var recognition = new webkitSpeechRecognition();

    recognition.continuous = false;
    recognition.interimResults = false;

    recognition.lang = "en-US";
    recognition.start();

    recognition.onresult = function(e) {
      console.log(typeof(e.results[0][0].transcript));
      // alert(e.results[0][0].transcript);
      textDisp.innerHTML = e.results[0][0].transcript;
      sendVoice(e.results[0][0].transcript);
      recognition.stop();
    };

    recognition.onerror = function(e) {
      recognition.stop();
    }

  }
}
    

document.querySelector("#mic").onclick = startVoice;