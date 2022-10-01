

# GlobeScanner



![Globescanner1](https://user-images.githubusercontent.com/56683039/164520810-3cbd8bdd-96c0-4f3c-b68f-05a40cd45059.gif)





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

## Link to Deployed application:
http://107.20.193.99:4200

The application has been deployed in AWS using EC2.

## Demo video functionality:
https://user-images.githubusercontent.com/31062202/164365612-a338115d-38d4-41fc-bcbe-124ba68bf686.mp4

## Cypress Testing:
https://user-images.githubusercontent.com/31062202/164363001-72815b66-8167-4342-bcfa-09477c2e969e.mp4

## Backend Testing:
https://user-images.githubusercontent.com/39860389/164360088-8496e629-7af3-482f-94ad-c6f33639f8f4.mp4

## Link to API Documentation:
https://github.com/arijitd97/GlobeScanner/wiki/GlobeScanner-API-documentation

## Link to Project Board:

Sprint 1: https://github.com/arijitd97/GlobeScanner/projects/1

Sprint 2: https://github.com/arijitd97/GlobeScanner/projects/2

Sprint 3: https://github.com/arijitd97/GlobeScanner/projects/3

Sprint 4: https://github.com/arijitd97/GlobeScanner/projects/4

## Link to Sprint4 deliverables:
[Sprint4.md](https://github.com/arijitd97/GlobeScanner/blob/main/Sprint4.md) -- [https://github.com/arijitd97/GlobeScanner/blob/main/Sprint4.md](https://github.com/arijitd97/GlobeScanner/blob/main/Sprint4.md)


## Team Members:
1. Harshitha Patel Velpula (harshitha0422)- Frontend
2. Arijit Dutta (arijitd97) - Frontend
3. Vishesha Sadu(Vishesha)- Backend
4. Meghamala Gupta (Meghamala26) -Backend
 

## Tech Stack:
1. Frontend - Angular JS
2. Backend - Go


# How to run the project

Clone the project from https://github.com/arijitd97/GlobeScanner.git

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
