# API FOR SEARCHING USERS

### This is a `go api` project searching for users, giving likes and searching for matches.

So far there are a few endpoints in this API
1) http://localhost/users
2) http://localhost/users/like?userid=1&likeduser=5
3) http://localhost/users/search?id=1
4) http://localhost/users/matches?id=1
5) http://localhost/users/preferences?id=1
6) http://localhost/users/life?id=1

To run the application, download the project and run with the following command
`go run main.go`

This will start the server and you can navigate to you web browser and use any of the endpoints above. 

FYI, I am still working on the tests, and also the matching routine. I may have over complicated it for now but it can be impoved

### Good coding practices
I have used abstractions quite a bit in this project, mainly because i am using `sqlite`. The main reason for this is to have a quick database to get the project going.
I am aware you make use of MariaDB, that's great, if i were to use that, i would make use of the same interface `UserRepository` in this project and then just create the implementation of the interface.

### To Improve
- Create tests for this project, **TEST ARE ESSENTIAL**
- Sort out the matching function, it's a bit tricky as i've tried to create my own algorithm and may have over complicated it in the meanttime, i could work on this later. 
- I am concious of the time and I do want to have something for you to look at.