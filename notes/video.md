### Managed Video

- videos list
- Filter videos by tag
- Edit videos
- Create a tagging job for a video
- Download tags for a video

### Creating a job

- Assigning a video to a user for tagging
- Configure the tagging job

### The Tagging

- Navigating frame by frame
- select areas on the frame, and tags for each area
- video controls
- review a tagging job
- sned a job for review by an Admin, or approve the tagging job:

### Video Controls

- Video Timebar
- Previouse frame
- Play/Pause
- Next frame
- Skip to the furst untagged frame
- frame index
- Play speed
- Timer
- Volume/Mute

## QNA service
---

### Setting up the service

- the bot app
- the bot framework
- the qan service
- ngrok
``` bash
ngrok http 5000
``` 

### Regiser Your Bot
### Buil the Bot


### Iot Simulator

- Compass Heading
- Geolocation(Lat，Long)
- X-Y-Z Accelerometer Values

``` c#
var eventBody = {
    Timetamp: new Date(),
    UserId: app.deviceId,
};

if (ENV.settings.sensors.geolocation && app.currentLocation) {
    eventBody["Latitude"] = app.currentLocation.latitude;
    eventBody["Longitude"] = app.currentLocation.longitude;
}

if (ENV.settings.sensors.accelerometer && app.accelerometer) {
    eventBody["X"] = app.accelerometer.x;
    eventBody["Y"] = app.accelerometer.y;
    eventBody["Z"] = app.accelerometer.z;
}

if (ENV.settings.sensors.compass && app.currentHeading) {
    eventBody["MagneticHeading"] = app.heading.magneticHeading;
    eventBody["TrueHeading"] = app.heading.trueHeading;
    eventBody["HeadingAccuracy"] = app.heading.headingAccuracy;
}

var message = new EventData(eventBody);
```

``` javascript
startPositionWatch: function() {
    if (app.locationWatchId) {
        app.stopPositionWatch();
    }
    // Watch foreground location
    app.locationWatchId = navigator.geolocation.watchPosition(function(position) {
        app.updateLocation(location.coords);
    }, function(){}, {
        enableHighAccuracy: true,
        timeout: 30000,
        maximumAge: 0
    });
},

stopPositionWatch: function() {
    if (app.locationWatchId) {
        navigator.geolocation.clearWatch(app.locationWatchId);
        app.locationWatchId = null;
    }
},

updateLocation: function(coords) {
    console.log("updateLocation");

    if (!app.mapLayers) {
        return;
    }