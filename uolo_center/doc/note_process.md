# notes


# test signup
curl -v -X POST -d 'username=gonghuan' -d 'password=abcde' -d 'email=gonghuan.dev@gmail.com' localhost:3002/signup
{"id":"6489acf2-4f86-43dc-83ff-f7a17ec26b03","username":"gonghuan","email":"gonghuan.dev@gmail.com","password":"$2a$10$rq30.I8K0rbo2yL/djcp5ew210RDlviazl5.lycBd1EbGjdTJ.s4y"}

curl -v -X POST -d 'username=gonghuan' -d 'password=abcde' localhost:3002/login
{"expire":"2018-04-15T18:14:01+08:00","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjM3ODcyNDEsImlkIjoiNjQ4OWFjZjItNGY4Ni00M2RjLTgzZmYtZjdhMTdlYzI2YjAzIn0.5GCyMlxm1xH1SjjAZhhEzHp3iY3WK5tQf_gJRm2GFzE"}

curl localhost:3002/Test -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjM3ODY0MDUsImlkIjoiNjQ4OWFjZjItNGY4Ni00M2RjLTgzZmYtZjdhMTdlYzI2YjAzIn0.B9Uvn_wGq1eH4oyFeOrTUzoLoKhyAF1o-juLz0HM5M0"
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjM3ODY0MDUsImlkIjoiNjQ4OWFjZjItNGY4Ni00M2RjLTgzZmYtZjdhMTdlYzI2YjAzIn0.B9Uvn_wGq1eH4oyFeOrTUzoLoKhyAF1o-juLz0HM5M0