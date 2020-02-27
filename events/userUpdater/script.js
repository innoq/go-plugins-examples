(function(){
    console.log("event: " + JSON.stringify(data));
    if(data["type"] === "user") {
        console.log("Ok, I will also update the User DB!");
    }
})();
