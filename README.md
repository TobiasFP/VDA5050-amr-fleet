# BotNana - A VDA5050 compliant amr fleet
Public
##### Currently, this is a work in progress, and basically just a bunch of my thoughts.

A VDA 5050 compliant fleet software engine

## Why

I am creating this project, to create my own job.
I love Robots, and I love hooking up the real world
with interesting software.
This is my attempt at creating a Fleet software engine,
that complies with VDA5050 for scalability, but it will
also allow for addition of robots who do not comply with
vda5050 in a "vendor to vda5050" mapper approach.
I release this as open source, but I really hope you will
hire me to help facilitating and integrating this software
with your organisation.

## Purpose

The purpose of BotNana is four fold:
* To create an engine/platform which Robot developers can build upon (Like ROS for fleet management but without all the bad decisions - oh well, I will probably make a few.)
* To make a complete, extendable, fleet management system that is ready for the VDA5050 revolution
* To show off my awesome GoLang skills
* To land me a consulting job at any interesting robotics firm 

## How

I am certain of my skills as a backend and robotics developer, but i am not good at UI/UX, so the overall approach is too make an awesome vda5050 compliant backend service, and little by little add vda5050 features. Meanwhile, i will design a fleet frontend in Angular, I will write proper, maintainable code, but the UI/UX design decisions will most likely be very bad. Therefore, the frontend will most likely look horrible, but have all the functionalities needed to prove that the fleet software is working and working well.

The code will be written test drivenly, and some integration tests will be added, especially with mqtt recordings to prove the state of the system. 

A github actions pipeline will be setup for ci/cd integration.

A closed source vda5050 simulator is also under development and the artifacts from this simulator will be used to do integration tests on a build/release basis. 


## General approach

In order for this software to work, I need three things:
A robot speaking VDA5050 (Simulated at first),
a MQTT broker between the robot and this software and to
facilitate easy use, a frontend for this software:
https://github.com/TobiasFP/BananaUI

(Also, an OAuth setup, as this is just currently
the right way to go in terms of security)

The frontend is meant to be an example frontend that will
be usable, but I hope you will hire me (Tobias) to do a
proper front-end for your end users.

### High level design decisions

#### Rest and mqtt, why both?

Plane and simple, to keep complexity low. With a fleet management system, it will be impossible to make a system void of complexity, but to lower the amount of complexity, separation of concern will be vital. Therefore, the mqtt setup is in place to manage the robots with the vda5050 compliant protocols, and REST will be the “human” facing interface that can be used to present data, and to initiate actions in the fleet. 

Therefore, there will also be two databases (both MySql, but can be swapped out with postgres or even a nosql database, as the ORM “GORM” will be used) in the system, one “mqtt” database to persist and manage mqtt data, and a user facing database, where user settings, etc. Can be accessed. 
The mqtt server will have read/write access to the mqtt database, but not have access to the REST database.
The REST server will have read/write access to the REST database, and will only have Read access to the MQTT database. 
This also means that the REST server is completely optional for the fleet to function. This means that if someone wants to take the Spartan/puritan approach and make everything with mqtt, this can easily be added on top. 

Having the REST api will make the design of a front end very easy, and will also make sure that the fleet itself will be able to facilitate fleet-to-fleet interoperability. For the REST api to do anything that sends mqtt messages, it has to call functions in the mqtt layer.


#### How do we integrate vda5050 non-compliant robots 

Well, we don't. That being said, we do allow integrations by having an integrations folder, where “vda5050 mqtt to vendor specific interfaces” can be implemented. 

#### Structure

VDA5050 will of course be implemented with MQTT, all MQTT handlers
will be easily monitored and inspected in the MQTT logs.
From the MQTT messaging, we will be able to extract data,
and serve this data via REST for end-user decisions.
An example of this is, from the MQTT messages, we will
be able to generate a list of active robots, keep that in memory,
and display to an end-user via REST.
This is the first feature to be implemented.

#### Why Golang

GoLang was chosen for its easy of readability, coder conformity
and because I just really love working in GoLang.
The benefits to you is that if you have developers who know
C++, C# or any C-like languages, GoLang is so easy to learn
and read, they can get up and running in no-time.

Other aspects of why GoLang is nice for this project is it's
way of handling concurrence, as I will write this in a very TDD
approach, everything will be able to be run synchronously,
for easy use and testing in a real life setting, but will be
made to run asynchronously as a "whole fleet package"

