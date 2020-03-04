(function(){

    request = {
        "host":"https://google.com"
    }
    response = GET(request);
    LOG("response: " + JSON.stringify(response));

    request = {
        "host":"http://example.com/api"
    }
    response = GET(request);
    var body = response["body"];
    LOG("response: " + body);
})();
