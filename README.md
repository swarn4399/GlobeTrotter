

# GlobeScanner


## Introduction:

GlobeScanner brings in a platform to connect the agencies providing tour packages with the people planning their trips. Using Globescanner, it will be easier for local tour guides to market their tour packages to the tourists directly and the tourists will also be able to make an informed choice on which tour package or tour guide will suit their budget and requirements to the best. It is also used to locate the places to visit in an area helping users plan their trips better. Know what other users have to say, plan your trip accordingly.

## Motivation:

The United Nations World Tourism Organization (UNWTO) estimates that internationally there were just 25 million tourist arrivals in 1950. 68 years later this number has increased to 1.4 billion international arrivals per year. This is a 56-fold increase. Thus, it is necessary that the tourists can plan their trip effectively and find the places they can visit and their reviews on a single platform. Also, with the boom in the tourism sector many agencies and small scale tourist guides need a platform to advertise their deals.

## User roles
1. Tourist - People who wants to plan a trip and are searching for places to visit along with tour packages.
2. Tour guide - People/Agencies that provide tour packages and aid tourists visiting a location.

## Functionality:
1. Users can register as a tourist or guide to the application with their username, email, password and role(Tourist/Guide).
2. Users have to login with their email,password and role(Tourist/Guide) to the application.
3. Tourists and Tour guides can look up for the places to see in an area.
4. Tourists can view tour packages available for the place they have searched.
5. Tourists can book any available tour package.
6. Guides can add tour packages for any location and also view tour packages already added for the same location.
7. Tourists and guides will both have access to see their account profile. 
8. Tourists and guides will both have access to edit details on their account profile.
9. Tourists can see their booked packages on their profile.
10. Guides can see the packages they have added on their profile.
11. Tourists can add reviews of any tour packages that they have booked.


## Tech Stack:
1. Frontend - Angular JS
2. Backend - Go

The application has been deployed in AWS using EC2. 


# How to run the project

Clone the project from https://github.com/swarn4399/GlobeTrotter.git

## Backend

Go to server folder and run the command 'go run .' to start the server.

Command to test backend in server folder - go test


## Frontend

run the below steps in the terminal-</br>
npm install</br>
ng serve --o

#### Cypress Tests

Steps:-

npm install cypress

In package.json :-

"scripts": { [....] "cypress:open": "cypress open", "cypress:run": "cypress run", "cypress": "cypress open" },

To Run :-

npm run cypress
