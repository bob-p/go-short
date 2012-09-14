A link shortner in Go
=====================

POST /create link=http:google.com (set content type as application/x-www-form-urlencoded)

Short url will be returned

GET /url-key

Deploy to heroku
----------------

This assumes you have go installed/setup aleady.

  git clone git@github.com:bob-p/go-short.git

  cd go-short

  go install
  
  git add *

  git commit -a -m 'Added binary'

  heroku create -s cedar --buildpack git://github.com/kr/heroku-buildpack-go.git

  git push heroku master
