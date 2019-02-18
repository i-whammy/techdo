{{define "hello"}}
<!DOCTYPE html>
<html>
<body>
    <form name="first">
    <input type="number" id="first"/>
    </form>
    <form name="second">
    <input type="number" id="second"/>
    </form>

    <button onClick="calc()">
        実行
    </button>
    <div id="result">
    </div>
</body>
<script>
function calc() {
    const url = '/plus';
    var first = document.getElementById('first');
    var second = document.getElementById('second');

    var newUrl = url + '?first=' + first.value + '&second=' + second.value

    fetch(newUrl).then(function(response) {
        return response.json();
    }).then(function(json) {
        var result = document.querySelector('#result');
        result.innerHTML = json.result;
    });
}

function kickApi() {
    const url = '/hello_api';
    fetch(url).then(function(response) {
        return response.json();
    }).then(function(json) {
        var result = document.querySelector('#result');
        var titleName = json.title_name;
        var textMessage = json.text_message;
        result.innerHTML = titleName + ", " + textMessage;
    });
}
</script>
</html>
{{end}}