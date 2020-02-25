(function(){
    console.log("event: " + JSON.stringify(event));
    if(event["updated_type"] === "user") {
        console.log("Ok, I will also update the User DB!");
    }
})();
