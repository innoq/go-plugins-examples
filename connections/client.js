(function(){
    request = {
        "oauth1": {
            "consumerKey": env("TWITTER_API_KEY"),
            "consumerSecret": env("TWITTER_API_SECRET_KEY"),
            "accessToken": env("TWITTER_ACCESS_TOKEN"),
            "accessSecret": env("TWITTER_ACCESS_TOKEN_SECRET")
        },
        "host":"https://api.twitter.com/1.1/search/tweets.json?q=from%3Atwitterdev&result_type=mixed&count=2"
    }
    response = GET(request);
    var body = response["body"];
    for(i in body["statuses"]) {
        var status = body["statuses"][i];
        LOG(status["created_at"] + " @"+status["user"]["screen_name"] + ": " + status["text"]);
    }

})();