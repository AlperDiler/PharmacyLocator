# Pharmacy Locator
Hi everyone! This project provides a function to find pharmacies closest to user's location and return them as an HTTP response. 

We have 3 files in the project. These files are controllers, models and routers .

In Models file i define the Pharmacy and UserCoords structs. While the Pharmacy struct provides my relationship with my database, our UserCoords structure holds the location data I receive from the user. 

Routers file sets up routing for a web application using the Gorilla Mux router and configures it to handle specific HTTP requests.

And the Controllers file defines a function for finding the nearest pharmacies to a user's location. The function uses a combination of the geo-golang and geodist packages to calculate distances between the user and various pharmacies. It performs the following tasks:

# Parse Coordinates:

The parseCoordinates function extracts latitude and longitude from the 'Koordinat' field of each pharmacy. It splits the coordinate string by a comma, trims whitespace, and converts the strings to float values. The parsed latitude and longitude are then stored in the Latidute and Longidute fields of the Pharmacy struct.

# Find Nearest Pharmacies:

The FindNearestPharmacies function is an HTTP handler that:

a)Verifies if the database connection (db) is initialized.
b) Reads the user's location data from the request body and parses it into latitude and longitude.
c) Queries the database to fetch all pharmacies.
d) Computes the distance between the user’s location and each pharmacy using the Haversine formula. It creates a slice of distancePharmacy structs that hold pharmacy details and their respective distances from the user. 
e) Sort and Limit Results: Sorts the pharmacies by distance and limits the result to the top 5 closest pharmacies.
f) Sets the response header to application/json and encodes the result into JSON format, which is sent back to the client.

# Output 

I used Postman API platform to see the output of my code. So here are the result : 

![github1](https://github.com/user-attachments/assets/ac1f6fc7-4480-435f-bba3-af9baf4494be)


I used SQLite3 database in the project. If you want to check it out : 

![github](https://github.com/user-attachments/assets/397831ca-636f-41e0-9ab7-ce57e37647f0)


Thanks for your time. 
