# Groupie Tracker

## Description

This project is using given API (https://groupietrackers.herokuapp.com/api) which includes 4 parts:
  - The first one, `artists`, containing information about some bands and artists like their name(s), image, in which year they began their activity, the date of their first album and the members.

  - The second one, `locations`, consists in their last and/or upcoming concert locations.

  - The third one, `dates`, consists in their last and/or upcoming concert dates.

  - And the last one, `relation`, links between the other 3 parts. (`artists`, `locations`, `dates`)

This data is used to build a website which allows users to access the information about the groups, specifically their name, creation date, first album release date, and the location and dates for their concerts.

## Usage

### Run a program:

```
go run ./cmd 
```

Follow the link in the terminal (ctrl + left mouse button):

```
Starting server on http://localhost:8080
```

### Run tests:

```
go test .
```

### Team:
  - [skabidul](https://01.alem.school/git/skabidul)  
  - [abolat](https://01.alem.school/git/abolat)  