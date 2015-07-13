(function() {
    // Set-up the speech synthesis API

    var msg = new SpeechSynthesisUtterance();
    var voices = window.speechSynthesis.getVoices();
    console.log(voices)
    msg.voice = voices[5]; // Note: some voices don't support altering params
    msg.voiceURI = 'native';
    msg.volume = 1; // 0 to 1
    msg.rate = 0.20; // 0.1 to 10
    msg.pitch = 1; //0 to 2
    msg.lang = 'pl-PL';


    var API_BASE = "api/1/"

    function apiCall(endpoint, params, cb) {
        var req = new XMLHttpRequest();
        req.open("GET", API_BASE + endpoint, true);
        req.onload = cb;
        req.send();
    }

    function getDate(cb) {
        apiCall("random_date", null, function() {
            var date = JSON.parse(this.responseText);
            cb(date.polish_day + " " + date.polish_year);
        });
    }

   var ListenToDate = React.createClass({displayName: "ListenToDate",
     speakDate: function() {
        msg.text = this.props.date;
        speechSynthesis.speak(msg);
     },
     render: function() {
       return (
        React.createElement("div", null, 
            React.createElement("button", {onClick: this.speakDate}, "Listen again"), 
            React.createElement("div", null, 
                React.createElement("label", null, "Enter the date you hear (dd/mm/yyyy)"), 
                React.createElement("input", {type: "text"})
            ), 

            React.createElement("button", null, "Check")
        ));
     }
   });


   getDate(function(d) {
         React.render(React.createElement(ListenToDate, {date: d}),
              document.getElementById('container')
         );
   });


}());
