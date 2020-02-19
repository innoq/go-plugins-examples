(function(){

    LOG("bad call!");
    request = {
        "host":"https://example.com"
    }
    response = POST(request);
    LOG("response: " + JSON.stringify(response));

    LOG("good call!");
    request = {
        "oauth1": {
            "consumerKey": env("TWITTER_API_KEY"),
            "consumerSecret": env("TWITTER_API_SECRET_KEY"),
            "accessToken": env("TWITTER_ACCESS_TOKEN"),
            "accessSecret": env("TWITTER_ACCESS_TOKEN_SECRET")
        },
        "host":"https://api.twitter.com/1.1/search/tweets.json?q=from%3Atwitterdev&result_type=mixed&count=2",
        "contentType":"application/json",
        "content":""
    }
    response = POST(request);
    LOG("response: " + JSON.stringify(response));

})();