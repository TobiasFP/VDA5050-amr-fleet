# BotNana

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

### Design decisions

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

Also, make sure to create a new user and a password for this user, under the tab "users".

Now, spin up the frontend for this project, found here:
https://github.com/TobiasFP/BananaUI

When you have spun it up, simply go to its frontend, on:
http://localhost:8100/home

This should first redirect you to the backend located at https://localhost:8002, which then redirects you to the keycloak Oauth realm botnana. Login and grant access. Now you have "securely" accessed the frontend, and we can now make secure rest requests to the backend.
