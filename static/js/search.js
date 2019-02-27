function httpGetAsync(theUrl, callback)
{
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.onreadystatechange = function() {
        if (xmlHttp.readyState === 4 && xmlHttp.status === 200)
            callback(xmlHttp.responseText);
    };
    xmlHttp.open("GET", theUrl, true); // true for asynchronous
    xmlHttp.send(null);

    return xmlHttp;
}

function suggest(url) {
    return function(term, response){
        httpGetAsync(url+'?q=' + term, function(data){
            data = JSON.parse(data);
            document.querySelector(".stats span").innerHTML = data.stats;

            response(data.suggest);
        });
    }
}

new autoComplete({
    selector: '#builder_name',
    minChars: 1,
    delay: 1,
    source: suggest("/suggest/builder_name"),
});

new autoComplete({
    selector: '#model_name',
    minChars: 1,
    delay: 1,
    source: suggest("/suggest/model_name"),
});