## Name origins

The word BotNana was first thought up as I really like Minions,
and therefore, the name is a tribute to them as the word sounds
like Banana. When the name was selected, it had not crossed my
mind that it could also be considered a "Nanny" of bots.
What a lovely coincidense.

A third niceness is it's tribute to one of the all time greats,
Boten Anna by Basshunter.

## Setup

### Get up and runnning

Copy ./config/.env.example into your root project and rename it to .env
copy ./config/example.yaml to ./config/development.yaml
Edit the two files with some sensible input

docker compose up -d

After this, run this project, either with "go run main.go" or hitting F5 in VSCode.

#### Setup of keycloak (If you have another OAuth provider, feel free to use that one.)

Now, open the following page in a webbrowser:
http://localhost:7080

Login with yuor admin credentials, create a new realm named botnana (refer to the keykloak documentation), and setup a new client under this realm just like it has been setup in the following image, that can be found here, in this repo:
info/realm-client-settings.png

Under the client, go to credentials, and use the client secret in the file "routes/rest/routes.go" and put in the var: "clientSecretDev".
Client Secret Dev should already be set to botnana, but if you want to have another name for this, just set it there.

This will be parametrized later.

Also, make sure to create a new user and a password for this user, under the tab "users".

Now, spin up the frontend for this project, found here:
https://github.com/TobiasFP/BananaUI

When you have spun it up, simply go to its frontend, on:
http://localhost:8100/home

This should first redirect you to the backend located at https://localhost:8002, which then redirects you to the keycloak Oauth realm botnana. Login and grant access. Now you have "securely" accessed the frontend, and we can now make secure rest requests to the backend.


## Business model

### So, how do we make money? You do this for the money, right? Right? 

Well, currently i have no idea which way I'm going to go with this. I have some ideas as to how, which are: 
1. All in open source, hope it grows and gets adopted by some robotics firms and hope build a consulting business
2. All in open source, no profits, just doing it because i think it's fun. 
3. The Odoo model: build the core features and release it properly, open source, but then make modules that are closed source that adds features
4. The Harness DRONE model: do everything in a restricted open source manner, that limits the commercial usage of the software, but keeps the code open source
5. Closed source, seeking investors: it will be hard to gain traction in such a competitive market. I am definitely open if some company wants to invest. 

I lean towards #1 or #2, as i like open source a lot and this means the project gets to be my hobby, and I can still make money by doing contract work. This means little to no chance of profiting but maybe, just maybe, my awesome code will show people how to do proper software and will serve as a foundation for some robotics platforms. 


#### Decision
I have decided on the following:
Keep the base GoLang project completely open source with the least restricted model possible (MIT).

The Frontend and vda5050 amr simulator however will be kept open to view, but very restricted, in that if you want to use these purposes you have to contact me and make a deal. 

### Acquiring partners

There are so many small robotics startups that simply don't have the time to write fleet software. By telling them to just support the vda5050 from the start and pay me to add fleet features must be a lot cheaper and easier than building everything from scratch. By contacting robotics firms and convincing them to let me do the hard part will surely be an easy way of getting partners.

### Work management 

I am a big believer in the agile manifesto, and I absolutely loathe everything that steals the good agile name, so one thing i will ban and write into any contract is that using Jira is banned and doing any wannabe agile framework like SAFe that is waterfall in disguise will be doubly banned. If somebody wants to work on BotNana and they need specs and contracts i will definitely support it by doing waterfall, as waterfall is completely fine.

While I am only myself working on it, the work will be done by simply doing what i want to do, to further the project. 



## Low level design decisions

### Maps

The maps are, like ROS, represented by PGM/Pbm "Netpbm grayscale image format" images (https://netpbm.sourceforge.net/doc/pgm.html).

This format is very easy to work with and therefore also quite versatile.

We therefore also assume that the maps the robot will use is in pgm, or we will create a converter for this.

The actual rendering of the maps and adding AMRs to the map will be purely handled in the frontend. 
I have chosen to use the game engine "phaser" to assist me in especially the math portion of illustrating the map and AMR's.

#### Resolution

The resolution unit of a pgm map is set to 0.05m/pixel


Read more about VDA5050 maps here (which are not related to the file format at all):

https://github.com/VDA5050/VDA5050/blob/main/VDA5050_EN.md#67-maps
