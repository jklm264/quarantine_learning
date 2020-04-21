# FlaskDB

## What does this do?

- This app was meant to explore flask's API with Flask_SQLAlchemy.
- THere are two routes:
    - Main ('/') which will list [aka `query.all()`] the information that is in the db.
    - '/new' will allow a user to POST a new man to the db and then will redirect you to the '/' page.

## Creating Env

1. `$virtualenv -p /usr/bin/python3 ./path_of_virtenv`
2. `$source ./path_of_virtenv/bin/activate`
3. `$export FLASK_APP=run.py`
4. `$export FLASK_ENV=development`
5. `$flask run` # If changing python (not html), will need to rerun this cmd
6. `$deactivate`

## Creating APIs

- You cannot use the -H "Authorization.." and the -u <user:pass> switches in curl when parsing like request.authorization.username. Or can find via -H "Authorization..." with request.headers.get('Authorization)
- [An api POST source](https://blog.miguelgrinberg.com/post/designing-a-restful-api-with-python-and-flask)
- https://programminghistorian.org/en/lessons/creating-apis-with-python-and-flask

### Troubleshooting

- Use the following command if Flask says port already in use: `lsof -i -P -n | grep LISTEN then kill <pid>`

## Running the app

- Flask is very name specific so make sure templates dir is in the root directory, else you will have to overwrite the path- which is a hassel.This includes images (serve static files), flask will look for the `/static` file
- Also, do not name your `{root}.py` app.py (especially if you have a subfolder called `app/`). The preference is to name it `run.py` and then make sure to run `$export FLASK_APP='run.py'`.
- To get app to reload on change, add the following to `run.app`: `app.run(debug=True)` and `$export FLASK_ENV=development`
- When running for first time, run the `runFlash.sh` script to export the needed variables.

## Bootstrap 4 navbar tips

- <https://getbootstrap.com/docs/4.0/components/navbar/>

## Jinja Templating

- <https://code.tutsplus.com/tutorials/templating-with-jinja2-in-flask-essentials--cms-25571>
- Note, that if the child template doesn’t define the footer block, the value from the parent template is used instead.

## Figure out later

- How to clean up database during shutdown. After every new run you need to delete the .sqlite3 db before running the next one again.
- How to add an existing db so you don't have to repopulate it every time!
- Add an error option to api return
- Check if topic is already there. Should be interactive at this point to allow user to then agree to overwrite the change, else dont bother updating OR query create a new entry with a modified topic name.
- Add security and permission for different users
    - Use secret key somehow

## Helpful Sources

- [Best for understanding POST](https://www.youtube.com/watch?v=kvux1SiRIJQ)
- [Another very helpful POST video](https://www.youtube.com/watch?v=qH--M56OsUg)