// This file is loaded by index.html and runs in the browser. 
// It fetches the weather data from the serverless function and updates the DOM with the temperature.
async function loadWeather() {

    // Read coords from HTML input fields
    const latitude = document.getElementById('latitude').value;
    const longitude = document.getElementById('longitude').value;

    // Fetch the weather data from the serverless function
    const weatherUrl = `/api/weather?lat=${latitude}&lon=${longitude}`
    console.log("weatherUrl =" + weatherUrl);
    const weatherResponse = await fetch(weatherUrl);

    // Parse the JSON response
    const weatherData = await weatherResponse.json();
    console.log("data from /api/weather=" +  JSON.stringify(weatherData));

    // Update the DOM with the temperature
    document.getElementById('temp').innerHTML = 
        'Latitude: ' + weatherData.latitude + '<br/>' + 
        'Longitude: ' + weatherData.longitude + '<br/>' +
        'Temperature: ' + weatherData.temperature_2m + '°C'
}

//can be called automatically but we'll call it on button click instead
// loadWeather();
