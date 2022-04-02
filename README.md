# alist-heroku-postgresql


## Deploy alist to heroku
[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

Use heroku's add-on postgres database, your settings will be persistent, don't worry about hibernate losing configuration.
If you can't deploy, and it says "We couldn't deploy your app because the source code violates the Salesforce Acceptable Use and External-Facing Services Policy.", you need to fork this repo and click the "Deploy" button in your own fork.

## Get Password
`More` -> `View logs` -> You will see your password, if it is scrolled to the top and out of view, click `Restart all dynos` and the log will be redisplayed.