A link shortner in Go
=====================

POST /create link=http:google.com (set content type as application/x-www-form-urlencoded)

Short url will be returned

GET /url-key

Deploy to heroku
----------------

  git clone git@github.com:bob-p/go-short.git 

  go install
  
  git add *

  git commit -a -m 'Added binary'

  heroku create

  git push heroku master
